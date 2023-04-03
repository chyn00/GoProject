package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"untitled/model"
)

func GetItem(c *gin.Context) {

	// 임시데이터
	items := []model.Item{
		{ID: 1, Name: "Item 1", Description: "This is item 1", Price: 1000},
		{ID: 2, Name: "Item 2", Description: "This is item 2", Price: 2000},
		{ID: 3, Name: "Item 3", Description: "This is item 3", Price: 3000},
	}

	// JSON 형태의 객체 반환
	c.JSON(http.StatusOK, items)
}
