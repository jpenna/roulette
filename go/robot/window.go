package robot

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-vgo/robotgo"
)

type Window struct {
	TopLeft     [2]int `json:"topLeft"`
	BottomRight [2]int `json:"bottomRight"`
}

func (w *Window) Capture() {
	fmt.Println("Posicione o mouse no canto superior esquerdo da janela e pressione Enter")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	x, y := robotgo.Location()
	w.TopLeft = [2]int{x, y}

	fmt.Println("Posicione o mouse no canto inferior direito da janela e pressione Enter")
	reader.ReadString('\n')

	x, y = robotgo.Location()
	w.BottomRight = [2]int{x, y}

	fmt.Println("Superior esquerdo: ", w.TopLeft)
	fmt.Println("Inferior direito: ", w.BottomRight)
}
