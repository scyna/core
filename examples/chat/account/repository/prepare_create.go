package repository

import (
	scyna "github.com/scyna/core"
)

func PrepareCreate(cmd *scyna.Command, user *User) {
	query := "INSERT INTO" + ACCOUNT_TABLE_NAME + "(id, name, email, password) VALUES(?,?,?,?)"
	cmd.Batch.Query(query, user.ID, user.Name, user.Email, user.Password)
}
