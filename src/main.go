package main

import (
	"os"

	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/controllers"
	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/middleware"
	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/model"
	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	router := setUp()
	router.Run()
}

/*ルーター等のセットアップ*/
func setUp() *gin.Engine {
	router := gin.Default()
	//DBのmigrate
	DSN := os.Getenv("LOCAL_DB_DSN")
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})

	//TODO: corsの設定
	//serviceの設定
	user_service := repository.NewUserService(db)
	//controllerの設定
	user_controller := controllers.NewUserController(user_service)
	//routerの設定
	api := router.Group("/api")
	v1 := api.Group("/v1")
	// /userと/authと/iiifだけはmiddlewareなしで通るように
	v1.POST("/user", user_controller.CreateUser)
	v1.POST("/auth", user_controller.Login)
	//以下middlwareを使用するところ
	service := v1.Group("/service")
	service.Use(middleware.AuthMiddleware)
	{
		service.GET("/manifest", controllers.GetAllManifest)
	}

	return router
}
