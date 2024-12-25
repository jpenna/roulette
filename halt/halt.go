package halt

import (
	"sync/atomic"
	"time"

	"elem.com/roulette/utils"
)

var halted atomic.Bool

func Stop() {
	if halted.Load() {
		return
	}

	halted.Store(true)

	go func() {
		for {
			if !IsHalted() {
				break
			}

			utils.Console.Warn().Msg("Jogo interrompido, aguardando [c]ontinue...")
			time.Sleep(10 * time.Second)
		}
	}()
}

func Continue() {
	halted.Store(false)
}

func IsHalted() bool {
	return halted.Load()
}
