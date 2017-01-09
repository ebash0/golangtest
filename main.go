package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Site       []string
	SearchText string
}

type Response struct {
	FoundAtSite string
}

func main() {
	router := gin.Default()
	router.POST("checkText", func(c *gin.Context) {
		var jsonRequest Request
		var jsonResponse Response

		if c.BindJSON(&jsonRequest) == nil {
			for _, site := range jsonRequest.Site {

				response, err := http.Get(site)
				if err != nil {
					log.Print(err)
					continue
				}

				defer response.Body.Close()

				html, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Print(err)
					continue
				}

				if strings.Contains(string(html), jsonRequest.SearchText) {
					jsonResponse.FoundAtSite = site
					c.JSON(http.StatusOK, jsonResponse)
					return
				}
			}
			c.Data(http.StatusNoContent, gin.MIMEHTML, nil)
		}
	})

	router.Run(":8080")
}
