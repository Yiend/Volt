package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"volt/server"
)

func main() {
	var state int32 = 1
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	//运行服务
	server.Run()

	select {
	  case sig := <-sc:
		   atomic.StoreInt32(&state, 0)
		   fmt.Printf("获取到退出信号[%s]", sig.String())
	}

	fmt.Printf("服务退出")
	os.Exit(int(atomic.LoadInt32(&state)))
}

