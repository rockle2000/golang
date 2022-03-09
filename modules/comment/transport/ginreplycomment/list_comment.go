package ginreplycomment

import (
	"instago2/common"
	"instago2/component"
	"instago2/modules/comment/commentbiz"
	"instago2/modules/comment/commentmodel"
	"instago2/modules/comment/commentstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListComment(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var filter commentmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := commentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := commentbiz.NewListCommentBiz(store)

		result, err := biz.ListComment(c.Request.Context(), int(uid.GetLocalID()), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
