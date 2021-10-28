package go_boardgame

// BoardGame is a representation of a board game allowing one to perform actions on the game as well as to retrieve game data
type BoardGame interface {
	// Do performs an action on the game
	Do(action BoardGameAction) error

	// GetSnapshot retrieves the current game state from 'team' view
	// Entering nothing returns a complete snapshot i.e. all hands, resources, etc.
	// Entering more than one team will error
	GetSnapshot(team ...string) (*BoardGameSnapshot, error)

	// GetSeed returns the seed used to generate randomness in the game
	// Useful for recreating games for replays or rollbacks using the seed and list of game actions
	// Some games do not require randomness in which case zero may be returned
	GetSeed() int64
}
