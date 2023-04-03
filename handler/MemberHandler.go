package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"untitled/model"
)

func GetMember(c *gin.Context) {

	// 임시 데이터
	member := model.Member{
		ID:   1,
		Name: "John",
	}

	c.JSON(http.StatusOK, member)
}
