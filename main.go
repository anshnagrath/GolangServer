package main

import (
	"fmt"
	"strings"
)

// simple struct
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

// methods attached  to  struct Car
func (c Car) Print() {
	fmt.Println("Function encapsulated inside an object")
}

// methods attached to struct Cat
func (c Car) Drive() string {
	return c.Name

}

// any struct which has Drive as a method will implement Car interface
type Carz interface {
	Drivez()
	Stopz()
}

type Lambo struct {
	lambModel string
}
type Chevy struct {
	chevModel string
}

// Enforcing the new model
func NewModel(arg string) Carz {
	return &Lambo{arg}
}

func (L *Lambo) Drivez() {
	fmt.Println("Lambod is on the move", L.lambModel)
}

func (L *Lambo) Stopz() {
	fmt.Println("Lambod stopped")

}

func (C *Chevy) Drivez() {

	fmt.Println("Chevy is on the move", C.chevModel)

}
func (C *Chevy) Stopz() {
	fmt.Println("Chevy Stopped")
}

func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

func pointers() {
	m1 := 2
	m2 := 34
	ptr := &m1

	fmt.Println(ptr, *ptr, "this is an pointer")
	fmt.Println(m1, m2)
	//swapped varibles using pointers
	swap(&m1, &m2)
	fmt.Println(m1, m2)

}

func abstractDatatypes() {
	//var arr []int
	arr := []int{4545, 34, 12312, 123}
	arr2 := []string{"dcasdcasdc", "cascas"}
	arr3 := append(arr2, "appended")
	c := Car{"Chevy", 1, 2}
	c.Print()
	fmt.Println(c.Drive(), "Drive")
	fmt.Println(c)
	fmt.Println(arr, arr2, arr3)
	pointers()

}

func main() {
	//atomic data types in go
	var m1 = 2
	//declaring variable in multiple linesss
	var (
		m2 = 3
		m3 = 3
	)
	fmt.Println("Hellooo go", m1, m2, m3)
	// type casting the one varible type to another so as to add
	var e1 int32
	var e2 int64
	// default value of an nuber is 0 so this will result in zero
	fmt.Println(int64(e1) + e2)

	//intializing variable in one line
	u1 := 45
	u5 := 78
	fmt.Println(u1+u5, "declared variable without specifing datatype")

	//strings in golang
	mz := "ansh"
	var mn string
	mn = " nagrath"
	fmt.Println(mz+mn, strings.Contains(mz+mn, "na"))
	fmt.Println(strings.ReplaceAll(mz+mn, "na", "Na"))
	fmt.Println(strings.Split(mn, " "))
	abstractDatatypes()

}
