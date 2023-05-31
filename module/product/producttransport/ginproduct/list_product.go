package ginproduct

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/product/productbiz"
	"TKPM-Go/module/product/productmodel"
	"TKPM-Go/module/product/productstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListProduct(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		var filter productmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		pagingData.Fulfill()

		var result []productmodel.Product
		store := productstorage.NewSQLStore(db)
		business := productbiz.NewListProductBusiness(store)
		result, err := business.ListProduct(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))

	}
}
