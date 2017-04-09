// https://github.com/karlseguin/golang-set-fun/blob/master/set.go

package set

type Set interface {
	Add(values []string)
	Contains(value string) bool
	Length() int
	RemoveDuplicates()
	Copy() *[]string
}
