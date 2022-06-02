package utils

// Error standard Error Template
type Error struct {
	Msg string
}

func (e Error) Error() string {
	return e.Msg
}
