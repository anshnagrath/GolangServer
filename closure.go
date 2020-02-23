package main
import "fmt"
func main(){
	ab:= addBy();
	mb:= multipleby();
	q:= ab(4)
	z:= ab(5)
	fmt.Println(q,z);
	m:= mb(6);
	l:= mb(9);
	fmt.Println(m,l)


}
func addBy() func (int) int{
	ab:= 1
	return func( s int) int{
		ab += s
		return ab
	}
}
func multipleby( ) func (int) int  {
	mb:= 1
	return func (b int) int {
		mb *= b 
		return mb

	}
}