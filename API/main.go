package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = make(map[string]Item)

func getItems(c *gin.Context) {
	values := make([]Item, 0, len(items))
	for _, v := range items {
		values = append(values, v)
	}
	c.JSON(http.StatusOK, values)
}

func getItem(c *gin.Context) {
	id := c.Param("id")
	if item, exists := items[id]; exists {
		c.JSON(http.StatusOK, item)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	}
}

func createItem(c *gin.Context) {
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	items[item.ID] = item
	c.JSON(http.StatusCreated, item)
}

func updateItem(c *gin.Context) {
	id := c.Param("id")
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = id
	items[id] = item
	c.JSON(http.StatusOK, item)
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")
	if _, exists := items[id]; exists {
		delete(items, id)
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	}
}

func main() {
	r := gin.Default()
	r.GET("/items", getItems)
	r.GET("/items/:id", getItem)
	r.POST("/items", createItem)
	r.PUT("/items/:id", updateItem)
	r.DELETE("/items/:id", deleteItem)
	r.Run(":8080")
}
