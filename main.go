package main

func main() {
	server := NewAPIServer(":3380")
	server.Run()
}
