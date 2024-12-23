package robot

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/go-vgo/robotgo"
)

type Window struct {
	TopLeft     [2]int `json:"topLeft"`
	BottomRight [2]int `json:"bottomRight"`

	TerminalPosition [2]int

	ReadyBarPosition [2]int
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
}

func (w *Window) setReadyBarPosition(offset int) {
	w.ReadyBarPosition = [2]int{
		int(20 + float64(w.TopLeft[0])*1.1),        // 10% to the right (add 20 in case the coordinate is 0)
		int(float64(w.TopLeft[1])*1.5084) + offset, // Found position
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

func (w *Window) IsReadyToBet() (bool, error) {
	img, err := robotgo.CaptureImg(w.ReadyBarPosition[0], w.ReadyBarPosition[1], 1, 1)
	if err != nil {
		return false, fmt.Errorf("error capturing image: %w", err)
	}

	pixel := img.At(0, 0)
	r, g, b, _ := pixel.RGBA()

	greenMargin := g - 50 // 50 is the margin of error

	// Check if green component is significantly higher than red and blue
	return greenMargin > r && greenMargin > b, nil
}
