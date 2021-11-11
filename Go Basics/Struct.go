package main

type Marks struct {
	name  string
	maths int
	chem  int
	phy   int
	csc   int
	avg   int
}

func main() {
	var per1 Marks

	per1.name = "Ashwin Kumar R"
	per1.maths = 90
	per1.chem = 50
	per1.phy = 78
	per1.avg = per1.maths + per1.phy + per1.chem
	println("The Result for", per1.name, "is", per1.avg)

}
