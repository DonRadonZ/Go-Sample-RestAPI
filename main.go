package main

import (
	"examples/go-sample-restapi/db"
	"examples/go-sample-restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
  db.InitDB()
  server := gin.Default()

  routes.RegisterRoutes(server)
  
  server.Run(":9000") // localhost:9000
}



