package models

import (
	"cmdb/st"
	"gorm.io/gorm"
)

type Host struct {
	gorm.Model
	st.Host
	MainBoardId uint
	MainBoard   MainBoard
}

type System struct {
	gorm.Model
	st.System
	MainBoardId uint
	MainBoard   MainBoard
}

type CPU struct {
	gorm.Model
	st.CPU
	MainBoardId uint
	MainBoard   MainBoard
}

type Memory struct {
	gorm.Model
	st.Memory
	MainBoardId uint
	MainBoard   MainBoard
}

type MainBoard struct {
	gorm.Model
	st.MainBoard
}

type NIC struct {
	gorm.Model
	st.NIC
	MainBoardId uint
	MainBoard   MainBoard
}

type Disk struct {
	gorm.Model
	st.Disk
	MainBoardId uint
	MainBoard   MainBoard
}
