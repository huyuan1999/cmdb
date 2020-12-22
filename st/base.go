package st

type Host struct {
	Address   string
	Port      uint
	UserName  string
	AuthType  string `gorm:"type:enum('password','key');default:password"`
	Password  string
	SecretKey string
}

type System struct {
	HostName string
	Release  string
	Kernel   string
}

type CPU struct {
	Number      uint   // 物理 CPU 数量
	Core        uint   // cup 核心数
	Sibling     uint   // cpu 线程数
	Processor   uint   // 服务器总线程数 = Number * Siblings
	ModelName   string // cpu 详细信息
}

type Memory struct {
	Total       uint   // 总内存大小
	Type        string // 内存条型号
	Number      uint   // 内存条数量
	Slot        uint   // 内存插槽数量
	MaxSize     string // 主板支持的最大内存大小
	FreeSlot    uint   // 剩余可用插槽
}

type MainBoard struct {
	SerialNumber string
	UUID         string
	Manufacturer string // 服务器厂商
	ProductName  string // 服务器型号
}

type NIC struct {
	Name    string
	Mac     string
	Address []string
}

type Disk struct {
	Name         string
	SerialNumber string
	Manufacturer string
	Size         uint
	//FormFactor   uint
}
