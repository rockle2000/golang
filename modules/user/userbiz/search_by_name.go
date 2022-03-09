package userbiz

import (
	"context"
	"instago2/common"
)

type SearchUserByName interface {
	FindUserByName(
		ctx context.Context,
		searchKey string,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type searchUserByNameBiz struct {
	store SearchUserByName
}

func SearchUserByNameBiz(store SearchUserByName) *searchUserByNameBiz {
	return &searchUserByNameBiz{store: store}
}

func (biz *searchUserByNameBiz) SearchUserByName(ctx context.Context, searchKey string) ([]common.SimpleUser, error) {
	data, err := biz.store.FindUserByName(ctx, searchKey)

	if err != nil {
		return nil, common.ErrCannotGetEntity("There is no result !", err)
	}
	return data, err
}
