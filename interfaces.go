package main

import "fmt"

//if we don't know what is the value that is going to passed with in the function so this is used
func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("main function")
	Anything(2.44)
	Anything("Ansh")
	Anything(struct{}{})
	mymap := make(map[string]interface{})
	mymap["name"] = "Ansh"
	mymap["age"] = 10
	fmt.Println(mymap)
	// conditional Statements
	f := true

	flag := &f

	if flag == nil {
		fmt.Println("value is nill")
	} else if *flag {
		fmt.Println("flag is present")
	} else {
		fmt.Println("false")
	}
	arr := []string{"ansh", "nagrath", "check"}

	for i, value := range arr {
		fmt.Println(i)
		fmt.Println(value, "value of an array ======")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	mymap = make(map[string]interface{})
	mymap["Name"] = "Ansh"
	mymap["LastName"] = "Nagrath"

	for k, v := range mymap {
		fmt.Printf("key: %s and value %v", k, v)
	}
	for i := 0; i < 10; i++ {
		if i%2 == 0 {

			continue
		} else {
			fmt.Println("Else Printed", i)
		}
	}
	day := "Sunday"
	switch day {
	case "Sunday":
		fmt.Println("kl office jana h")
	default:
		fmt.Println("Learn go lang")

	}

}
