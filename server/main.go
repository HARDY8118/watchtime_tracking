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
	// router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/list/images", CORSMiddleware(), func(c *gin.Context) {
		listFile, err := os.Open("./assets/imagelist.json")
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"error": "failed to load file",
			})
		} else {

			var imagelist []struct {
				Title        string `json:"title"`
				File         string `json:"filename"`
				Category     string `json:"category"`
				Photographer string `json:"photographer"`
				Source       string `json:"source"`
			}
			if err = json.NewDecoder(listFile).Decode(&imagelist); err != nil {
				fmt.Println(err)
				c.JSON(500, gin.H{
					"error": "failed to load data",
				})
			} else {
				c.JSON(200, imagelist)
			}
		}
	})

	router.Use(static.Serve("/videos", static.LocalFile("./assets/videos", false)))
	router.Use(static.Serve("/images", static.LocalFile("./assets/images", false)))

	router.POST("/switch", func(c *gin.Context) {
		// fmt.Println(c.Request.GetBody())
		// fmt.Println(c.Request.Body.)
		d, e := c.GetRawData()

		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(string(d))
		}

		fmt.Println()
	})

	router.Run(":8080")
	// router.RunTLS(":8080", "./ssl/server.pem", "./ssl/server.key")
}
