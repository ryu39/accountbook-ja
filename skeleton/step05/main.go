package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Category string
	Price    int
}

func main() {
	file, err := os.Create("accountbook.txt")
	if err != nil {
		log.Fatal(err)
	}

	var n int
	fmt.Print("何件入力しますか>")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		if err := inputItem(file); err != nil {
			log.Fatal(err)
		}
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	if err := showItems(); err != nil {
		log.Fatal(err)
	}
}

func inputItem(file *os.File) error {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	line := fmt.Sprintf("%s %d\n", item.Category, item.Price)
	if _, err := file.WriteString(line); err != nil {
		return err
	}

	return nil
}

func showItems() error {
	file, err := os.Open("accountbook.txt")
	if err != nil {
		return err
	}

	fmt.Println("===========")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		splited := strings.Split(line, " ")
		if len(splited) != 2 {
			return fmt.Errorf("パースに失敗しました")
		}

		category := splited[0]

		price, err := strconv.Atoi(splited[1])
		if err != nil {
			return err
		}

		fmt.Printf("%s:%d円\n", category, price)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("===========")

	return nil
}
