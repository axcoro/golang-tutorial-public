package kvs

import (
	"encoding/json"
	"fmt"
	"gosanluis/kvs"
	"io/ioutil"
	"net/http"

	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/go-meli-toolkit/gingonic/mlhandlers"
	"github.com/mercadolibre/go-meli-toolkit/gokvsclient"
)

var router *gin.Engine

type data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	router = mlhandlers.DefaultMeliRouter()
	router.POST("/save", saveKVS)
	router.GET("/get/:key", getKVS)
	router.Run(":8080")
}

func getKVS(c *gin.Context) {
	key := c.Param("key")
	item, err := kvs.Get(key)
	if err != nil {
		panic(err)
	}
	logger.Infof("item: %v", item)
	var value data
	err = item.GetValue(&value)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, fmt.Sprintf("key: %v value: %v", key, value.Value))
}

func saveKVS(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	var dataSave data
	err = json.Unmarshal(body, &dataSave)
	if err != nil {
		panic(err)
	}
	item := gokvsclient.MakeItem(dataSave.Key, dataSave)
	if err := kvs.Save(item); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, fmt.Sprintf("key: %v value: %v", dataSave.Key, dataSave.Value))
}
