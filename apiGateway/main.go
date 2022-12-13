package main

func main() {
	app, err := BuildApplication()
	if err != nil {
		panic(err)
	}
	app.Start()
}
