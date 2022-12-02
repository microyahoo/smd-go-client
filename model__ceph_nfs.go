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

type CephNfs struct {
	Instances int64 `json:"instances"`
	Name string `json:"name"`
	NfsName string `json:"nfs_name"`
}
