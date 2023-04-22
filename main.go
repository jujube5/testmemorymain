package main

import (
	"fmt"
	"time"
	"strconv"
	"runtime"
	"github.com/jujube5/testmemory"
)

func CheckMemAndTime(s string) {
	fmt.Println(s, testmemory.ReadMemInfo().Available/1024, "MB")
	t := time.Now()
	formatted := fmt.Sprintf("%02d:%02d:%02d",
        t.Hour(), t.Minute(), t.Second())
	fmt.Println(formatted)
}

func OutputRes(m string, s string) {
	testmemory.Colorize(testmemory.BlueBrightFont, s)
	CheckMemAndTime("Memory Available before test -- ")
	for i := 0; i < 5000000; i++ {
		s := strconv.Itoa(i)
		switch m {
			case "rset":
				testmemory.RedisSet(s, s)
			case "rget":
				testmemory.RedisGet(s)
			case "rdelete":
				testmemory.RedisDelete(s)
			case "nset":
				testmemory.Set(s, s)
			case "nget":
				testmemory.Get(s)
			case "ndelete":
				testmemory.Delete(s)
		}
	}
	CheckMemAndTime("Memory Available after test -- ")
}

func main() {
	OutputRes("rset", "Redis Set 5 000 000 records")
	OutputRes("rget", "Redis Get 5 000 000 records")
	OutputRes("rdelete", "Redis Delete 5 000 000 records")
	OutputRes("nset", "Native Set 5 000 000 records")
	OutputRes("nget", "Native Get 5 000 000 records")
	OutputRes("ndelete", "Native Delete 5 000 000 records")
	if testmemory.Get("8") == nil {
		testmemory.Colorize(testmemory.MagentaBrightFont, "Native Cache removed")
	}
	OutputRes("nset", "Native Set 5 000 000 records")
	OutputRes("ndelete", "Native Delete 5 000 000 records")
	testmemory.Colorize(testmemory.BlueBrightFont, "----------")
	runtime.GC()
	testmemory.Colorize(testmemory.MagentaBrightFont, "Garbage collector has started")
	time.Sleep(60 * time.Second)
	CheckMemAndTime("Check Memory Available -- ")
	time.Sleep(60 * time.Second)
	CheckMemAndTime("Check Memory Available -- ")
}
