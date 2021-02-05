package main

import (
	"fmt"
)

type Cat struct {
	Name  string
	Sound string
}

func (c Cat) method1() string {
	return "I am the methods inherited by struct and my values are %s +%s" + c.Name + c.Sound
}
func main() {
	CatObj := Cat{Name: "Dog", Sound: "dogg"}
	fmt.Println(CatObj)
	// var CatObj2 = new(Cat)
	// CatObj2.Name = "abcd"
	// CatObj2.Sound = "ABCD"
	// fmt.Println(CatObj2)
	fmt.Println(CatObj.method1())

}
