/*

        hellostringutil.go
        Contoh sederhana untuk menggambarkan cara menggunakan lib
        Diambil dari https://golang.org/doc/code.html

*/
package main

import (
	"fmt"
	"github.com/andikristiant86/latihanGo/src/reverse"
)

func main() {
	fmt.Printf(stringutil.Reverse("Hello, World!"))
}