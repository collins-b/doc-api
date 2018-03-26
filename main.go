package main

func main() {
	a := App{}
	a.Initialize("db_user", "db_pass", "dbname")

	a.Run(":3000")
}
