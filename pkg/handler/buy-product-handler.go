package handler

import (
	"dechdev/pkg/model"
	"dechdev/pkg/rule"
	"dechdev/pkg/util"
	"net/http"
	"time"

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
	productList = mockData()
	return productList, nil
}

func mockData() []model.ProductType {
	layout := time.RFC3339
	const recordDateStr = "2025-01-19T00:00:00Z"
	return model.ProductListType{
		{
			Id:               "1",
			Name:             "Tissue Paper Small Size",
			NumUnit:          24,
			UnitName:         "roll",
			MinimumUnit:      4,
			RecordDate:       parseTime(layout, "2024-01-19T00:00:00Z"),
			ExpireDate:       nil,
			MinimumExpireDay: 0,
			BuyFlag:          false,
		},
		{
			Id:               "2",
			Name:             "Tissue Paper Big Size",
			NumUnit:          1,
			UnitName:         "pcs",
			MinimumUnit:      2,
			RecordDate:       parseTime(layout, "2024-01-19T00:00:00Z"),
			ExpireDate:       nil,
			MinimumExpireDay: 0,
			BuyFlag:          false,
		},
		{
			Id:               "3",
			Name:             "Fish Sauce",
			NumUnit:          1,
			UnitName:         "bottle",
			RecordDate:       parseTime(layout, recordDateStr),
			ExpireDate:       parseTime(layout, "2025-04-30T00:00:00Z"),
			MinimumExpireDay: 30,
			BuyFlag:          false,
		},
		{
			Id:               "4",
			Name:             "Tomato Sauce",
			NumUnit:          1,
			UnitName:         "bottle",
			RecordDate:       parseTime(layout, recordDateStr),
			ExpireDate:       parseTime(layout, "2025-12-15T00:00:00Z"),
			MinimumExpireDay: 30,
			BuyFlag:          false,
		},
		{
			Id:               "5",
			Name:             "Soy Sauce",
			NumUnit:          1,
			UnitName:         "bottle",
			MinimumUnit:      0,
			RecordDate:       parseTime(layout, "2025-01-19T00:00:00Z"),
			ExpireDate:       parseTime(layout, "2025-12-15T00:00:00Z"),
			MinimumExpireDay: 30,
			BuyFlag:          true,
		},
		{
			Id:               "6",
			Name:             "Chili Sauce",
			NumUnit:          1,
			UnitName:         "bottle",
			RecordDate:       parseTime(layout, recordDateStr),
			ExpireDate:       parseTime(layout, "2025-04-25T00:00:00Z"),
			MinimumExpireDay: 10,
			BuyFlag:          false,
		},
	}
}

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}
