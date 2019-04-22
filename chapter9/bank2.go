// p. 309
package main

var (
	sema    = make(chan struct{}, 1) // marker
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // catch marker
	balance = balance + amount
	<-sema // free marker
}

func Balance() int {
	sema <- struct{}{} // catch marker
	b := balance
	<-sema // free marker
	return b
}
