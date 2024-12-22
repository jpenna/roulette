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

	TerminalPosition [2]int `json:"terminalPosition"`
}

func (w *Window) Capture() {
	fmt.Print("Posicione o mouse no canto superior esquerdo da janela e pressione Enter")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	x, y := robotgo.Location()
	w.TopLeft = [2]int{x, y}

	fmt.Print("Posicione o mouse no canto inferior direito da janela e pressione Enter")
	reader.ReadString('\n')

	x, y = robotgo.Location()
	w.BottomRight = [2]int{x, y}

	fmt.Println("Superior esquerdo: ", w.TopLeft)
	fmt.Println("Inferior direito: ", w.BottomRight)
}

func (w *Window) CaptureTerminal() {
	fmt.Print("Posicione o mouse no TERMINAL e pressione Enter")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	x, y := robotgo.Location()
	w.TerminalPosition = [2]int{x, y}

	fmt.Println("Terminal position: ", w.TerminalPosition)
}

func (w *Window) ClickTerminal() {
	Click(w.TerminalPosition[0], w.TerminalPosition[1])
}
