package tictactoe

import (
	"fmt"
	"strings"

	bg "github.com/quibbble/go-boardgame"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
)

const size = 3

var (
	indexToMark = map[int]string{0: "O", 1: "X"}
	markToIndex = map[string]int{"O": 0, "X": 1}
)

// state handles all the internal game logic for the game
type state struct {
	turn    string
	teams   []string
	winners []string
	board   [size][size]string
}

func newState(teams []string) *state {
	return &state{
		turn:    teams[0],
		teams:   teams,
		winners: make([]string, 0),
		board:   [size][size]string{},
	}
}

func (s *state) MarkLocation(team string, row, column int) error {
	index := indexOf(s.teams, team)
	if index < 0 {
		return &bgerr.Error{
			Err:    fmt.Errorf("%s not playing the game", team),
			Status: bgerr.StatusInvalidActionDetails,
		}
	}
	if team != s.turn {
		return &bgerr.Error{
			Err:    fmt.Errorf("%s cannot play on %s turn", team, s.turn),
			Status: bgerr.StatusInvalidAction,
		}
	}
	if row < 0 || row >= size || column < 0 || column >= size {
		return &bgerr.Error{
			Err:    fmt.Errorf("row or column out of bounds"),
			Status: bgerr.StatusInvalidActionDetails,
		}
	}
	if s.board[row][column] != "" {
		return &bgerr.Error{
			Err:    fmt.Errorf("%d,%d already marked", row, column),
			Status: bgerr.StatusInvalidAction,
		}
	}

	// mark index
	s.board[row][column] = indexToMark[index]

	// check and update winner
	if winner(s.board) != "" {
		s.winners = []string{s.teams[markToIndex[winner(s.board)]]}
	} else if draw(s.board) {
		s.winners = s.teams
	}

	// update turn
	s.turn = s.teams[(index+1)%2]
	return nil
}

func (s *state) targets() []*bg.BoardGameAction {
	targets := make([]*bg.BoardGameAction, 0)
	for r, row := range s.board {
		for c, loc := range row {
			if loc == "" {
				targets = append(targets, &bg.BoardGameAction{
					Team:       s.turn,
					ActionType: ActionMarkLocation,
					MoreDetails: MarkLocationActionDetails{
						Row:    r,
						Column: c,
					},
				})
			}
		}
	}
	return targets
}

func (s *state) message() string {
	message := fmt.Sprintf("%s must mark a location", s.turn)
	if len(s.winners) > 0 {
		message = fmt.Sprintf("%s tie", strings.Join(s.winners, " and "))
		if len(s.winners) == 1 {
			message = fmt.Sprintf("%s wins", s.winners[0])
		}
	}
	return message
}

func winner(board [size][size]string) string {
	for i := 0; i < size; i++ {
		// check rows
		if board[i][0] != "" && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return board[i][0]
		}
		// check columns
		if board[0][i] != "" && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			return board[0][i]
		}
	}
	// check diagonal
	if board[0][0] != "" && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		return board[0][0]
	}
	// check diagonal
	if board[2][0] != "" && board[2][0] == board[1][1] && board[2][0] == board[0][2] {
		return board[2][0]
	}
	return ""
}

func draw(board [size][size]string) bool {
	for _, row := range board {
		for _, loc := range row {
			if loc == "" {
				return false
			}
		}
	}
	return true
}

func indexOf(items []string, item string) int {
	for i, it := range items {
		if it == item {
			return i
		}
	}
	return -1
}
