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

type ModelsDisk struct {
	Cache bool `json:"Cache"`
	CreatedAt time.Time `json:"CreatedAt"`
	DeletedAt string `json:"DeletedAt"`
	DeviceModel string `json:"DeviceModel"`
	Empty bool `json:"Empty"`
	Encrypted bool `json:"Encrypted"`
	Filesystem string `json:"Filesystem"`
	HasChildren bool `json:"HasChildren"`
	Host *ModelsHost `json:"Host"`
	HostID int64 `json:"HostID"`
	ID int64 `json:"ID"`
	IsRoot bool `json:"IsRoot"`
	MBytes int64 `json:"MBytes"`
	Name string `json:"Name"`
	Osd *ModelsOsd `json:"Osd"`
	PathID string `json:"PathID"`
	RealPath string `json:"RealPath"`
	Removed bool `json:"Removed"`
	Rotational bool `json:"Rotational"`
	Serial string `json:"Serial"`
	Type_ string `json:"Type"`
	UUID string `json:"UUID"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	Vendor string `json:"Vendor"`
	WWN string `json:"WWN"`
}
