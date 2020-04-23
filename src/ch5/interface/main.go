package main

import "fmt"

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	// 型アサーション
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}

	vs := []Stringify{
		&Person{Name: "Taro", Age: 22},
		&Car{Number: "111112", Model: "Audi"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	Println(&Person{Age: 23, Name: "Hanako"})
	Println(&Car{Number: "xyz-00001", Model: "rt-38"})

	// fmt.Stringerインターフェースを実装していれば、fmt.Printlnを実行する際に実装したメソッドを実行してくれる
	t := T{Id: 10, Name: "Taro"}
	fmt.Println(&t)

	i := I2(&T2{Id: 1})
	fmt.Println(i.Method1(), i.Method2(), i.Method3())
}

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{
		Message: "エラーが発生しました",
		ErrCode: 1234,
	}
}

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s]%s", c.Number, c.Model)
}

func Println(s Stringify) {
	fmt.Println(s.ToString())
}

type T struct {
	Id   int
	Name string
}

func (t *T) String() string {
	return fmt.Sprintf("<<%d, %s>>", t.Id, t.Name)
}

type I0 interface {
	Method1() string
}

type I1 interface {
	I0
	Method2() string
}

type I2 interface {
	I1
	Method3() string
}

type T2 struct {
	Id int
}

func (t *T2) Method1() string {
	return fmt.Sprintf("Method1: id: %d", t.Id)
}

func (t *T2) Method2() string {
	return fmt.Sprintf("Method2: id: %d", t.Id)
}

func (t *T2) Method3() string {
	return fmt.Sprintf("Method3 id: %d", t.Id)
}
