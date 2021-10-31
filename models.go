package go_boardgame

// BoardGameOptions are the options used to create a new game
type BoardGameOptions struct {
	// Teams is the list of teams that will be playing the game - required
	Teams []string

	// Seed is the number used to generate deterministic randomness - required
	// This allows for features such as replays or rollbacks of games
	Seed int64

	// MoreOptions allows for additional game options to be passed that are unique to each game
	// Additional options are not required for every game so this field can be ignored if desired
	MoreOptions interface{} `json:",omitempty"`
}

// BoardGameAction represents an action that is performed on the game state
type BoardGameAction struct {
	// Team is the team performing the action - required
	Team string

	// ActionType is the key that determines what action to perform - required
	ActionType string

	// MoreDetails allows for additional action details to be passed that are unique to the action type
	// Additional details are not required for every action so this field can be ignored if desired
	MoreDetails interface{} `json:",omitempty"`
}

// BoardGameSnapshot represents the current state of the game
type BoardGameSnapshot struct {
	// Turn is the turn of the current team - required
	Turn string

	// Teams is a list of all teams playing the game - required
	Teams []string

	// Winners is a list of teams that have won the game - required
	Winners []string

	// MoreData allows for additional game data to be returned that is unique to each game
	// Typically more data such as boards, decks, etc. are needed for a game but this field can be ignored if desired
	MoreData interface{} `json:",omitempty"`

	// Targets are a list of actions that can be performed on the game state
	// This can be a helpful feature when displaying valid actions to a player through a GUI
	// Optional feature not required to play a game and can be ignored if desired
	Targets []*BoardGameAction `json:",omitempty"`

	// Actions is a list of past game actions that have lead to the current game state
	// This can be a helpful feature for undoing a past action or providing a replay of a game
	// Optional feature not required to play a game and can be ignored if desired
	Actions []*BoardGameAction `json:",omitempty"`
}
