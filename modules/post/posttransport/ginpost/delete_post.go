package ginpost

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/post/postbiz"
	"instago2/modules/post/postmodel"
	"instago2/modules/post/poststorage"
	"net/http"
)

func DeletePost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := postmodel.PostDelete{
			PostId: int(uid.GetLocalID()),
			UserId: requester.GetUserId(),
		}

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewDeletePostBiz(store, appCtx.GetPubsub())

		if err := biz.DeletePost(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
