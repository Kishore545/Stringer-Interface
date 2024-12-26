package StringerInterface

import (
	"fmt"
	"strconv"
)

type Book struct {
	title string
}

func (b Book) String() string {
	return fmt.Sprint("The title of the Book is ", b.title)
}

type Count int

func (c Count) String() string {
	return fmt.Sprint("Tis is te Number ", strconv.Itoa(int(c)))
}

/*func main() {
	b := Book{title: "ADFSRF"}
	var c count = 42
	fmt.Println(b)
	fmt.Println(c)
}*/
