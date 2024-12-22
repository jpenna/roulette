package robot

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-vgo/robotgo"
)

type Window struct {
	topLeft     [2]int
	bottomRight [2]int
}

func (w *Window) Capture() {
	fmt.Println("Posicione o mouse no canto superior esquerdo da janela e pressione Enter")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	x, y := robotgo.Location()
	fmt.Println("Superior esquerdo: ", x, y)
	w.topLeft = [2]int{x, y}

	fmt.Println("Posicione o mouse no canto inferior direito da janela e pressione Enter")
	reader.ReadString('\n')

	x, y = robotgo.Location()
	fmt.Println("Inferior direito: ", x, y)
	w.bottomRight = [2]int{x, y}
}
