package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/pitchat/finalexam/database"
	"github.com/pitchat/finalexam/customer"
)

func main() {
	database.InitDB()
	defer database.Close()
	r := setupRouter()
	r.Run(":2019") //listen and serve on 0.0.0.0:2019
}

func setupRouter() *gin.Engine{
	r := gin.Default()

	r.Use(authMiddleware)

	api := r.Group("/")
	api.GET("/customers", customer.GetHandler)
	api.GET("/customers/:id", customer.GetByIDHandler)
	api.POST("/customers", customer.CreateHandler)
	api.PUT("/customers/:id", customer.UpdateHandler)
	api.DELETE("/customers/:id", customer.DeleteByIDHandler)	
	return r
}

func authMiddleware(c *gin.Context){
	token := c.GetHeader("Authorization")
   if token != "token2019"{
	   c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
	   c.Abort()
	   return
   }
   c.Next()
}
