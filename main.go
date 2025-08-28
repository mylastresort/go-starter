package main

import "server/internal"

func main() {
	internal.Init()

	internal.StartServer()
	defer internal.StopServer()
}
