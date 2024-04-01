package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Generate JWT 32Byte string... ")
	fmt.Println(generate32BytesPwd())
	fmt.Println(generate32BytesPwd())
}

func generate32BytesPwd() string {

	runes := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, 32)
	for i := range b {

		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)

}
