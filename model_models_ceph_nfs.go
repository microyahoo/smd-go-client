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

type ModelsCephNfs struct {
	CreatedAt time.Time `json:"CreatedAt"`
	DeletedAt string `json:"DeletedAt"`
	ID int64 `json:"ID"`
	Instances int64 `json:"Instances"`
	NFSName string `json:"NFSName"`
	Name string `json:"Name"`
	Passive bool `json:"Passive"`
	Status string `json:"Status"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
