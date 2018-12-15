import QtQuick 2.10				//ListView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//CustomListModel

Item {
    width: 640
    height: 480

    ColumnLayout {
        anchors.fill: parent

        ListView {
            id: listview
            height: 350
            Layout.fillWidth: true
            highlightFollowsCurrentItem: false
            /* Layout.fillHeight: true */

            model: CustomListModel{}
            delegate: Text {
                text: display
                MouseArea {
                    anchors.fill: parent
                    onClicked: listview.currentIndex = index;
                }
            }
            highlight: Rectangle {
                /* border.color: "yellow" */
                /* border.width: 3 */
                color: "steelblue"
                height: 12
                width: ListView.view.width
                y:  listview.currentItem.y      // highlighting direct binding to selected item!
                z: Infinity
                opacity: 0.25
            }
        }

        TextField {
            id: textField
            Layout.fillWidth: true
            height: 20
            text: qsTr("Text Input")
            font.pixelSize: 12
            activeFocusOnPress: true
        }

        Button {
            Layout.fillWidth: true
            height: 100
            text: "Update with Tags"
            onClicked: {
                console.log(textField.text);
                listview.model.change(textField.text);
            }
        }

    }
}
