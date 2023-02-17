package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/weather", func(c *gin.Context) {
		city := c.Query("city") //获取URL中的查询字符串参数
		if city == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "city is required"})
			return
		}
		//调用第三方的天气API，这里使用的是https://www.tianqiapi.com/
		resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=e61f5aacf4fbb8f77f3fd460d6f6b34b")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		//将第三ss方的响应转发给客户端
		c.DataFromReader(http.StatusOK, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	})
	r.Run()

}
