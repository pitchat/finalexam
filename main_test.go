package main

import (
	"fmt"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"github.com/pitchat/finalexam/database"
	"github.com/pitchat/finalexam/customer"
)

func TestHandler(t *testing.T) {
	database.InitDB()
	router := setupRouter()
	fmt.Println("Begin test")
	cu := customer.Customer{}

	//Case 1 Create
	c1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/customers", bytes.NewBufferString(`{"name":"anuchito", "email":"anuchito@imail.com","status": "active"}`))
	req1.Header.Set("Authorization", "token2019")
	router.ServeHTTP(c1, req1)
	if !(c1.Code ==  201 ||  c1.Code ==  202) {
		t.Errorf("create customer http response code is incorrect(201,202). result value: %d.", c1.Code)
	}
	json.Unmarshal([]byte(c1.Body.String()), &cu)
	if cu.ID ==  0 {
		t.Errorf("create customer ID is zero.")
	}
	if cu.Name !=  "anuchito" {
		t.Errorf("customer Name is incorrect. expect value: %q, result value: %q.", "anuchito", cu.Name)
	}
	if cu.Email !=  "anuchito@imail.com" {
		t.Errorf("customer Email is incorrect. expect value: %q, result value: %q.", "active", cu.Email)
	}
	if cu.Status !=  "active" {
		t.Errorf("customer Status is incorrect. expect value: %q, result value: %q.", "active", cu.Status)
	}
	
	//Case 2 Get by id
	id := cu.ID
	var url = fmt.Sprintf("/customers/%d",id)
	c2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", url, nil)
	req2.Header.Set("Authorization", "token2019")
	router.ServeHTTP(c2, req2)
	if c2.Code !=  200 {
		t.Errorf("http response code is incorrect(200). result value: %d.", c2.Code)
	}
	json.Unmarshal([]byte(c2.Body.String()), &cu)
	if cu.ID != int64(id) {
		t.Errorf("customer ID is incorrect. expect value: %d, result value: %d.", id, cu.ID)
	}
	if cu.Name == "" {
		t.Errorf("customer name is incorrect. result value is empty")
	}
	if cu.Email == "" {
		t.Errorf("customer email is incorrect. result value is empty")
	}


	//Case 3 Get all
	list := []customer.Customer{}
	c3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("GET", "/customers", nil)
	req3.Header.Set("Authorization", "token2019")
	router.ServeHTTP(c3, req3)
	if c3.Code !=  200 {
		t.Errorf("http response code is incorrect(200). result value: %d.", c3.Code)
	}
	json.Unmarshal([]byte(c3.Body.String()), &list)
	if len(list) == 0 {
		t.Errorf("customer list is incorrect. expect length more than 0")
	}
	

	//Case 4 Update
	var upBody = `{"id": `+fmt.Sprintf("%d",id)+`,"name": "nong","email": "nong@imail.com","status": "inactive" }`
	c4 := httptest.NewRecorder()
	req4, _ := http.NewRequest("PUT", url, bytes.NewBufferString(upBody))
	req4.Header.Set("Authorization", "token2019")
	router.ServeHTTP(c4, req4)
	if c4.Code !=  200 {
		t.Errorf("update customer http response code is incorrect(200). result value: %d.", c4.Code)
	}
	json.Unmarshal([]byte(c4.Body.String()), &cu)
	if cu.ID != int64(id) {
		t.Errorf("update customer ID is incorrect. expect value: %d, result value: %d.", id, cu.ID)
	}
	if cu.Name != "nong" {
		t.Errorf("update customer name is incorrect. expect value: %q, result value is %q", "nong",cu.Status)
	}
	if cu.Email != "nong@imail.com" {
		t.Errorf("update customer email is incorrect. expect value: %q, result value is %q", "nong@imail.com",cu.Status)
	}
	if cu.Status != "inactive" {
		t.Errorf("update customer status is incorrect. expect value: %q, result value is %q", "inactive",cu.Status)
	}
	

	//Case 5 delete
	c5 := httptest.NewRecorder()
	req5, _ := http.NewRequest("DELETE", url, nil)
	req5.Header.Set("Authorization", "token2019")
	router.ServeHTTP(c5, req5)
	if c5.Code !=  200 {
		t.Errorf("http response code is incorrect(200). result value: %d.", c5.Code)
	}
	var checkDel string = `{"message":"customer deleted"}`
	if c5.Body.String() != checkDel {
		t.Errorf("delete customer response is incorrect. expect value: %q, result value: %q.", checkDel, c5.Body.String())
	}

	//Case 6 Authen fail
	c6 := httptest.NewRecorder()
	req6, _ := http.NewRequest("GET", "/customers", nil)
	req6.Header.Set("Authorization", "token2019wrong_token")
	router.ServeHTTP(c6, req6)
	if c6.Code !=  401 {
		t.Errorf("authen fail http response code is incorrect(401). result value: %d.", c6.Code)
	}
}

