package protoconv

import (
	"time"

	"github.com/aws/smithy-go/ptr"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToBoolPtr(value *wrapperspb.BoolValue) *bool {
	if value == nil {
		return nil
	}
	return ptr.Bool(value.GetValue())
}

func ToStringPtr(value *wrapperspb.StringValue) *string {
	if value == nil {
		return nil
	}
	return ptr.String(value.GetValue())
}

func ToInt32Ptr(value *wrapperspb.Int32Value) *int32 {
	if value == nil {
		return nil
	}
	return ptr.Int32(value.GetValue())
}

func ToInt64Ptr(value *wrapperspb.Int64Value) *int64 {
	if value == nil {
		return nil
	}
	return ptr.Int64(value.GetValue())
}

func ToFloat32Ptr(value *wrapperspb.FloatValue) *float32 {
	if value == nil {
		return nil
	}
	return ptr.Float32(value.GetValue())
}

func ToFloat64Ptr(value *wrapperspb.DoubleValue) *float64 {
	if value == nil {
		return nil
	}
	return ptr.Float64(value.GetValue())
}

func ToUInt32Ptr(value *wrapperspb.UInt32Value) *uint32 {
	if value == nil {
		return nil
	}
	return ptr.Uint32(value.GetValue())
}

func ToUInt64Ptr(value *wrapperspb.UInt64Value) *uint64 {
	if value == nil {
		return nil
	}
	return ptr.Uint64(value.GetValue())
}

func ToBytesPtr(value *wrapperspb.BytesValue) *[]byte {
	if value == nil {
		return nil
	}
	b := value.GetValue()
	return &b
}

func ToBoolValue(v *bool) *wrapperspb.BoolValue {
	if v == nil {
		return nil
	}
	return wrapperspb.Bool(*v)
}

func ToStringValue(v *string) *wrapperspb.StringValue {
	if v == nil {
		return nil
	}
	return wrapperspb.String(*v)
}

func ToInt32Value(v *int32) *wrapperspb.Int32Value {
	if v == nil {
		return nil
	}
	return wrapperspb.Int32(*v)
}

func ToInt64Value(v *int64) *wrapperspb.Int64Value {
	if v == nil {
		return nil
	}
	return wrapperspb.Int64(*v)
}

func ToUInt32Value(v *uint32) *wrapperspb.UInt32Value {
	if v == nil {
		return nil
	}
	return wrapperspb.UInt32(*v)
}

func ToUint64Value(v *uint64) *wrapperspb.UInt64Value {
	if v == nil {
		return nil
	}
	return wrapperspb.UInt64(*v)
}

func ToFloatValue(v *float32) *wrapperspb.FloatValue {
	if v == nil {
		return nil
	}
	return wrapperspb.Float(*v)
}

func ToDoubleValue(v *float64) *wrapperspb.DoubleValue {
	if v == nil {
		return nil
	}
	return wrapperspb.Double(*v)
}

func ToBytesValue(v *[]byte) *wrapperspb.BytesValue {
	if v == nil {
		return nil
	}
	return wrapperspb.Bytes(*v)
}

func ToTimePtr(v *timestamppb.Timestamp) *time.Time {
	if v == nil {
		return nil
	}
	return ptr.Time(v.AsTime())
}

func ToTimestamp(v *time.Time) *timestamppb.Timestamp {
	if v == nil {
		return nil
	}
	return timestamppb.New(*v)
}
