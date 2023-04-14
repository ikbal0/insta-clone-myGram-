package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"insta-clone/database"
	"insta-clone/helpers"
	"insta-clone/models"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func DeleteImage(ctx *gin.Context) {
	db := database.GetDB()
	PhotoDelete := models.Photo{}
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	err := db.First(&PhotoDelete, "Id = ?", photoId).Error

	imageId := strconv.Itoa(int(PhotoDelete.ImageID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:3000/image/"+imageId, nil)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request 4",
			"message": err.Error(),
		})
		return
	}

	rsp, err := client.Do(req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "get rsp",
			"message": err.Error(),
		})
		return
	}

	defer rsp.Body.Close()

	resBody, err := io.ReadAll(rsp.Body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "io read all",
			"message": err.Error(),
		})
		return
	}

	errDelete := db.Debug().Delete(&PhotoDelete).Error

	if errDelete != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "delete fail",
			"message": err.Error(),
		})
		return
	}
	// fmt.Println("body:", string(resBody))

	if rsp.StatusCode != http.StatusOK {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Status code",
			"message": fmt.Sprintf("Request failed with response code: %d", rsp.StatusCode),
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "delete success!",
		"resp":    string(resBody),
	})
}

func UploadFile(ctx *gin.Context) {
	db := database.GetDB()
	fileIn, err := ctx.FormFile("file")
	caption := ctx.PostForm("caption")
	title := ctx.PostForm("title")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request 1",
			"message": err.Error(),
		})
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	timeStamp := helpers.MakeTimeStamp()
	newName := strconv.Itoa(int(timeStamp)) + fileIn.Filename

	writer.CreateFormFile("file", fileIn.Filename)
	writer.WriteField("userId", "1")
	writer.WriteField("name", newName)

	f, err := os.CreateTemp("", fileIn.Filename)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request 2",
			"message": err.Error(),
		})
	}

	writer.Close()

	defer f.Close()

	req, err := http.NewRequest("POST", "http://localhost:3000/image", bytes.NewReader(body.Bytes()))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request 4",
			"message": err.Error(),
		})
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	rsp, err := client.Do(req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "get rsp",
			"message": err.Error(),
		})
		return
	}

	defer rsp.Body.Close()

	resBody, err := io.ReadAll(rsp.Body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "io read all",
			"message": err.Error(),
		})
		return
	}

	fmt.Println("body:", string(resBody))

	if rsp.StatusCode != http.StatusOK {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request 5",
			"message": fmt.Sprintf("Request failed with response code: %d", rsp.StatusCode),
		})
		return
	}

	type Image struct {
		ImageID  uint   `json:"image_id"`
		UserID   uint   `json:"user_id"`
		ImageUrl string `json:"image_url"`
	}

	var responseObj Image
	photo := models.Photo{}

	json.Unmarshal(resBody, &responseObj)

	photo.Caption = caption
	photo.Title = title
	photo.PhotoUrl = responseObj.ImageUrl
	photo.UserID = responseObj.UserID
	photo.ImageID = responseObj.ImageID

	errCreate := db.Debug().Create(&photo).Error

	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"message":           "file sended",
		"image server resp": responseObj,
		"photo":             photo,
	})
}
