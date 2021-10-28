package go_boardgame

// BoardGameBuilder builds a game given a set of options
type BoardGameBuilder interface {
	// Create creates a game with desired options and an optional random seed
	// Leaving seed unset will typically result in the seed being set to time.Now().UnixNano()
	// Seed is used for deterministic randomness allowing for features such as replays or rollbacks of games
	Create(options BoardGameOptions, seed ...int64) (BoardGame, error)

	// Key gets the game's unique key, typically the name of the game
	Key() string
}
