package main

import "fmt"

func main() {
	a := &artist{}
	fmt.Println(a)
	a.Unmarshal("queen", 2000, 1960)
	fmt.Println(a)
}

type artist struct {
	name       string
	Created    int
	firstalbum int
}

func (a *artist) Unmarshal(name string, c, f int) {
	a.name = name
	a.Created = c
	a.firstalbum = f
}
