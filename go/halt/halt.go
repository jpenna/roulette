package halt

import (
	"fmt"
	"sync/atomic"

	hook "github.com/robotn/gohook"
)

var IsHalted atomic.Bool

func ListenForHalt() {
	hook.Register(hook.KeyDown, []string{"s"}, func(e hook.Event) {
		fmt.Print("\033[41m")
		fmt.Print("\nJogada interrompida!")
		fmt.Print("\033[0m\n")

		IsHalted.Store(true)
	})

	s := hook.Start()

	<-hook.Process(s)
}

func StopListenForHalt() {
	hook.End()
}
