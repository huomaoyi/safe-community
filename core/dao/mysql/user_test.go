/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 17:07
 */

package mysql

import (
	"log"
	"safe-community/core/dao/models"
	"testing"
)

func TestStore_CreateUserAccount(t *testing.T) {

	userInfo := &models.UserAccount{
		AliaName:"lengfeng2",
		Community:"展春园",
		BuildingNumber:15,
		BuildingUint:5,
		HouseNumber:301,
		Province:"beijing",
		City:"beijing",
	}

	err := SingleStore().CreateUserAccount(userInfo)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStore_CreateUserTemperature(t *testing.T) {

	temp := &models.UserTemperature{
		AliaName:"冷锋$^_^",
		Temperature: 36.3,
	}

	log.Fatal(SingleStore().CreateUserTemperature(temp))
}

func TestStore_GetUserAccountByAliaName(t *testing.T) {
	userInfo, err := SingleStore().GetUserAccountByAliaName("lengfeng")
	if err != nil {
		t.Fatal(err)
	}

	log.Println(userInfo)
}

func TestStore_GetTemperatureByAliaName(t *testing.T) {
	temperature, err := SingleStore().GetTemperatureByAliaName("冷锋$^_^")
	if err != nil {
		t.Fatal(err)
	}

	log.Println(temperature)
}