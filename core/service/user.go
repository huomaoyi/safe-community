/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 18:49
 */

package service

import (
	"safe-community/core/controller/vo"
	"safe-community/core/dao/models"
	"safe-community/core/dao/mysql"
)

func UserInfoGet(alia string) (*vo.UserInfo, error) {
	a, err := mysql.SingleStore().GetUserAccountByAliaName(alia)
	if err != nil {
		return nil, err
	}

	return &vo.UserInfo{
		RealName: a.RealName,
		AliaName: a.AliaName,
		Phone: a.Phone,
		Email: a.Email,
		Community: a.Community,
		BuildingNumber: a.BuildingNumber,
		BuildingUint: a.BuildingUint,
		HouseNumber: a.HouseNumber,
		Province: a.Province,
		City: a.City,
	}, nil
}

func UserSignupPost(info vo.UserInfo) error {
	acc := &models.UserAccount{
		RealName: info.RealName,
		AliaName: info.AliaName,
		Phone: info.Phone,
		Email: info.Email,
		Community:info.Community,
		BuildingNumber: info.BuildingNumber,
		BuildingUint: info.BuildingUint,
		HouseNumber: info.HouseNumber,
		Province: info.Province,
		City: info.City,
	}

	return mysql.SingleStore().CreateUserAccount(acc)
}

func UserTemperaturePost(user vo.UserTemperature) error {
	t := &models.UserTemperature{
		AliaName:    user.AliaName,
		Temperature: user.Temperature,
	}

	return mysql.SingleStore().CreateUserTemperature(t)
}

func UserTemperatureGet(alia string) ([]float64, error) {
	temperatures, err := mysql.SingleStore().GetTemperatureByAliaName(alia)
	if err != nil {
		return nil, err
	}

	var temps []float64
	temps = make([]float64, len(temperatures))
	for i, t := range temperatures {
		temps[i] = t.Temperature
	}

	return temps, nil
}