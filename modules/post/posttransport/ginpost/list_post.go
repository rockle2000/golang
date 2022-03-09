package ginpost

import (
	"instago2/common"
	"instago2/component"
	"instago2/modules/post/postbiz"
	"instago2/modules/post/poststorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListPost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewListPostBiz(store)

		result, err := biz.ListPost(c.Request.Context(), &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
