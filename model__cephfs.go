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

type Cephfs struct {
	DataPools []RoutesPoolRequestParams `json:"data_pools"`
	FsName string `json:"fs_name"`
	MetadataPool *RoutesPoolRequestParams `json:"metadata_pool"`
	Name string `json:"name"`
	PreserveFilesystemOnDelete bool `json:"preserve_filesystem_on_delete"`
	PreservePoolsOnDelete bool `json:"preserve_pools_on_delete"`
}
