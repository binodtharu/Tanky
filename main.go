
package main

import (
	"fmt"
	"github.com/binodtharu/01-06-2020/tanky"
)


func main() { 
	t := tanky.Tanky{Capacity:1000}

	var m string
	
	 for {
		fmt.Print("what do you want fill or drain:")
		fmt.Scan(&m)
		if m == "fill" {
			var a int
			fmt.Scan(&a)
			t.Fill(a)
		} else if m == "drain" {
			var d int
			fmt.Scan(&d)
			t.Drain(d)
		}else{
			fmt.Println("invalid")
		}
		l := t.Levels()
		fmt.Println("water level",l)
	}
}
