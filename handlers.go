package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/gin-gonic/gin"
)

func GetDynJsonObjectWithIdAdmin(c *gin.Context) {
	id := c.Params.ByName("eid")

	var dynentry DynJsonEntry
	var result map[string]interface{}

	if err := Db.Where("id = ?", id).First(&dynentry).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		_ = json.Unmarshal([]byte(dynentry.Extra), &result)
		jsonout := ResultToJson("admin", result, dynentry)

		c.JSON(200, jsonout)
	}
}

func GetDynJsonObjectWithIdUser(c *gin.Context) {
	id := c.Params.ByName("eid")

	var dynentry DynJsonEntry
	var result map[string]interface{}

	if err := Db.Where("id = ?", id).First(&dynentry).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		_ = json.Unmarshal([]byte(dynentry.Extra), &result)
		jsonout := ResultToJson("user", result, dynentry)

		c.JSON(200, jsonout)
	}
}

func CreateDynJsonObject(c *gin.Context) {
	bytesVal, _ := ioutil.ReadAll(c.Request.Body)

	var jsonObj map[string]interface{}
	json.Unmarshal(bytesVal, &jsonObj)

	dynEntry := JsonToDynObj(jsonObj)

	Db.Create(&dynEntry)

	c.JSON(200, dynEntry)
}

func DynStoreFileInDb(c *gin.Context) {
	var filepath FilePath
	c.BindJSON(&filepath)

	Db.Create(&filepath)

	pathToReadFile := "/Users/vedanganandwaradpande/"
	pathToReadFile += filepath.FPath

	data, err := ioutil.ReadFile(pathToReadFile)

	if err != nil {
		fmt.Print(err)
	}

	var jsonInt []map[string]interface{}
	_ = json.Unmarshal(data, &jsonInt)

	var wg sync.WaitGroup

	for i := 0; i < len(jsonInt); i++ {

		wg.Add(1)
		go func(jsonMap map[string]interface{}) {
			var dynTemp DynJsonEntry
			dynTemp = JsonToDynObj(jsonMap)
			Db.Create(&dynTemp)
			wg.Done()
		}(jsonInt[i])
	}

	wg.Wait()

	c.JSON(200, "Created.")
}
