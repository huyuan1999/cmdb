package Linux

import (
	"cmdb/st"
	"cmdb/utils"
	"strings"
)

type MainBoard struct {
	st.MainBoard
	dmidecode string
}

func NewMainBoard() (*MainBoard, error) {
	mainBoard := &MainBoard{}
	dmidecode, err := utils.Dmidecode()
	if err != nil {
		return nil, err
	}
	mainBoard.dmidecode = dmidecode
	utils.Call(mainBoard)
	return mainBoard, nil
}

func (m *MainBoard) GetSerialNumber() {
	matched, err := utils.LoopMatchString(m.dmidecode, []string{"(?s)(?U)System\\s+Information\\n+.*Handle", "Serial\\s+Number:\\s+.*"})
	if err != nil {
		return
	}
	split := strings.Split(matched, ":")
	if len(split) == 2 {
		m.SerialNumber = utils.Trim(split[1])
	}
}

func (m *MainBoard) GetUUID() {
	matched, err := utils.LoopMatchString(m.dmidecode, []string{"(?s)(?U)System\\s+Information\\n+.*Handle", "UUID:\\s+.*"})
	if err != nil {
		return
	}
	split := strings.Split(matched, ":")
	if len(split) == 2 {
		m.UUID = utils.Trim(split[1])
	}
}

func (m *MainBoard) GetManufacturer() {
	matched, err := utils.LoopMatchString(m.dmidecode, []string{"(?s)(?U)System\\s+Information\\n+.*Handle", "Manufacturer:\\s+.*"})
	if err != nil {
		return
	}
	split := strings.Split(matched, ":")
	if len(split) == 2 {
		m.Manufacturer = utils.Trim(split[1])
	}
}

func (m *MainBoard) GetProductName() {
	matched, err := utils.LoopMatchString(m.dmidecode, []string{"(?s)(?U)System\\s+Information\\n+.*Handle", "Product\\s+Name:\\s+.*"})
	if err != nil {
		return
	}
	split := strings.Split(matched, ":")
	if len(split) == 2 {
		m.ProductName = utils.Trim(split[1])
	}
}
