package client

import (
	"cmdb/models"
	"cmdb/utils"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type CPU struct {
	models.CPU
	document string
}

func NewCPU(path string) (*CPU, error) {
	cpu := &CPU{}
	text, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cpu.document = string(text)
	return cpu, nil
}

func (c *CPU) GetNumber() {
	compile, err := regexp.Compile("physical id.*")
	if err != nil {
		return
	}
	physical := compile.FindAllString(c.document, -1)
	c.Number = uint(len(utils.RemoveDuplicate(physical)))
}

func (c *CPU) GetCore() {
	matched, err := utils.LoopMatchString(c.document, []string{"cpu cores.*", "\\d+"})
	if err != nil {
		return
	}

	if core, err := strconv.Atoi(matched); err == nil {
		c.Core = uint(core)
	}
}

func (c *CPU) GetSibling() {
	matched, err := utils.LoopMatchString(c.document, []string{"siblings.*", "\\d+"})
	if err != nil {
		return
	}

	if sibling, err := strconv.Atoi(matched); err == nil {
		c.Sibling = uint(sibling)
	}
}

func (c *CPU) GetProcessor() {
	c.GetNumber()
	c.GetSibling()
	c.Processor = c.Number * c.Sibling
}

func (c *CPU) GetModelName() {
	compile, err := regexp.Compile("model name.*")
	if err != nil {
		return
	}
	modelName := compile.FindString(c.document)

	modelArr := strings.Split(modelName, ":")
	if len(modelArr) == 2 {
		c.ModelName = strings.Trim(modelArr[1], " ")
	}
}
