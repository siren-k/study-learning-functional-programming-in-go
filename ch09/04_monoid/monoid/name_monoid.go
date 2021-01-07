package monoid

type NameMonoid interface {
	Append(s string) NameMonoid
	Zero() string
}

type nameContainer struct {
	Name string
}

func (box nameContainer) Append(s string) NameMonoid {
	box.Name += s
	return box
}

func (box nameContainer) Zero() string {
	return ""
}

func (box nameContainer) String() string {
	return box.Name
}

func WrapName(s string) NameMonoid {
	return nameContainer{Name: s}
}
