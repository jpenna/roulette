package robot

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

type Window struct {
	TopLeft     [2]int `json:"topLeft"`
	BottomRight [2]int `json:"bottomRight"`

	TerminalPosition [2]int

	ReadyBarPosition [2]int

	NumberArea image.Rectangle
	WinArea    image.Rectangle
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
}

func (w *Window) SetReadyBarPosition(offset int) {
	height := w.BottomRight[1] - w.TopLeft[1]

	w.ReadyBarPosition = [2]int{
		int(20 + float64(w.TopLeft[0])*1.1),                 // 10% to the right (add 20 in case the coordinate is 0)
		w.TopLeft[1] + int(float64(height)*0.5084) + offset, // Found position
	}

	w.confirmReadyBar()
}

func (w *Window) confirmReadyBar() {
	MoveTo(w.ReadyBarPosition[0], w.ReadyBarPosition[1])

	fmt.Print("O mouse está sobre a barra de PRONTO? (`y` para continuar): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1] // Remove newline

	if input == "y" {
		fmt.Println("Barra de PRONTO encontrada")
		return
	}

	offset, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Por favor digite 'y' para confirmar ou um número para ajustar a posição da barra (+ para abaixar e - para subir)")
		w.SetReadyBarPosition(offset)
		return
	}

	w.SetReadyBarPosition(offset)
}

func (w *Window) SetNumberAreas() {
	width := w.BottomRight[0] - w.TopLeft[0]
	height := w.BottomRight[1] - w.TopLeft[1]

	topLeftX := w.TopLeft[0] + int(float64(width)*0.47)
	topLeftY := w.TopLeft[1] + int(float64(height)*0.27)

	bottomRightX := w.TopLeft[0] + int(float64(width)*0.52)
	bottomRightY := w.TopLeft[1] + int(float64(height)*0.35)

	w.NumberArea = image.Rectangle{
		Min: image.Point{X: topLeftX, Y: topLeftY},
		Max: image.Point{X: bottomRightX, Y: bottomRightY},
	}

	diffY := bottomRightY - topLeftY
	w.WinArea = image.Rectangle{
		Min: image.Point{X: topLeftX, Y: topLeftY - diffY},
		Max: image.Point{X: bottomRightX, Y: bottomRightY - diffY},
	}

	for {
		// window
		MoveTo(w.TopLeft[0], w.TopLeft[1])
		time.Sleep(1 * time.Second)
		MoveTo(w.BottomRight[0], w.BottomRight[1])
		time.Sleep(1 * time.Second)

		// number
		MoveTo(w.NumberArea.Min.X, w.NumberArea.Min.Y)
		time.Sleep(1 * time.Second)
		MoveTo(w.NumberArea.Max.X, w.NumberArea.Max.Y)
		time.Sleep(1 * time.Second)

		// win
		MoveTo(w.WinArea.Min.X, w.WinArea.Min.Y)
		time.Sleep(1 * time.Second)
		MoveTo(w.WinArea.Max.X, w.WinArea.Max.Y)
		time.Sleep(1 * time.Second)
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
