# Go CLI App Layout

The `go install` command looks for a go file with a `package main`
in the directory specified (or current dir by default).
It compiles and places the executable in the `$GOPATH/bin` folder _with
the same name as the directory containing the main `.go` file_.

``` 
greet/
  cmd/
    greet/
      main.go // imports goapplayout/greet
  greet/    
    greet.go // lib package
  README.md
  go.mod
```

Running `go install cmd/greet/` from the root of the repo
will create a `greet.exe` in the `$GOPATH/bin` folder.
If this directory is on the system path, `greet` will be executable
from anywhere on the system.

This command can fetch and install the app from github
on a brand new machine with the Go CLI installed.
```
go get github.com/user/goapplayout/cmd/greet
```



# Go Library Layout

```
libname/
    greet/
        greet.go // package greet
        greet_test.go
    config.go // package libname (if config files)
    go.mod
    README.md
    LICENCE
```

See https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/ for another example.