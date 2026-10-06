package main

import (
	"log"
	"main/Controllers"
	"main/Models"
	"main/Repository"
	"main/Services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func initDB() *gorm.DB {
	url, _ := os.LookupEnv("DB_URL")
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Models.Product{})

	return db
}

func initProductAPI(db *gorm.DB) Controllers.MainStruct {
	productRepository := Repository.ProvideProductRepository(db)
	productService := Services.ProvideProductService(productRepository)
	mainAPI := Controllers.MainStructAPI(productService)
	return mainAPI
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	db := initDB()
	defer db.Close()

	MainStructAPI := initProductAPI(db)

	router := gin.Default()
	router.GET("/api/v1/wallets/:walletId", MainStructAPI.MainIndex)
	router.POST("api/v1/wallet", MainStructAPI.MainPost)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
