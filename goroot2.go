package main

import (
	"fmt"
	"time"
)

func firstMess(mess1 string) {
	fmt.Println(mess1)
	time.Sleep(1 * time.Second)
}

func firstMess2(mess1 string) {
	fmt.Println(mess1)
	time.Sleep(1 * time.Second)
}

func firstMess3(mess1 string) {
	fmt.Println(mess1)
	time.Sleep(1 * time.Second)
}

func firstMess4(mess1 string) {
	fmt.Println(mess1)
	time.Sleep(1 * time.Second)
}

func firstMess5(mess1 string) {
	fmt.Println(mess1)
	time.Sleep(2 * time.Second)
}

func main() {
	go firstMess("Hello World")
	go firstMess2("Hello World")
	go firstMess3("Hello World")
	go firstMess4("Hello World")
	go firstMess5("Hello World")
	time.Sleep(3 * time.Second)
}
