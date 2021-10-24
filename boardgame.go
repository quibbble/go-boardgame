package go_boardgame

// BoardGame is a representation of a board game allowing one to perform actions on the game as well as to retrieve game data
type BoardGame interface {
	// Do performs an action on the game
	Do(action BoardGameAction) error

	// GetSnapshot retrieves the current game state from 'team' view
	// entering nothing returns a top-down snapshot meaning everything in the game is seen
	// entering more than one team will error
	GetSnapshot(team ...string) (*BoardGameSnapshot, error)
}
