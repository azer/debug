## deprecated

I'm maintaining [logger](http://github.com/azer/logger) now.

## debug

Golang equivalent of [visionmedia/debug](http://github.com/visionmedia/debug)

foo.go:

```go
Debug("doing some work")
```

bar.go:
```go
Debug("doing more work")
```

qux.go:
```go
Debug("keeps doing more work")
```

**Example Outputs**

`DEBUG=*`

![](https://dl.dropboxusercontent.com/s/ebi3qvza1twrplk/debug1.png)

`DEBUG=foo,bar`

![](https://dl.dropboxusercontent.com/s/58yizd4nbvq7q74/debug2.png)

See `examples/` for more info.
