package gotool

import (
	"context"
	"time"
	"fmt"
)

type Result struct{

}

type CallbackFunc func(signal chan Result) 

func process(seconds int,cbFunc CallbackFunc) (status int) {
	ctx,cancel := context.WithTimeout(context.Background(),time.Duration(seconds)*time.Second)
	defer cancel()

	//do something init
	//end
	finish := make(chan Result)
	go cbFunc(finish)
	select {
	//timeout
	case <-ctx.Done():
		//do something
		fmt.Println("timeout happen")
		return 1
	//finish	
	case <- finish:
		fmt.Println("finish")
		return 0
	}
}