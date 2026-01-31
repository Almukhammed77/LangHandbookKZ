package concurrency

import (
	"fmt"
	"sync"
)

var viewsMap = make(map[uint]int)

var viewsMutex sync.RWMutex

var viewChannel = make(chan uint, 100) // буфер на 100, чтобы не ждать долго

func StartViewCounter() {
	fmt.Println("Запускаю фоновый счётчик просмотров...")

	go func() {
		for langID := range viewChannel {
			viewsMutex.Lock()

			viewsMap[langID] = viewsMap[langID] + 1
			viewsMutex.Unlock()

		}
		fmt.Println("Канал закрыт, счётчик просмотров остановлен")
	}()

	fmt.Println("Фоновый счётчик просмотров запущен")
}

func AddView(langID uint) {
	select {
	case viewChannel <- langID:
	default:
		fmt.Println("Канал переполнен, пропускаем просмотр языка", langID)
	}
}

func GetViewsCount(langID uint) int {
	viewsMutex.RLock()
	defer viewsMutex.RUnlock()

	count, exists := viewsMap[langID]
	if !exists {
		return 0
	}
	return count
}
