package robot

func Screen() {
	window := Window{}

	window.Capture()

	MoveTo(window.topLeft[0], window.topLeft[1])
}
