package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
	})

	ctx := context.Background()

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
		d, e := c.GetRawData()

		if e != nil {
			fmt.Println(e)
		} else {
			data := strings.Split(string(d), ":")
			fingerprint := data[0]
			action := data[1]
			imageId := data[2]
			ts := data[3]

			fmt.Printf("%s %s watching %s at %s\n", fingerprint, action, imageId, ts)
			if strings.Compare(action, "started") == 0 {
				rdb.HSet(ctx, fingerprint, imageId, ts).Result()
				rdb.Expire(ctx, fingerprint, 60*time.Second)
			} else if strings.Compare(action, "ended") == 0 {
				t, e := rdb.HGet(ctx, fingerprint, imageId).Result()

				startts, _ := strconv.Atoi(t)
				endts, _ := strconv.Atoi(ts)

				if e != nil {
					if e == redis.Nil {
						fmt.Printf("Start time not available for %s", imageId)
					} else {
						fmt.Println(e)
					}
				} else {
					fmt.Printf("Started: %d, Ended: %d, Total: %d\n", startts, endts, endts-startts)
					rdb.HDel(ctx, fingerprint, imageId)
				}
			}
		}
	})

	router.Run(":8080")
	// router.RunTLS(":8080", "./ssl/server.pem", "./ssl/server.key")
}
