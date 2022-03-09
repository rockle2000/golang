package userstorage

import (
	"context"
	"gorm.io/gorm"
	"instago2/common"
	"instago2/modules/user/usermodel"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}

func (s *sqlStore) FindUserByName(
	ctx context.Context,
	searchKey string,
	moreInfo ...string,
) ([]common.SimpleUser, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user []common.SimpleUser

	if err := db.Where("? LIKE user_name", searchKey).Or("? LIKE first_name", searchKey).Or("? LIKE last_name", searchKey).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return user, nil
}
