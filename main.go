package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Animal struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()

	// GET /person 엔드포인트
	router.GET("/person", getPerson)

	// GET /animal 엔드포인트
	router.GET("/animal", getAnimal)

	router.Run(":8080")
}

func getPerson(c *gin.Context) {
	// JSON 형태의 객체 생성
	person := Person{
		Id:   1,
		Name: "John",
	}

	// JSON 형태의 객체 반환
	c.JSON(200, person)
}

func getAnimal(c *gin.Context) {
	// JSON 형태의 객체 생성
	animal := Animal{
		Id:   1,
		Name: "Dog",
	}

	// JSON 형태의 객체 반환
	c.JSON(200, animal)
}
