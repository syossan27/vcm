package main

type (
	Selection struct {
		Y int
	}
)

func NewSelection() Selection {
	return Selection{
		Y: 15,
	}
}

func (s *Selection) Exec() {
	switch s.Y {
	case 15:

	case 16:
	}
}
