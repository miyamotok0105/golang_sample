
package main

import "fmt"

type Person struct {
    name string
    age int
}

type Person2 struct {
    Name string	// 外部からアクセス可能
    Age int		// 外部からアクセス可能
    status int		// 外部からアクセス不可
}

//パーソンの拡張メソッド
func (p *Person) SetPerson(name string, age int) {
    p.name = name
    p.age = age
}

func (p *Person) GetPerson() (string, int) {
    return p.name, p.age
}

func main() {
    var p1 Person
    var name string
    var age int
    p1.SetPerson("Yamada", 26)
    name, age = p1.GetPerson()
    fmt.Printf(name, age)
    fmt.Printf(string(age))

    a1 := Person2{"Yamada", 26, 1}			// 順序通りに初期化
    a2 := Person2{Name: "Tanaka", Age: 32, status: 1}		// 名前で初期化
    fmt.Printf("%v\n", a1)
    fmt.Printf("%v\n", a2)
}
