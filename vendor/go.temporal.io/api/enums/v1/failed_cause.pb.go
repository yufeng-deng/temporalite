// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/failed_cause.proto

package enums

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Workflow tasks can fail for various reasons. Note that some of these reasons can only originate
// from the server, and some of them can only originate from the SDK/worker.
type WorkflowTaskFailedCause int32

const (
	WORKFLOW_TASK_FAILED_CAUSE_UNSPECIFIED WorkflowTaskFailedCause = 0
	// Between starting and completing the workflow task (with a workflow completion command), some
	// new command (like a signal) was processed into workflow history. The outstanding task will be
	// failed with this reason, and a worker must pick up a new task.
	WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_COMMAND                                         WorkflowTaskFailedCause = 1
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_ACTIVITY_ATTRIBUTES                          WorkflowTaskFailedCause = 2
	WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_ACTIVITY_ATTRIBUTES                    WorkflowTaskFailedCause = 3
	WORKFLOW_TASK_FAILED_CAUSE_BAD_START_TIMER_ATTRIBUTES                                WorkflowTaskFailedCause = 4
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_TIMER_ATTRIBUTES                               WorkflowTaskFailedCause = 5
	WORKFLOW_TASK_FAILED_CAUSE_BAD_RECORD_MARKER_ATTRIBUTES                              WorkflowTaskFailedCause = 6
	WORKFLOW_TASK_FAILED_CAUSE_BAD_COMPLETE_WORKFLOW_EXECUTION_ATTRIBUTES                WorkflowTaskFailedCause = 7
	WORKFLOW_TASK_FAILED_CAUSE_BAD_FAIL_WORKFLOW_EXECUTION_ATTRIBUTES                    WorkflowTaskFailedCause = 8
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_WORKFLOW_EXECUTION_ATTRIBUTES                  WorkflowTaskFailedCause = 9
	WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_ATTRIBUTES WorkflowTaskFailedCause = 10
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CONTINUE_AS_NEW_ATTRIBUTES                            WorkflowTaskFailedCause = 11
	WORKFLOW_TASK_FAILED_CAUSE_START_TIMER_DUPLICATE_ID                                  WorkflowTaskFailedCause = 12
	// The worker wishes to fail the task and have the next one be generated on a normal, not sticky
	// queue. Generally workers should prefer to use the explicit `ResetStickyTaskQueue` RPC call.
	WORKFLOW_TASK_FAILED_CAUSE_RESET_STICKY_TASK_QUEUE                  WorkflowTaskFailedCause = 13
	WORKFLOW_TASK_FAILED_CAUSE_WORKFLOW_WORKER_UNHANDLED_FAILURE        WorkflowTaskFailedCause = 14
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_WORKFLOW_EXECUTION_ATTRIBUTES WorkflowTaskFailedCause = 15
	WORKFLOW_TASK_FAILED_CAUSE_BAD_START_CHILD_EXECUTION_ATTRIBUTES     WorkflowTaskFailedCause = 16
	WORKFLOW_TASK_FAILED_CAUSE_FORCE_CLOSE_COMMAND                      WorkflowTaskFailedCause = 17
	WORKFLOW_TASK_FAILED_CAUSE_FAILOVER_CLOSE_COMMAND                   WorkflowTaskFailedCause = 18
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_INPUT_SIZE                    WorkflowTaskFailedCause = 19
	WORKFLOW_TASK_FAILED_CAUSE_RESET_WORKFLOW                           WorkflowTaskFailedCause = 20
	WORKFLOW_TASK_FAILED_CAUSE_BAD_BINARY                               WorkflowTaskFailedCause = 21
	WORKFLOW_TASK_FAILED_CAUSE_SCHEDULE_ACTIVITY_DUPLICATE_ID           WorkflowTaskFailedCause = 22
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SEARCH_ATTRIBUTES                    WorkflowTaskFailedCause = 23
	// The worker encountered a mismatch while replaying history between what was expected, and
	// what the workflow code actually did.
	WORKFLOW_TASK_FAILED_CAUSE_NON_DETERMINISTIC_ERROR                   WorkflowTaskFailedCause = 24
	WORKFLOW_TASK_FAILED_CAUSE_BAD_MODIFY_WORKFLOW_PROPERTIES_ATTRIBUTES WorkflowTaskFailedCause = 25
	// We send the below error codes to users when their requests would violate a size constraint
	// of their workflow. We do this to ensure that the state of their workflow does not become too
	// large because that can cause severe performance degradation. You can modify the thresholds for
	// each of these errors within your dynamic config.
	//
	// Spawning a new child workflow would cause this workflow to exceed its limit of pending child
	// workflows.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_CHILD_WORKFLOWS_LIMIT_EXCEEDED WorkflowTaskFailedCause = 26
	// Starting a new activity would cause this workflow to exceed its limit of pending activities
	// that we track.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_ACTIVITIES_LIMIT_EXCEEDED WorkflowTaskFailedCause = 27
	// A workflow has a buffer of signals that have not yet reached their destination. We return this
	// error when sending a new signal would exceed the capacity of this buffer.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_SIGNALS_LIMIT_EXCEEDED WorkflowTaskFailedCause = 28
	// Similarly, we have a buffer of pending requests to cancel other workflows. We return this error
	// when our capacity for pending cancel requests is already reached.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_REQUEST_CANCEL_LIMIT_EXCEEDED WorkflowTaskFailedCause = 29
)

