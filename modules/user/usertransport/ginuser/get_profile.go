package ginuser

import (
	"instago2/common"
	"instago2/component"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		//data := c.MustGet(common.CurrentUser)
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
