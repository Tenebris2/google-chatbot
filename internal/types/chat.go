package types

type EventType int

const (
	Message EventType = iota
	CardClicked
	AddedToSpace
)

var EventTypeName = map[string]EventType{
	"MESSAGE":        Message,
	"CARD_CLICKED":   CardClicked,
	"ADDED_TO_SPACE": AddedToSpace,
}

type CommmandType int

const (
	Start CommmandType = iota
)

var CommmandTypeName = map[string]CommmandType{
	"START": Start,
}
