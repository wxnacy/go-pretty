package main

import (
	"strconv"

	"github.com/wxnacy/go-pretty"
)

type Demo struct {
	ID   string
	Name string
	Age  int
}

func (d Demo) BuildPretty() []pretty.Field {
	fields := make([]pretty.Field, 0)
	fields = append(fields, pretty.Field{Name: "ID", Value: d.ID, IsFillLeft: true})
	fields = append(fields, pretty.Field{Name: "Name", Value: d.Name})
	fields = append(fields, pretty.Field{Name: "Age", Value: strconv.Itoa(d.Age)})
	return fields
}

func main() {
	items := make([]Demo, 0)
	items = append(items, Demo{ID: "1", Name: "wxnacy", Age: 18})
	items = append(items, Demo{ID: "12", Name: "李四", Age: 18})
	items = append(items, Demo{ID: "123", Name: "王五", Age: 18})

	// print list 1
	l := &pretty.List{}
	for _, item := range items {
		l.Add(item)
	}
	l.Print()

	// or print list 2
	// prettyList := make([]pretty.Pretty, 0)
	// for _, item := range items {
	// prettyList = append(prettyList, item)
	// }
	// pretty.List(prettyList).Print()
}
