// https://github.com/karlseguin/golang-set-fun/blob/master/ondemandarrayset.go

package set

type OnDemandArraySet struct {
	data []string
}

func (this *OnDemandArraySet) Add(values []string) {
	this.data = append(this.data, values...)
}

func (this *OnDemandArraySet) Contains(value string) (exists bool) {
	for _, v := range this.data {
		if v == value {
			return true
		}
	}
	return false
}

func (this *OnDemandArraySet) Length() int {
	return len(this.data)
}

func (this *OnDemandArraySet) RemoveDuplicates() {
	length := len(this.data) - 1
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			if this.data[i] == this.data[j] {
				this.data[j] = this.data[length]
				this.data = this.data[0:length]
				length--
				j--
			}
		}
	}
}

func (this *OnDemandArraySet) Copy() *[]string {
	var result = make([]string, this.Length())
	copy(result, this.data)
	return &result
}

func NewSet(capacity int) Set {
	return &OnDemandArraySet{make([]string, 0, capacity)}
}
