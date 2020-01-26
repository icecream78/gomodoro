package cmd

type BarWidgeter interface {
	Run()
}

func NewBar() *Bar {
	return &Bar{}
}

type Bar struct {
	widget BarWidgeter
}

func (b *Bar) SetTimer(ticker BarWidgeter) *Bar {
	b.widget = ticker
	return b
}

func (b *Bar) Run() {
	b.widget.Run()
}
