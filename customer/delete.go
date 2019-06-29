package customer

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pitchat/finalexam/database"
	"net/http"
	"strconv"
)

//Delete customer
func (cu Customer) Delete(conn *sql.DB) error {

	stmt, err := conn.Prepare("DELETE FROM customers WHERE id=$1;")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cu.ID)
	return err
}

//DeleteByIDHandler gin api
func DeleteByIDHandler(c *gin.Context) {

	cu := Customer{}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cu.ID = id

	err = database.Delete(cu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
