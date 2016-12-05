package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func hashhash(length uint, prefix string) {
	ts := time.Now()
	defer func(ts time.Time) {
		log.Println("finished", time.Since(ts))
	}(ts)
	fmtstr := fmt.Sprintf("%%0%dx", length)
	maxv := uint64(1) << (length * 4)
	if maxv == 0 {
		log.Panic("maxv is zero")
	}
	for i := uint64(0); i < maxv; i++ {
		v := fmt.Sprintf(fmtstr, i)
		b := sha1.Sum([]byte(prefix + v))
		r := fmt.Sprintf("%x", b[len(b)-int(length)/2:])
		if r == v {
			log.Println("found", prefix, v, fmt.Sprintf("%x", b))
		}
	}
}

func main() {
	var (
		prefix string
		len    uint
	)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
	flag.StringVar(&prefix, "prefix", "", "prefix")
	flag.UintVar(&len, "len", 0, "len")
	flag.Parse()
	log.Printf("prefix=%s, len=%d\n", prefix, len)
	hashhash(len, prefix)
}
