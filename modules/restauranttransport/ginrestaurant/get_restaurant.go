package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component"
	"food-delivery/modules/restaurant/restaurantbiz"
	"food-delivery/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(err)
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), id)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
