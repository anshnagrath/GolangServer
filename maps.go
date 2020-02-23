package main
import  ("fmt"
 		 "sort") 
func main ()  {
	days := make(map[string]int)
	days["Sun"] = 3
	days["Sun"] -= 3
	fmt.Println(days["Noday"])

	//using maps without make

	var ps  map[string]int
	if ps == nil{
		fmt.Println("ps is nil",ps)
		ps = make(map[string]int)
	}
	ps["scsd"] = 1
	delete(ps,"scsd");
	fmt.Println("ps is nil",len(ps))
	
	sal:= map[string]float64{
		"Znsh":1.34,
		"Lakshman":4232342.234,
		"Yuv":3453453453.00909,
	}
	if val,ok:= sal["Ansh"];ok{
		fmt.Printf("ok %v \n",val)
	}else{
		fmt.Println("Does not Exist")
	}

	for key,val := range sal{
		fmt.Println(key,val);
	}

	// sorting this unordered map
	sl := make([]string,len(sal))
	for n := range sal{
		 sl = append(sl,n)

	}
	sort.Strings(sl)
	fmt.Println(sl);
	nm:= make(map[string]float64)
	for _,val:=range sl{
		nm[val] = sal[val]
	}
	fmt.Println(nm);




	
	





}