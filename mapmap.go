package main
import "fmt"

var s  = make(map[string]map[string]bool)
func main(){
 addPlayer("John","Cena");
 addPlayer("Ansh","Nagrath");
fmt.Println(s ,hasPlayer("Ansh","Nagrath"))


}
func addPlayer( f ,l string){
	if s[f] == nil{
		s[f] = make(map[string]bool)
	}
	s[f][l] = true
	

}
func hasPlayer(f,l string) bool{
	return s[f][l]
}