package go_boardgame

// BoardGame is a representation of a board game allowing one to perform actions on the game as well as to retrieve game data
type BoardGame interface {
	// Do performs an action on the game
	Do(action BoardGameAction) error

	// GetSnapshot retrieves the current game state from 'team' view
	GetSnapshot(team string) (BoardGameSnapshot, error)
}
