package subscriber

import (
	"context"
	"instago2/component"
	"instago2/modules/comment_likes/comment_likesstorage"
	"instago2/pubsub"
)

type HasCommentDeleteId interface {
	GetCommmentDeleteId() int
}

func DeleteCommentLikeAfterDeleteComment(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Delete comment like after delete comment",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := comment_likesstorage.NewSQLStore(appCtx.GetMainDBConnection())
			deleteData := message.Data().(HasCommentDeleteId)
			return store.DeleteAfterDeleteComment(ctx, deleteData.GetCommmentDeleteId())
		},
	}
}
