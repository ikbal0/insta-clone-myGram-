package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
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
)

func UpdatePhoto(ctx *gin.Context) {
	var db = database.GetDB()
	var photo entities.Photo
	var input entities.Photo
	photoId := ctx.Param("photoId")
	err := db.First(&photo, "Id = ?", photoId).Error

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

	imageId := strconv.Itoa(int(photo.ImageID))
	req, err := http.NewRequest("PATCH", "http://localhost:3000/image/"+imageId, bytes.NewReader(body.Bytes()))

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

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "file sended",
		"photo":   photo,
	})

	file.Close()

	// delete temporary file
	path := "temp/tempFile" + fileIn.Filename
	defer utils.DeleteTempFile(path)
}

func (h httpHandlerImpl) GetAllPhoto(ctx *gin.Context) {
	photos, err := h.PhotoService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "err get Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": photos})
}

func (h httpHandlerImpl) GetOnePhoto(ctx *gin.Context) {
	getId := ctx.Param("photoId")
	id, errConv := strconv.Atoi(getId)

	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errConv.Error()})
	}

	photo, err := h.PhotoService.GetByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func (h httpHandlerImpl) UploadFile(ctx *gin.Context) {
	userID := utils.GetUserID(ctx)

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
	if err := ctx.SaveUploadedFile(fileIn, "temp/"+fileIn.Filename); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save file",
		})
		return
	}

	photo, err := fileServerPostReq(userID, caption, title, fileIn)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when send file to file server",
			"error":   err.Error(),
		})
		return
	}

	savedPhoto, savedErr := h.PhotoService.Input(photo)

	if savedErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "file sended",
		"photo":   savedPhoto,
	})
}

func fileServerPostReq(userID uint, caption string, title string, fileIn *multipart.FileHeader) (entities.Photo, error) {
	photo := entities.Photo{}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("userId", strconv.Itoa(int(userID)))

	fw, err := writer.CreateFormFile("file", "temp/"+fileIn.Filename)
	if err != nil {
		return photo, err
	}

	file, err := os.Open("temp/" + fileIn.Filename)
	if err != nil {
		return photo, err
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		return photo, err
	}

	writer.Close()

	// send save request to image server
	newBody := bytes.NewReader(body.Bytes())
	req, err := http.NewRequest("POST", "http://localhost:3000/image", newBody)

	if err != nil {
		return photo, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)

	if rsp.StatusCode != http.StatusOK {
		statusCode := strconv.Itoa(rsp.StatusCode)
		message := "Request failed with response code: " + statusCode
		err := errors.New(message)
		return photo, err
	}

	resBody, err := io.ReadAll(rsp.Body)

	if err != nil {
		return photo, err
	}

	type Image struct {
		ImageID  uint   `json:"image_id"`
		UserID   uint   `json:"user_id"`
		ImageUrl string `json:"image_url"`
	}

	var responseObj Image

	json.Unmarshal(resBody, &responseObj)

	file.Close()

	// delete temporary file
	path := "temp/" + fileIn.Filename
	defer utils.DeleteTempFile(path)

	photo.Caption = caption
	photo.Title = title
	photo.PhotoUrl = responseObj.ImageUrl
	photo.UserID = responseObj.UserID
	photo.ImageID = responseObj.ImageID

	return photo, nil
}
