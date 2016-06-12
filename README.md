PETL: A Pipeline ETL process and receive data from pipe in Golang
==============

[![GoDoc](https://godoc.org/github.com/kkdai/maglev?status.svg)](https://godoc.org/github.com/kkdai/petl)  [![Build Status](https://travis-ci.org/kkdai/petl.svg?branch=master)](https://travis-ci.org/kkdai/petl) [![](https://goreportcard.com/badge/github.com/kkdai/petl)](https://goreportcard.com/badge/github.com/kkdai/petl)



Installation and Usage
=============


Install
---------------
```
go get github.com/kkdai/petl
```

Usage
---------------



```go

func main() {
	ret := PipeProcess(Extract("a  s", " smm", "sss"), TransformRemoveSpace, TransformRemoveSpace)
	log.Println("ret:", <-ret) //"as"
	log.Println("ret2:", <-ret) //"smm"
}
```

Install Console App
---------------
```
go get github.com/kkdai/petl/petl_cli
```

Usage
---------------

Use pipeline for pipe string



```
//Remove all space in output
cat somefile.txt | petl_cli pipeline="r"

//Make upper case
cat somefile.txt | petl_cli pipeline="u"

//Make lower case 
cat somefile.txt | petl_cli pipeline="l"

//Pipeline it
//Will make lower case first then upper case, final remove all space.
cat somefile.txt | petl_cli pipeline="lur"
```


Inspired By
---------------

- [Experimenting with Go pipelines](http://www.gmarik.info/blog/2016/experimenting-with-golang-pipelines/)
- [Go Blog:  Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)
- [Comparing Golang, Scala, Elixir, Ruby, and now Python3 for ETL: Part 2](http://blog.dimroc.com/2015/05/07/etl-language-showdown-pt2/)
- [stackoverflow: Can Functions be passed as parameters in Go?](http://stackoverflow.com/questions/12655464/can-functions-be-passed-as-parameters-in-go)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

etcd is under the Apache 2.0 [license](LICENSE). See the LICENSE file for details.