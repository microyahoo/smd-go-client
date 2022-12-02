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

type OsUser struct {
	BucketCapability string `json:"bucket_capability"`
	BucketQuotaMaxBuckets int32 `json:"bucket_quota_max_buckets"`
	BucketQuotaMaxObjects int64 `json:"bucket_quota_max_objects"`
	BucketQuotaMaxSize string `json:"bucket_quota_max_size"`
	MetadataCapability string `json:"metadata_capability"`
	Name string `json:"name"`
	ObjectStoreId int64 `json:"object_store_id"`
	UsageCapability string `json:"usage_capability"`
	UserCapability string `json:"user_capability"`
	UserName string `json:"user_name"`
	ZoneCapability string `json:"zone_capability"`
}
