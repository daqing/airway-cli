package tpl

func Table() string {
	return `package tables

type {{.Name}} struct{}

func ({{.Name}}) TableName() string {
	return "{{.TableName}}"
}
	`
}
