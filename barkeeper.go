package main

import "fmt"

type product struct {
	name     string
	category string
	price    float64
}

func main() {
	products := []product{
		{"Juniper", "Gin", 4.50},
		{"Tanqueray", "Gin", 5.50},
		{"Bulldog", "Gin", 6.00},
		{"Monkey 47", "Gin", 8.50},
		{"Hendrick's", "Gin", 7.00},

		{"Jack Daniel's", "Whiskey", 5.00},
		{"Glenfiddich", "Whiskey", 7.50},
		{"Jim Beam", "Whiskey", 4.50},
		{"Macallan", "Whiskey", 9.00},

		{"Plantation dark", "Rum", 4.50},
		{"Plantation 3 Stars", "Rum", 4.50},
		{"Bumbu", "Rum", 6.50},
		{"Diplomatico", "Rum", 8.50},

		{"Clase Azul", "Tequila", 12.00},
		{"Don Ramon", "Tequila", 6.50},

		{"Finlandia", "Vodka", 4.00},
		{"Beluga", "Vodka", 6.00},
		{"Belvedere", "Vodka", 6.50},
		{"Grey Goose", "Vodka", 7.00},
	}

	var choice int

	fmt.Println("===== BAR POS =====")
	fmt.Println("1. Gin")
	fmt.Println("2. Whiskey")
	fmt.Println("3. Rum")
	fmt.Println("4. Tequila")
	fmt.Println("5. Vodka")

	var selectedCategory string
	fmt.Print("Select category: ")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		selectedCategory = "Gin"
	case 2:
		selectedCategory = "Whiskey"
	case 3:
		selectedCategory = "Rum"
	case 4:
		selectedCategory = "Vodka"
	case 5:
		selectedCategory = "Tequila"
		return
	default:
		fmt.Println("Invalid choice")
		return
	}
	fmt.Println()
	fmt.Println("=====", selectedCategory, "=====")

	for _, product := range products {
		if product.category == selectedCategory {
			fmt.Printf("%s - %.2f EUR\n", product.name, product.price)
		}
	}
}
