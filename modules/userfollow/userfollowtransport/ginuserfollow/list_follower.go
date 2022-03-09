package ginuserfollow

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/userfollow/userfollowbiz"
	"instago2/modules/userfollow/userfollowmodel"
	"instago2/modules/userfollow/userfollowstorage"
	"net/http"
)

func ListFollower(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		userId := requester.GetUserId()
		var filter userfollowmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := userfollowstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := userfollowbiz.NewListFollowerBiz(store)
		result, err := biz.ListFollower(c.Request.Context(), userId, &filter, &paging)
		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
