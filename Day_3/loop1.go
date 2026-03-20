package main
import "fmt"

func main(){
	
	fmt.Println("Hello")
	
	// classic loop 
	
	for i:=0 ; i<=5 ; i++ {
		fmt.Println(i)
	}
	
	// iteration - range => array/slice(in Go language)
	var servers []string
	servers = []string{"s1" , "s2" }
	
	for i:=0 ; i< len(servers) ; i++ {
		fmt.Println(i , " => " ,servers[i])
	}
}
