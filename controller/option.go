package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/model"
	"net/http"
)

func GetOptions(c *gin.Context) {
	var options []*model.Option
	for k, v := range common.OptionMap {
		options = append(options, &model.Option{
			Key:   k,
			Value: common.Interface2String(v),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    options,
	})
	return
}

func UpdateOption(c *gin.Context) {
	var option model.Option
	err := json.NewDecoder(c.Request.Body).Decode(&option)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	model.UpdateOption(option.Key, option.Value)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}
