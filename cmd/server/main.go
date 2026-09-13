package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.New()
  r.GET("/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
  r.POST("/commitments", func(c *gin.Context) { c.JSON(http.StatusNotImplemented, gin.H{"error": "commitment workflow pending"}) })
  _ = r.Run(":8080")
}
