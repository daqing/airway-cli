package tpl

func Cmd() string {
	return `package cmd

  func create{{.Name}}(args []string) {
    if len(args) == 0 {
      fmt.Println("Usage:  create_{{.LowerName}} {{.Fields}}")
      return
    }

		services.Create{{.Name}}({{.Args}})
  }

  func delete{{.Name}}(args []string) {
    if len(args) != 1 {
      log.Fatal("Usage: delete_{{.LowerName}} <id>")
    }

    id, err := strconv.ParseInt(args[0], 10, 64)
    if err != nil {
      log.Fatalf("Invalid id: %s", args[0])
    }

    services.Delete{{.Name}}(sql.IdType(id))
  }

  func update{{.Name}}(args []string) {
    if len(args) < 2 {
      log.Fatal("Usage: update_{{.LowerName}} <id> {{.Fields}}")
    }

    id, err := strconv.ParseInt(args[0], 10, 64)
    if err != nil {
      log.Fatalf("Invalid id: %s", args[0])
    }

    services.Update{{.Name}}(sql.IdType(id), {{.Args1}})
  }

  func find{{.Name}}(args []string) (*models.{{.Name}}, error) {
    if len(args) != 1 {
      log.Fatal("Usage: find_{{.LowerName}} <id>")
    }

    id, err := strconv.ParseInt(args[0], 10, 64)
    if err != nil {
      log.Fatalf("Invalid id: %s", args[0])
    }

    return services.Find{{.Name}}(sql.IdType(id))
  }
`
}
