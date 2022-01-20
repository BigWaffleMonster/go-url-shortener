package handler

import (
	"fmt"
	"net/http"

	"log"

	urlshortener "gihtub.com/BigWaffleMonster/go-url-shortener"
	"gihtub.com/BigWaffleMonster/go-url-shortener/pkg/helperFunctions"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func shortenUrl(c *gin.Context, db *sqlx.DB) {
  var url urlshortener.FULL_URL 

  if err := c.BindJSON(&url); err !=nil {
    return
  }
  
  var newUrl = "http://localhost:8000/" + helperfunctions.GenerateRandomString()

  tx := db.MustBegin()
  tx.MustExec("INSERT INTO urls (full_url, short_url) VALUES ($1, $2)", url.FullUrl, newUrl)
  tx.Commit()

  c.IndentedJSON(http.StatusCreated, newUrl)
}

func getUrl(c *gin.Context, db *sqlx.DB) {
  var url string = c.Param("url")
  var query string = "http://localhost:8000/" + url

  var full_url string

  err := db.Get(&full_url,fmt.Sprintf("SELECT full_url FROM urls where short_url='%s'", query))

  if err != nil {
    log.Fatalf("Error with get: %s", err.Error()) 
  }

  c.String(http.StatusOK, "Full Url: %s", full_url)
}
