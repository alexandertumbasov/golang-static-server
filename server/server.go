package server

import (
	"fmt"
	"golang-static-server/db"
)

func Init() {
	r := NewRouter()
	client := db.Init()
	fmt.Println(client)
	r.Run(":8080")
}
