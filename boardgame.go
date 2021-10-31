package go_boardgame

// BoardGame is a representation of a board game allowing one to perform actions on the game as well as to retrieve game data
type BoardGame interface {
	// Do performs an action on the game
	Do(action BoardGameAction) error

	// GetSnapshot retrieves the current game state from 'team' view
	// Entering nothing returns a complete snapshot i.e. all hands, resources, etc.
	// Entering more than one team will error
	GetSnapshot(team ...string) (*BoardGameSnapshot, error)

	// GetNotation gets the game notation required to load this game at a later time
	// A notation is the game seed and a simplified list of actions done to reach the current game state
	// Notation - "'seed':'team index','action index','details one','details two',...;..."
	// EX: "123:0,0,0,F;1,1;" seed is 123, team 0 played action 0 with details 0 and F, team 1 played action 1
	// todo need a way to store BordGameOptions in this notation
	GetNotation() string

	// GetSeed returns the seed used to generate randomness in the game
	// Useful for recreating games for replays or rollbacks using the seed and list of game actions
	// Some games do not require randomness in which case zero may be returned
	GetSeed() int64
}
