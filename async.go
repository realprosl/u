package u

type Promise struct {
	status bool
	data   string
	err    string
	Then   func()
	Cath   func()
}

func (r *Promise) Async(f func() string) {
	r.reset()
	go r.runThen(r.Then)
	go r.runCath(r.Cath)
	r.data = f()
	for {
		if r.data != "" {
			r.status = true
			return
		}
	}
}
func (r *Promise) runCath(f func()) {
	cont := 0
	for {
		if cont > 10000000000 && !r.status {
			r.err = "Error : fail response"
			f()
			return
		}
		cont++
	}
}
func (r *Promise) runThen(f func()) *Promise {
	for {
		if r.status && r.err == "" {
			f()
			return r
		}
	}
}
func (r *Promise) reset() {
	r.data = ""
	r.err = ""
	r.status = false
}
