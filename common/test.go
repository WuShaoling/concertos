package common;

import "fmt"
import "reflect"
import "unsafe"

type T struct {
	Id  int
	Ips []string
}

func main() {
	s := []T{{Id: 1,
		Ips: []string{"123", "456"}},}
	var sk []T
	var i interface{} = &sk

	svi := reflect.Indirect(reflect.ValueOf(i))
	ss := reflect.Indirect(reflect.NewAt(svi.Type(), unsafe.Pointer(reflect.ValueOf(&s).Pointer())))
	svi.Set(ss)

	fmt.Println(svi)
}
