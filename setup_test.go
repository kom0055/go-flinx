package flinx

type foo struct {
	f1 int
	f2 bool
	f3 string
}

func getF1(f foo) int {
	return f.f1
}

func getF2(f foo) bool {
	return f.f2
}
