package go_boardgame

import "github.com/quibbble/go-boardgame/pkg/bgn"

// BoardGameBuilder builds a game given a set of options
type BoardGameBuilder interface {
	// Create creates a game with desired options
	Create(options *BoardGameOptions) (BoardGame, error)

	// Key gets the game's unique key, typically the name of the game
	Key() string
}

// BoardGameWithBGNBuilder builds a game with additional bgn functionality
type BoardGameWithBGNBuilder interface {
	BoardGameBuilder

	// CreateWithBGN creates a game with desired options
	CreateWithBGN(options *BoardGameOptions) (BoardGameWithBGN, error)

	// Load takes bgn and rebuilds the game up to the point defined in the notation
	Load(bgn *bgn.BGN) (BoardGameWithBGN, error)
}
