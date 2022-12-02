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

type RoutesPoolRequestParams struct {
	CrushRootId int64 `json:"crush_root_id"`
	EcCodingChunks int64 `json:"ec_coding_chunks"`
	EcDataChunks int64 `json:"ec_data_chunks"`
	Name string `json:"name"`
	Replicated int64 `json:"replicated"`
	Type_ string `json:"type"`
}
