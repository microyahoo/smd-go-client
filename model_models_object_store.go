/*
 * SmdService
 *
 * Resource for managing storage
 *
 * API version: 1.0.0
 * Contact: canzhu@deeproute.ai
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package smdclient

import (
	"time"
)

type ModelsObjectStore struct {
	CreatedAt time.Time `json:"CreatedAt"`
	DeletedAt string `json:"DeletedAt"`
	ID int32 `json:"ID"`
	Name string `json:"Name"`
	PoolProperties []ModelsPoolProperty `json:"PoolProperties"`
	Port int32 `json:"Port"`
	PreservePoolsOnDelete bool `json:"PreservePoolsOnDelete"`
	Status string `json:"Status"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
