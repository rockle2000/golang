package gincomment

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/comment/commentbiz"
	"instago2/modules/comment/commentmodel"
	"instago2/modules/comment/commentstorage"
	"net/http"
)

func DeleteComment(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		data := commentmodel.CommentDelete{
			CommentId: int(uid.GetLocalID()),
		}
		store := commentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := commentbiz.NewDeleteCommentBiz(store, appCtx.GetPubsub())

		if err := biz.DeleteComment(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
