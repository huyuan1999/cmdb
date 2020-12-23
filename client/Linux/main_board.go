package Linux

import (
	"cmdb/st"
	"cmdb/utils"
)

type MainBoard struct {
	st.MainBoard
	dmidecode *dmidecode
}

func NewMainBoard() (*MainBoard, error) {
	mainBoard := &MainBoard{}
	dmidecode, err := NewDmidecode()
	if err != nil {
		return nil, err
	}
	mainBoard.dmidecode = dmidecode
	utils.Call(mainBoard)
	return mainBoard, nil
}

func (m *MainBoard) GetSerialNumber() {
	m.SerialNumber = m.dmidecode.SystemSerialNumber()
}

func (m *MainBoard) GetUUID() {
	m.UUID = m.dmidecode.SystemUuid()
}

func (m *MainBoard) GetManufacturer() {
	m.Manufacturer = m.dmidecode.SystemManufacturer()
}

func (m *MainBoard) GetProductName() {
	m.ProductName = m.dmidecode.SystemProductName()
}
