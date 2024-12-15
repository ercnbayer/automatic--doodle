package main

import (
	"automatic-doodle/di"
	_ "automatic-doodle/ent/runtime"
)

func main() {
	db := di.DBBuilder()
	server := di.Wire(db)

	server.Listen(4000)
}
