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

func CreatePost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data postmodel.PostCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = requester.GetUserId()

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewCreatePostBiz(store)

		if err := biz.CreatePost(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DbTypePost)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
