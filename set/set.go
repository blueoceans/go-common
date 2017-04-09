package set

type Set interface {
	Add(values []string)
	Contains(value string) bool
	Length() int
	RemoveDuplicates()
	Copy() *[]string
}
