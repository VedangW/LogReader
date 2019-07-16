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
		adminapi.POST("/store", DynStoreFileInDb)
		adminapi.GET("/findid/:eid", GetDynJsonObjectWithIdAdmin)
		adminapi.GET("/finduser/:uid", GetDynJSONObjectWithUserIDAdmin)
		adminapi.GET("/findvalid/:ov", GetDynJSONObjectWithOrderValidAdmin)
	}

	userapi := router.Group("/user")
	{
		userapi.POST("/", CreateDynJsonObject)
		userapi.POST("/store", DynStoreFileInDb)
		userapi.GET("/findid/:eid", GetDynJsonObjectWithIdUser)
		userapi.GET("/finduser/:uid", GetDynJSONObjectWithUserIDUser)
		userapi.GET("/findvalid/:ov", GetDynJSONObjectWithOrderValidUser)
	}

	router.Run(":3000")
}
