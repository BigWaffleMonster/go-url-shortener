package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRoutes(db *sqlx.DB) *gin.Engine {
  router := gin.Default()

  router.POST("/", func (c *gin.Context) {
    shortenUrl(c, db)
  })

  router.GET("/:url", func (c *gin.Context) {
    getUrl(c, db)
  })

  return router
}
