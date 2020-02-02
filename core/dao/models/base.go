/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 17:00
 */

package models

import "time"

type Base struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}