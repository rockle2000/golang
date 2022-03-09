package ginuserfollow

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/user/userstorage"
	"instago2/modules/userfollow/userfollowbiz"
	"instago2/modules/userfollow/userfollowstorage"
	"net/http"
)

func UserUnfollowUser(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		store := userfollowstorage.NewSQLStore(ctx.GetMainDBConnection())
		decStore := userstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := userfollowbiz.NewUserUnfollowUserBiz(store, decStore)

		if err := biz.UserUnfollowUser(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
