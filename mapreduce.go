package main

import "fmt"

func Map(sw func(string) []Words, data string) []Words {
	fmt.Println(sw(data))
	return sw(data)
}

func Reduce() {

}
