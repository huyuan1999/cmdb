package Linux

import (
	"cmdb/st"
	"cmdb/utils"
	"errors"
	"github.com/shirou/gopsutil/v3/disk"
	"golang.org/x/sys/unix"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"syscall"
	"unsafe"
)

type Disk []struct {
	st.Disk
	diskNameArr []string
}

const diskstats = "/proc/diskstats"

func isDisk(array []string) (diskArray []string) {
	for _, base := range array {
		for _, dev := range array {
			if dev == base {
				continue
			}
			if strings.HasPrefix(dev, base) {
				diskArray = append(diskArray, base)
				goto loop
			}
		}
	loop:
	}
	return
}

func NewDisk() (Disk, error) {
	text, err := ioutil.ReadFile(diskstats)
	if err != nil {
		return nil, err
	}
	compile, err := regexp.Compile("\\s+8\\s+\\d+\\s+\\w+")
	if err != nil {
		return nil, err
	}

	diskArr := compile.FindAllString(string(text), -1)
	var nameArr []string
	for _, dev := range diskArr {
		split := strings.Split(utils.DeleteExtraSpace(utils.Trim(dev)), " ")
		if len(split) >= 3 {
			name := path.Join("/dev/", split[2])
			nameArr = append(nameArr, name)
		}
	}
	diskNameArr := isDisk(nameArr)
	diskNumber := len(diskNameArr)
	if diskNumber < 1 {
		return nil, errors.New("获取硬盘信息错误: 硬盘数量小于 1")
	}

	d := make(Disk, diskNumber)
	d[0].diskNameArr = diskNameArr
	utils.Call(d)
	return d, nil
}

func (d Disk) GetName() {
	for index, dev := range d[0].diskNameArr {
		d[index].Name = dev
	}
}

func (d Disk) GetSerialNumber() {
	for index, dev := range d[0].diskNameArr {
		serial, err := disk.SerialNumber(dev)
		if err != nil || serial == "" {
			continue
		}
		s := strings.Split(serial, "_")
		sn := s[len(s)-1]
		d[index].SerialNumber = sn
	}
}

func (d Disk) GetManufacturer() {
	for index, dev := range d[0].diskNameArr {
		serial, err := disk.SerialNumber(dev)
		if err != nil || serial == "" {
			continue
		}
		s := strings.Split(serial, "_")
		if (len(s) - 2) >= 0 {
			sn := s[0 : len(s)-2]
			d[index].Manufacturer = strings.Join(sn, "_")
		}
	}
}

func (d Disk) GetSize() {
	var size uint64
	for index, dev := range d[0].diskNameArr {
		fd, err := unix.Open(dev, os.O_RDONLY, 0660)
		if err != nil {
			continue
		}
		_, _, ErrOn := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), unix.BLKGETSIZE64, uintptr(unsafe.Pointer(&size)))
		if unix.ErrnoName(ErrOn) != "" {
			continue
		}
		d[index].Size = uint(size >> 30)
	}
}

func (d Disk) GetFormFactor() {

}
