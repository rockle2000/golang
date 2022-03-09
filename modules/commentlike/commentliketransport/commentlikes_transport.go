package commentliketransport

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/comment/commentstorage"
	"instago2/modules/commentlike/commentlikebusiness"
	"instago2/modules/commentlike/commentlikemodel"
	"instago2/modules/commentlike/commentlikestorage"
	"net/http"
)

func CreateCommentLikes(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		cid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := commentlikemodel.CommentLikes{
			CommentId: int(cid.GetLocalID()),
			UserId:    requester.GetUserId(),
		}

		store := commentlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := commentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := commentlikebusiness.NewCreateCommentLikesBiz(store, incStore)

		if err := biz.CreateCommentLikes(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
