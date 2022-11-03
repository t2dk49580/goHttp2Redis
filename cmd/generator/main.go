package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"goHttp2Redis/connector"
	"os"
)

func main() {
	var (
		redisaddr string
		redispass string
		redisargs string
		redisset  string
	)

	var cmdl = flag.NewFlagSet("", flag.ExitOnError)
	cmdl.StringVar(&redisaddr, "redis", "localhost:6666", "redis address")
	cmdl.StringVar(&redispass, "pass", "", "redis password (default: \"\")")
	cmdl.StringVar(&redisargs, "get", "", "get args from redis key (default: \"\")")
	cmdl.StringVar(&redisset, "set", "", "set value to redis key (default: \"\")")
	cmdl.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of params:\n")
		cmdl.PrintDefaults()
	}
	cmdl.Parse(os.Args[1:])

	key, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("err: crypto.GenerateKey()," + err.Error())
		return
	}

	args, err := connector.Get(redisaddr, redispass, redisargs)
	if redisargs != "" {
		if err != nil {
			fmt.Println("err: connector.Get()," + err.Error())
			return
		}
		fmt.Println("redisargs: " + args)
	} else {
		fmt.Println("redisargs: not set")
	}

	if redisset != "" {
		result, err := connector.Set(redisaddr, redispass, redisset, hex.EncodeToString(key.D.Bytes()))
		if err != nil {
			fmt.Println("err: connector.Set()," + err.Error())
			return
		}
		fmt.Println("redisset: " + result)
	} else {
		fmt.Println("redisset: not set")
	}
	fmt.Println("privateKey: " + hex.EncodeToString(key.D.Bytes()))
	fmt.Println("address: " + crypto.PubkeyToAddress(key.PublicKey).Hex())
}
