package main

import (
	"os"
	"strconv"
	"strings"
	"time"
)

/*
sleep 1
sleep 1s
sleep 1m
sleep 1h
sleep 2.5
sleep 0.5m
*/
func main() {
	if len(os.Args) <= 1 {
		return
	}
	s := os.Args[1][0 : len(os.Args[1])-1]
	f0, err0 := strconv.ParseFloat(os.Args[1], 32)
	if len(s) < 1 && err0 != nil {
		return
	}
	var f1 float64
	var err2 error
	if err0 != nil {
		f1, err2 = strconv.ParseFloat(s, 32)
		if err2 != nil {
			return
		}
	}
	if strings.HasSuffix(os.Args[1], "s") {
		time.Sleep(time.Duration(int(f1*1000)) * time.Millisecond)
	} else if strings.HasSuffix(os.Args[1], "m") {
		time.Sleep(time.Duration(int(f1*1000*60)) * time.Millisecond)
	} else if strings.HasSuffix(os.Args[1], "h") {
		time.Sleep(time.Duration(int(f1*1000*60*60)) * time.Millisecond)
	} else if err0 == nil {
		time.Sleep(time.Duration(int(f0*1000)) * time.Millisecond)
	}
}
