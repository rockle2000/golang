package ginuser

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/user/userbiz"
	"instago2/modules/user/userstorage"
	"net/http"
)

func GetOtherProfile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrEntityNotFound("Not found the user!", err))
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewGetOtherProfileBiz(store)

		data, err := biz.GetProfile(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
