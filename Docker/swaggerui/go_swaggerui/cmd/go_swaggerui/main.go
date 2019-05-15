package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/specs", "/specs")
	r.Static("/swaggerui", "/swaggerui")
	r.Run()
}
