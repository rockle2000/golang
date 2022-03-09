package postliketransport

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/post/poststorage"
	"instago2/modules/postlike/postlikebusiness"
	"instago2/modules/postlike/postlikemodel"
	"instago2/modules/postlike/postlikestorage"
	"net/http"
)

func CreatePostLikes(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		pid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := postlikemodel.PostLikes{
			PostId: int(pid.GetLocalID()),
			UserId: requester.GetUserId(),
		}

		store := postlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postlikebusiness.NewCreatePostLikesBiz(store, incStore)

		if err := biz.CreatePostLikes(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
