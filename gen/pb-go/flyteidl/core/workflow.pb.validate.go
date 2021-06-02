// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/core/workflow.proto

package core

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _workflow_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on IfBlock with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *IfBlock) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCondition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IfBlockValidationError{
				field:  "Condition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetThenNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IfBlockValidationError{
				field:  "ThenNode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// IfBlockValidationError is the validation error returned by IfBlock.Validate
// if the designated constraints aren't met.
type IfBlockValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IfBlockValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IfBlockValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IfBlockValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IfBlockValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IfBlockValidationError) ErrorName() string { return "IfBlockValidationError" }

// Error satisfies the builtin error interface
func (e IfBlockValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIfBlock.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IfBlockValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IfBlockValidationError{}

// Validate checks the field values on IfElseBlock with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *IfElseBlock) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCase()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IfElseBlockValidationError{
				field:  "Case",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetOther() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IfElseBlockValidationError{
					field:  fmt.Sprintf("Other[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.Default.(type) {

	case *IfElseBlock_ElseNode:

		if v, ok := interface{}(m.GetElseNode()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IfElseBlockValidationError{
					field:  "ElseNode",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *IfElseBlock_Error:

		if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IfElseBlockValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// IfElseBlockValidationError is the validation error returned by
// IfElseBlock.Validate if the designated constraints aren't met.
type IfElseBlockValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IfElseBlockValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IfElseBlockValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IfElseBlockValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IfElseBlockValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IfElseBlockValidationError) ErrorName() string { return "IfElseBlockValidationError" }

// Error satisfies the builtin error interface
func (e IfElseBlockValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIfElseBlock.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IfElseBlockValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IfElseBlockValidationError{}

// Validate checks the field values on BranchNode with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *BranchNode) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetIfElse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BranchNodeValidationError{
				field:  "IfElse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// BranchNodeValidationError is the validation error returned by
// BranchNode.Validate if the designated constraints aren't met.
type BranchNodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BranchNodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BranchNodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BranchNodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BranchNodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BranchNodeValidationError) ErrorName() string { return "BranchNodeValidationError" }

// Error satisfies the builtin error interface
func (e BranchNodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBranchNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BranchNodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BranchNodeValidationError{}

// Validate checks the field values on TaskNode with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TaskNode) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Reference.(type) {

	case *TaskNode_ReferenceId:

		if v, ok := interface{}(m.GetReferenceId()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TaskNodeValidationError{
					field:  "ReferenceId",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TaskNodeValidationError is the validation error returned by
// TaskNode.Validate if the designated constraints aren't met.
type TaskNodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskNodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskNodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskNodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskNodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskNodeValidationError) ErrorName() string { return "TaskNodeValidationError" }

// Error satisfies the builtin error interface
func (e TaskNodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskNodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskNodeValidationError{}

// Validate checks the field values on WorkflowNode with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *WorkflowNode) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Reference.(type) {

	case *WorkflowNode_LaunchplanRef:

		if v, ok := interface{}(m.GetLaunchplanRef()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WorkflowNodeValidationError{
					field:  "LaunchplanRef",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *WorkflowNode_SubWorkflowRef:

		if v, ok := interface{}(m.GetSubWorkflowRef()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WorkflowNodeValidationError{
					field:  "SubWorkflowRef",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// WorkflowNodeValidationError is the validation error returned by
// WorkflowNode.Validate if the designated constraints aren't met.
type WorkflowNodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowNodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowNodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowNodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowNodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowNodeValidationError) ErrorName() string { return "WorkflowNodeValidationError" }

// Error satisfies the builtin error interface
func (e WorkflowNodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowNodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowNodeValidationError{}

// Validate checks the field values on ArrayNode with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ArrayNode) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTaskReference()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ArrayNodeValidationError{
				field:  "TaskReference",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Parallelism

	// no validation rules for Size

	// no validation rules for MinSuccessRatio

	return nil
}

// ArrayNodeValidationError is the validation error returned by
// ArrayNode.Validate if the designated constraints aren't met.
type ArrayNodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArrayNodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArrayNodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArrayNodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArrayNodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArrayNodeValidationError) ErrorName() string { return "ArrayNodeValidationError" }

// Error satisfies the builtin error interface
func (e ArrayNodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArrayNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArrayNodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArrayNodeValidationError{}

// Validate checks the field values on NodeMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *NodeMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeMetadataValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRetries()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeMetadataValidationError{
				field:  "Retries",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.InterruptibleValue.(type) {

	case *NodeMetadata_Interruptible:
		// no validation rules for Interruptible

	}

	return nil
}

// NodeMetadataValidationError is the validation error returned by
// NodeMetadata.Validate if the designated constraints aren't met.
type NodeMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeMetadataValidationError) ErrorName() string { return "NodeMetadataValidationError" }

// Error satisfies the builtin error interface
func (e NodeMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeMetadataValidationError{}

// Validate checks the field values on Alias with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Alias) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Var

	// no validation rules for Alias

	return nil
}

// AliasValidationError is the validation error returned by Alias.Validate if
// the designated constraints aren't met.
type AliasValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AliasValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AliasValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AliasValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AliasValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AliasValidationError) ErrorName() string { return "AliasValidationError" }

// Error satisfies the builtin error interface
func (e AliasValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAlias.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AliasValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AliasValidationError{}

// Validate checks the field values on Node with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Node) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetInputs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeValidationError{
					field:  fmt.Sprintf("Inputs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetOutputAliases() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeValidationError{
					field:  fmt.Sprintf("OutputAliases[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.Target.(type) {

	case *Node_TaskNode:

		if v, ok := interface{}(m.GetTaskNode()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeValidationError{
					field:  "TaskNode",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Node_WorkflowNode:

		if v, ok := interface{}(m.GetWorkflowNode()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeValidationError{
					field:  "WorkflowNode",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Node_BranchNode:

		if v, ok := interface{}(m.GetBranchNode()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeValidationError{
					field:  "BranchNode",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Node_ArrayNode:

		if v, ok := interface{}(m.GetArrayNode()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeValidationError{
					field:  "ArrayNode",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// NodeValidationError is the validation error returned by Node.Validate if the
// designated constraints aren't met.
type NodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeValidationError) ErrorName() string { return "NodeValidationError" }

// Error satisfies the builtin error interface
func (e NodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeValidationError{}

// Validate checks the field values on WorkflowMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *WorkflowMetadata) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetQualityOfService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowMetadataValidationError{
				field:  "QualityOfService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for OnFailure

	return nil
}

// WorkflowMetadataValidationError is the validation error returned by
// WorkflowMetadata.Validate if the designated constraints aren't met.
type WorkflowMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowMetadataValidationError) ErrorName() string { return "WorkflowMetadataValidationError" }

// Error satisfies the builtin error interface
func (e WorkflowMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowMetadataValidationError{}

// Validate checks the field values on WorkflowMetadataDefaults with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *WorkflowMetadataDefaults) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Interruptible

	return nil
}

// WorkflowMetadataDefaultsValidationError is the validation error returned by
// WorkflowMetadataDefaults.Validate if the designated constraints aren't met.
type WorkflowMetadataDefaultsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowMetadataDefaultsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowMetadataDefaultsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowMetadataDefaultsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowMetadataDefaultsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowMetadataDefaultsValidationError) ErrorName() string {
	return "WorkflowMetadataDefaultsValidationError"
}

// Error satisfies the builtin error interface
func (e WorkflowMetadataDefaultsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowMetadataDefaults.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowMetadataDefaultsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowMetadataDefaultsValidationError{}

// Validate checks the field values on WorkflowTemplate with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *WorkflowTemplate) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowTemplateValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowTemplateValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetInterface()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowTemplateValidationError{
				field:  "Interface",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetNodes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WorkflowTemplateValidationError{
					field:  fmt.Sprintf("Nodes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetOutputs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WorkflowTemplateValidationError{
					field:  fmt.Sprintf("Outputs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetFailureNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowTemplateValidationError{
				field:  "FailureNode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadataDefaults()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowTemplateValidationError{
				field:  "MetadataDefaults",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkflowTemplateValidationError is the validation error returned by
// WorkflowTemplate.Validate if the designated constraints aren't met.
type WorkflowTemplateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowTemplateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowTemplateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowTemplateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowTemplateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowTemplateValidationError) ErrorName() string { return "WorkflowTemplateValidationError" }

// Error satisfies the builtin error interface
func (e WorkflowTemplateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowTemplate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowTemplateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowTemplateValidationError{}
