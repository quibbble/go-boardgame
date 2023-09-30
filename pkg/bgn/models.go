package bgn

const (
	GameTag    string = "Game"
	TeamsTag   string = "Teams"
	VariantTag string = "Variant"
)

// RequiredTags are the list of tags any game must include
var RequiredTags = []string{
	GameTag,  // the name of the game being played
	TeamsTag, // list of teams playing the game
}
