package tpl

func Service() string {
	return `package services

	func Find{{.Name}}(id sql.IdType) (*models.{{.Name}}, error) {
		return db.FindOne[models.{{.Name}}](sql.H{"id": id})
	}

	func Create{{.Name}}({{.Fields}}) (*models.{{.Name}}, error) {
    return db.Create[models.{{.Name}}](sql.H{{.SQLH}})
	}

	func Update{{.Name}}(id sql.IdType, {{.Fields}}) error {
		return db.Update[models.{{.Name}}](sql.H{{.SQLH}}, sql.Eq("id", id))
	}

	func Delete{{.Name}}(id sql.IdType) error {
		return db.DeleteById[models.{{.Name}}](id)
	}
`
}
