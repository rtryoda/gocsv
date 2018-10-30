package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("tmp/test.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	if err := file.Truncate(0); err != nil {
		fmt.Println(err)
		return
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{"name", "age"})
	writer.Write([]string{"Alice", "20"})
	writer.Write([]string{"Bob", "21"})
	writer.Write([]string{"Carol", "22"})
	writer.Flush()

	fmt.Println("success creating csv")
}
