package repo

type Requester interface {
	Ask(field string, input string) (string, error)
}
