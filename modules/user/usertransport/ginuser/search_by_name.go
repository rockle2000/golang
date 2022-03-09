package ginuser

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/user/userbiz"
	"instago2/modules/user/userstorage"
	"net/http"
)

func SearchUserByName(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		searchKey := c.Param("searchKey")
		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.SearchUserByNameBiz(store)

		data, err := biz.SearchUserByName(c.Request.Context(), searchKey)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)

		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
