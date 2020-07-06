package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var (
	client    *redis.Client
	uri       string
	listened  string
	redisaddr string
	redispass string
	redisdb   int
	luascript string
)

func favicon(w http.ResponseWriter, r *http.Request) {

}

func rsp(element interface{}, w http.ResponseWriter, r *http.Request) {
	switch element.(type) {
	case int:
		cur := strconv.Itoa(element.(int))
		w.Write([]byte(cur))
	case int64:
		cur := strconv.FormatInt(element.(int64), 10)
		w.Write([]byte(cur))
	case string:
		w.Write([]byte(element.(string)))
	default:
		w.Write([]byte("unknow element type"))
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Form)
	fmt.Println(r.PostFormValue("data"))
	result, err := client.Do("EVAL", luascript, "0", r.PostFormValue("data")).Result()
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		rsp(result, w, r)
	}
}

func main() {
	var cmdl = flag.NewFlagSet("", flag.ExitOnError)
	cmdl.StringVar(&uri, "uri", "/", "handle request field")
	cmdl.StringVar(&listened, "listen", ":9999", "listen port, default for internet")
	cmdl.StringVar(&redisaddr, "redis", "localhost:6379", "redis address")
	cmdl.StringVar(&redispass, "pass", "", "redis password (default: \"\")")
	cmdl.IntVar(&redisdb, "db", 0, "redis database id (default: 0)")
	cmdl.StringVar(&luascript, "lua", "return ARGV[1]..' '..ARGV[2]", "lua handle function body")
	cmdl.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of params:\n")
		cmdl.PrintDefaults()
	}
	cmdl.Parse(os.Args[1:])

	client = redis.NewClient(&redis.Options{Addr: redisaddr, Password: redispass, DB: redisdb})
	http.HandleFunc(uri, handle)
	http.HandleFunc("/favicon.ico", favicon)
	fmt.Println("uri:", uri, "listen:", listened, "redis:", redisaddr, "db:", redisdb, "lua:", luascript)
	err := http.ListenAndServe(listened, nil)
	if err != nil {
		fmt.Println(err)
	}
}
