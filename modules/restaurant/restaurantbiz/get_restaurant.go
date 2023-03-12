package restaurantbiz

import (
	"context"
	"food-delivery/modules/restaurant/restaurantmodel"
)

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		morekeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	return result, err
}
