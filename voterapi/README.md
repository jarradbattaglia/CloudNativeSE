# How to run

The program runs on port 8080 instead of 1080, because macs block that port and it kept asking for permission.

Extra credit parts 1 and 2 were done, (json and puts/deletes)

But run in one tab:

```
go run .
```

and then run `go test ./... -v. ` to run the test suite.

Which will start up a bare server (with no info), you may run /voters/populate to give some hard coded data.

But dont need to do that! test will delete all data on start and cleanup and runs through each call once
