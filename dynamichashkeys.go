package main

import ( "fmt"
		"math" )

func main(){
values:=[]string{"test","for","unique","hash","keys","for","this","element","again","test"}

for val := range values{
	 hash :=  generateHash(values[val])
	fmt.Println(hash)
}

}

func generateHash(s string) int {

result:= 0
idx := 0
for i := len(s)-1 ; i  > -1 ; i-- {
	result += int(math.Pow10(i)) * int(s[idx]);
	idx++
} 

return result;
}


