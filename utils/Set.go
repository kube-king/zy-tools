package utils

type Set map[interface{}]struct{}

func (s *Set) Add(k interface{}) {
	(*s)[k] = struct{}{}
}

func (s *Set) Remove(k interface{}) {
	delete((*s), k)
}

func (s *Set) Has(k interface{}) bool {
	_, ok := (*s)[k]
	return ok
}

func (s *Set) IsEmpty() bool {
	return 0 == len(*s)
}

func (s *Set) Len() int {
	return len(*s)
}

func (s *Set) All() []interface{} {
	var list []interface{}
	for i, _ := range *s {
		list = append(list, i)
	}
	return list
}
