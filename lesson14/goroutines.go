package main
import "sync"
import "fmt"

func main() {
    ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(name string) {
		for i := 1; i <= 5; i++ {
			ch <- fmt.Sprintf("%s: %d", name, i)
		}
		wg.Done()
	}("Горутина 1")

	go func(name string) {
        for i := 1; i <= 5; i++ {
            ch <- fmt.Sprintf("%s: %d", name, i)
        }
        wg.Done()
    }("Горутина 2")
    
	go func() {
        wg.Wait()
        close(ch)
    }()

    for msg := range ch {
        fmt.Println(msg)
    }
}
