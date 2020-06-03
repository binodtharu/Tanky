package Tanky
import "fmt"
 

type TANKY struct{
	Capacity	int
	Level		int

}
func (t *TANKY) Fill() {
	
	var a int 
	fmt.Print("fill water in tanky:")
	fmt.Scan(&a)
	if a <= t.Capacity {
		if t.Level+a <= t.Capacity{
			t.Level = t.Level+a
			
		fmt.Println("level of water liter:",t.Level,"liter")

		}else {
	
			fmt.Println("not enough space in tanky")
		}
	
	}else {
		fmt.Println("not enough space in tanky")
	}
}
func (t *TANKY) Drain() {
	var b int
	fmt.Print("drain water in takny:")
	fmt.Scan(&b)
	if t.Level - b < 0 {
		fmt.Println("not enough water in the tanky")
		
	}else {
		t.Level -= b
		fmt.Println("level of water",t.Level,"liter")
	}
}