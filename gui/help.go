package gui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	helpHeaders      = []string{"KEY", "DESCRIPTION"}
	filesKeybindings = []map[string]string{
		{"tab": "focus to files"},
		{"j": "move next"},
		{"k": "move previous"},
		{"g": "move to top"},
		{"G": "move to bottom"},
		{"ctrl-b": "move previous page"},
		{"ctrl-f": "move next page"},
		{"h": "move to parent path"},
		{"l": "move to specified path"},
		{"y": "copy selected file or directory"},
		{"p": "paste file or directory"},
		{"d": "delete selected file or directory"},
		{"m": "make a new directory"},
		{"n": "make a new file"},
		{"r": "rename a directory or file"},
		{"e": "edit file with $EDITOR"},
		{"o": "open file or dierectory"},
		{"f or /": "search files or directories"},
		{"ctrl-j": "scroll preview panel down"},
		{"ctrl-k": "scroll preview panel up"},
		{"c or :": "focus cmdline"},
		{".": "edit config.yaml"},
		{"b": "bookmark directory"},
		{"B": "open bookmarks panel"},
	}

	pathKeybindings = []map[string]string{
		{"enter": "change directory"},
	}

	cmdKeybindings = []map[string]string{
		{"enter": "execute command"},
	}

	bookmarkKeybindings = []map[string]string{
		{"a": "add bookmark"},
		{"d": "delete bookmark"},
		{"q": "close bookmarks panel"},
		{"ctrl-g": "go to bookmark"},
		{"f or /": "search bookmarks"},
	}
)

type Help struct {
	*tview.Table
}

func NewHelp() *Help {
	h := &Help{
		Table: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1),
	}

	h.SetBorder(true).SetTitle("help").
		SetTitleAlign(tview.AlignLeft)

	return h
}

func (h *Help) UpdateView(panel Panel) {
	table := h.Table.Clear()

	for i, h := range helpHeaders {
		table.SetCell(0, i, &tview.TableCell{
			Text:            h,
			NotSelectable:   true,
			Align:           tview.AlignLeft,
			Color:           tcell.ColorYellow,
			BackgroundColor: tcell.ColorDefault,
		})
	}

	var keybindings []map[string]string

	switch panel {
	case PathPanel:
		keybindings = pathKeybindings
	case FilesPanel:
		keybindings = filesKeybindings
	case CmdLinePanel:
		keybindings = cmdKeybindings
	case BookmarkPanel:
		keybindings = bookmarkKeybindings
	}

	for i, keybind := range keybindings {
		for k, d := range keybind {
			table.SetCell(i+1, 0, tview.NewTableCell(k).SetTextColor(tcell.ColorWhite))
			table.SetCell(i+1, 1, tview.NewTableCell(d).SetTextColor(tcell.ColorWhite))
		}
	}
}