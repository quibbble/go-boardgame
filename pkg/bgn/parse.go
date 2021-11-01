package bgn

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// RequiredTags are the list of tags any game must include
var RequiredTags = []string{
	"Game",  // the name of the game being played
	"Teams", // list of teams playing the game
	"Seed",  // seed for deterministic randomness
}

func ParseGame(s *scanner.Scanner) (*BGN, error) {
	g := BGN{Tags: map[string]string{}, Actions: []Action{}}
	err := ParseTags(s, &g)
	if err != nil {
		return nil, err
	}
	for _, required := range RequiredTags {
		if g.Tags[required] == "" {
			return nil, fmt.Errorf("missing required %s tag", required)
		}
	}
	err = ParseActions(s, &g)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func ParseTags(s *scanner.Scanner, bgn *BGN) error {
	run := s.Peek()
	inside := false
	for run != scanner.EOF {
		switch run {
		case '[':
			if inside {
				return fmt.Errorf("cannot nest right bracket for tags")
			}
			inside = true
			run = s.Next()
		case ']':
			if !inside {
				return fmt.Errorf("missing starting right bracket for tags")
			}
			inside = false
			run = s.Next()
		case '\n', '\r', ' ':
			run = s.Next()
		default:
			if !inside {
				return nil
			}
			s.Scan()
			tag := s.TokenText()
			s.Scan()
			val := s.TokenText()
			bgn.Tags[tag] = strings.Trim(val, "\"")
		}
		run = s.Peek()
	}
	return nil
}

func ParseActions(s *scanner.Scanner, bgn *BGN) error {
	run := s.Peek()
	var action *Action
	for run != scanner.EOF {
		switch run {
		case ' ':
			run = s.Next()
		case '{':
			for run != '}' && run != scanner.EOF {
				run = s.Next()
			}
		default:
			s.Scan()
			if s.TokenText() == "-" {
				if action == nil {
					return fmt.Errorf("cannot have action details before base")
				}
				s.Scan()
				details := s.TokenText()
				for s.Peek() != ' ' && s.Peek() != scanner.EOF {
					s.Scan()
					if s.TokenText() == "-" {
						return fmt.Errorf("mutple dashes found in action")
					}
					details += s.TokenText()
				}
				split := strings.Split(details, ".")
				action.Details = split
				bgn.Actions = append(bgn.Actions, *action)
				action = nil
			} else {
				if action != nil {
					bgn.Actions = append(bgn.Actions, *action)
				}
				base := s.TokenText()
				for s.Peek() != ' ' && s.Peek() != '-' && s.Peek() != scanner.EOF {
					s.Scan()
					base += s.TokenText()
				}
				if len(base) < 2 {
					return fmt.Errorf("invalid action base")
				}
				team, err := strconv.Atoi(base[:len(base)-1])
				if err != nil {
					return err
				}
				action = &Action{
					TeamIndex: team,
					ActionKey: rune(base[len(base)-1]),
				}
			}
		}
		run = s.Peek()
		if run == scanner.EOF && action != nil {
			bgn.Actions = append(bgn.Actions, *action)
		}
	}
	return nil
}
