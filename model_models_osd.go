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

type ModelsOsd struct {
	ClusterID int32 `json:"ClusterID"`
	CreatedAt time.Time `json:"CreatedAt"`
	CrushRootID int32 `json:"CrushRootID"`
	DeletedAt string `json:"DeletedAt"`
	Disk *ModelsDisk `json:"Disk"`
	DiskID int32 `json:"DiskID"`
	ID int32 `json:"ID"`
	In bool `json:"In"`
	Name string `json:"Name"`
	OsdID int32 `json:"OsdID"`
	Status string `json:"Status"`
	Up bool `json:"Up"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
