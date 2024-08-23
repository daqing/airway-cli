package tpl

func Model() string {
	return `
		package models

		type {{.Name}} struct {
			ID IdType ` +
		"`gorm:\"primarykey\" json:\"id\"`\n\n" +
		`{{.Fields}}` +
		"\n\nCreatedAt Timestamp `json:\"created_at\"`" +
		"\nUpdatedAt Timestamp `json:\"updated_at\"`" +
		`
    }

		func ({{.Name}}) TableName() string { return "{{.TableName}}" }
	`
}
