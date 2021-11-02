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

	// Load loads a game in board game notation into a normal game
	Load(bgn *bgn.Game) (BoardGameWithBGN, error)
}
