package repository

import (
	scyna "github.com/scyna/core"
)

const _QUERY = "INSERT INTO" + ACCOUNT_TABLE_NAME + "(id, name, email, password) VALUES(?,?,?,?)"

func PrepareCreate(cmd *scyna.Command, user *User) {
	cmd.Batch.Query(_QUERY, user.ID, user.Name, user.Email, user.Password)
}
