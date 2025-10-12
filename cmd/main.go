package main

import (
	"log"

	"github.com/fiqryomaratala/simple-forum/internal/configs"
	"github.com/fiqryomaratala/simple-forum/internal/handlers/memberships"
	"github.com/gin-gonic/gin"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	//Function inisiasi
	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	membershipsHandler := memberships.NewHandler(r)
	membershipsHandler.RegisterRoute()
	r.Run(cfg.Service.Port)
}
