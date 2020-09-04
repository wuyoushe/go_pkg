
package main


func Hello() {
	fmt.Println("Hello go mod")
}

func Bye() {
	fmt.Println("Bye go Mod!")
	fmt.Println(quote.Hello())
}