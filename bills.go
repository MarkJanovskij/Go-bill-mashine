package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//make new bills

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

//format the bill "soup": 4.99, "pie": 7.99, "salad": 6.99, "toffee pudding": 3.55

func (b *bill) format() string {
	fs := "Bill breackdown: \n"
	var total float64 = 0

	//list of items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ....$%v \n", k+":", v)
		total += v
	}
	//Total
	fs += fmt.Sprintf("%-25v ... $%0.2f", "total:", total)

	// add a tip
	fs += fmt.Sprintf("%-25v ....$%v\n", "tip:", b.tip)
	//total
	fs += fmt.Sprintf("%-25v ....$%0.2f", "total:", total+b.tip)
	return fs
}

// update the tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// Save Bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("billscopy/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)

	}
}
