package internal

type Iterator interface {
	Next() (value string, ok bool)
}

const INVALID_INT_VAL = -1
const INVALID_STRING_VAL = ""

type collection struct {
	index int
	List  []string
}

func (collection *collection) Next() (value string, ok bool) {
	collection.index++
	if collection.index >= len(collection.List) {
		return INVALID_STRING_VAL, false
	}
	return collection.List[collection.index], true
}

func NewSlice(s []string) *collection {
	return &collection{INVALID_INT_VAL, s}
}
