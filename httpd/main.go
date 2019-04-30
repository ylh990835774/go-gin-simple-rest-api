package main

import (
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {

	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	/*
		// use js send post request
		await fetch('/newsfeed', {
		    method: 'POST',
		    headers: {'content-type': 'application/json'},
			body: JSON.stringify({
				title: "Hello",
				post: "World!!!!!"
			})
		});
	*/

	r.Run()
}
