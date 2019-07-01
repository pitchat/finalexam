package customer

import (
	"log"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pitchat/finalexam/database"
	"net/http"
	"strconv"
)

//GetAll get all data in table
func (cu Customer) GetAll(conn *sql.DB) ([]database.DataLayer, error) {

	dd := []database.DataLayer{}
	rows, err := conn.Query("SELECT id, name, email, status FROM customers")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var cu Customer
		if err := rows.Scan(&cu.ID, &cu.Name, &cu.Email, &cu.Status); err != nil {
			return nil, err
		}
		dd = append(dd, database.IConv(cu))
	}
	return dd, err
}

//GetHandler gin api
func GetHandler(c *gin.Context) {
	cu := Customer{}
	cu2, err := database.GetAll(cu)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error":http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, cu2)
}

//GetByKey get record by key
func (cu Customer) GetByKey(conn *sql.DB) (database.DataLayer, error) {

	row := conn.QueryRow("SELECT id, name, email, status FROM customers where id = $1", cu.ID)
	err := row.Scan(&cu.ID, &cu.Name, &cu.Email, &cu.Status)
	if err != nil {
		return cu, err
	}
	return database.IConv(cu), err
}

//GetByIDHandler for retrive data by ID
func GetByIDHandler(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error":http.StatusText(http.StatusBadRequest)})
		return
	}
	cu1 := Customer{}
	cu1.ID = id

	cu2, err :=  database.GetByKey(cu1)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error":http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, cu2)
}
