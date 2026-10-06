package Controllers

import (
	"main/Services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MainStruct struct {
	ProductService Services.ProductService
}

func MainStructAPI(p Services.ProductService) MainStruct {
	return MainStruct{ProductService: p}
}

func (m *MainStruct) MainIndex(c *gin.Context) {

	walletStr := c.Param("walletId")
	walletId64, err := strconv.ParseUint(walletStr, 10, 64)
	if err != nil {
		// Если клиент передал не число, возвращаем ошибку 400 Bad Request
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат walletId"})
		return
	}

	var walletId uint = uint(walletId64)

	product := m.ProductService.Find(walletId)

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (m *MainStruct) MainPost(c *gin.Context) {

	valletStr := c.PostForm("valletId")
	valletId64, err := strconv.ParseUint(valletStr, 10, 64)
	if err != nil {
		// Если клиент передал не число, возвращаем ошибку 400 Bad Request
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат valletId"})
		return
	}

	operationType := c.PostForm("operationType")

	if operationType != "deposit" && operationType != "withdraw" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат operationType"})
		return
	}

	amountStr := c.PostForm("amount")
	amount64, err := strconv.ParseUint(amountStr, 10, 64)
	if err != nil {
		// Если клиент передал не число, возвращаем ошибку 400 Bad Request
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат amount"})
		return
	}

	var valletId uint = uint(valletId64)
	var amount uint = uint(amount64)

	product := m.ProductService.Save(valletId, operationType, amount)

	c.JSON(http.StatusOK, gin.H{"valletId": product.ValletId})
}
