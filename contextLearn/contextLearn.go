package main

import (
	"context"
	"fmt"
	"time"
)

func dosomething(ctx context.Context) {
	//  监听ctx有没有被cancel
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("finish do thing")
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}

	}

}

func main() {
	// 创建context
	ctx := context.Background() //返回一个空的context ,不能被cancel

	// todo := context.TODO()  // 和background类似  ,当你不确定用 那种context

	ctx, cancelFunc := context.WithCancel(ctx)
	start := time.Now()
	go func() {
		time.Sleep(6 * time.Second)
		cancelFunc()
	}()
	// context 作为函数的第一个参数使用 命名为ctx
	dosomething(ctx)
	end := time.Since(start)
	print(end)
}
