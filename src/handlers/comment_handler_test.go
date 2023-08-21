package handlers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_httpHandlerImpl_PostComment(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		h    httpHandlerImpl
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.PostComment(tt.args.ctx)
		})
	}
}
