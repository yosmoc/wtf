package jenkins

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("j", widget.next)
	widget.SetKeyboardChar("k", widget.prev)
	widget.SetKeyboardChar("o", widget.openJob)
	widget.SetKeyboardChar("r", widget.Refresh)

	widget.SetKeyboardKey(tcell.KeyDown, widget.next)
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openJob)
	widget.SetKeyboardKey(tcell.KeyEsc, widget.unselect)
	widget.SetKeyboardKey(tcell.KeyUp, widget.prev)
}
