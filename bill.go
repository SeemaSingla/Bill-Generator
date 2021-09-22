package main

type bill struct {
	name string
	item map[string]float64
	tip  float64
}

//make new bills
func newBill(name string) bill {
	b := bill{
		name: name,
		item: map[string]float64{},
		tip:  0,
	}
	return b
}

//Receiver functions
//format bill

func (b bill) format() string {

}

//now the above can only be called as:
// billObject = newBill()
//billObject.format()
