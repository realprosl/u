package u

import (
	"fmt"
	"reflect"
	"strings"
)
var tags []string = []string{"body","div","h1","h2","h3","h4","h5","h6","button","input"}
var signals []string = []string{"<",">","/"}
var asignatures[]string = []string{"=","'"}
var keywords []string = []string{"class","name","id","value"}

func PrintHtml(ele string) {

	for _,tag := range tags{
		ele = strings.ReplaceAll(ele,tag,InyectStyle(0,0,5)+tag+InyectReset())
	}
	for _,signal := range signals{
		ele = strings.ReplaceAll(ele,signal,InyectStyle(1,0,5)+signal+InyectReset())
	}
	for _,asignature := range asignatures{
		ele = strings.ReplaceAll(ele,asignature,InyectStyle(0,0,4)+asignature+InyectReset())
	}
	for _,keyword := range keywords {
		ele = strings.ReplaceAll(ele,keyword,InyectStyle(1,0,4)+keyword+InyectReset())
	}

	fmt.Print(ele,"\n")
}
func PrintStruct(s interface{}){
	rv := reflect.ValueOf(s)
	numField := rv.Elem().NumField()
	field := rv.Elem().Field
	param := reflect.TypeOf(s).Elem().Field


	for i:=0 ; i < numField ; i++{
		Style(0,0,2)
		nameField := param(i).Name
		fmt.Print(nameField)
		Style(1,0,2)
		fmt.Print(":")
		Reset()
		content := field(i)
		if nameField == "innerHtml" || nameField == "OuterHtml" {
			fmt.Print("\n")
			PrintHtml(fmt.Sprint(content))
		}else if fmt.Sprint(content) == "<nil>"{
			Style(0,0,1)
			fmt.Print(" ",content,"\n")
		}else if fmt.Sprint(content) == ""{
			Style(0,0,1)
			fmt.Print(" undefined","\n")
		}else if nameField == "ParentNode"{
			Reset()
			fmt.Print("\n")
			PrintHtml(fmt.Sprint(content))
		}else{
			fmt.Print("	",content,"\n")
		}
	}
}