package main

import (
	"fmt"
	"github.com/andlabs/ui"
	"strconv"
)

type modelHandler struct {
}

func newModelHandler() *modelHandler {
	// You can throw meta data in here
	m := new(modelHandler)
	return m
}

func (mh *modelHandler) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return []ui.TableValue{
		ui.TableString("Job"),		// column 0 text
		ui.TableString("Location"),		// column 1 text
		ui.TableString("Company"),		// column 2 text
	}
}

func (mh *modelHandler) NumRows(m *ui.TableModel) int {
	return 15
}


func (mh *modelHandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	fmt.Println(strconv.Itoa(row), strconv.Itoa(column))
	switch column {
	case 0:
		return ui.TableString(fmt.Sprintf("Row %d", row))
	case 2:
		return ui.TableString("Editing this won't change anything")
	case 1:
		return ui.TableString("Colors!")
	}
	panic("At the disco")
}

func (mh *modelHandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
}

func setupUI() {
	mainwin := ui.NewWindow("libui Control Gallery", 640, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	mh := newModelHandler()
	model := ui.NewTableModel(mh)

	table := ui.NewTable(&ui.TableParams{
		Model:	model,
		RowBackgroundColorModelColumn:	-1,
	})

	mainwin.SetChild(table)
	mainwin.SetMargined(true)

	table.AppendTextColumn("Job Title",
		0, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("Company",
		1, ui.TableModelColumnNeverEditable, nil);
	table.AppendTextColumn("Location",
		2, ui.TableModelColumnNeverEditable, nil)

	mainwin.Show()
}

func main() {
	err := ui.Main(setupUI)
	if err != nil {
		panic(err)
	}
}

