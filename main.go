package botmain

import (
	bot "qme/Bot"
	"sync"
)

func TestingStart() {
	main()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		go bot.Init()
	}()

	wg.Wait()
}
