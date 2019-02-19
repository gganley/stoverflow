import QtQuick 2.10				//ListView
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//CustomListModel

Item {
    width: 960
    height: 480

    RowLayout {
        anchors.fill: parent

        ColumnLayout{
            id: leftColumn
            width: 160
            y: 0
            height: 480
            ListView {
                id: listview
                width: parent.width
                height: 300
                highlightFollowsCurrentItem: false
                /* Layout.fillHeight: true */

                model: CustomListModel {
                    id: listModel
                }
                delegate: Text {
                    text: display[0]
                    MouseArea {
                        anchors.fill: parent
                        onClicked: {
                            listview.currentIndex = index;
                            companyText.text = display[0];
                            postingDate.text = display[1];
                            locationText.text = display[2];
                            tagsText.text = display[3];
                            descriptionText.text = display[4];
                        }
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
                Layout.fillWidth: false
                width: parent.width
                height: 20
                text: qsTr("Text Input")
                font.pixelSize: 12
                activeFocusOnPress: true
            }

            Button {
                height: 100
                width: parent.width
                text: "Update with Tags"
                onClicked: {
                    console.log(textField.text);
                    listview.model.change(textField.text);
                }
            }
        }

        ColumnLayout {
            id: centerColumn
            x: 160
            y: 0
            width: 160
            height: 480

            Text{
                id: locationText
                text: qsTr("Text")
            }

            Text{
                id: companyText
                text: qsTr("Text")
            }

            Text{
                id: postingDate
                text: qsTr("Text Area")
            }
            Text{
                id: tagsText
                width: parent.width
                wrapMode: Text.WordWrap
                text: qsTr("Text Area")
            }
        }
        ColumnLayout {
            id: rightColumn
            x: 320
            y: 0
            width: 320
            height: 480
            Flickable {
                id: scrollView
                clip: true
                Layout.fillHeight: true
                Layout.fillWidth: true
                contentWidth: parent.width
                contentHeight: descriptionText.height
                Text{
                    id: descriptionText
                    width: parent.width
                    /* height: parent.height */
                    text: qsTr("Text Area")
                    textFormat: Text.StyledText
                    wrapMode: Text.WordWrap
                    onLinkActivated: Qt.openUrlExternally(link)

                }
            }
        }
    }
}
