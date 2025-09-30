package main

import (
	"github.com/fiqryomaratala/simple-forum/internal/handlers/memberships"
	"github.com/gin-gonic/gin"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	membershipsHandler := memberships.NewHandler(r)
	membershipsHandler.RegisterRoute()
	r.Run(":9999")
}
