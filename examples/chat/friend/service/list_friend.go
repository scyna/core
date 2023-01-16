package friend

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	scyna "github.com/scyna/core"
	"github.com/scyna/core/examples/chat/friend/proto"
)

func ListFriend(c *scyna.Endpoint, request *proto.ListFriendRequest) scyna.Error {
	c.Logger.Info("Receive ListFriendRequest")

	if validation.Validate(request.Email, validation.Required, is.Email) != nil {
		return scyna.REQUEST_INVALID
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
	return scyna.OK
}
