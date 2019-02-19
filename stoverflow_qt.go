package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
	"github.com/sirupsen/logrus"
	sd "github.com/gganley/stoverflow_data"
	"strings"
)

var Log = logrus.New()
func init() {
	CustomListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomListModel")
	Log.Out = os.Stderr
	Log.Level = logrus.InfoLevel
}

// type ListItem struct {
// 	firstName string 
// 	lastName  string
// }
// type Job struct {
// 	company_name string
// 	publish_date string
// 	location string
// 	tags []string
// }

type CustomListModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ func()                                  `signal:"remove,auto"`
	// _ func() `signal:"get,auto"`
	_ func(obj []*core.QVariant)              `signal:"add,auto"`
	_ func(company_name, publish_date, location string, tags []string, description string) `signal:"edit,auto"`
	_ func(item *core.QVariant) `signal:"change,auto"`
	modelData []sd.Job
}

func (m *CustomListModel) init() {
	m.modelData = sd.GetData([]string{}, []string{})

	m.ConnectRowCount(m.rowCount)
	m.ConnectData(m.data)

	// REMEMBER TO CONNECT YOUR SIGNALS
	// Q: why is it saying the function is not defined when I have it in my source code
}

// func (m *CustomListModel) get(i int) *core.QVariant {
// 	ret := []string{m.modelData[i].CompanyName, m.modelData[i].PublishDate, m.modelData[i].Location, strings.Join(m.modelData[i].Tags, ",")}
// 	return core.NewQVariant19(ret)
// }

func (m *CustomListModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

// Why on God's green earth do you call this DATA when it's linked to DISPLAY in QML
func (m *CustomListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := m.modelData[index.Row()]
	// This is what is bound to `display` in main.qml::Item::RowLayout::ColumnLayout::ListView
	return core.NewQVariant19([]string{item.CompanyName, item.PublishDate, item.Location, strings.Join(item.Tags, ","), item.Description})
}

func (m *CustomListModel) remove() {
	if len(m.modelData) == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.modelData)-1, len(m.modelData)-1)
	m.modelData = m.modelData[:len(m.modelData)-1]
	m.EndRemoveRows()
}

func (m *CustomListModel) add(item []*core.QVariant) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, sd.Job{item[0].ToString(), item[1].ToString(), item[2].ToString(), item[3].ToStringList(), item[4].ToString()})
	m.EndInsertRows()
}

func (m *CustomListModel) edit(company_name, publish_date, location string, tags []string, description string) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[len(m.modelData)-1] = sd.Job{company_name, publish_date, location, tags, description}
	m.DataChanged(m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()), []int{int(core.Qt__DisplayRole)})
}

func (m *CustomListModel) change(item *core.QVariant) {
	spliced_tags := strings.Split(item.ToString(), ",")
	Log.Info(fmt.Sprintf("%+v\n", spliced_tags))
	positiveTags := make([]string, 0)
	negativeTags := make([]string, 0)

	for _, tag := range spliced_tags {
		if strings.HasPrefix(tag, "-") {
			negativeTags = append(negativeTags, tag)
		} else {
			positiveTags = append(positiveTags, tag)
		}
	}
	jobdata := sd.GetData(positiveTags, negativeTags)
	m.BeginResetModel()
	m.modelData = jobdata
	m.EndResetModel()
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetTitle("listview Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	app.Exec()
}
