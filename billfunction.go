package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prommt string, r *bufio.Reader) (string, error) {
	fmt.Print(prommt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func creatBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Creat new customer's bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Creat new customer's bill name: ", reader)

	b := newBill(name)
	fmt.Print("Create the bill -", b.name)

	return b
}

func promtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choos from options (a - add item, s - save the bill, t - add a tip) : ", reader)

	switch opt {
	case "a":
		name, _ := getInput("item name", reader)
		price, _ := getInput("item price", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promtOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("item is added - ", name, price)
		promtOptions(b)

	case "s":
		//sav, _ := getInput("Bill is saved", reader)
		b.save()
		fmt.Println("you chose to save the bill", b.name)
		promtOptions(b)
	case "t":
		tip, _ := getInput("Enter Tip amount ($):", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			//promtOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip added - ", tip)
		promtOptions(b)
	default:
		fmt.Println("that was not a valid option ....")
		promtOptions(b)
	}
	fmt.Println(opt)
}

func main() {
	mybill := creatBill()
	promtOptions(mybill)
	// mybill.addItem("onion soup", 4.50)
	// mybill.addItem("veg pie", 8.50)
	// mybill.addItem("toffee pudding", 4.95)
	// mybill.addItem("coffee", 2.50)
	// mybill.addItem("pelmeni", 4.25)

	// mybill.updateTip(10)

	//fmt.Println(mybill.format())

	fmt.Print(mybill)
}
