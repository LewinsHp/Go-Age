package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Umur Kamu :")

	var umur string
	_, err := fmt.Scanln(&umur)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("----\n")

	inputInt, _ := strconv.Atoi(umur)
	fmt.Printf("Kamu Lahir Pada Tahun Berapa %d\n",
		2022-inputInt)
}
