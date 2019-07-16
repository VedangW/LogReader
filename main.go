package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	var err error

	openLink := "Hi"

	Db, err = gorm.Open("mysql", openLink)
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&FilePath{})
	Db.AutoMigrate(&DynJsonEntry{})
}

func main() {
	router := gin.Default()

	adminapi := router.Group("/admin")
	{
		adminapi.POST("/", CreateDynJsonObject)
		adminapi.GET("/:eid", GetDynJsonObjectWithIdAdmin)
		adminapi.POST("/store", DynStoreFileInDb)
	}

	userapi := router.Group("/user")
	{
		userapi.POST("/", CreateDynJsonObject)
		userapi.GET("/:eid", GetDynJsonObjectWithIdUser)
		userapi.POST("/store", DynStoreFileInDb)
	}

	router.Run(":3000")
}
