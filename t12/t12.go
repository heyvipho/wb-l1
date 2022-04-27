package main

import "fmt"

type Set struct {
	m map[string][]string
}

func CreateSet() Set {
	return Set{
		m: make(map[string][]string),
	}
}

func (v *Set) Add(s ...string) {
	for _, val := range s {
		firstLetter := string(val[0])

		if _, ok := v.m[firstLetter]; !ok {
			v.m[firstLetter] = make([]string, 0)
		}

		v.m[firstLetter] = append(v.m[firstLetter], val)
	}
}

func (v *Set) Get() map[string][]string {
	return v.m
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	set := CreateSet()
	set.Add(arr...)

	fmt.Println(set.Get())
}
