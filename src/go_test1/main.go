package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func Sum(a int, b int) int {
	return a + b

}

func main() {
	mali := "MaliKarupin"

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(Hello(mali))

	fmt.Print("Age of Mali")
	scanner.Scan()
	j, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Age of Mali", "__", j)
}
