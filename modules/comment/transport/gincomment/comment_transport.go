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

func CreateComment(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		pid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var content commentmodel.CommentCreate
		if err := c.ShouldBind(&content); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data := commentmodel.CommentCreate{
			PostId:  int(pid.GetLocalID()),
			UserId:  requester.GetUserId(),
			Content: content.Content,
		}

		store := commentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := commentbiz.NewCreateCommentBiz(store)

		if err := biz.CreateComment(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
