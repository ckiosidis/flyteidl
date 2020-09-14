/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type AdminPluginOverride struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `json:"task_type,omitempty"`
	// The unique name of the plugin which should handle tasks of this type instead of the default registered plugin.
	PluginId string `json:"plugin_id,omitempty"`
	// OPTIONAL The unique names of the plugin implementations, in order of decreasing preference, which should handle tasks of this type when the override plugin id is unavailable on the flyte platform.
	FallbackPluginIds []string `json:"fallback_plugin_ids,omitempty"`
	MissingPluginBehavior *PluginOverrideMissingPluginBehavior `json:"missing_plugin_behavior,omitempty"`
}
