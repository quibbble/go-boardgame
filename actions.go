package go_boardgame

import (
	"fmt"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
	"strconv"
)

// Common actions every game may utilize if desired
const (
	ActionSetWinners = "SetWinners" // overrides the winner(s) of a game
)

// SetWinnersActionDetails overrides the current winner(s) of the game with Winners
type SetWinnersActionDetails struct {
	Winners []string
}

// EncodeBGN converts SetWinnersActionDetails into bgn.Action Details object given the list of teams in game
func (o *SetWinnersActionDetails) EncodeBGN(teams []string) ([]string, error) {
	indices := make([]string, 0)
	for _, winner := range teams {
		index := -1
		for i, team := range teams {
			if winner == team {
				index = i
			}
		}
		if index == -1 {
			return []string{}, &bgerr.Error{
				Err:    fmt.Errorf("winner not found in teams when encoding action %s", ActionSetWinners),
				Status: bgerr.StatusBGNEncodingFailure,
			}
		}
		indices = append(indices, strconv.Itoa(index))
	}
	return indices, nil
}

// DecodeSetWinnersActionDetailsBGN converts a bgn.Action Details object into the SetWinnersActionDetails object
func DecodeSetWinnersActionDetailsBGN(details, teams []string) (*SetWinnersActionDetails, error) {
	winners := make([]string, 0)
	for _, i := range details {
		index, err := strconv.Atoi(i)
		if err != nil {
			return nil, &bgerr.Error{
				Err:    fmt.Errorf("cannot convert string to int when decoding action %s", ActionSetWinners),
				Status: bgerr.StatusBGNDecodingFailure,
			}
		}
		if index >= len(teams) {
			return nil, &bgerr.Error{
				Err:    fmt.Errorf("index out of bounds when decoding action %s", ActionSetWinners),
				Status: bgerr.StatusBGNDecodingFailure,
			}
		}
		winners = append(winners, teams[index])
	}
	return &SetWinnersActionDetails{
		Winners: winners,
	}, nil
}
