package main

import (
	"kanbanboard/database"
	_ "kanbanboard/initializer"
)

func init() {
	database.Connect()
}

func main() {

}
