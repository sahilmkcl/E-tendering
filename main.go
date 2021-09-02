package main

func main() {
	r := registerRoutes()
	r.Run("127.0.0.1:8080")
}
