package u

import "fmt"

type print struct{
	label string 
	input string
}
func Print()print{
	return print{}
}
func ( p print) Label(s string)print{
	p.label = s
	return p
}
func (p print) Input(s string){
	p.input = s
	Style(0,10,1)
	fmt.Print(p.label, ":")
	Style(4,10,3)
	fmt.Print(p.input,"\n")
}

var colors map[string]int = map[string]int{
	"black": 0,
	"red" : 1,
	"green" : 2,
	"yellow" : 3,
	"blue" : 4,
	"purple" : 5,
	"cyan" : 6,
	"gray" : 7,
	"white" : 8,
}
var types map[string]int = map[string]int{
	"normal" : 0,
	"light" : 1,
	"subline" : 4,
	"cursive" : 3,
	"parpant" : 5,
	"inverse" : 7,
	"ocult" : 8,
}
func InyectStyle(t int, b int, c int)string{
	return 	fmt.Sprint("\033[", t, ";4", b, ";3", c, "m")
}
func InyectReset()string{
	return 	fmt.Sprint("\033[;0m")

}
func Style(t int, b int, c int) {
	fmt.Print("\033[", t, ";4", b, ";3", c, "m")
}
func StyleS( t string , b string , c string ){
	fmt.Println(types[t],colors[b],colors[c])
	if b == "transparent"{
		fmt.Print("\033[", types[t] ,";30",";3",colors[c],"m")
	}else{
		fmt.Print("\033[", types[t] ,";4",colors[b],";3",colors[c],"m")
	}
}
func Reset(){
	fmt.Print("\033[;0m")
}
