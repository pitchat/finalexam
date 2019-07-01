package customer

import (
	"log"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pitchat/finalexam/database"
	"net/http"
	"strconv"
)

//Update record in database
func (cu Customer) Update(conn *sql.DB) error {

	stmt, err := conn.Prepare("UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1;")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cu.ID, cu.Name, cu.Email, cu.Status)
	return err
}

//UpdateHandler gin api
func UpdateHandler(c *gin.Context) {

	cu := Customer{}
	if err := c.ShouldBindJSON(&cu); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	cu.ID = id

	err = database.Update(cu)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error":http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, cu)
}