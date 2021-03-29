package main

import (
	"image"
)

type Connotation int

const (
	VERY_GOOD Connotation = iota
	GOOD      Connotation = iota
	NEUTRAL   Connotation = iota
	BAD       Connotation = iota
	VERY_BAD  Connotation = iota
)

type Alignment int

const (
	LAWFUL_GOOD     Alignment = iota
	NEUTRAL_GOOD    Alignment = iota
	CHAOTIC_GOOD    Alignment = iota
	LAWFUL_NEUTRAL  Alignment = iota
	TRUE_NEUTRAL    Alignment = iota
	CHAOTIC_NEUTRAL Alignment = iota
	LAWFUL_EVIL     Alignment = iota
	NEUTRAL_EVIL    Alignment = iota
	CHAOTIC_EVIL    Alignment = iota
	UNALIGNED       Alignment = iota
)

type Entity struct {
	tag         string
	name        string
	description string
	pic         image.Image
}

type Player struct {
	name        string
	isDm        bool
	pic         image.Image
	description string
	alignment   Alignment
}

type PlayerAction struct {
	player      Player
	description string
}

type Note struct {
	text string
}

type PlayerEvent struct {
	description string
	eventType   string
}

type Session struct {
	name              string
	realWorldDate     int8
	modifiedDate      int8
	gameStartDateTime string
	gameEndDateTime   string
	notes             []Note
	events            []PlayerEvent
	playerActions     []PlayerAction
	entities          []Entity
	creator           Player
}
