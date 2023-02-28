package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 匹配新增的增量更新

type AccessToken struct {
	Token string `json:"token"`
}

func restartBsc(c *gin.Context) {
	var accessToken AccessToken
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		Logger.Sugar().Error(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err = json.Unmarshal(jsonData, &accessToken)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if accessToken.Token != viper.GetString("token") {
		Logger.Sugar().Error(errors.New("invalid token: " + accessToken.Token))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "invalid token"})
		return
	}
	res := RestartDocker(viper.GetString("docker-name"))
	c.IndentedJSON(http.StatusOK, res)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(ginzap.Ginzap(Logger, time.RFC3339, false))

	r.POST("/restart_bsc", restartBsc)
	port := viper.GetInt("port")

	r.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
