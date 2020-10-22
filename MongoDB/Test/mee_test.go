package Test

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/kriangkrai/Mee/MongoDB/Models"
	"github.com/kriangkrai/Mee/MongoDB/Router"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGet(t *testing.T) {

	r, _ := Router.SetupRouter()

	req := httptest.NewRequest("GET", "/get", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	response := res.Result()
	body, _ := ioutil.ReadAll(response.Body)
	result := []Models.DataModel{}

	err := json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	expect := Models.DataModel{
		ID:     primitive.NilObjectID,
		Device: "Sensor3",
		Date:   "2020-10-22 09:39:37.1459882 +0700 +07",
		Value:  "30.6",
	}

	assert.Equal(t, expect.Device, result[0].Device)
	assert.Equal(t, expect.Value, result[0].Value)
	assert.Equal(t, expect.Date, result[0].Date)

}

func TestAdd(t *testing.T) {

	r, _ := Router.SetupRouter()

	data := `{
				"device" : "Sensor1",
				"value" : "29.9"
			}`

	req := httptest.NewRequest("POST", "/insert", strings.NewReader(data))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	response := res.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var result string

	err := json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	expect := "Insert Success"

	assert.Equal(t, expect, result)

}

func TestUpdate(t *testing.T) {

	r, _ := Router.SetupRouter()

	data := `{
				"device" : "Sensor1",
				"value" : "20.0"
			}`

	req := httptest.NewRequest("PUT", "/update", strings.NewReader(data))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	response := res.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var result string

	err := json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	expect := "Update Success"

	assert.Equal(t, expect, result)

}

func TestDelete(t *testing.T) {

	r, _ := Router.SetupRouter()
	//Delete

	req := httptest.NewRequest("DELETE", "/delete/Sensor1", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	response := res.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var result string

	err := json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	expect := "Delete Success"

	assert.Equal(t, expect, result)

}
