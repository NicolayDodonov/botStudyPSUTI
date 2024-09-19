package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Event struct {
	Type Type
	Text string
	Meta interface{}
}

type Type int

const (
	Unknow Type = iota
	Message
)
