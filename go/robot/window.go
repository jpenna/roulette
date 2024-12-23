package robot

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"

	"github.com/go-vgo/robotgo"
)

type Window struct {
	TopLeft     [2]int `json:"topLeft"`
	BottomRight [2]int `json:"bottomRight"`

	TerminalPosition [2]int

	ReadyBarPosition [2]int

	NumberArea image.Rectangle
}

func (w *Window) CaptureSize() {
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

	w.setReadyBarPosition(0)
	w.setNumberArea()
}

func (w *Window) setReadyBarPosition(offset int) {
	height := w.BottomRight[1] - w.TopLeft[1]

	w.ReadyBarPosition = [2]int{
		int(20 + float64(w.TopLeft[0])*1.1),  // 10% to the right (add 20 in case the coordinate is 0)
		int(float64(height)*0.5084) + offset, // Found position
	}

	MoveTo(w.ReadyBarPosition[0], w.ReadyBarPosition[1])

	fmt.Print("O mouse está sobre a barra de PRONTO? (`y` para continuar): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1] // Remove newline

	if input == "y" {
		fmt.Println("Barra de PRONTO encontrada")
		return
	}

	if input == "y" {
		return
	}

	offset, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Por favor digite 'y' para confirmar ou um número para ajustar a posição da barra (+ para abaixar e - para subir)")
		w.setReadyBarPosition(offset)
		return
	}

	w.setReadyBarPosition(offset)
}

func (w *Window) setNumberArea() {
	width := w.BottomRight[0] - w.TopLeft[0]
	height := w.BottomRight[1] - w.TopLeft[1]

	topLeftX := float64(width) * 0.473919523
	topLeftY := float64(height) * 0.276257723

	bottomRightX := float64(width) * 0.52260308
	bottomRightY := float64(height) * 0.355692851

	w.NumberArea = image.Rectangle{
		Min: image.Point{X: int(topLeftX), Y: int(topLeftY)},
		Max: image.Point{X: int(bottomRightX), Y: int(bottomRightY)},
	}
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
