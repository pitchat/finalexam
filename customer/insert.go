package customer

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pitchat/finalexam/database"
	"net/http"
)

//Insert customer
func (cu Customer) Insert(conn *sql.DB) (database.DataLayer, error) {

	row := conn.QueryRow("INSERT INTO customers (name, email, status) VALUES ($1, $2, $3) RETURNING id", cu.Name, cu.Email, cu.Status)
	err := row.Scan(&cu.ID)
	return database.IConv(cu), err
}

//CreateHandler gin api
func CreateHandler(c *gin.Context) {

	c1 := Customer{}
	if err := c.ShouldBindJSON(&c1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c2, err := database.Insert(c1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, c2)
}
