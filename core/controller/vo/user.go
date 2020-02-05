/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 18:39
 */

package vo

type UserInfo struct {
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
	AliaName string
	Temperature float64
	HealthStatus int
	Others string
}