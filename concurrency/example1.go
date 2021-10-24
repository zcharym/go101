package main

func main() {
	ch := make(chan int)
	go a(ch)
	go b(ch)
	select {} // prevent the program from terminating, ignore for now
}
func a(ch chan int) {
	println("a before")
	ch <- 5
	println("a after")
}
func b(ch chan int) {
	println("b before")
	println(<-ch)
	println("b after")
}
