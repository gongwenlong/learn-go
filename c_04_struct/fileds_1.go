package main
import (
	"fmt"
	"strings"
	"unsafe"
)

// 树的结构体
type Node struct {
	pr      *Node
	data    float64
	su      *Node
}

type Tree struct {
	le      *Tree
	data    float64
	ri      *Tree
}

type Person struct {
	firstName   string
	lastName    string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1-struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"

	fmt.Print("person 1:  ")
	fmt.Println(pers1)
	fmt.Println(&pers1)

	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)


	// 2—struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward"  // 这是合法的

	fmt.Print("person 2:  ")
	fmt.Println(pers2)
	fmt.Println(&pers2)

	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3—struct as a literal:
	pers3 := &Person{"Chris","Woodward"}
	fmt.Print("person 3:  ")
	fmt.Println(pers3)
	fmt.Println(&pers3)

	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)

	size := unsafe.Sizeof(Person{"casn", "asf"})
	fmt.Print(size)
}
