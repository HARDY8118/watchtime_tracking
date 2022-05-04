package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// https://stackoverflow.com/a/29439630/8411160
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.New()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/list/videos", CORSMiddleware(), func(c *gin.Context) {
		listFile, err := os.Open("./assets/videolist.json")
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"error": "failed to load file",
			})
		} else {

			var videolist []struct {
				Videographer string `json:"videographer"`
				File         string `json:"file"`
				Title        string `json:"title"`
			}
			if err = json.NewDecoder(listFile).Decode(&videolist); err != nil {
				fmt.Println(err)
				c.JSON(500, gin.H{
					"error": "failed to load data",
				})
			} else {
				c.JSON(200, videolist)
			}
		}
	})

	router.Use(static.Serve("/videos", static.LocalFile("./assets/videos", false)))
	router.Use(static.Serve("/images", static.LocalFile("./assets/images", false)))

	router.Run(":8080")
}