var WorkflowTaskFailedCause_name = map[int32]string{
	0:  "Unspecified",
	1:  "UnhandledCommand",
	2:  "BadScheduleActivityAttributes",
	3:  "BadRequestCancelActivityAttributes",
	4:  "BadStartTimerAttributes",
	5:  "BadCancelTimerAttributes",
	6:  "BadRecordMarkerAttributes",
	7:  "BadCompleteWorkflowExecutionAttributes",
	8:  "BadFailWorkflowExecutionAttributes",
	9:  "BadCancelWorkflowExecutionAttributes",
	10: "BadRequestCancelExternalWorkflowExecutionAttributes",
	11: "BadContinueAsNewAttributes",
	12: "StartTimerDuplicateId",
	13: "ResetStickyTaskQueue",
	14: "WorkflowWorkerUnhandledFailure",
	15: "BadSignalWorkflowExecutionAttributes",
	16: "BadStartChildExecutionAttributes",
	17: "ForceCloseCommand",
	18: "FailoverCloseCommand",
	19: "BadSignalInputSize",
	20: "ResetWorkflow",
	21: "BadBinary",
	22: "ScheduleActivityDuplicateId",
	23: "BadSearchAttributes",
	24: "NonDeterministicError",
	25: "BadModifyWorkflowPropertiesAttributes",
	26: "PendingChildWorkflowsLimitExceeded",
	27: "PendingActivitiesLimitExceeded",
	28: "PendingSignalsLimitExceeded",
	29: "PendingRequestCancelLimitExceeded",
}

var WorkflowTaskFailedCause_value = map[string]int32{
	"Unspecified":                                         0,
	"UnhandledCommand":                                    1,
	"BadScheduleActivityAttributes":                       2,
	"BadRequestCancelActivityAttributes":                  3,
	"BadStartTimerAttributes":                             4,
	"BadCancelTimerAttributes":                            5,
	"BadRecordMarkerAttributes":                           6,
	"BadCompleteWorkflowExecutionAttributes":              7,
	"BadFailWorkflowExecutionAttributes":                  8,
	"BadCancelWorkflowExecutionAttributes":                9,
	"BadRequestCancelExternalWorkflowExecutionAttributes": 10,
	"BadContinueAsNewAttributes":                          11,
	"StartTimerDuplicateId":                               12,
	"ResetStickyTaskQueue":                                13,
	"WorkflowWorkerUnhandledFailure":                      14,
	"BadSignalWorkflowExecutionAttributes":                15,
	"BadStartChildExecutionAttributes":                    16,
	"ForceCloseCommand":                                   17,
	"FailoverCloseCommand":                                18,
	"BadSignalInputSize":                                  19,
	"ResetWorkflow":                                       20,
	"BadBinary":                                           21,
	"ScheduleActivityDuplicateId":                         22,
	"BadSearchAttributes":                                 23,
	"NonDeterministicError":                               24,
	"BadModifyWorkflowPropertiesAttributes":               25,
	"PendingChildWorkflowsLimitExceeded":                  26,
	"PendingActivitiesLimitExceeded":                      27,
	"PendingSignalsLimitExceeded":                         28,
	"PendingRequestCancelLimitExceeded":                   29,
}

func (WorkflowTaskFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{0}
}

type StartChildWorkflowExecutionFailedCause int32

