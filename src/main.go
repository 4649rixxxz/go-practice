package main

import (
	"fmt"
)

func main() {
	// 文字列はダブルクォートで囲む
	conferenceName := "Go Conference"
	// 短縮記法は使えない
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// slice
	bookings := []string{}
	fmt.Println(bookings)
	// 変数宣言
	var x, y = 1, 3
	fmt.Println(x, y)


	// 型を確認することができる
	fmt.Printf("conferenceTickets is %T\n", conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Remaining Tickets are %v \n", remainingTickets)

	// スタックに積まれる
	defer fmt.Println("end")
	defer fmt.Println("the")
	// ループ
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
	}

	numbers := []int {1,2,3,4,5,6}
	// sliceの一部を切り取る
	// part := numbers[0:2]
	part := numbers[2:]
	fmt.Println(part)
	part = numbers[:3]
	fmt.Println(part)
	

	makeBy := make([]int, 5)
	fmt.Println(makeBy)
	makeBy = append(makeBy, 1)
	fmt.Println(makeBy)

	for index, value := range makeBy {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}

	// インデックスいらない場合
	for _, value := range makeBy {
		fmt.Println("no index:")
		fmt.Println("value:", value)
	}

	// 連想配列
	m := map[string]int{
		"key1": 100,
		"key2": 200,
		"key3": 300,
	}

	fmt.Println(m["key1"])

	// キーがある場合はokがtrue
	elem, ok := m["key1"]


	fmt.Println(elem, ok)

	plus := increment()
	fmt.Println(plus())
	fmt.Println(plus())
	fmt.Println(plus())
	fmt.Println(plus())
	fmt.Println(plus())
	fmt.Println(plus())

	tom := Person{
		"Tom",
		22,
	}

	fmt.Println(tom.Introduce())
	// 上記と等価
	fmt.Println((&tom).Introduce())

	tom.changeAge()
	// 上記と等価
	(&tom).changeAge()
	// 参照先での変更が反映されている
	fmt.Println(tom.Age)


	var animal Animal

	lion := Lion{"lion"}
	// interfaceで定義したメソッドと同じメソッドを定義していない場合はエラー
	animal = lion
	animal.bite()
	animal.walk()

	var t *T
	if t == nil {
		fmt.Println("###nil###")
	}
}

// クロージャ
func increment() func() int {
	x := 1
	return func () int {
		x++
		return x
	}
}

type Person struct {
	Name string
	Age int
}

func (p Person) Introduce() string {
	return "My Name is " + p.Name
}

// ポインタレシーバ
func (p *Person) changeAge() {
	p.Age = 50
}

type Animal interface {
	bite()
	walk()
}

type Lion struct {
	name string
}


func (l Lion) bite() {
	fmt.Println("woooo")
}

func (l Lion) walk() {
	fmt.Println("walk")
}

type T struct {
	S string
}

