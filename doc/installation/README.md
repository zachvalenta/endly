# Installation

1) [Download latest binary](https://github.com/viant/endly/releases/)
```bash
 tar -xvzf endly_xxx.tar.gz
 cp endly /usr/local/bin
 endly -h
 endly -v
```

 
2) Build from source
   a) install go 1.11+
   b) run the following commands:

```bash
git clone https://github.com/viant/endly.git
cd endly/endly
go build endly.go
cp endly /usr/local/bin
```


3) Custom build, in case you need additional drivers, dependencies or UDF with additional imports:

@endly.go
```go

package main

//import your udf package  or other dependencies here

import "github.com/viant/endly/bootstrap"

func main() {
	bootstrap.Bootstrap()
}

```       





