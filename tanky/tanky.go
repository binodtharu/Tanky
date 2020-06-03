package tanky
import "fmt"
 

type Tanky struct{
	Capacity	int
	Level		int

}
func (t *Tanky) Fill(a int) {
	
	// var a int 
	// fmt.Print("fill water in tanky:")
	// fmt.Scan(&a)
	if a <= t.Capacity {
		if t.Level+a <= t.Capacity{
			t.Level = t.Level+a
			
		fmt.Println("fill water ", a,"liter")

		}else {
		fmt.Println("not enough space in tanky")
		}
	
	}else {
		fmt.Println("not enough space in tanky")
	}
}
func (t*Tanky) Drain(d int) {
	// var d int
	// fmt.Print("drain water in takny:")
	// fmt.Scan(&b)
	if t.Level - d < 0 {
		fmt.Println("not enough water in the tanky")
		
	}else {
		t.Level -= d
		fmt.Println("drain water ",d,"liter")
	}
}
func (t Tanky) Levels() int{
	// fmt.Println(t.Level)
	return t.Level
}