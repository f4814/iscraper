package models

import "github.com/ahmdrz/goinsta/v2"

type Follows struct {
	From string `json:"_from"`
	To   string `json:"_to"`
}

type Posts struct {
	From string `json:"_from"`
	To   string `json:"_to"`
}

type Likes struct {
	From       string `json:"_from"`
	To         string `json:"_to"`
	IsTopliker bool   `json:"is_topliker"`
}

type Tags struct {
	From string `json:"_from"`
	To   string `json:"_to"`
	// TODO FB Tags
	// TODO Coordinates
}

type Mentions struct {
	From string `json:"_from"`
	To   string `json:"_to"`

	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Z        int64   `json:"z"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Rotation float64 `json:"rotation"`
	IsPinned int     `json:"is_pinned"`
}

type Comments struct {
	From string `json:"_from"`
	To   string `json:"_to"`
}

type Child struct {
	From string `json:"_from"`
	To   string `json:"_to"`
}

func NewMention(i goinsta.Mentions) Mentions {
	return Mentions{
		X:        i.X,
		Y:        i.Y,
		Z:        i.Z,
		Width:    i.Width,
		Height:   i.Height,
		Rotation: i.Rotation,
		IsPinned: i.IsPinned,
	}
}
