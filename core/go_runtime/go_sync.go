package go_runtime

import (
	"fmt"
	"strconv"
)

var x int64

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}
func GoAdd() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println("done:" + strconv.FormatInt(x, 10))
}
