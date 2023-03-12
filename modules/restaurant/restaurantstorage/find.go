package restaurantstorage

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	morekeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var result restaurantmodel.Restaurant

	db := s.db

	for i := range morekeys {
		db = db.Preload(morekeys[i])
	}

	if err := db.Where(conditions).
		First(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
