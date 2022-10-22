package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	externalUrl := os.Getenv("EXTERNAL_URL")
	if externalUrl == "" {
		fmt.Println("EXTERNAL_URL not set")
		os.Exit(1)
	}

	r := gin.Default()

	cssHandler := func(c *gin.Context) {
		client := &http.Client{}

		originHost := "https://fonts.googleapis.com"

		req, err := http.NewRequest("GET", originHost+c.Request.URL.Path+"?"+c.Request.URL.RawQuery, nil)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		req.Header.Set("User-Agent", c.Request.UserAgent())

		res, err := client.Do(req)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		replaced := strings.ReplaceAll(string(body), "https://fonts.gstatic.com", externalUrl+"/fonts")

		c.Header("Cache-Control", res.Header.Get("Cache-Control"))

		c.Data(res.StatusCode, res.Header.Get("Content-Type"), []byte(replaced))
	}

	fontsHandler := func(c *gin.Context) {
		client := &http.Client{}

		originHost := "https://fonts.gstatic.com"

		originPath := c.Param("path")

		req, err := http.NewRequest("GET", originHost+"/"+originPath+"?"+c.Request.URL.RawQuery, nil)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		req.Header.Set("User-Agent", c.Request.UserAgent())

		res, err := client.Do(req)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		defer res.Body.Close()

		c.DataFromReader(res.StatusCode, res.ContentLength, res.Header.Get("Content-Type"), res.Body, map[string]string{
			"Cache-Control": res.Header.Get("Cache-Control"),
		})
	}

	r.GET("/css2", cssHandler)
	r.GET("/css", cssHandler)

	r.GET("/fonts/*path", fontsHandler)

	r.Run()
}
