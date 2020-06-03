package main

import (
	"fmt"
	"github.com/hive-repo/01-06-2020/Tanky"
)


func main() { 
t := Tanky.TANKY{500,0}
	
	var m string
	
	for {
		fmt.Print("what do you want fill or drain:")
		fmt.Scan(&m)
		if m == "Fill" {
			t.Fill()
		} else if m == "Drain" {
			t.Drain()
		}else{
			fmt.Println("invalid")
		}
	
	}
}
