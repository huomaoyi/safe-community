/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 16:58
 */

package models

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
}