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

type ModelsCephCluster struct {
	CreatedAt time.Time `json:"CreatedAt"`
	DeletedAt string `json:"DeletedAt"`
	ID int64 `json:"ID"`
	Name string `json:"Name"`
	Namespace string `json:"Namespace"`
	Osds []ModelsOsd `json:"Osds"`
	State string `json:"State"`
	Status string `json:"Status"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
