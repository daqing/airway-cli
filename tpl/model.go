package tpl

func Model() string {
	return `package models

import "time"

type {{.Name}} struct {
	ID sql.IdType ` +
		"`json:\"id\"`\n" +
		"\n  CreatedAt time.Time `json:\"created_at\"`" +
		"\n  UpdatedAt time.Time `json:\"updated_at\"`" +
		`
}`
}
