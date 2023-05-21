package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"insta-clone/database"
	"insta-clone/internals/utils"
	"insta-clone/src/modules/photo/entities"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UpdatePhoto(ctx *gin.Context) {
	var db = database.GetDB()
	var photo entities.Photo
	var input entities.Photo
	err := db.First(&photo, "Id = ?", ctx.Param("photoId")).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	fileIn, err := ctx.FormFile("file")
	if err != nil {
		ctx.ShouldBindJSON(&input)

		db.Model(&photo).Updates(&input)

		ctx.JSON(http.StatusOK, gin.H{"data": photo})
		return
	}

	// save file to temp folder
	if err := ctx.SaveUploadedFile(fileIn, "temp/tempFile"+fileIn.Filename); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save file",
		})
		return
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", "temp/tempFile"+fileIn.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error create form file",
			"message": err.Error(),
		})
		return
	}

	file, err := os.Open("temp/tempFile" + fileIn.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error Open file",
			"message": err.Error(),
		})
		return
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error Copy",
			"message": err.Error(),
		})
		return
	}

	writer.Close()

	req, err := http.NewRequest("PATCH", "http://localhost:3000/image", bytes.NewReader(body.Bytes()))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)

	if rsp.StatusCode != http.StatusOK {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": fmt.Sprintf("Request failed with response code: %d", rsp.StatusCode),
		})
		// return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "io read all",
			"message": err.Error(),
		})
		return
	}

	ctx.ShouldBindJSON(&input)
	db.Model(&photo).Updates(&input)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "file sended",
		"photo":   photo,
	})

	file.Close()

	// delete temporary file
	path := "temp/tempFile" + fileIn.Filename
	defer utils.DeleteTempFile(path, ctx)
}

func GetAllPhoto(ctx *gin.Context) {
	var db = database.GetDB()

	var photo []entities.Photo

	err := db.Preload("Comments").Find(&photo).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "err get Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": photo})
}

func GetOnePhoto(ctx *gin.Context) {
	var db = database.GetDB()

	var photo []entities.Photo

	err := db.First(&photo, "Id = ?", ctx.Param("photoId")).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data one": photo})
}

func DeleteImage(ctx *gin.Context) {
	db := database.GetDB()
	PhotoDelete := entities.Photo{}
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
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	db := database.GetDB()

	// getting form post
	caption := ctx.PostForm("caption")
	title := ctx.PostForm("title")
	// getting file form user input
	fileIn, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No File is received",
			"err":     err.Error(),
		})
		return
	}
	// save file to temp folder
	if err := ctx.SaveUploadedFile(fileIn, "temp/tempFile"+fileIn.Filename); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save file",
		})
		return
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("userId", strconv.Itoa(int(userID)))
	fw, err := writer.CreateFormFile("file", "temp/tempFile"+fileIn.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error create form file",
			"message": err.Error(),
		})
		return
	}

	file, err := os.Open("temp/tempFile" + fileIn.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error Open file",
			"message": err.Error(),
		})
		return
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error Copy",
			"message": err.Error(),
		})
		return
	}

	writer.Close()

	req, err := http.NewRequest("POST", "http://localhost:3000/image", bytes.NewReader(body.Bytes()))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)

	if rsp.StatusCode != http.StatusOK {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": fmt.Sprintf("Request failed with response code: %d", rsp.StatusCode),
		})
		// return
	}

	resBody, err := io.ReadAll(rsp.Body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "io read all",
			"message": err.Error(),
		})
		return
	}

	type Image struct {
		ImageID  uint   `json:"image_id"`
		UserID   uint   `json:"user_id"`
		ImageUrl string `json:"image_url"`
	}

	var responseObj Image
	photo := entities.Photo{}

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

	ctx.JSON(http.StatusCreated, gin.H{
		"message":           "file sended",
		"image server resp": responseObj,
		"photo":             photo,
	})

	file.Close()

	// delete temporary file
	path := "temp/tempFile" + fileIn.Filename
	defer utils.DeleteTempFile(path, ctx)
}
