// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proto/payment/payment.proto

package payment

import (
	transaction "github.com/hilmiikhsan/my-grpc-proto/protogen/go/transaction"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart        *transaction.Cart `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	Currency    string            `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	TotalAmount uint32            `protobuf:"varint,3,opt,name=total_amount,proto3" json:"total_amount,omitempty"`
	Tax         uint32            `protobuf:"varint,4,opt,name=tax,proto3" json:"tax,omitempty"`
	PromoCode   string            `protobuf:"bytes,16,opt,name=promo_code,proto3" json:"promo_code,omitempty"`
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentRequest) GetCart() *transaction.Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

func (x *PaymentRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *PaymentRequest) GetTotalAmount() uint32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *PaymentRequest) GetTax() uint32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *PaymentRequest) GetPromoCode() string {
	if x != nil {
		return x.PromoCode
	}
	return ""
}

type PaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentUuid string `protobuf:"bytes,1,opt,name=payment_uuid,proto3" json:"payment_uuid,omitempty"`
	Confirmed   bool   `protobuf:"varint,2,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_proto_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentResponse) GetPaymentUuid() string {
	if x != nil {
		return x.PaymentUuid
	}
	return ""
}

func (x *PaymentResponse) GetConfirmed() bool {
	if x != nil {
		return x.Confirmed
	}
	return false
}

var File_proto_payment_payment_proto protoreflect.FileDescriptor

var file_proto_payment_payment_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x61, 0x78,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x53, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x65, 0x64, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x6c, 0x6d, 0x69, 0x69, 0x6b, 0x68, 0x73, 0x61, 0x6e, 0x2f,
	0x6d, 0x79, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_payment_payment_proto_rawDescOnce sync.Once
	file_proto_payment_payment_proto_rawDescData = file_proto_payment_payment_proto_rawDesc
)

func file_proto_payment_payment_proto_rawDescGZIP() []byte {
	file_proto_payment_payment_proto_rawDescOnce.Do(func() {
		file_proto_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_payment_payment_proto_rawDescData)
	})
	return file_proto_payment_payment_proto_rawDescData
}

var file_proto_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_payment_payment_proto_goTypes = []any{
	(*PaymentRequest)(nil),   // 0: payment.PaymentRequest
	(*PaymentResponse)(nil),  // 1: payment.PaymentResponse
	(*transaction.Cart)(nil), // 2: transaction.Cart
}
var file_proto_payment_payment_proto_depIdxs = []int32{
	2, // 0: payment.PaymentRequest.cart:type_name -> transaction.Cart
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_payment_payment_proto_init() }
func file_proto_payment_payment_proto_init() {
	if File_proto_payment_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_payment_payment_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PaymentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_payment_payment_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PaymentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_payment_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_payment_payment_proto_goTypes,
		DependencyIndexes: file_proto_payment_payment_proto_depIdxs,
		MessageInfos:      file_proto_payment_payment_proto_msgTypes,
	}.Build()
	File_proto_payment_payment_proto = out.File
	file_proto_payment_payment_proto_rawDesc = nil
	file_proto_payment_payment_proto_goTypes = nil
	file_proto_payment_payment_proto_depIdxs = nil
}
