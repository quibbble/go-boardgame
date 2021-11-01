package go_boardgame

// BoardGameBuilder builds a game given a set of options
type BoardGameBuilder interface {
	// Create creates a game with desired options
	Create(options *BoardGameOptions) (BoardGame, error)

	// Key gets the game's unique key, typically the name of the game
	Key() string
}

// BoardGameWithNotationBuilder builds a game with additional notation functionality
type BoardGameWithNotationBuilder interface {
	BoardGameBuilder

	// CreateWithNotation creates a game with desired options
	CreateWithNotation(options *BoardGameOptions) (BoardGameWithNotation, error)

	// Load takes a list of teams and game notation and rebuilds the game up to the point defined in the notation
	// Length of teams must match the number of teams defined in notation
	Load(teams []string, notation string) (BoardGameWithNotation, error)
}
