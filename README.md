# golang packages
* [Everything you need to know about Packages in Go](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc)

```bash
$ cd $GOPATH

# main package
$ vi src/app/entry.go
$ go install app
$ app
app/entry.go ==> main()

# add package greet
$ vi src/greet/day.go
$ go run src/app/entry.go 
Hey, Good Morning

# nested package
$ vi src/greet/de/day.go
$ go run src/app/entry.go 
app/entry.go ==> main()
Hey, Good Morning
Hallo, Guten Morgen

$ ls -lah pkg/darwin_amd64/g
golang.org/ greet/      greet.a     

# run all
$ vi src/app/version.go
$ go run src/app/*.go 
app/entry.go ==> main()
Hey, Good Morning
Hallo, Guten Morgen
version ===>  1.0.0
$ go run src/app/version.go src/app/entry.go 
app/entry.go ==> main()
Hey, Good Morning
Hallo, Guten Morgen
version ===>  1.0.0

# implicit init() method
$ vi src/app/f.go
$ go run src/app/*.go
app/entry.go ==> init() [1]
app/entry.go ==> init() [2]
app/entry.go ==> main()
Hey, Good Morning
Hallo, Guten Morgen
version ===>  1.0.0
2 2 1

# package alias
$ vi src/app/entry.go
$ go run src/app/*.go
greet/greet/child.go ===> init()
app/entry.go ==> init() [1]
app/entry.go ==> init() [2]
app/entry.go ==> main()
Hey, Good Morning
Hallo, Guten Morgen
version ===>  1.0.0
2 2 1
Hey there. I am child.

# import package from github
$ go get -u github.com/jinzhu/gorm
$ cat src/app/entry.go
```
