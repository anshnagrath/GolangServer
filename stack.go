package main
import ( "fmt" )
func main(){
	s:=[]int{1,2}
	s = push(s,3,4)
	m := pop(s)
	d:=top(s)
	fmt.Println(s,m,d);
}
func push(s []int , ds ...int) []int{
return append(s,ds...)
}

func pop(s []int ) []int {
	return s[:len(s)-1]

}
func top(s []int) int{
	return s[len(s)-1]

}
