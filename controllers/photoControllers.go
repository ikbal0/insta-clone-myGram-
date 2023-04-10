package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status bool     `json:"status"`
	Data   DataPool `json:"data"`
}

type DataPool struct {
	Files    []File    `json:"file"`
	Products []Product `json:"product"`
}

type File struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type Product struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
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

	var responseObject Response

	json.Unmarshal(body, &responseObject)

	ctx.JSON(http.StatusCreated, gin.H{
		"file": responseObject.Data.Files,
	})
}
