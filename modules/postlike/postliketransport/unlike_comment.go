package postliketransport

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/post/poststorage"
	"instago2/modules/postlike/postlikebusiness"
	"instago2/modules/postlike/postlikestorage"
	"net/http"
)

func UnlikePost(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := postlikestorage.NewSQLStore(ctx.GetMainDBConnection())
		decStore := poststorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := postlikebusiness.NewUnlikePostBiz(store, decStore)

		if err := biz.UnlikePost(c.Request.Context(), int(uid.GetLocalID()), requester.GetUserId()); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
