package handler

import (
	"dechdev/pkg/config"
	"dechdev/pkg/model"
	"dechdev/pkg/rule"
	"dechdev/pkg/util"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuyProduct(c *gin.Context) {
	var requestBody model.BuyProductRequestBody

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultList := model.ResultListBuyProduct{}

	productList, err := initData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, v := range productList {
		// fmt.Printf("Product ID: %s\n", v.Id)
		isPurchasable, reason := rule.IsProductPurchasable(&v, *util.DateStringToTime(requestBody.ShoppingDate))
		//fmt.Printf("Is Product Purchasable: %t\nReason: %s\n", isPurchasable, reason)
		//fmt.Println("------------------------------")

		result := model.ResultBuyProduct{
			ProductId:   v.Id,
			ProductName: v.Name,
			IsBuy:       isPurchasable,
			Reason:      reason,
		}

		resultList = append(resultList, result)
	}

	c.JSON(200, gin.H{
		"shoppingDate": requestBody.ShoppingDate,
		"result":       resultList,
		"error":        nil,
	})
}

func initData() (productList model.ProductListType, err error) {
	productList = model.ProductListType{}

	// Read value from .env
	dataJsonFile := config.DataJsonFile

	// Read Data from JSON file
	data, err := util.ReadJsonFile(dataJsonFile)
	if err != nil {
		return productList, err
	}

	err = json.Unmarshal(data, &productList)
	if err != nil {
		return productList, err
	}

	fmt.Println("Data loaded successfully")

	return productList, nil
}
