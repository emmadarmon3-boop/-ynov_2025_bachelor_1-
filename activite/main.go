package main 

import (
	"fmt"
    "time"
)
func main(){
	Start:=time.Now()
list_prenom:=[1000]string{"Lea","Louis"}
for i:=0; i<len(list_prenom); i++{
	fmt.Println(list_prenom[i])
}
fmt.Println(time.Since(Start))
}