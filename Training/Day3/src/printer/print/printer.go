package print

type Printer interface {
	PageSetup()
	Preview()
	Print()
}

type abc interface {
}

func Print(p Printer) {
	p.Print()
}

func PageSetup(p Printer) {
	p.PageSetup()
}

func Preview(p Printer) {
	p.Preview()
}
