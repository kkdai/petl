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
	sizeN := 5
	lookupSizeM := 13 //(must be prime number)

	var names []string
	for i := 0; i < sizeN; i++ {
		names = append(names, fmt.Sprintf("backend-%d", i))
	}
	//backend-0 ~ backend-4 

	mm := NewMaglev(names, lookupSizeM)
	v, err := mm.Get("IP1")
	fmt.Println("node1:", v)
	//node1: backend-2
	v, _ = mm.Get("IP2")
	log.Println("node2:", v)
	//node2: backend-1
	v, _ = mm.Get("IPasdasdwni2")
	log.Println("node3:", v)
	//node3: backend-0

	if err := mm.Remove("backend-0"); err != nil {
		log.Fatal("Remove failed", err)
	}
	v, _ = mm.Get("IPasdasdwni2")
	log.Println("node3-D:", v)
	//node3-D: Change from "backend-0" to "backend-1"
}
```

Install Console App
---------------
```
go get github.com/kkdai/petl/petl_cli
```

Usage
---------------


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