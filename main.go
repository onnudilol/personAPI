package main

func main() {
	a := App{}
	a.Initialize(host, port, user, dbname)
	a.Run(":8080")
}
