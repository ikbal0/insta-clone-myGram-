package controllers

import (
	"bytes"
	"encoding/json"
	"insta-clone/helpers"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Chart struct {
	Name  string `json:"name" form:"name"`
	Icon  string `json:"icon" form:"icon"`
	Price int    `json:"price" form:"price"`
	Total int    `json:"total" form:"total"`
	Type  string `json:"type" form:"type"`
	Qty   int    `json:"qty" form:"qty"`
}

func PostTest(ctx *gin.Context) {

	contentType := helpers.GetContentType(ctx)
	Chart := Chart{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&Chart)
	} else {
		ctx.ShouldBind(&Chart)
	}

	data := map[string]interface{}{
		"name":  Chart.Name,
		"icon":  Chart.Icon,
		"price": Chart.Price,
		"total": Chart.Total,
		"type":  Chart.Type,
		"qty":   Chart.Qty,
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}

	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:3000/create/cart", bytes.NewBuffer(requestJson))

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("body:", string(body))

	type Response struct {
		Status bool `json:"status"`
	}

	var responseObject Response

	json.Unmarshal(body, &responseObject)

	if !responseObject.Status {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"body": responseObject,
	})
	// fmt.Println("chart bind:", Chart)
	// fmt.Println("json marshal:", requestJson)
}

func GetAllPhoto(ctx *gin.Context) {
	res, err := http.Get("https://server-image.up.railway.app/view")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	// fmt.Println(res.Body)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	defer res.Body.Close()

	// sb := string(body)

	type File struct {
		ID   string `json:"_id"`
		Name string `json:"name"`
	}

	type Product struct {
		ID   string `json:"_id"`
		Name string `json:"name"`
	}

	type DataPool struct {
		Files    []File    `json:"file"`
		Products []Product `json:"product"`
	}

	type Response struct {
		Status bool     `json:"status"`
		Data   DataPool `json:"data"`
	}

	var responseObject Response

	json.Unmarshal(body, &responseObject)

	ctx.JSON(http.StatusCreated, gin.H{
		"file": responseObject.Data.Files,
	})
}
