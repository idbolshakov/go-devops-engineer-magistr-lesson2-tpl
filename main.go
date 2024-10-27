package main

import (
  "os"
  "fmt"
  "gopkg.in/yaml.v3"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println(fmt.Errorf("yaml file path is not set"))
    os.Exit(-1)
    return
  }

  yamlFilePath := os.Args[1]

  // 0: read file
  content, err := os.ReadFile(yamlFilePath)
  if err != nil {
    fmt.Println(fmt.Errorf("cannot read file content: %w", err))
    os.Exit(-1)
    return
  }

  // 1: parse yaml
  var root yaml.Node
  if err := yaml.Unmarshal(content, &root); err != nil {
    fmt.Println(fmt.Errorf("cannot unmarshal file content: %w", err))
    os.Exit(-1)
    return
  }


  for _, doc := range root.Content {
    // поле `.Content` содержит вложенные дочерние элементы
    for _, node := range doc.Content {
      fmt.Printf("line: %d, value: %s\n", node.Line, node.Value)
    }
  } 
}
