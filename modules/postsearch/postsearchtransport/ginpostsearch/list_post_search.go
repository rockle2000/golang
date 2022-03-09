package ginpostsearch

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/postsearch/postsearchbiz"
	"instago2/modules/postsearch/postsearchmodel"
	"instago2/modules/postsearch/postsearchstorage"
	"net/http"
)

func ListPostSearch(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter postsearchmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		data := postsearchmodel.DataSearch{}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := postsearchstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postsearchbiz.NewListPostBiz(store)

		result, err := biz.ListSearch(c.Request.Context(), &data, &filter, &paging)

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
