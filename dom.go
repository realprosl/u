package u

import (
	"strings"
)

// types and variables
type Element struct {
	TagName    string
	innerHtml  string
	OuterHtml  string
	Value      string
	ClassName string
	Id      string
	Name string
	ParentNode *Element
	Children   []*Element
}
var document *Element = &Element{
		TagName: "document",
}
var Dom []*Element
// instanciate
func NewElement(t string) (e *Element) {
	if t == "body" {
		var body Element = Element{
			TagName:   "body",
			OuterHtml: "<body></body>",
			ParentNode: document,
		}
		Dom = append(Dom,&body)
		e = &body
	}
	if t == "div" {
		var div Element = Element{
			TagName:   "div",
			OuterHtml: "<div></div>",
		}
		Dom = append(Dom,&div)
		e = &div
	}
	if t == "button" {
		var button Element = Element{
			TagName:   "button",
			OuterHtml: "<button></button>",
		}
		Dom = append(Dom,&button)
		e = &button
	}
	if t == "h1" {
		var h1 Element = Element{
			TagName:   "h1",
			OuterHtml: "<h1></h1>",
		}
		Dom = append(Dom,&h1)
		e = &h1
	}
	if t == "h2" {
		var h2 Element = Element{
			TagName:   "h2",
			OuterHtml: "<h2></h2>",
		}
		Dom = append(Dom,&h2)
		e = &h2
	}
	if t == "h3" {
		var h3 Element = Element{
			TagName:   "h3",
			OuterHtml: "<h3></h3>",
		}
		Dom = append(Dom,&h3)
		e = &h3
	}
	if t == "h4" {
		var h4 Element = Element{
			TagName:   "h4",
			OuterHtml: "<h4></h4>",
		}
		Dom = append(Dom,&h4)
		e = &h4
	}
	return
}
func NewElementL(html string)(ele *Element){
	var id string
	var className string
	var name string
	var tagName string
	var value string
	var innerHtml string

	i := strings.Index(html, "<")
	i2 := strings.Index(html, ">")
	iTag := strings.Index(html, " ")
	if i > iTag || iTag > i2  {
		iTag = i2
	}
	open := html[i:i2+1]
	tagName = html[i+1:iTag]
	attrs := strings.Split(html[iTag:i2]," ")
	// obtener innerHtml descartando tags del mismo tipo dentro del elemento
	innerHtml = html[i2+1:]
	tagEnd := "</"+ tagName + ">"
	tagInit := "<"+ tagName 
	end := strings.Index(innerHtml,tagEnd)
	init := strings.Index(innerHtml,tagInit)
	iClose := 0
	iTemp := 0

	for strings.Contains(innerHtml,tagEnd) {
			iTemp = strings.Index(innerHtml,tagEnd)+len(tagEnd)
			iClose += iTemp
			innerHtml = innerHtml[iTemp:]
			if init > end{
				break
			}
	}
	innerHtml = html[i2+1:iClose+i2+1-len(tagEnd)]

	//fmt.Println("tagname:",tagName)
	//fmt.Println("attrs:",attrs)
	//fmt.Println("inner:",innerHtml)
	// fin de obtencion
	for _,v := range attrs {
		if strings.Contains(v , "id"){
			id = strings.ReplaceAll(v," ","")[4:len(v)-1]
		}
		if strings.Contains(v , "class"){
			className = strings.ReplaceAll(v," ","")[7:len(v)-1]
		}
		if strings.Contains(v , "name"){
			name = strings.ReplaceAll(v," ","")[6:len(v)-1]
		}
		if strings.Contains(v,"value"){
			value = strings.ReplaceAll(v," ","")[7:len(v)-1]
		}
	}

	ele = NewElement(tagName)
	ele.Id = id
	ele.ClassName = className
	ele.Name = name
	ele.Value = value
	ele.innerHtml = innerHtml
	ele.OuterHtml = open + innerHtml + tagEnd

	for strings.Contains(innerHtml,"<") && strings.Contains(innerHtml,"</"){
		child := NewElementL(innerHtml)
		existe := false
		for _,v := range Dom{
			if v == child{
				existe = true
			}
		}
		if !existe{
			Dom = append(Dom, child)
		}
		ele.Children = append(ele.Children ,child)
		child.ParentNode = ele
		innerHtml = strings.Replace(innerHtml,child.OuterHtml,"",1)
	}
	return 
}
// methods
func (e *Element) SetInnerHTML(html string) {
	if e.TagName != "input"{
		clouse := "</" + e.TagName + ">"
		e.innerHtml = html
		e.OuterHtml = strings.Replace(e.OuterHtml , clouse , html + clouse, 1)
	}
}
func (e *Element) GetInnerHTML() string {
	if e.TagName != "input"{
		return e.innerHtml
	}
	return ""
}
func (e *Element) Append(ele *Element) {
	close := "</" + e.TagName + ">"
	open := e.OuterHtml[:strings.Index(e.OuterHtml,">")+1]
	e.Children = append(e.Children, ele)
	e.innerHtml += ele.OuterHtml
	e.OuterHtml = open + e.innerHtml + close
	ele.ParentNode = e
}
func (e *Element) SetAttribute( t string, v string ){

	if t == "Name"{
		e.SetName(v)
	}
	if t == "ClassName"{
		e.SetClassName(v)
	}
	if t == "Id"{
		e.SetId(v)
	}
	if t == "value"{
		e.Value = v
	}
}
func (e *Element) SetId(v string) {
	e.Id = v 
	e.OuterHtml = strings.Replace(e.OuterHtml,e.TagName , e.TagName + " id='" + v + "'",1)
}
func (e *Element) SetClassName( v string ) {
	e.ClassName = v
	e.OuterHtml = strings.Replace(e.OuterHtml,e.TagName , e.TagName + " class='" + v + "'" ,1)
}
func (e *Element) SetName( v string ) {
	e.Name = v
	e.OuterHtml = strings.Replace(e.OuterHtml,e.TagName , e.TagName + " name='" + v + "'" ,1)
}
func (e *Element) SetValue(v string){
	e.Value = v
}

