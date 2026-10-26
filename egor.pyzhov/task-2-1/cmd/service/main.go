package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("Error: Enter a whole number for the departments")
		return
	}

	if n < 1 || n > 1000 {
		fmt.Println("Error: the number of departments must be between 1 and 1000")
		return
	}

	for i := 1; i <= n; i++ {

		var (
			k       int
			minTemp int = 15
			maxTemp int = 30
		)

		_, err := fmt.Scan(&k)

		if err != nil {
			fmt.Println("Error: Enter a whole number for the employees")
			return
		}

		if k < 1 || k > 1000 {
			fmt.Println("Error: the number of employees must be between 1 and 1000")
			return
		}

		for j := 1; j <= k; j++ {

			var (
				operation string
				temp      int
			)

			_, err := fmt.Scan(&operation, &temp)

			if err != nil {
				fmt.Println("Error: Enter an operator (>= or <=) and a whole number for the temp")
				return
			}

			switch operation {
			case ">=":
				if temp > minTemp {
					minTemp = temp
				}
			case "<=":
				if temp < maxTemp {
					maxTemp = temp
				}
			default:
				fmt.Println("Error: operator must be >= or <=")
				return
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
