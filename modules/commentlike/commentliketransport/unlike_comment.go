package commentliketransport

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/comment/commentstorage"
	"instago2/modules/commentlike/commentlikebusiness"
	"instago2/modules/commentlike/commentlikestorage"
	"net/http"
)

func UnlikeComment(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := commentlikestorage.NewSQLStore(ctx.GetMainDBConnection())
		decStore := commentstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := commentlikebusiness.NewUnlikeCommentBiz(store, decStore)

		if err := biz.UnlikeComment(c.Request.Context(), int(uid.GetLocalID()), requester.GetUserId()); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
