package go_boardgame

// BoardGameOptions are the options used to create a new game
type BoardGameOptions struct {
	// Teams is the list of teams that will be playing the game - required
	Teams []string

	// MoreOptions allows for additional game options to be passed that are unique to each game
	// Additional options are not required for every game so this field can be ignored if desired
	MoreOptions interface{} `json:",omitempty"`
}

// BoardGameInfo provides additional details about the game
type BoardGameInfo struct {
	// GameKey is the unique key that differentiates this game from others
	GameKey string

	// MinTeams and MaxTeams represents the min and max teams allowed when creating a new game
	MinTeams, MaxTeams int

	// MoreInfo allows for additional game specific info
	// Additional info is not required for every game so this field can be ignored if desired
	MoreInfo interface{} `json:",omitempty"`
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

// BoardGameSnapshot represents the current state of the game that will be viewed by a player
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

	// Targets are typically a list of BoardGameAction that can be performed on the game state
	// This can allow players to view all valid actions through a GUI
	// This field is left as an interface to allow for different targeting approaches as well
	// Optional feature not required to play a game and can be ignored if desired
	Targets interface{} `json:",omitempty"`

	// Actions is a list of past game actions that have lead to the current game state
	// This can allow players to view game logs of past actions
	// Optional feature not required to play a game and can be ignored if desired
	Actions []*BoardGameAction `json:",omitempty"`

	// Message provides players with information about what to do next
	// Optional feature not required to play a game and can be ignored if desired
	Message string `json:",omitempty"`
}
