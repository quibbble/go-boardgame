package go_boardgame

// BoardGameBuilder builds a game given a set of options
type BoardGameBuilder interface {
	// Create creates a game with desired options
	Create(options BoardGameOptions) (BoardGame, error)

	// Key gets the game's unique key, typically the name of the game
	Key() string
}
