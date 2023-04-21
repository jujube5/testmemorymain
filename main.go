package main

import (
	"fmt"
	"time"
	"strconv"
	"github.com/jujube5/testmemory"
)

func OutputRes(m string, s string) {
	testmemory.Colorize(testmemory.YellowBrightFont, s)
	fmt.Println("Memory Available before test -- ", testmemory.ReadMemInfo().Available/1024, "MB")
	t := time.Now()
	formatted := fmt.Sprintf("%02d:%02d:%02d",
        t.Hour(), t.Minute(), t.Second())
	fmt.Println(formatted)

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
	
	fmt.Println("Memory Available after test -- ", testmemory.ReadMemInfo().Available/1024, "MB")
	t2 := time.Now()
	formatted2 := fmt.Sprintf("%02d:%02d:%02d",
        t2.Hour(), t2.Minute(), t2.Second())
	fmt.Println(formatted2)
}

func main() {
	OutputRes("rset", "Redis Set 5 000 000 records")
	OutputRes("rget", "Redis Get 5 000 000 records")
	OutputRes("rdelete", "Redis Delete 5 000 000 records")
	OutputRes("nset", "Native Set 5 000 000 records")
	OutputRes("nget", "Native Get 5 000 000 records")
	OutputRes("ndelete", "Native Delete 5 000 000 records")
}
