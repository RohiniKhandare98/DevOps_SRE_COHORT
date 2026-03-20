package main
import "fmt"

func main(){
	
	fmt.Println("Hello")
	
	// iteration - range => array/slice(in Go language)
	var servers []string
	servers = []string{"s1" , "s2" , "s3"}
	
	// classic for loop
	fmt.Println("using classic loop")
	
	for i:=0 ; i< len(servers) ; i++ {
		fmt.Println(i , " => " ,servers[i])
	}
	
	// using range 
	fmt.Println("\nusing RANGE for loop")
	for i , r:= range servers {
		fmt.Printf("%d => %s \n" ,i , r)
	}
	
	
	// range without i --> i will be replaced by underscore (_)
	fmt.Println("Skip index with _ ")
	
	for _, r:= range servers{
		fmt.Println(r)
	}
	
	
}
