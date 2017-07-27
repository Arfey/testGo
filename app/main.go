package main

func main() {
	app := App{}
	app.InitDB(
		"postgres",
		"postgres",
		"postgres",
	)
	app.Run(":8081")
}
