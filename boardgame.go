package go_boardgame

import "github.com/quibbble/go-boardgame/pkg/bgn"

// BoardGame is a representation of a board game allowing one to perform actions on the game as well as to retrieve game data
type BoardGame interface {
	// Do performs an action on the game
	Do(action *BoardGameAction) error

	// GetSnapshot retrieves the current game state from 'team' view
	// Entering nothing returns a complete snapshot with no data hidden i.e. all hands, resources, etc.
	// Entering more than one team will error
	GetSnapshot(team ...string) (*BoardGameSnapshot, error)
}

// BoardGameWithBGN provides extra bgn functionality that does not necessarily need to be implemented to play a game
type BoardGameWithBGN interface {
	BoardGame

	// GetBGN returns the board game notation of the game
	GetBGN() *bgn.BGN
}
