// 参考https://www.cnblogs.com/landv/p/11114508.html

package main

import (
	"flag"
	"fmt"
	"os"
)

// 定义需要解析的参数变量
var (
	h bool

	v, V bool
	t, T bool
	q    *bool

	s string
	p string
	c string
	g string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")

	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure option then exit")

	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration dump it and exit")

	q = flag.Bool("q", false, "suppress non-error messages during confuaration testing")

	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration file")

	// 改变默认Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}
}
