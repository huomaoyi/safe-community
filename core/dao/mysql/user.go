/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 17:04
 */

package mysql

import "safe-community/core/dao/models"

func (s *Store) CreateUserAccount(userAccount *models.UserAccount) error {
	return s.db.Create(userAccount).Error
}

func (s *Store) CreateUserInfo(userInfo *models.UserInfo) error {
	return s.db.Create(userInfo).Error
}

func (s *Store) CreateUserTemperature(tem *models.UserTemperature) error {
	return s.db.Create(tem).Error
}

func (s *Store) GetUserAccountByAliaName(alia string) (*models.UserAccount, error) {
	var acc models.UserAccount
	err := s.db.Where("alia_name=?", alia).Find(&acc).Error
	if err != nil {
		return nil, err
	}

	return &acc, nil
}

func (s *Store) GetTemperatureByAliaName(alia string)([]*models.UserTemperature, error) {
	var temps []*models.UserTemperature
	err := s.db.Where("alia_name=?", alia).Find(&temps).Error
	if err != nil {
		return nil, err
	}

	return temps, nil
}