package Week06

import (
	"fmt"
	"testing"
	"time"
)

func TestLimiter(t *testing.T){
	limiter,err := NewLimiter(time.Second*5,time.Second,50000)
	if err != nil {
		panic(err)
	}

	go func() {
		for i := 0;i < 30000; i ++ {
			ok := limiter.Add()
			if ok {
				//限速之后的逻辑
			}
		}
	}()
	go func() {
		for i := 0;i < 30000; i ++ {
			ok := limiter.Add()
			if ok {
			}
		}
	}()
	time.Sleep(6*time.Second)
	limiter.Stop()
}


func TestLoop(t *testing.T){
	var name string
	fmt.Scanln(&name)
	fmt.Println(name)
}

