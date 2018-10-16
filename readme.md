# Quote for quick string constructions

Construct a string slice as Ruby's quote
- q.Word() ~ Ruby's %w
- q.Heredoc() ~ Ruby's <<~ EOF
- q.Template() ~ Ruby's embedded string #{var}
- q.Line()

## Synposis

Get and import
```
go get github.com/zhiminwen/q
```

import "github.com/zhiminwen/q"

Sample:

```golang
var list []string

// slice of word
list = q.Word(`master1 master2 master3`) 

// slice from lines
list = q.Line(`
  // Lines are trimmed with Left and Right Spaces
  # Empty line and comment line (started with # or //) are ignored
  line 1
  line 2
`)

// a heredoc and docf
doc := q.HereDoc(`
  A block of heredoc with proper indent. (similar as ruby's <<~ )
    Simply a wrapper of github.com/MakeNowJust/heredoc heredoc.Doc(text)
`)

docf := q.HereDocf(`
  A block of heredoc with proper indent. (similar as ruby's <<~ )
    Simply a wrapper of github.com/MakeNowJust/heredoc heredoc.Doc(text)
    Variable = %s
`, "Value")

// a shortcut for join commands
cmd := q.Cmd(`
  // construct command lines. Same as Line(), join lines with " && " or other connector like ";" 
  ls -ltr

  hostname
`, " && ")
//cmd == "ls -ltr && hostname"


//Construct a quick template based string using map[string]string
text := q.Template(`This is a test for {{ .Name}} at {{ .Address }}`, map[string]string{
  Name: "test",
  Address: "Street 1", 
})

``` 


