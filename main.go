package main

import (
	"notion/model"
)

func main() {

	client := model.NotionInit()
	model.Update(client)
}
