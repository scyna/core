package friend

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/friend/proto"
)

func ListFriend(c *scyna.Endpoint, request *proto.ListFriendRequest) {
	c.Logger.Info("Receive ListFriendRequest")

	if validation.Validate(request.Email, validation.Required, is.Email) != nil {
		c.Error(scyna.REQUEST_INVALID)
		return
	}

	// if err, user := model.Repository.GetByEmail(c.Logger, request.Email); err != nil {
	// 	c.Error(model.USER_NOT_EXISTED)
	// } else {
	// 	if err, users := model.Repository.ListFriend(c.Logger, user.ID); err != nil {
	// 		c.Error(err)
	// 	} else {
	// 		result := make([]*proto.User, len(users))
	// 		for i, u := range users {
	// 			result[i] = u.ToDTO()
	// 		}
	// 		c.Done(&proto.ListFriendResponse{
	// 			Items: result,
	// 		})
	// 	}
	// }
}
