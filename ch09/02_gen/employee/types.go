package employee

// +gen slice:"Aggregate[string]"
type Employee struct {
	Name       string
	Department string
}
