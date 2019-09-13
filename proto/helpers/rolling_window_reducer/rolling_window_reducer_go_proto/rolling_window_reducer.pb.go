// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/helpers/rolling_window_reducer/rolling_window_reducer.proto

package mako_helpers

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	mako_go_proto "github.com/google/mako/spec/proto/mako_go_proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RWRConfig_WindowOperation int32

const (
	RWRConfig_COUNT       RWRConfig_WindowOperation = 1
	RWRConfig_SUM         RWRConfig_WindowOperation = 2
	RWRConfig_MEAN        RWRConfig_WindowOperation = 3
	RWRConfig_PERCENTILE  RWRConfig_WindowOperation = 4
	RWRConfig_RATIO_SUM   RWRConfig_WindowOperation = 5
	RWRConfig_ERROR_COUNT RWRConfig_WindowOperation = 6
)

var RWRConfig_WindowOperation_name = map[int32]string{
	1: "COUNT",
	2: "SUM",
	3: "MEAN",
	4: "PERCENTILE",
	5: "RATIO_SUM",
	6: "ERROR_COUNT",
}

var RWRConfig_WindowOperation_value = map[string]int32{
	"COUNT":       1,
	"SUM":         2,
	"MEAN":        3,
	"PERCENTILE":  4,
	"RATIO_SUM":   5,
	"ERROR_COUNT": 6,
}

func (x RWRConfig_WindowOperation) Enum() *RWRConfig_WindowOperation {
	p := new(RWRConfig_WindowOperation)
	*p = x
	return p
}

func (x RWRConfig_WindowOperation) String() string {
	return proto.EnumName(RWRConfig_WindowOperation_name, int32(x))
}

func (x *RWRConfig_WindowOperation) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RWRConfig_WindowOperation_value, data, "RWRConfig_WindowOperation")
	if err != nil {
		return err
	}
	*x = RWRConfig_WindowOperation(value)
	return nil
}

func (RWRConfig_WindowOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dffccf54472da3d3, []int{0, 0}
}

