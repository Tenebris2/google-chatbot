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

type Command int

const (
	Start  Command = iota
	Create Command = iota
)

var CommandType = map[string]Command{
	"START":  Start,
	"CREATE": Create,
}
