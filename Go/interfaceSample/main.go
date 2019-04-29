package main

import (
	"fmt"
	"reflect"
)

func main() {
	A("ABCDEFG")
	A(12345)

	B("This is test.")

	C(MySQL{})
	C(MongoDB{})
}

func A(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("value 的值是 string： %s\n", v)
	case int:
		fmt.Printf("value 的值是 int： %d\n", v)
	}
}

func B(value interface{}) {
	fmt.Println("value 的值是: " + value.(string) + "\n")
}

type Database interface {
	Read(name string) string
	Write(data string)
}

// MySQL 會實作 Database 介面。
type MySQL struct{}

func (m MySQL) Read(n string) string {
	// return fmt.Sprintf("%T done.\n", MySQL{})
	return fmt.Sprintf("%s done.\n", reflect.TypeOf(m).Name())
}
func (m MySQL) Write(d string) {
	fmt.Printf("Write to %s %s\n", reflect.TypeOf(m).Name(), d)
}

// MongoDB 會實作 Database 介面。
type MongoDB struct{}

func (m MongoDB) Read(n string) string {
	// return fmt.Sprintf("%T done.\n", MongoDB{})
	return fmt.Sprintf("%s done.\n", reflect.TypeOf(m).Name())
}
func (m MongoDB) Write(d string) {
	fmt.Printf("Write to %s %s\n", reflect.TypeOf(m).Name(), d)
}

func C(db Database) {
	s := "db connection string"
	fmt.Printf("db read from %s", db.Read(s))
	db.Write("successfully.")
}
