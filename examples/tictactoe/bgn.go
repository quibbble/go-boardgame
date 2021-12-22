package tictactoe

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
	bg "github.com/quibbble/go-boardgame"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
	"github.com/quibbble/go-boardgame/pkg/bgn"
)

// converts action types to correct bgn notation
var (
	actionToNotation = map[string]string{ActionMarkLocation: "m"}
	notationToAction = map[string]string{"m": ActionMarkLocation}
)

// encodeBGN writes MarkLocationActionDetails into bgn.Action Details format
func (p *MarkLocationActionDetails) encodeBGN() []string {
	return []string{strconv.Itoa(p.Row), strconv.Itoa(p.Column)}
}

// decodeMarkLocationActionDetailsBGN writes bgn.Action Details into MarkLocationActionDetails
func decodeMarkLocationActionDetailsBGN(details []string) (*MarkLocationActionDetails, error) {
	if len(details) != 2 {
		return nil, &bgerr.Error{
			Err:    fmt.Errorf("invalid details for action %s", ActionMarkLocation),
			Status: bgerr.StatusBGNDecodingFailure,
		}
	}
	row, err := strconv.Atoi(details[0])
	if err != nil {
		return nil, &bgerr.Error{
			Err:    fmt.Errorf("invalid details for action %s", ActionMarkLocation),
			Status: bgerr.StatusBGNDecodingFailure,
		}
	}
	column, err := strconv.Atoi(details[1])
	if err != nil {
		return nil, &bgerr.Error{
			Err:    fmt.Errorf("invalid details for action %s", ActionMarkLocation),
			Status: bgerr.StatusBGNDecodingFailure,
		}
	}
	return &MarkLocationActionDetails{
		Row:    row,
		Column: column,
	}, nil
}

// GetBGN converts a TicTacToe game into a bgn.Game
func (t *TicTacToe) GetBGN() *bgn.Game {
	tags := map[string]string{
		"Game":  key,
		"Teams": strings.Join(t.state.teams, ", "),
	}
	actions := make([]bgn.Action, 0)
	for _, action := range t.actions {
		bgnAction := bgn.Action{
			TeamIndex: indexOf(t.state.teams, action.Team),
			ActionKey: rune(actionToNotation[action.ActionType][0]),
		}
		switch action.ActionType {
		case ActionMarkLocation:
			var details MarkLocationActionDetails
			_ = mapstructure.Decode(action.MoreDetails, &details)
			bgnAction.Details = details.encodeBGN()
		}
		actions = append(actions, bgnAction)
	}
	return &bgn.Game{
		Tags:    tags,
		Actions: actions,
	}
}

// CreateWithBGN creates a game with BGN functionality
func (b *Builder) CreateWithBGN(options *bg.BoardGameOptions) (bg.BoardGameWithBGN, error) {
	return NewTicTacToe(options)
}

// Load loads a game given a bgn.Game
func (b *Builder) Load(game *bgn.Game) (bg.BoardGameWithBGN, error) {
	if game.Tags["Game"] != key {
		return nil, &bgerr.Error{
			Err:    fmt.Errorf("game tag does not match game key"),
			Status: bgerr.StatusBGNDecodingFailure,
		}
	}
	teamsStr, ok := game.Tags["Teams"]
	if !ok {
		return nil, &bgerr.Error{
			Err:    fmt.Errorf("team tag missing"),
			Status: bgerr.StatusBGNDecodingFailure,
		}
	}
	teams := strings.Split(teamsStr, ", ")
	g, err := b.CreateWithBGN(&bg.BoardGameOptions{
		Teams: teams,
	})
	if err != nil {
		return nil, err
	}
	for _, action := range game.Actions {
		if action.TeamIndex >= len(teams) {
			return nil, &bgerr.Error{
				Err:    fmt.Errorf("team index %d out of range", action.TeamIndex),
				Status: bgerr.StatusBGNDecodingFailure,
			}
		}
		team := teams[action.TeamIndex]
		actionType := notationToAction[string(action.ActionKey)]
		if actionType == "" {
			return nil, &bgerr.Error{
				Err:    fmt.Errorf("invalid action key %s", string(action.ActionKey)),
				Status: bgerr.StatusBGNDecodingFailure,
			}
		}
		var details interface{}
		switch actionType {
		case ActionMarkLocation:
			result, err := decodeMarkLocationActionDetailsBGN(action.Details)
			if err != nil {
				return nil, err
			}
			details = result
		}
		if err := g.Do(&bg.BoardGameAction{
			Team:        team,
			ActionType:  actionType,
			MoreDetails: details,
		}); err != nil {
			return nil, err
		}
	}
	return g, nil
}
