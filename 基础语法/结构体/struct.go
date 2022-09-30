package main

import "fmt"

type Books struct {
	title   string
	author  string
	time    string
	book_id int
}

type MyBook struct {
	Books
	price int
	name  string
}

func main() {
	// 创建一个结构体
	fmt.Println(Books{"Go学习笔记", "zuojt", "2020-09-29", 93844552324})

	// 使用key - value 格式
	fmt.Println(Books{title: "Go学习笔记", author: "zuojt", time: "2020-09-29", book_id: 93844552324})

	// 忽略某些字段
	fmt.Println(Books{title: "Go学习笔记", book_id: 93844552324})

	var book Books

	book.title = "Hello Go"
	book.author = "zuojt"
	// 访问结构体成员
	fmt.Println("结构体book成员,title:", book.title)

	// 使用指针
	printBook(&book)

	// 结构体继承
	var my_book MyBook
	my_book.price = 78
	my_book.name = "三体"
	fmt.Println("结构体继承测试:", my_book.name, my_book.price)
}

func printBook(book *Books) {
	fmt.Println("Book title: %s\n", book.title)
	fmt.Println("Book author: %s\n", book.author)
}
