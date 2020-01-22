package main

import (
	"fmt"
	"github.com/KennyChenFight/goChatter/lib/auth"
	"github.com/KennyChenFight/goChatter/lib/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
	"xorm.io/xorm"
)

func init() {
	connectStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSL_MODE"))
	db, err := xorm.NewEngine("postgres", connectStr)
	if err != nil {
		log.Panic("DB connection initialization failed", err)
	}

	secretKey := os.Getenv("SECRET_KEY")
	tokenLifeTime, err := time.ParseDuration(os.Getenv("TOKEN_LIFE_TIME"))
	if err != nil {
		log.Panic("JWT life time parse failed", err)
	}

	auth.Init([]byte(secretKey), tokenLifeTime)
	middleware.Init(db)

	log.Println("init dependency success")
}

func main() {
	router := gin.Default()

	router.Run()
}