const (
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED             StartChildWorkflowExecutionFailedCause = 0
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS StartChildWorkflowExecutionFailedCause = 1
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND     StartChildWorkflowExecutionFailedCause = 2
)

var StartChildWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "WorkflowAlreadyExists",
	2: "NamespaceNotFound",
}

var StartChildWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":           0,
	"WorkflowAlreadyExists": 1,
	"NamespaceNotFound":     2,
}

func (StartChildWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{1}
}

type CancelExternalWorkflowExecutionFailedCause int32

const (
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED                           CancelExternalWorkflowExecutionFailedCause = 0
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND CancelExternalWorkflowExecutionFailedCause = 1
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND                   CancelExternalWorkflowExecutionFailedCause = 2
)

var CancelExternalWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "ExternalWorkflowExecutionNotFound",
	2: "NamespaceNotFound",
}

var CancelExternalWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":                       0,
	"ExternalWorkflowExecutionNotFound": 1,
	"NamespaceNotFound":                 2,
}

func (CancelExternalWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{2}
}

type SignalExternalWorkflowExecutionFailedCause int32

const (
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED                           SignalExternalWorkflowExecutionFailedCause = 0
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND SignalExternalWorkflowExecutionFailedCause = 1
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND                   SignalExternalWorkflowExecutionFailedCause = 2
)

var SignalExternalWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "ExternalWorkflowExecutionNotFound",
	2: "NamespaceNotFound",
}

var SignalExternalWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":                       0,
	"ExternalWorkflowExecutionNotFound": 1,
	"NamespaceNotFound":                 2,
}

func (SignalExternalWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{3}
}

type ResourceExhaustedCause int32

const (
	RESOURCE_EXHAUSTED_CAUSE_UNSPECIFIED ResourceExhaustedCause = 0
	// Caller exceeds request per second limit.
	RESOURCE_EXHAUSTED_CAUSE_RPS_LIMIT ResourceExhaustedCause = 1
	// Caller exceeds max concurrent request limit.
	RESOURCE_EXHAUSTED_CAUSE_CONCURRENT_LIMIT ResourceExhaustedCause = 2
	// System overloaded.
	RESOURCE_EXHAUSTED_CAUSE_SYSTEM_OVERLOADED ResourceExhaustedCause = 3
	// Namespace exceeds persistence rate limit.
	RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_LIMIT ResourceExhaustedCause = 4
)

var ResourceExhaustedCause_name = map[int32]string{
	0: "Unspecified",
	1: "RpsLimit",
	2: "ConcurrentLimit",
	3: "SystemOverloaded",
	4: "PersistenceLimit",
}

var ResourceExhaustedCause_value = map[string]int32{
	"Unspecified":      0,
	"RpsLimit":         1,
	"ConcurrentLimit":  2,
	"SystemOverloaded": 3,
	"PersistenceLimit": 4,
}

func (ResourceExhaustedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{4}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.WorkflowTaskFailedCause", WorkflowTaskFailedCause_name, WorkflowTaskFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.StartChildWorkflowExecutionFailedCause", StartChildWorkflowExecutionFailedCause_name, StartChildWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.CancelExternalWorkflowExecutionFailedCause", CancelExternalWorkflowExecutionFailedCause_name, CancelExternalWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.SignalExternalWorkflowExecutionFailedCause", SignalExternalWorkflowExecutionFailedCause_name, SignalExternalWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.ResourceExhaustedCause", ResourceExhaustedCause_name, ResourceExhaustedCause_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/failed_cause.proto", fileDescriptor_b293cf8d1d965f2d)
}

