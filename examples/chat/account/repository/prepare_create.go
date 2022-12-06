package repository

import (
	scyna "github.com/scyna/core"
)

const _QUERY = "INSERT INTO " + ACCOUNT_TABLE + "(id, name, email, password) VALUES(?,?,?,?)"

func PrepareCreate(cmd *scyna.Command, user *Account) {
	cmd.Batch.Query(_QUERY, user.ID, user.Name, user.Email, user.Password)
}
