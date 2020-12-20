package models

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	Address     string
	Port        uint
	UserName    string
	AuthType    string `gorm:"type:enum('password','key');default:password"`
	Password    string
	SecretKey   string
	MainBoardId uint
	MainBoard   MainBoard
}

type System struct {
	HostName    string
	Release     string
	Kernel      string
	MainBoardId uint
	MainBoard   MainBoard
}

type CPU struct {
	Number      uint   // 物理 CPU 数量
	Core        uint   // cup 核心数
	Sibling     uint   // cpu 线程数
	Processor   uint   // 服务器总线程数 = Number * Siblings
	ModelName   string // cpu 详细信息
	MainBoardId uint
	MainBoard   MainBoard
}

type Memory struct {
	Total       uint   // 总内存大小
	Type        string // 内存条型号
	Number      uint   // 内存条数量
	Slot        uint   // 内存插槽数量
	MaxSize     string // 主板支持的最大内存大小
	FreeSlot    uint   // 剩余可用插槽
	MainBoardId uint
	MainBoard   MainBoard
}

// 主板信息
type MainBoard struct {
	// 一对一关联 host
	SerialNumber string
	UUID         string
	Manufacturer string // 服务器厂商
	ProductName  string // 服务器型号
	//Manufacturer: Dell Inc.
	//Product Name: PowerEdge C6220
	//Serial Number: B2BRKY1

	//Physical Memory Array
	//Location: System Board Or Motherboard
	//Use: System Memory
	//Error Correction Type: Multi-bit ECC
	//Maximum Capacity: 512 GB
	//Error Information Handle: 0x0034
	//Number Of Devices: 16

	//Memory Device
	//Array Handle: 0x001B
	//Error Information Handle: 0x002C
	//Total Width: 72 bits
	//Data Width: 72 bits
	//Size: 8192 MB
	//Form Factor: DIMM
	//Set: None
	//Locator: DIMM_A1
	//Bank Locator: CPU1
	//Type: DDR3
	//Type Detail: Synchronous Registered (Buffered)
	//Speed: Unknown
	//Manufacturer: 2C00B3002C00
	//Serial Number: DFC55F35
	//Asset Tag: 0F151463
	//Part Number: 36JSF1G72PZ-1G9K1
	//Rank: 2
	//Configured Clock Speed: 1333 MHz
}

type Nic struct {
	// 一对多关联 host
	// /proc/net/dev 获取网卡名称
	// https://blog.csdn.net/ygm_linux/article/details/24661839 获取IP
	Name        string
	Mac         string
	Address     string
	MainBoardId uint
	MainBoard   MainBoard
}

type Disk struct {
	// 一对多关联 host
	gorm.Model
	Name         string
	SerialNumber string
	Manufacturer string
	Size         uint
	FormFactor   uint
	MainBoardId  uint
	MainBoard    MainBoard
}
