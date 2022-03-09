package commentbiz

import (
	"context"
	"instago2/common"
	"instago2/modules/comment/commentmodel"
)

type CreateComment interface {
	CreateComment(ctx context.Context, data *commentmodel.CommentCreate) error
}

type CreateCommentBiz struct {
	store CreateComment
}

func (biz *CreateCommentBiz) CreateComment(ctx context.Context, data *commentmodel.CommentCreate) error {
	err := biz.store.CreateComment(ctx, data)

	if err != nil {
		return common.ErrCannotCreateEntity("Can not create comment!", err)
	}
	return err
}

func NewCreateCommentBiz(store CreateComment) *CreateCommentBiz {
	return &CreateCommentBiz{store: store}
}
