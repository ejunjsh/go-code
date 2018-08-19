package main

import (
    "context"
    "log"
    "os"
    "time"
)

var logg *log.Logger

func someHandler() {
    ctx, cancel := context.WithCancel(context.Background())
    go doStuff(ctx)

//10秒后取消doStuff
    time.Sleep(10 * time.Second)
    cancel()
    time.Sleep(1 * time.Second)
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
    for {
        time.Sleep(1 * time.Second)
        select {
        case <-ctx.Done():
            logg.Printf("done")
            return
        default:
            logg.Printf("work")
        }
    }
}

func main() {
    logg = log.New(os.Stdout, "", log.Ltime)
    someHandler()
    logg.Printf("end")
}