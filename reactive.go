package u

// string
type RString struct {
	value    string
	OnChange func()
}

func NewString(initial string) *RString {
	r := RString{value: initial}
	return &r
}
func (this *RString) SetValue(v string) {
	this.value = v
	this.OnChange()
}
func (this *RString) Value() string {
	return this.value
}

// int
type RInt struct {
	value    int
	OnChange func()
}

func NewInt(initial int) *RInt {
	r := RInt{value: initial}
	return &r
}
func (this *RInt) SetValue(v int) {
	this.value = v
	this.OnChange()
}
func (this *RInt) Value() int {
	return this.value
}
// bool
type RBool struct {
	value    bool
	OnChange func()
}

func NewBool(initial bool) *RBool {
	r := RBool{value: initial}
	return &r
}
func (this *RBool) SetValue(v bool) {
	this.value = v
	this.OnChange()
}
func (this *RBool) Value() bool {
	return this.value
}

