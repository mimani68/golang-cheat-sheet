package mathmodule

type Result struct {
	Result int
	Error  bool
}

func Calc(number int, ch chan Result) {
	a := Result{
		Result: number + 1,
		Error:  false,
	}
	ch <- a
}
