package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"strings"
	"time"
)

func statHelper(ip string, url string) {
	ctx := context.Background()
	rdb := common.RDB
	ipKey := "statIP:" + ip
	urlKey := "statURL:" + url
	t := time.Now()
	//_, offset := t.Local().Zone()
	reqKey := "statReq:" + t.In(time.Local).Format("2006-01-02 15")
	rdb.Incr(ctx, ipKey)
	rdb.Expire(ctx, ipKey, time.Duration(common.StatCacheTimeout)*time.Hour)
	if !strings.HasPrefix(urlKey, "statURL:/public") {
		rdb.Incr(ctx, urlKey)
		rdb.Expire(ctx, urlKey, time.Duration(common.StatCacheTimeout)*time.Hour)
	}
	rdb.Incr(ctx, reqKey)
	rdb.Expire(ctx, reqKey, time.Duration(common.StatReqTimeout)*time.Hour*24)
}

func AllStat() func(c *gin.Context) {
	return func(c *gin.Context) {
		if common.StatEnabled {
			go statHelper(c.ClientIP(), c.Request.URL.String())
		}
		c.Next()
	}
}
