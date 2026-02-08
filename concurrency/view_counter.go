package concurrency

import (
	"fmt"
	"sync"

	"github.com/Almukhammed77/LangHandbookKZ/storage"
)

var viewsMap = make(map[uint]int)
var viewsMutex sync.RWMutex
var viewChannel = make(chan uint, 100)

func StartViewCounter() {
	fmt.Println("Background view counter is enabled")
	go func() {
		for langID := range viewChannel {
			viewsMutex.Lock()
			viewsMap[langID]++
			viewsMutex.Unlock()
			storage.UpdateViews(langID, viewsMap[langID])
		}
	}()
}

func AddView(langID uint) {
	viewChannel <- langID
}

func GetViewsCount(langID uint) int {
	viewsMutex.RLock()
	defer viewsMutex.RUnlock()
	return viewsMap[langID]
}
