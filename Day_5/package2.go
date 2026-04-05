package main

import (
  "fmt"
  "os"
  str  "strings"
 "strconv"
)


func main(){

	fmt.Println("welcome to package")
	port := 8080
        
        msg := fmt.Sprintf("Port: %d", port)
        fmt.Println(msg)
     
      
       //  get environmetal variables  -- get integer port as String

       portStr := os.Getenv("MY_PORT")
       fmt.Println("Port:", portStr)

          // conversion of string port into  integer port
	portInt, err := strconv.Atoi(portStr)
        if err != nil {
          fmt.Println("ERROR:", err)
	  os.Exit(1)
	}

	fmt.Println("Port Integer:", portInt)
	

      // using host variable
      host := os.Getenv("MY_HOST")
      fmt.Println("Host name :", host)

      
     // accept our service name
     serviceName := os.Getenv("MYAPP")
     fmt.Println("Service Name:", serviceName)

    
    //  service name into uppercase
    fmt.Println(str.ToUpper(serviceName))   

}
