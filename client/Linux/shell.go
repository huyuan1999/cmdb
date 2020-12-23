package Linux

import (
	"cmdb/utils"
	"regexp"
)

type smartctl struct {
}

type dmidecode struct {
}

func NewSmart() (*smartctl, error) {
	if err := utils.Which("smartctl"); err != nil {
		return nil, err
	} else {
		return &smartctl{}, nil
	}
}

func NewDmidecode() (*dmidecode, error) {
	if err := utils.Which("dmidecode"); err != nil {
		return nil, err
	} else {
		return &dmidecode{}, nil
	}
}

func (s *smartctl) re(device, compile string) string {
	result := utils.Shell("smartctl", "-i", device)
	reg := regexp.MustCompile(compile)
	matched := reg.FindStringSubmatch(result.Stdout)
	if len(matched) == 2 {
		return matched[1]
	}
	return ""
}

func (s *smartctl) Info(device string) string {
	result := utils.Shell("smartctl", "-i", device)
	return result.Stdout
}

func (s *smartctl) ModelFamily(device string) string {
	return utils.Trim(s.re(device, "Model\\s+Family:\\s+(.*)"))
}

func (s *smartctl) DeviceModel(device string) string {
	return utils.Trim(s.re(device, "Device\\s+Model:\\s+(.*)"))
}

func (s *smartctl) SerialNumber(device string) string {
	return utils.Trim(s.re(device, "Serial\\s+Number:\\s+(.*)"))
}

func (s *smartctl) FormFactor(device string) string {
	return utils.Trim(s.re(device, "Form\\s+Factor:\\s+(.*)"))
}

func (s *smartctl) RotationRate(device string) string {
	return utils.Trim(s.re(device, "Rotation\\s+Rate:\\s+(.*)"))
}

func (s *smartctl) UserCapacity(device string) string {
	return utils.Trim(s.re(device, "User\\s+Capacity:\\s+(.*)"))
}

func (d *dmidecode) Info() string {
	result := utils.Shell("dmidecode")
	return utils.Trim(result.Stdout)
}

func (d *dmidecode) SystemManufacturer() string {
	result := utils.Shell("dmidecode", "-s", "system-manufacturer")
	return utils.Trim(result.Stdout)
}

func (d *dmidecode) SystemProductName() string {
	result := utils.Shell("dmidecode", "-s", "system-product-name")
	return utils.Trim(result.Stdout)
}

func (d *dmidecode) SystemSerialNumber() string {
	result := utils.Shell("dmidecode", "-s", "system-serial-number")
	return utils.Trim(result.Stdout)
}

func (d *dmidecode) SystemUuid() string {
	result := utils.Shell("dmidecode", "-s", "system-uuid")
	return utils.Trim(result.Stdout)
}
