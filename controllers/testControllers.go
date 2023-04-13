package controllers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	fileIn, err := ctx.FormFile("file")
	name := ctx.PostForm("name")

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

	writer.CreateFormFile("file", fileIn.Filename)
	writer.WriteField("userId", "1")
	writer.WriteField("name", name)

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

	// fmt.Println("file:", f)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request 6",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "file sended",
		"resp":    string(resBody),
	})
}
