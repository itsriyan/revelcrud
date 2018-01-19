package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"revelcrud/app/models"
	"revelcrud/app/routes"

	"github.com/revel/revel"
)

type Employed struct {
	*revel.Controller
}
type trd struct {
	Err  string            `json:"err"`
	Data []models.Employed `json:"data"`
}
type trds struct {
	Err  string          `json:"err"`
	Data models.Employed `json:"data"`
}

func (c Employed) New() revel.Result {
	return c.Render()
}

func (c Employed) SaveEmployed() revel.Result {
	var err error
	id := c.Params.Get("id")
	name := c.Params.Get("nameemployed")
	email := c.Params.Get("email")
	phone := c.Params.Get("phone")
	address := c.Params.Get("address")
	data := models.Employed{
		Id:           id,
		NameEmployed: name,
		Email:        email,
		Phone:        phone,
		Address:      address,
	}
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}
	urls := "http://127.0.0.1:8082/api/v1/e/employed/"
	res, err := http.Post(urls, "application/x-www-form-urlencoded", bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		panic(res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var rd trds
	if err := json.Unmarshal(body, &rd); err != nil {
		panic(err)
	}
	if rd.Err != "" {
		panic(err)
	}
	return c.Redirect(routes.Employed.Index())
}
func (c Employed) Index() revel.Result {
	urls := "http://127.0.0.1:8082/api/v1/e/employed/"
	res, err := http.Get(urls)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		panic(res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	var rd trd
	if err := json.Unmarshal(data, &rd); err != nil {
		panic(err)
	}
	if rd.Err != "" {
		panic(rd.Err)
	}
	employeds := rd.Data
	return c.Render(employeds)
}

func (c Employed) Employed() revel.Result {
	var employed = models.Employed{}
	id := c.Params.Get("id")
	if id == "" {
		return c.Render(employed)
	}
	urls := "http://127.0.0.1:8082/api/v1/e/employed/" + id
	res, err := http.Get(urls)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		panic(res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	var rd trds
	if err := json.Unmarshal(data, &rd); err != nil {
		panic(err)
	}
	if rd.Err != "" {
		panic(rd.Err)
	}
	employed = rd.Data
	return c.Render(employed)
}

func (c Employed) EditEmployed() revel.Result {
	var err error
	id := c.Params.Get("id")
	name := c.Params.Get("nameemployed")
	email := c.Params.Get("email")
	phone := c.Params.Get("phone")
	address := c.Params.Get("address")
	data := models.Employed{
		Id:           id,
		NameEmployed: name,
		Email:        email,
		Phone:        phone,
		Address:      address,
	}
	fmt.Println(data)
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
	}
	urls := "http://127.0.0.1:8082/api/v1/e/employed/" + data.Id
	req, err := http.NewRequest("PUT", urls, bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		panic(res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var rd trds
	if err := json.Unmarshal(body, &rd); err != nil {
		panic(err)
	}
	if rd.Err != "" {
		panic(rd.Err)
	}
	return c.Redirect(routes.Employed.Index())
}
func (c Employed) DeleteEmployed() revel.Result {
	idStr := c.Params.Get("id")
	urls := "http://127.0.0.1:8082/api/v1/e/employed/" + idStr
	req, err := http.NewRequest("DELETE", urls, nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		panic(res.StatusCode)
	}
	defer res.Body.Close()

	return c.Redirect(routes.Employed.Index())
}
