package bgn

import "fmt"

/*
BGN or Board Game Notation is a spin on PGN or Portable Game Notation that can be applied to virtually any board game

A BGN is made of a set of tags and actions

Tags describe information used to create the game as well as other helpful data
Required tags include Game, Teams, and Seed
Example Tags: [Game "Carcassonne"][Teams "A, B"][Seed "123"][Date "10-31-2021"][Halloween "Spooky"]

Actions are the list of steps teams have taken to reach the current game state
An action requires the team index - index of team in "Teams" tag - and the action key
Additional details such i.e. x,y locations, card value, etc. may be included in the details section of the action
Example Action: 0c {team A does action c} 0a-1.0 {team A does action a with details 1 and 0}

Putting it all together:
[Game "Carcassonne"][Teams "A, B"][Seed "123"][Date "10-31-2021"][Halloween "Spooky"]0c {comment like this} 0a-1.0
*/
type BGN struct {
	Tags    map[string]string
	Actions []Action
}

func (b *BGN) String() string {
	bgn := ""
	for key, value := range b.Tags {
		bgn += fmt.Sprintf("[%s \"%s\"]\n", key, value)
	}
	bgn += "\n"
	line := ""
	for _, action := range b.Actions {
		line += fmt.Sprintf("%s ", action.String())
		if len(line) > 70 {
			bgn += fmt.Sprintf("%s\n", line[:len(line)-1])
			line = ""
		}
	}
	if line != "" {
		bgn += line[:len(line)-1]
	}
	return bgn
}

type Action struct {
	TeamIndex int      // the index of team in Teams Tag
	ActionKey rune     // single character key indicating the action to perform
	Details   []string // additional details that can be optionally used when describing an action
}

func (a *Action) String() string {
	bgn := fmt.Sprintf("%d%s", a.TeamIndex, string(a.ActionKey))
	details := ""
	for _, detail := range a.Details {
		details += fmt.Sprintf("%s.", detail)
	}
	if len(details) > 0 {
		details = details[:len(details)-1]
		bgn = fmt.Sprintf("%s-%s", bgn, details)
	}
	return bgn
}
