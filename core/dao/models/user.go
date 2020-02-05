/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 16:58
 */

package models

var Health = []string{
	"正常",
	"有发热、咳嗽、感冒等症状",
	"医院确定疑似",
	"医院确定感染新型肺炎",
}

type UserAccount struct {
	Base
	RealName string
	AliaName string
	Phone string
	Email string
	Community string
	BuildingNumber int
	BuildingUint int
	HouseNumber int
	Province string
	City string
}

type UserTemperature struct {
	Base
	AliaName string
	Temperature float64
	HealthStatus int
	Others string
}

