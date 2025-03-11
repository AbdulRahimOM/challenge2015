package status

import (
	"fmt"
	"runtime"
	"test/internal/config"
	"time"
)

func init() {
	if config.LogGoroutineCount {
		go func() {
			for {
				time.Sleep(3 * time.Second)
				GetStatistics()
			}
		}()
	}
}

func GetStatistics() {

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Println("=====================================")
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine(), "@", time.Now().Format("15:04:05"))
	// fmt.Println("Heap Allocation (bytes):", memStats.HeapAlloc)
	// fmt.Println("Stack System (bytes):", memStats.StackSys)
	// fmt.Println("Stack In Use (bytes):", memStats.StackInuse)

	// buf := make([]byte, 10*1024)
	// n := runtime.Stack(buf, true)
	// fmt.Println("result:", string(buf[:n]))

	// fmt.Println("=====================================")
}
