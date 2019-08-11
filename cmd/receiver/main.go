package main

import . "github.com/ingtranet/nriy-receiver"

func main() {
	receiver := NewApp()
	receiver.Run()
}
