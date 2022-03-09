package ginreplycomment

import (
	"github.com/gin-gonic/gin"
	"instago2/common"
	"instago2/component"
	"instago2/modules/comment/commentbiz"
	"instago2/modules/comment/commentmodel"
	"instago2/modules/comment/commentstorage"
	"net/http"
)

func CreateReply(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		dataRequest := commentmodel.CommentCreateRequest{}
		if err := c.ShouldBind(&dataRequest); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		postUid, err := common.FromBase58(dataRequest.PostId)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		commentUid, err := common.FromBase58(dataRequest.CommentId)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data := commentmodel.CommentCreate{
			UserId:   requester.GetUserId(),
			PostId:   int(postUid.GetLocalID()),
			ParentId: int(commentUid.GetLocalID()),
			Content:  dataRequest.Content,
		}

		store := commentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := commentbiz.NewCreateReplyBiz(store)

		if err := biz.CreateReply(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DbTypeComment)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
