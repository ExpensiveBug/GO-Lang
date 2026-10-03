package main

import "fmt"

func main() {
	var n1, n2 float64
	fmt.Println("Enter two Numbers : ")
	fmt.Scan(&n1, &n2)

	var op int
	fmt.Println("1. Addition \n2. Subtraction \n3. Multiplication \n4. Division \nChoose an Option : ")
	fmt.Scan(&op)

	switch op {
	case 1:
		fmt.Println("Addition : ", n1+n2)
	case 2:
		fmt.Println("Subtraction : ", n1-n2)
	case 3:
		fmt.Println("Multiplication : ", n1*n2)
	case 4:
		if n2 == 0 {
			fmt.Println("Cannot Divide by Zero!")
		} else {
			fmt.Println("Division : ", n1/n2)
		}
	default:
		fmt.Println("You have not chosen an operation!")
	}

}
