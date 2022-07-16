package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"net/http"
	"sort"
	"strconv"
)

type ipItem struct {
	Ip    string `json:"name"`
	Count int    `json:"value"`
}

type urlItem struct {
	Url   string `json:"name"`
	Count int    `json:"value"`
}

type reqItem struct {
	Time  string `json:"name"`
	Count int    `json:"value"`
}

func GetIPs(c *gin.Context) {
	var ips []ipItem
	ctx := context.Background()
	rdb := common.RDB
	iter := rdb.Scan(ctx, 0, "statIP:*", 0).Iterator()
	for iter.Next(ctx) {
		value, err := rdb.Get(ctx, iter.Val()).Result()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}
		count, _ := strconv.Atoi(value)
		ips = append(ips, ipItem{
			Ip:    iter.Val()[7:],
			Count: count,
		})
	}
	sort.Slice(ips, func(i, j int) bool {
		return ips[i].Count > ips[j].Count
	})
	if len(ips) >= common.StatIPNum {
		ips = ips[:common.StatIPNum]
	}

	if err := iter.Err(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    ips,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    ips,
	})
}

func GetURLs(c *gin.Context) {
	var urls []urlItem
	ctx := context.Background()
	rdb := common.RDB
	iter := rdb.Scan(ctx, 0, "statURL:*", 0).Iterator()
	for iter.Next(ctx) {
		value, err := rdb.Get(ctx, iter.Val()).Result()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}
		count, _ := strconv.Atoi(value)
		urls = append(urls, urlItem{
			Url:   iter.Val()[8:],
			Count: count,
		})
	}
	sort.Slice(urls, func(i, j int) bool {
		return urls[i].Count > urls[j].Count
	})
	if len(urls) >= common.StatURLNum {
		urls = urls[:common.StatIPNum]
	}

	if err := iter.Err(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    urls,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    urls,
	})
}

func GetReqs(c *gin.Context) {
	var reqs []reqItem
	ctx := context.Background()
	rdb := common.RDB
	iter := rdb.Scan(ctx, 0, "statReq:*", 0).Iterator()
	for iter.Next(ctx) {
		value, err := rdb.Get(ctx, iter.Val()).Result()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}
		count, _ := strconv.Atoi(value)
		reqs = append(reqs, reqItem{
			Time:  iter.Val()[8:],
			Count: count,
		})
	}

	if err := iter.Err(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    reqs,
	})
}
