package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func capchalol() {
	rand.Seed(time.Now().UnixNano())
	randomNum := random(1, 20)
	randomNum2 := random(1, 20)
	fmt.Printf("\x1b[96m%d + %d? \n", randomNum, randomNum2)
	//fmt.Printf("%d \n", randomNum+randomNum2) output
	var as = randomNum + randomNum2
	var capcha int
	fmt.Print("\x1b[96mAwnser: ")
	fmt.Scanln(&capcha)
	if capcha != as {
		fmt.Printf("\rCaptcha Incorrect\r\n")
		return
	}
	fmt.Printf("\r\x1b[96mUwU u got it g\r\n")
}

func main() {
	capchalol()
}
