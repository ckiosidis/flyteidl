/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// A Workflow graph Node. One unit of execution in the graph. Each node can be linked to a Task, a Workflow or a branch node.
type CoreNode struct {
	// A workflow-level unique identifier that identifies this node in the workflow. \"inputs\" and \"outputs\" are reserved node ids that cannot be used by other nodes.
	Id string `json:"id,omitempty"`
	// Extra metadata about the node.
	Metadata *CoreNodeMetadata `json:"metadata,omitempty"`
	// Specifies how to bind the underlying interface's inputs. All required inputs specified in the underlying interface must be fullfilled.
	Inputs []CoreBinding `json:"inputs,omitempty"`
	// +optional Specifies execution depdendency for this node ensuring it will only get scheduled to run after all its upstream nodes have completed. This node will have an implicit depdendency on any node that appears in inputs field.
	UpstreamNodeIds []string `json:"upstream_node_ids,omitempty"`
	// +optional. A node can define aliases for a subset of its outputs. This is particularly useful if different nodes need to conform to the same interface (e.g. all branches in a branch node). Downstream nodes must refer to this nodes outputs using the alias if one's specified.
	OutputAliases []CoreAlias `json:"output_aliases,omitempty"`
	// Information about the Task to execute in this node.
	TaskNode *CoreTaskNode `json:"task_node,omitempty"`
	// Information about the Workflow to execute in this mode.
	WorkflowNode *CoreWorkflowNode `json:"workflow_node,omitempty"`
	// Information about the branch node to evaluate in this node.
	BranchNode *CoreBranchNode `json:"branch_node,omitempty"`
	// Information about the array job to evaluate in this node.
	ArrayNode *CoreArrayNode `json:"array_node,omitempty"`
}
