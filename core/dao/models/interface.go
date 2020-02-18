/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 14:28
 */

package models

type IStore interface {
	BeginTx() (IStore, error)
	Rollback() error
	CommitTx() error

	IUser
}

type IUser interface {
	CreateUserAccount(user *UserAccount) error
	CreateUserTemperature(tem *UserTemperature) error
	CreateUserInfo(userInfo *UserInfo) error
	GetUserAccountByAliaName(alia string) (*UserAccount, error)
	GetTemperatureByAliaName(alia string)([]*UserTemperature, error)
}