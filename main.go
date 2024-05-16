package main

import "fmt"

//func main() {
// var nameOne string = "mario"
// var nameTwo = "luigi"
// var nameThree string
// fmt.Println(nameOne, nameTwo, nameThree)

// nameOne = "peach"
// nameThree = "boser"

// fmt.Println(nameOne, nameTwo, nameThree)

// nameFour := "yoshi"
// fmt.Println(nameFour)

// //ints
// var ageOne int = 20
// var ageTwo = 30
// ageThree := 40

// fmt.Println(ageOne, ageTwo, ageThree)

// //bit & memory
// var numOne int8 = 25
// var numTwo = -128
// var numThree uint16 = 266

// var scoreOne float32 = 25.6963
// var scoreTwo float64 = 14851456126481
// scoreThree := 1.5
// fmt.Println(numOne, numTwo, numThree, scoreOne, scoreTwo, scoreThree)
// age := 33
// name := "Mark Janovskij"
// //Print
// fmt.Print("hello,  ")
// fmt.Print("world!! \n")
// fmt.Print("new line \n")

// fmt.Println("hello world! ")
// fmt.Println("goodbye Pirats!")
// fmt.Println("my age is ", age, "and my name is", name)

// // PrintF (formated strings) %_ = format specifier
// fmt.Printf(" my age is %v and my name is %v \n", age, name)
// fmt.Printf(" my age is %q and my name is %q \n", age, name)
// fmt.Printf(" my age is of typr %T \n ", age)
// fmt.Printf("you scored %f points! \n", 225.55)
// fmt.Printf("you scored %0.1f points! \n", 255.55)

// //Sprintf (save formatted string)
// var str = fmt.Sprintf(" my age is %v and my name is %v \n", age, name)
// fmt.Println("the saved string is: ", str)

// var ages [3]int = [3]int{20, 25, 30}
// names := [4]string{"yoshi", "mario", "peach", "bowser"}
// fmt.Println(ages, len(ages))
// fmt.Println(names, len(names))

// //slices( use array under the hood)

// var scores = []int{100, 50, 60}
// scores[2] = 25
// scores = append(scores, 85)

// fmt.Println(scores, len(scores))
// //slice renges
// rangeOne := names[1:3]
// rangeTwo := names[2:]
// rangeThree := names[:3]
// fmt.Println(rangeOne, rangeTwo, rangeThree)

// rangeOne = append(rangeOne, "koopa")
// fmt.Println(rangeOne)

// greeting := "hello there friends!"
// fmt.Println(strings.Contains(greeting, "hello"))
// fmt.Println(strings.ReplaceAll(greeting, "hello", "Hi"))
// fmt.Println("original string value =", greeting)
// fmt.Println(strings.ToUpper(greeting))
// fmt.Println(strings.Index(greeting, "th"))
// fmt.Println(strings.Split(greeting, " "))

// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

// sort.Ints(ages)
// fmt.Println(ages)

// index := sort.SearchInts(ages, 30)
// fmt.Println(index)

// // names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

// // sort.Strings(names)
// // fmt.Println(names)

// // fmt.Println(sort.SearchStrings(names, "bowser"))

// x := 0
// for x < 5 {
// 	fmt.Println("value of x is:", x)
// 	x++
// }
// for i := 0; i < 5; i++ {
// 	fmt.Println("value of i is:", i)
// }

// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
// // for i := 0; i < len(names); i++ {
// // 	fmt.Println(names[i])
// // }
// //for _, value:= range names {fmt.Printf("the value at index %v  \n", value)

// for index, value := range names {
// 	fmt.Printf("the value at index %v is %v \n", index, value)
// 	value = "new string"
// }
// fmt.Println(names)

// age3 := 45

// fmt.Println(age3 <= 50)
// fmt.Println(age3 >= 50)
// fmt.Println(age3 <= 50)
// fmt.Println(age3 == 45)
// fmt.Println(age3 != 50)

// if age3 < 30 {
// 	fmt.Println("age is less that 30")
// } else if age3 < 40 {
// 	fmt.Println("age is less that 40")
// } else {
// 	fmt.Println("age is not less that 45")
// }

// names2 := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

// for index, value2 := range names2 {
// 	if index == 1 {
// 		fmt.Println("counting at pos", index)
// 		continue
// 	}
// 	if index > 2 {
// 		fmt.Println("bracking at pos ", index)
// 		break
// 	}
// 	fmt.Printf("the value at pos %v is %v \n", index, value2)
// }

// functions
// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v \n", n)
// }

// func sayBye(n string) {
// 	fmt.Printf("Goodbye %v \n", n)
// }

// func cycleNames(n []string, f func(string)) {
// 	for _, v := range n {
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
//}

// func getInitials(n string) (string, string) {
// 	s := strings.ToUpper(n)
// 	names := strings.Split(s, " ")

// 	var initials []string
// 	for _, v := range names {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], "_"

// }

var score = 99.5

// func updateName(x string) string {
// 	x = "wedge"
// 	return x
// }

func updateName(x *string) {
	*x = "wedge"

}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main2() {
	// sayGreeting("mario")
	// sayGreeting("luigi")
	// sayBye("mario")
	// sayBye("luigi")
	// cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	// cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	// a1 := circleArea(10.5)
	// a2 := circleArea(15)

	// fmt.Println(a1, a2)
	// fmt.Printf("cirlce 1 is %0.3f and circle 2 is %0.3f", a1, a2)
	// fn1, sn1 := getInitials("tifa lockhart")
	// fmt.Println(fn1, sn1)

	// fn2, sn2 := getInitials("cloud strife")
	// fmt.Println(fn2, sn2)

	// fn3, sn3 := getInitials("barret")
	// fmt.Println(fn3, sn3)
	sayHello("mario")
	for _, v := range points {
		fmt.Println(v)
	}
	showScore()

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}
	updateMenu(menu)
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping maps

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}
	// ints as key type
	phonebook := map[int]string{
		541756:   "mario",
		2345362:  "luigi",
		56347568: "peach",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[541756])

	phonebook[2345362] = "bowser"
	fmt.Println(phonebook)

	name := "tifa"
	//name = updateName(name)

	// updateName(name)

	//fmt.Println("memory address of name is ", &name)

	m := &name
	fmt.Println("memory adress", m)
	fmt.Println("value at memory adress:", *m)

	fmt.Println(name)

	updateName(m)
	fmt.Println(name)
}

// func main2() {
// 	greeting := "hello friends!"
// 	fmt.Println(strings.Contains(greeting, "hello"))
// }
