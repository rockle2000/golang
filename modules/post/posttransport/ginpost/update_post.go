package ginpost

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/post/postbiz"
	"instago2/modules/post/postmodel"
	"instago2/modules/post/poststorage"
	"net/http"
	"strconv"
)

func UpdatePost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//uid, err := common.FromBase58(c.Param("id"))
		uid, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data postmodel.PostUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewUpdatePostBiz(store)

		if err := biz.UpdatePost(c.Request.Context(), int(uid), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
