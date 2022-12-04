package race

import "sync"

func getCounter() int {
	counter := 0
	wg := sync.WaitGroup{}
	wg.Add(5)
	mx := sync.Mutex{}
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				mx.Lock() // without locking, data race will happen
				counter++
				mx.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}
