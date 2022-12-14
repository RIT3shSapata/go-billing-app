package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	mybill := newBill("turtle's bill")

// 	mybill.addItem("onion soup",4.50)
// 	mybill.addItem("veg pie",8.95)
// 	mybill.addItem("toffee pudding",4.95)
// 	mybill.addItem("coffee",3.25)

// 	mybill.updateTip(10)

// 	fmt.Println(mybill.format())
// }

func getInput(prompt string,r *bufio.Reader) (string,error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)
	return input,err
}

func createBill() bill{
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader) 
	b:=newBill(name)
	return b
}

func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)

	opt,_ := getInput("Choose option (a - add item, s - save bill, t - add tip):",reader)

	switch opt{
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p,err := strconv.ParseFloat(price,64)
		if err != nil{
			fmt.Println("The price must be a number")
		promptOptions(b)
		}
		b.addItem(name,p)
		fmt.Println("item added - ",name,price)
		promptOptions(b)

	case "t":
		tip,_ := getInput("Tip amount: ", reader)
		t,err := strconv.ParseFloat(tip,64)
		if err != nil{
			fmt.Println("The tip must be a number")
		promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ",tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("you saved the file - ", b.name)
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}

func main(){
	mybill:=createBill()
	promptOptions(mybill)
}