type RWRConfig struct {
	InputMetricKeys            []string                   `protobuf:"bytes,1,rep,name=input_metric_keys,json=inputMetricKeys" json:"input_metric_keys,omitempty"`
	DenominatorInputMetricKeys []string                   `protobuf:"bytes,9,rep,name=denominator_input_metric_keys,json=denominatorInputMetricKeys" json:"denominator_input_metric_keys,omitempty"`
	ErrorSamplerNameInputs     []string                   `protobuf:"bytes,10,rep,name=error_sampler_name_inputs,json=errorSamplerNameInputs" json:"error_sampler_name_inputs,omitempty"`
	OutputMetricKey            *string                    `protobuf:"bytes,2,opt,name=output_metric_key,json=outputMetricKey" json:"output_metric_key,omitempty"`
	WindowOperation            *RWRConfig_WindowOperation `protobuf:"varint,3,opt,name=window_operation,json=windowOperation,enum=mako.helpers.RWRConfig_WindowOperation" json:"window_operation,omitempty"`
	WindowSize                 *float64                   `protobuf:"fixed64,4,opt,name=window_size,json=windowSize" json:"window_size,omitempty"`
	StepsPerWindow             *int32                     `protobuf:"varint,5,opt,name=steps_per_window,json=stepsPerWindow,def=1" json:"steps_per_window,omitempty"`
	ZeroForEmptyWindow         *bool                      `protobuf:"varint,6,opt,name=zero_for_empty_window,json=zeroForEmptyWindow" json:"zero_for_empty_window,omitempty"`
	PercentileMilli            *int32                     `protobuf:"varint,7,opt,name=percentile_milli,json=percentileMilli" json:"percentile_milli,omitempty"`
	MaxSampleSize              *int32                     `protobuf:"varint,8,opt,name=max_sample_size,json=maxSampleSize,def=-1" json:"max_sample_size,omitempty"`
	ErrorMatcher               *string                    `protobuf:"bytes,11,opt,name=error_matcher,json=errorMatcher" json:"error_matcher,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                   `json:"-"`
	XXX_unrecognized           []byte                     `json:"-"`
	XXX_sizecache              int32                      `json:"-"`
}

func (m *RWRConfig) Reset()         { *m = RWRConfig{} }
func (m *RWRConfig) String() string { return proto.CompactTextString(m) }
func (*RWRConfig) ProtoMessage()    {}
func (*RWRConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_dffccf54472da3d3, []int{0}
}

func (m *RWRConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RWRConfig.Unmarshal(m, b)
}
func (m *RWRConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RWRConfig.Marshal(b, m, deterministic)
}
func (m *RWRConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RWRConfig.Merge(m, src)
}
func (m *RWRConfig) XXX_Size() int {
	return xxx_messageInfo_RWRConfig.Size(m)
}
func (m *RWRConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RWRConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RWRConfig proto.InternalMessageInfo

const Default_RWRConfig_StepsPerWindow int32 = 1
const Default_RWRConfig_MaxSampleSize int32 = -1

func (m *RWRConfig) GetInputMetricKeys() []string {
	if m != nil {
		return m.InputMetricKeys
	}
	return nil
}

func (m *RWRConfig) GetDenominatorInputMetricKeys() []string {
	if m != nil {
		return m.DenominatorInputMetricKeys
	}
	return nil
}

func (m *RWRConfig) GetErrorSamplerNameInputs() []string {
	if m != nil {
		return m.ErrorSamplerNameInputs
	}
	return nil
}

func (m *RWRConfig) GetOutputMetricKey() string {
	if m != nil && m.OutputMetricKey != nil {
		return *m.OutputMetricKey
	}
	return ""
}

func (m *RWRConfig) GetWindowOperation() RWRConfig_WindowOperation {
	if m != nil && m.WindowOperation != nil {
		return *m.WindowOperation
	}
	return RWRConfig_COUNT
}

func (m *RWRConfig) GetWindowSize() float64 {
	if m != nil && m.WindowSize != nil {
		return *m.WindowSize
	}
	return 0
}

func (m *RWRConfig) GetStepsPerWindow() int32 {
	if m != nil && m.StepsPerWindow != nil {
		return *m.StepsPerWindow
	}
	return Default_RWRConfig_StepsPerWindow
}

func (m *RWRConfig) GetZeroForEmptyWindow() bool {
	if m != nil && m.ZeroForEmptyWindow != nil {
		return *m.ZeroForEmptyWindow
	}
	return false
}

func (m *RWRConfig) GetPercentileMilli() int32 {
	if m != nil && m.PercentileMilli != nil {
		return *m.PercentileMilli
	}
	return 0
}

func (m *RWRConfig) GetMaxSampleSize() int32 {
	if m != nil && m.MaxSampleSize != nil {
		return *m.MaxSampleSize
	}
	return Default_RWRConfig_MaxSampleSize
}

func (m *RWRConfig) GetErrorMatcher() string {
	if m != nil && m.ErrorMatcher != nil {
		return *m.ErrorMatcher
	}
	return ""
}

type RWRAddPointsInput struct {
	PointList            []*mako_go_proto.SamplePoint `protobuf:"bytes,1,rep,name=point_list,json=pointList" json:"point_list,omitempty"`
	ErrorList            []*mako_go_proto.SampleError `protobuf:"bytes,2,rep,name=error_list,json=errorList" json:"error_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RWRAddPointsInput) Reset()         { *m = RWRAddPointsInput{} }
func (m *RWRAddPointsInput) String() string { return proto.CompactTextString(m) }
func (*RWRAddPointsInput) ProtoMessage()    {}
func (*RWRAddPointsInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_dffccf54472da3d3, []int{1}
}

func (m *RWRAddPointsInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RWRAddPointsInput.Unmarshal(m, b)
}
func (m *RWRAddPointsInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RWRAddPointsInput.Marshal(b, m, deterministic)
}
func (m *RWRAddPointsInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RWRAddPointsInput.Merge(m, src)
}
func (m *RWRAddPointsInput) XXX_Size() int {
	return xxx_messageInfo_RWRAddPointsInput.Size(m)
}
func (m *RWRAddPointsInput) XXX_DiscardUnknown() {
	xxx_messageInfo_RWRAddPointsInput.DiscardUnknown(m)
}

var xxx_messageInfo_RWRAddPointsInput proto.InternalMessageInfo

func (m *RWRAddPointsInput) GetPointList() []*mako_go_proto.SamplePoint {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (m *RWRAddPointsInput) GetErrorList() []*mako_go_proto.SampleError {
	if m != nil {
		return m.ErrorList
	}
	return nil
}

type RWRCompleteOutput struct {
	PointList            []*mako_go_proto.SamplePoint `protobuf:"bytes,1,rep,name=point_list,json=pointList" json:"point_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RWRCompleteOutput) Reset()         { *m = RWRCompleteOutput{} }
func (m *RWRCompleteOutput) String() string { return proto.CompactTextString(m) }
func (*RWRCompleteOutput) ProtoMessage()    {}
func (*RWRCompleteOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_dffccf54472da3d3, []int{2}
}

func (m *RWRCompleteOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RWRCompleteOutput.Unmarshal(m, b)
}
func (m *RWRCompleteOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RWRCompleteOutput.Marshal(b, m, deterministic)
}
func (m *RWRCompleteOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RWRCompleteOutput.Merge(m, src)
}
func (m *RWRCompleteOutput) XXX_Size() int {
	return xxx_messageInfo_RWRCompleteOutput.Size(m)
}
func (m *RWRCompleteOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_RWRCompleteOutput.DiscardUnknown(m)
}

var xxx_messageInfo_RWRCompleteOutput proto.InternalMessageInfo

func (m *RWRCompleteOutput) GetPointList() []*mako_go_proto.SamplePoint {
	if m != nil {
		return m.PointList
	}
	return nil
}

func init() {
	proto.RegisterEnum("mako.helpers.RWRConfig_WindowOperation", RWRConfig_WindowOperation_name, RWRConfig_WindowOperation_value)
	proto.RegisterType((*RWRConfig)(nil), "mako.helpers.RWRConfig")
	proto.RegisterType((*RWRAddPointsInput)(nil), "mako.helpers.RWRAddPointsInput")
	proto.RegisterType((*RWRCompleteOutput)(nil), "mako.helpers.RWRCompleteOutput")
}

func init() {
	proto.RegisterFile("proto/helpers/rolling_window_reducer/rolling_window_reducer.proto", fileDescriptor_dffccf54472da3d3)
}

var fileDescriptor_dffccf54472da3d3 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x4f, 0xdb, 0x30,
	0x14, 0x57, 0x5a, 0x0a, 0xe4, 0x15, 0x68, 0xb0, 0xc4, 0x14, 0x90, 0xa6, 0x55, 0xdd, 0x61, 0x1d,
	0x68, 0x61, 0x70, 0xdb, 0x6e, 0xa5, 0xca, 0x34, 0x34, 0xfa, 0x47, 0x06, 0xc4, 0xd1, 0x8a, 0xd2,
	0x47, 0xb1, 0x88, 0xe3, 0xc8, 0x36, 0x2a, 0xf0, 0x35, 0xf6, 0x85, 0x27, 0xdb, 0x81, 0x51, 0xd0,
	0x0e, 0xbb, 0xb5, 0xbf, 0x7f, 0xef, 0xe5, 0x67, 0x1b, 0x06, 0x95, 0x92, 0x46, 0x1e, 0xde, 0x60,
	0x51, 0xa1, 0xd2, 0x87, 0x4a, 0x16, 0x05, 0x2f, 0xe7, 0x6c, 0xc1, 0xcb, 0x99, 0x5c, 0x30, 0x85,
	0xb3, 0xbb, 0x1c, 0xd5, 0x3f, 0xe0, 0xc4, 0x79, 0xc9, 0x86, 0xc8, 0x6e, 0x65, 0x52, 0x27, 0xec,
	0xed, 0xe8, 0x0a, 0xf3, 0x43, 0x9f, 0xea, 0x08, 0xf7, 0xb3, 0xf7, 0xbb, 0x05, 0x21, 0xbd, 0xa2,
	0x43, 0x59, 0x5e, 0xf3, 0x39, 0xd9, 0x87, 0x6d, 0x5e, 0x56, 0x77, 0x86, 0x09, 0x34, 0x8a, 0xe7,
	0xec, 0x16, 0x1f, 0x74, 0x1c, 0x74, 0x9b, 0xfd, 0x90, 0x76, 0x1c, 0x31, 0x72, 0xf8, 0x2f, 0x7c,
	0xd0, 0x64, 0x00, 0xef, 0x67, 0x58, 0x4a, 0xc1, 0xcb, 0xcc, 0x48, 0xc5, 0xde, 0xfa, 0x42, 0xe7,
	0xdb, 0x7b, 0x21, 0x3a, 0x7d, 0x15, 0xf1, 0x0d, 0x76, 0x51, 0x29, 0xa9, 0x98, 0xce, 0x44, 0x55,
	0xa0, 0x62, 0x65, 0x26, 0xd0, 0x27, 0xe9, 0x18, 0x9c, 0xfd, 0x9d, 0x13, 0x9c, 0x7b, 0x7e, 0x9c,
	0x09, 0x74, 0x19, 0xda, 0x6e, 0x2a, 0xef, 0xcc, 0xf2, 0xc8, 0xb8, 0xd1, 0x0d, 0xec, 0xa6, 0x9e,
	0x78, 0x9e, 0x43, 0x28, 0x44, 0x75, 0x41, 0xb2, 0x42, 0x95, 0x19, 0x2e, 0xcb, 0xb8, 0xd9, 0x0d,
	0xfa, 0x5b, 0xc7, 0x9f, 0x92, 0x97, 0x1d, 0x25, 0xcf, 0x45, 0x24, 0x57, 0x4e, 0x3f, 0x79, 0x92,
	0xd3, 0xce, 0x62, 0x19, 0x20, 0x1f, 0xa0, 0x5d, 0x67, 0x6a, 0xfe, 0x88, 0xf1, 0x4a, 0x37, 0xe8,
	0x07, 0x14, 0x3c, 0x74, 0xce, 0x1f, 0x91, 0x1c, 0x40, 0xa4, 0x0d, 0x56, 0x9a, 0x55, 0xa8, 0xea,
	0xf3, 0x89, 0x5b, 0xdd, 0xa0, 0xdf, 0xfa, 0x1e, 0x1c, 0xd1, 0x2d, 0x47, 0x4d, 0x51, 0xf9, 0x39,
	0xe4, 0x08, 0x76, 0x1e, 0x51, 0x49, 0x76, 0x2d, 0x15, 0x43, 0x51, 0x99, 0x87, 0x27, 0xc7, 0x6a,
	0x37, 0xe8, 0xaf, 0x53, 0x62, 0xc9, 0x1f, 0x52, 0xa5, 0x96, 0xaa, 0x2d, 0x9f, 0x21, 0xaa, 0x50,
	0xe5, 0x58, 0x1a, 0x5e, 0x20, 0x13, 0xbc, 0x28, 0x78, 0xbc, 0x66, 0xf3, 0x69, 0xe7, 0x2f, 0x3e,
	0xb2, 0x30, 0xd9, 0x87, 0x8e, 0xc8, 0xee, 0xeb, 0x92, 0xfd, 0xbe, 0xeb, 0x6e, 0x93, 0xc6, 0x97,
	0x23, 0xba, 0x29, 0xb2, 0x7b, 0x5f, 0xaf, 0x5b, 0xfb, 0x23, 0x6c, 0xfa, 0x23, 0x11, 0x99, 0xc9,
	0x6f, 0x50, 0xc5, 0x6d, 0xd7, 0xe9, 0x86, 0x03, 0x47, 0x1e, 0xeb, 0x31, 0xe8, 0xbc, 0x2a, 0x88,
	0x84, 0xd0, 0x1a, 0x4e, 0x2e, 0xc7, 0x17, 0x51, 0x40, 0xd6, 0xa0, 0x79, 0x7e, 0x39, 0x8a, 0x1a,
	0x64, 0x1d, 0x56, 0x46, 0xe9, 0x60, 0x1c, 0x35, 0xc9, 0x16, 0xc0, 0x34, 0xa5, 0xc3, 0x74, 0x7c,
	0x71, 0x7a, 0x96, 0x46, 0x2b, 0x64, 0x13, 0x42, 0x3a, 0xb8, 0x38, 0x9d, 0x30, 0x2b, 0x6c, 0x91,
	0x0e, 0xb4, 0x53, 0x4a, 0x27, 0x94, 0xf9, 0x88, 0xd5, 0xde, 0x02, 0xb6, 0xe9, 0x15, 0x1d, 0xcc,
	0x66, 0x53, 0xc9, 0x4b, 0xa3, 0xdd, 0x99, 0x93, 0xaf, 0x00, 0x95, 0xfd, 0xcb, 0x0a, 0xae, 0x8d,
	0xbb, 0x95, 0xed, 0xe3, 0x6d, 0x7f, 0x80, 0xfe, 0x03, 0x9c, 0x98, 0x86, 0x4e, 0x74, 0xc6, 0xb5,
	0x73, 0xf8, 0x8f, 0x71, 0x8e, 0xc6, 0x5b, 0x47, 0x6a, 0x59, 0x1a, 0x3a, 0x91, 0x75, 0xf4, 0x52,
	0x37, 0x78, 0x28, 0x2d, 0x69, 0x70, 0xe2, 0x2e, 0xd2, 0xff, 0x0f, 0x3e, 0xf9, 0x09, 0x07, 0xb9,
	0x14, 0xc9, 0x5c, 0xca, 0x79, 0x81, 0x89, 0x41, 0x6d, 0x78, 0x39, 0x4f, 0x2a, 0x54, 0xd7, 0x52,
	0x89, 0xac, 0xcc, 0x71, 0xe9, 0xde, 0x9d, 0xec, 0x52, 0xff, 0x8e, 0x7d, 0xa9, 0xd4, 0xbf, 0xe2,
	0xa9, 0x7d, 0x9f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x98, 0x3c, 0x1f, 0xdc, 0x08, 0x04, 0x00,
	0x00,
}
