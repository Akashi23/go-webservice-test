package main

func main() {
	ConnectDatabase()

	e := InitRouter()
	e.Logger.Fatal(e.Start(":8000"))
}
