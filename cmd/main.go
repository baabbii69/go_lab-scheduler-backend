package main

import (
	"fmt"
	"log"

	"github.com/baabbii69/go_lab-scheduler-backend/config"
	"github.com/baabbii69/go_lab-scheduler-backend/database"
	"github.com/baabbii69/go_lab-scheduler-backend/routes"
	"github.com/gin-gonic/gin"
)

// func main() {
// 	cfg, err := config.LoadConfig()
// 	if err != nil {
// 		log.Fatalf("Error loading config: %v", err)
// 	}

// 	fmt.Println("Configuration Loaded:")
// 	fmt.Println("APP_ENV:", cfg.AppEnv)
// 	fmt.Println("APP_PORT:", cfg.AppPort)
// 	fmt.Println("DB_HOST:", cfg.DBConfig.Host)
// 	fmt.Println("DB_PORT:", cfg.DBConfig.Port)
// 	fmt.Println("DB_USER:", cfg.DBConfig.User)
// 	fmt.Println("DB_PASSWORD:", cfg.DBConfig.Password) // Only for debugging, remove later!
// 	fmt.Println("DB_NAME:", cfg.DBConfig.DBName)
// 	fmt.Println("DB_SSLMODE:", cfg.DBConfig.SSLMode)
// }

func main() {
	// Loading configuration
	cfg, err:= config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initializing database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	_ = db

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	fmt.Printf("Configuration loaded:\nAppEnv: %s\nAppPort: %s\nDB Host: %s\n",
        cfg.AppEnv, cfg.AppPort, cfg.DBConfig.Host)

	router := gin.Default()
	routes.RegisterRoutes(router)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
