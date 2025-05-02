package rule

import (
	"dechdev/pkg/model"
	"dechdev/pkg/util"
	"fmt"
	"time"
)

func IsProductPurchasable(product *model.ProductType, shoppingDate time.Time) (bool, string) {
	if !isProductDefined(product) {
		return false, ""
	}

	if buy, reason := IsProductMarkedForPurchase(product); buy {
		return true, reason
	}
	if buy, reason := IsProductExpiredOrNearExpiry(product, shoppingDate); buy {
		return true, reason
	}
	if buy, reason := IsProductLowInStock(product); buy {
		return true, reason
	}

	return false, ""
}

func IsProductMarkedForPurchase(product *model.ProductType) (bool, string) {
	if !isProductDefined(product) {
		return false, ""
	}

	if product.BuyFlag {
		// fmt.Println("Buy From Flag")
		return true, "Product is marked for purchase"
	}

	return false, ""
}

func IsProductExpiredOrNearExpiry(product *model.ProductType, shoppingDate time.Time) (bool, string) {
	if !isProductDefined(product) {
		return false, ""
	}

	if product.ExpireDate == nil || product.MinimumExpireDay == 0 {
		return false, ""
	}

	diffDay := util.DiffDay(&shoppingDate, product.ExpireDate)

	if diffDay <= product.MinimumExpireDay {
		// fmt.Println("Buy Is Near Expired")
		if diffDay < 0 {
			diffDay = -diffDay
			return true, fmt.Sprintf("Product is past expiration (%d days)", diffDay)
		}

		return true, fmt.Sprintf("Product is near expiration (%d days remaining)", diffDay)
	}

	return false, ""
}

func IsProductLowInStock(product *model.ProductType) (bool, string) {
	if !isProductDefined(product) {
		return false, ""
	}

	if product.NumUnit == 0 || product.MinimumUnit == 0 {
		return false, ""
	}

	if product.NumUnit <= product.MinimumUnit {
		// fmt.Println("Buy Is Near Out Of Stock")
		return true, fmt.Sprintf("Product is near out of stock (%d/%d units)", product.NumUnit, product.MinimumUnit)
	}

	return false, ""
}

func isProductDefined(product *model.ProductType) bool {
	return product != nil
}
