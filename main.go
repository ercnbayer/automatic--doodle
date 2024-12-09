package main

import "automatic-doodle/di"

func main() {
	db := di.DBBuilder()
	server := di.Wire(db)

	server.Listen(4000)
}