var fileDescriptor_b293cf8d1d965f2d = []byte{
	// 982 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0x4d, 0x73, 0xdb, 0x44,
	0x18, 0xb6, 0xdc, 0x52, 0x60, 0x81, 0x22, 0x16, 0xda, 0x86, 0x2f, 0xcd, 0xc0, 0x40, 0xa6, 0xf5,
	0x80, 0xd3, 0xb4, 0xb4, 0x99, 0xda, 0x30, 0xe9, 0x7a, 0xf5, 0x3a, 0xde, 0x89, 0xb4, 0x52, 0x77,
	0x57, 0x89, 0xdd, 0xcb, 0x8e, 0x48, 0xdd, 0x56, 0x53, 0xd7, 0xf6, 0xf8, 0xa3, 0xe4, 0xc8, 0x4f,
	0x80, 0x7f, 0xc1, 0xf0, 0x0b, 0xf8, 0x09, 0x1c, 0x73, 0xec, 0x91, 0x38, 0x97, 0x0e, 0xa7, 0xce,
	0x70, 0xe2, 0xc6, 0xc8, 0xb5, 0x53, 0x25, 0xb1, 0x25, 0x9b, 0x9b, 0xad, 0x7d, 0x9e, 0x67, 0xf7,
	0x7d, 0xf6, 0xfd, 0x90, 0xd0, 0xd5, 0x41, 0xf3, 0x69, 0xb7, 0xd3, 0x0b, 0x5b, 0x6b, 0x61, 0x37,
	0x5a, 0x6b, 0xb6, 0x87, 0x4f, 0xfb, 0x6b, 0xcf, 0xd6, 0xd7, 0x1e, 0x86, 0x51, 0xab, 0xf9, 0x40,
	0xef, 0x85, 0xc3, 0x7e, 0xb3, 0xd8, 0xed, 0x75, 0x06, 0x1d, 0x7c, 0x69, 0x8a, 0x2c, 0x86, 0xdd,
	0xa8, 0x38, 0x46, 0x16, 0x9f, 0xad, 0x17, 0x5e, 0x5c, 0x44, 0x57, 0x76, 0x3b, 0xbd, 0x27, 0x0f,
	0x5b, 0x9d, 0x9f, 0x54, 0xd8, 0x7f, 0x52, 0x1d, 0x33, 0x69, 0x4c, 0xc4, 0x05, 0xb4, 0xba, 0xeb,
	0x89, 0xed, 0xaa, 0xe3, 0xed, 0x6a, 0x45, 0xe4, 0xb6, 0xae, 0x12, 0xe6, 0x80, 0xad, 0x29, 0x09,
	0x24, 0xe8, 0x80, 0x4b, 0x1f, 0x28, 0xab, 0x32, 0xb0, 0xcd, 0x1c, 0xbe, 0x8e, 0xbe, 0x49, 0xc5,
	0xd6, 0x08, 0xb7, 0xc7, 0xff, 0x3d, 0xd7, 0x25, 0xdc, 0x36, 0x0d, 0xbc, 0x89, 0xca, 0x29, 0x8c,
	0x0a, 0xb1, 0xb5, 0xa4, 0x35, 0xb0, 0x03, 0x07, 0x34, 0xa1, 0x8a, 0xed, 0x30, 0xd5, 0xd0, 0x44,
	0x29, 0xc1, 0x2a, 0x81, 0x02, 0x69, 0xe6, 0x31, 0x20, 0x92, 0x21, 0x20, 0xe0, 0x5e, 0x00, 0x52,
	0x69, 0x4a, 0x38, 0x05, 0x67, 0xa6, 0xcc, 0x39, 0x7c, 0x07, 0xdd, 0xca, 0x3a, 0x87, 0x22, 0x42,
	0x69, 0xc5, 0x5c, 0x10, 0x49, 0xea, 0x79, 0x5c, 0x42, 0xb7, 0x33, 0xa8, 0x93, 0x9d, 0xcf, 0x70,
	0xdf, 0xc0, 0x65, 0xb4, 0x91, 0x79, 0x7a, 0xea, 0x09, 0x5b, 0xbb, 0x44, 0x6c, 0x9f, 0x24, 0x5f,
	0xc0, 0x0c, 0x41, 0xd6, 0xc6, 0x9e, 0xeb, 0x3b, 0xa0, 0x40, 0x1f, 0xe3, 0xa0, 0x0e, 0x34, 0x50,
	0xcc, 0xe3, 0x49, 0xa9, 0x37, 0x17, 0x70, 0x31, 0x7e, 0x90, 0x21, 0xf3, 0x16, 0xde, 0x42, 0x74,
	0x31, 0x2b, 0xd2, 0x85, 0xde, 0xc6, 0x75, 0xa4, 0x96, 0xbb, 0x55, 0xa8, 0x2b, 0x10, 0x9c, 0x64,
	0x29, 0x23, 0xfc, 0x03, 0xba, 0x93, 0x69, 0x1a, 0x57, 0x8c, 0x07, 0xa0, 0x89, 0xd4, 0x1c, 0x76,
	0x93, 0xf4, 0x77, 0xf0, 0x06, 0xba, 0x99, 0x42, 0x4f, 0xe6, 0x88, 0x1d, 0xf8, 0x0e, 0xa3, 0x44,
	0x81, 0x66, 0xb6, 0xf9, 0x2e, 0xbe, 0x8d, 0x6e, 0xa4, 0x10, 0x05, 0x48, 0x50, 0x5a, 0x2a, 0x46,
	0xb7, 0x1b, 0xaf, 0x96, 0xef, 0x05, 0x10, 0x80, 0xf9, 0x1e, 0xbe, 0x8b, 0xbe, 0x4f, 0xe1, 0x1d,
	0x2f, 0xc5, 0x3f, 0x40, 0x24, 0x4a, 0x2c, 0x86, 0x05, 0x02, 0xcc, 0x8b, 0x0b, 0x5c, 0x8a, 0x64,
	0x5b, 0xd9, 0xd6, 0xbd, 0x8f, 0x29, 0xda, 0x5c, 0xa8, 0x46, 0x68, 0x8d, 0x39, 0xf6, 0x6c, 0x11,
	0x13, 0xdf, 0x40, 0xc5, 0x14, 0x91, 0xaa, 0x27, 0x28, 0x68, 0xea, 0x78, 0x12, 0x8e, 0x9b, 0xc4,
	0x07, 0xf8, 0x16, 0x5a, 0x4f, 0xe3, 0x10, 0xe6, 0x78, 0x3b, 0x20, 0x4e, 0xd1, 0x30, 0xfe, 0x0e,
	0x5d, 0x5f, 0x2c, 0x70, 0xc6, 0xfd, 0x40, 0x69, 0xc9, 0xee, 0x83, 0xf9, 0x21, 0xfe, 0x16, 0x5d,
	0xcb, 0xbc, 0xa8, 0x29, 0xc0, 0xfc, 0x08, 0x5f, 0x43, 0x5f, 0x67, 0x6c, 0x52, 0x61, 0x9c, 0x88,
	0x86, 0x79, 0x29, 0x23, 0xf5, 0xce, 0xf6, 0xb9, 0x13, 0x19, 0x74, 0x79, 0x91, 0x70, 0x80, 0x08,
	0x5a, 0x4b, 0xfa, 0x7d, 0x25, 0x23, 0xef, 0xb8, 0xc7, 0xb5, 0x0d, 0x0a, 0x84, 0xcb, 0x38, 0x8b,
	0xd3, 0x4f, 0x83, 0x10, 0x9e, 0x30, 0x57, 0x70, 0x0d, 0xd9, 0x19, 0xbb, 0xb9, 0x9e, 0xcd, 0xaa,
	0x8d, 0xd7, 0x59, 0xe3, 0x0b, 0xcf, 0x07, 0xa1, 0x18, 0xc8, 0xe4, 0x09, 0x3e, 0xce, 0xe8, 0x2d,
	0x3e, 0x70, 0x9b, 0xf1, 0xad, 0x49, 0xd2, 0x4c, 0x81, 0x52, 0x3b, 0xcc, 0x65, 0x4a, 0x43, 0x9d,
	0x02, 0xd8, 0x60, 0x9b, 0x9f, 0x64, 0x14, 0xc2, 0x54, 0x66, 0x62, 0x5e, 0x7c, 0x88, 0x53, 0x0a,
	0x9f, 0x66, 0xf8, 0x3f, 0x55, 0x78, 0x95, 0x13, 0x67, 0xe8, 0x9f, 0x61, 0x1b, 0xdd, 0x5d, 0x80,
	0x7e, 0xaa, 0x2f, 0x9d, 0x52, 0xf9, 0xbc, 0xf0, 0x8f, 0x81, 0x56, 0xe5, 0x20, 0xec, 0x0d, 0xe8,
	0xe3, 0xa8, 0xf5, 0x60, 0x3a, 0x74, 0x61, 0xbf, 0xb9, 0x37, 0x1c, 0x44, 0x9d, 0x76, 0x72, 0xf2,
	0x96, 0xd1, 0x46, 0xb2, 0xa0, 0x66, 0x94, 0x67, 0xca, 0x28, 0xde, 0x42, 0x74, 0x19, 0xf2, 0xf1,
	0x3a, 0x71, 0x04, 0x10, 0xbb, 0xa1, 0xa1, 0xce, 0xa4, 0x92, 0xa6, 0x11, 0x57, 0xfd, 0x32, 0x42,
	0x9c, 0xb8, 0x20, 0x7d, 0x42, 0xe3, 0xdc, 0x52, 0xba, 0xea, 0x05, 0xdc, 0x36, 0xf3, 0x85, 0x5f,
	0xf3, 0xa8, 0x40, 0xc3, 0xf6, 0x5e, 0xb3, 0x05, 0xfb, 0x83, 0x66, 0xaf, 0x1d, 0xb6, 0x52, 0x23,
	0xdf, 0x44, 0xe5, 0x05, 0xfa, 0x7a, 0x4a, 0xf4, 0x0d, 0x14, 0x2c, 0x2b, 0x90, 0x06, 0x7c, 0x1d,
	0x8a, 0x11, 0x1b, 0xbb, 0xac, 0xf4, 0x7c, 0x4f, 0x64, 0xf4, 0xa8, 0x1d, 0x2e, 0xec, 0xc9, 0xa4,
	0x5d, 0xfd, 0x7f, 0x4f, 0x96, 0x15, 0x58, 0xc2, 0x93, 0x65, 0xa5, 0x67, 0x7b, 0xf2, 0xaf, 0x81,
	0x2e, 0x8b, 0x66, 0xbf, 0x33, 0xec, 0xed, 0x35, 0x61, 0xff, 0x71, 0x38, 0xec, 0x0f, 0xa6, 0xf1,
	0x5f, 0x45, 0x5f, 0x09, 0x90, 0x5e, 0x10, 0x0f, 0x08, 0xa8, 0xd7, 0x48, 0x20, 0xd5, 0x9c, 0x40,
	0x57, 0xd1, 0x97, 0x73, 0x91, 0xc2, 0x9f, 0x54, 0xb6, 0x69, 0xc4, 0x9d, 0x7e, 0x2e, 0x8e, 0x7a,
	0x9c, 0x06, 0x42, 0x00, 0x57, 0x13, 0x78, 0x1e, 0x17, 0x51, 0x61, 0x2e, 0x5c, 0x36, 0xa4, 0x02,
	0x57, 0xc7, 0x63, 0xc8, 0xf1, 0x48, 0x5c, 0xe9, 0xe7, 0x52, 0xf1, 0x3e, 0x08, 0xc9, 0xa4, 0x02,
	0x4e, 0x61, 0xa2, 0x7f, 0xbe, 0xf2, 0x87, 0x71, 0x70, 0x68, 0xe5, 0x9e, 0x1f, 0x5a, 0xb9, 0x97,
	0x87, 0x96, 0xf1, 0xf3, 0xc8, 0x32, 0x7e, 0x1b, 0x59, 0xc6, 0x9f, 0x23, 0xcb, 0x38, 0x18, 0x59,
	0xc6, 0x5f, 0x23, 0xcb, 0x78, 0x31, 0xb2, 0x72, 0x2f, 0x47, 0x96, 0xf1, 0xcb, 0x91, 0x95, 0x3b,
	0x38, 0xb2, 0x72, 0xcf, 0x8f, 0xac, 0x1c, 0x5a, 0x89, 0x3a, 0xc5, 0x99, 0x6f, 0xf5, 0x15, 0x33,
	0x91, 0x3e, 0x7e, 0xfc, 0xfa, 0xef, 0x1b, 0xf7, 0xbf, 0x78, 0x94, 0x40, 0x47, 0x9d, 0x13, 0x1f,
	0x0c, 0xe5, 0xf1, 0x8f, 0xdf, 0xf3, 0x97, 0xd4, 0x14, 0x40, 0xba, 0x51, 0x11, 0xc6, 0x72, 0x3b,
	0xeb, 0x7f, 0xe7, 0x57, 0xa6, 0xcf, 0x4b, 0x25, 0xd2, 0x8d, 0x4a, 0xa5, 0xf1, 0x4a, 0xa9, 0xb4,
	0xb3, 0xfe, 0xe3, 0x85, 0xf1, 0xd7, 0xc5, 0xcd, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x47,
	0x79, 0x43, 0x89, 0x0c, 0x00, 0x00,
}

func (x WorkflowTaskFailedCause) String() string {
	s, ok := WorkflowTaskFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x StartChildWorkflowExecutionFailedCause) String() string {
	s, ok := StartChildWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x CancelExternalWorkflowExecutionFailedCause) String() string {
	s, ok := CancelExternalWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x SignalExternalWorkflowExecutionFailedCause) String() string {
	s, ok := SignalExternalWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ResourceExhaustedCause) String() string {
	s, ok := ResourceExhaustedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}