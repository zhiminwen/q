# Quote for a String Slice

Construct a string slice as Ruby's quote

## Synposis

Get and import
```
go get github.com/zhiminwen/q
```

Sample:

```golang
  var list []string
  list = q.Word(`master1 master2 master3 worker1 worker2 worker3`)

  list = q.Line(`
    // Lines are trimmed with Left and Right Spaces
    // Empty line and comment line (started with # or //) are ignored
    line 1
    line 2
  `
  )
``` 


