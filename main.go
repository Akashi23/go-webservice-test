package main

func main() {
	e := InitRouter()
	e.Logger.Fatal(e.Start(":8000"))
}
