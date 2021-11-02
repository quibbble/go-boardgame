package bgn

import "fmt"

// Game is a representation of a game in Board Game Notation (BGN)
type Game struct {
	Tags    map[string]string
	Actions []Action
}

func (g *Game) String() string {
	bgn := ""
	for key, value := range g.Tags {
		bgn += fmt.Sprintf("[%s \"%s\"]\n", key, value)
	}
	bgn += "\n"
	line := ""
	for _, action := range g.Actions {
		line += fmt.Sprintf("%s ", action.String())
		if len(line) > 70 {
			bgn += fmt.Sprintf("%s\n", line[:len(line)-1])
			line = ""
		}
	}
	if line != "" {
		bgn += line[:len(line)-1]
	}
	return bgn
}

type Action struct {
	TeamIndex int      // the index of team in Teams Tag
	ActionKey rune     // single character key indicating the action to perform
	Details   []string // additional details that can be optionally used when describing an action
}

func (a *Action) String() string {
	bgn := fmt.Sprintf("%d%s", a.TeamIndex, string(a.ActionKey))
	details := ""
	for _, detail := range a.Details {
		details += fmt.Sprintf("%s.", detail)
	}
	if len(details) > 0 {
		details = details[:len(details)-1]
		bgn = fmt.Sprintf("%s&%s", bgn, details)
	}
	return bgn
}
