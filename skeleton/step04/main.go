package main

import "fmt"

type Item struct {
	Category string
	Price    int
}

func main() {
	var n int
	fmt.Print("品目数>")
	fmt.Scan(&n)

	items := make([]Item, 0, n)

	for i := 0; i < n; i++ {
		items = inputItem(items)
	}

	showItems(items)
}

func inputItem(items []Item) []Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	items = append(items, item)

	return items
}

func showItems(items []Item) {
	fmt.Println("===========")

	for i := 0; i < len(items); i++ {
		item := items[i]
		fmt.Printf("%s:%d\n", item.Category, item.Price)
	}

	fmt.Println("===========")
}
