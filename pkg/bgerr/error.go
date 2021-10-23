package bgerr

import (
	"fmt"
)

// Error provides a simple method of returning common board game errors to reduce duplicate code
type Error struct {
	Err    error
	Status int
}

func (i *Error) Error() string {
	return fmt.Sprintf("%s: %s", StatusText(i.Status), i.Err.Error())
}
