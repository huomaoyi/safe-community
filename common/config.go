/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 14:22
 */

package common

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"
)

const (
	FileDir = "/Users/liteng/dev/src/github.com/huomaoyi/safe-community/config/"
)

var conf *config
var once sync.Once

//GetConfig ...
func GetConfig() *config {
	once.Do(func() {
		conf = setConfig(FileDir + "conf.ini")
		conf.readList()
	})
	return conf
}

type config struct {
	filePath string
	confList []map[string]map[string]string
}

func setConfig(path string) *config {
	c := new(config)
	c.filePath = path
	return c
}

//GetValue ...
func (c *config) GetValue(section, name string) string {
	for _, v := range c.confList {
		for key, value := range v {
			if key == section {
				return value[name]
			}
		}
	}
	return ""
}

func (c *config) GetSlice(section, name string) []string {
	res := c.GetValue(section, name)
	if len(res) > 0 {
		return strings.Split(res, ",")
	}
	return []string{}
}

func (c *config) readList() []map[string]map[string]string {
	file, err := os.Open(c.filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var data map[string]map[string]string
	var section string
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				continue
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case string(line[0]) == "#":
		case line[0] == '[' && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1 : len(line)-1])
			data = make(map[string]map[string]string)
			data[section] = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			if i == -1 {
				continue
			}
			value := strings.TrimSpace(line[i+1:])
			data[section][strings.TrimSpace(line[0:i])] = value
			if c.uniqueAppend(section) == true {
				c.confList = append(c.confList, data)
			}
		}
	}
	return c.confList
}

func (c *config) uniqueAppend(conf string) bool {
	for _, v := range c.confList {
		for k, _ := range v {
			if k == conf {
				return false
			}
		}
	}
	return true
}