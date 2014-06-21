// Code generated by protoc-gen-go.
// source: basic_protocol.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	basic_protocol.proto

It has these top-level messages:
	BasicOps
	BasicResp
	PagedListReq
	PagedListResp
*/
package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type BasicOps struct {
	Key              []byte `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Sorted           *bool  `protobuf:"varint,3,opt,name=sorted" json:"sorted,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BasicOps) Reset()         { *m = BasicOps{} }
func (m *BasicOps) String() string { return proto.CompactTextString(m) }
func (*BasicOps) ProtoMessage()    {}

func (m *BasicOps) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *BasicOps) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *BasicOps) GetSorted() bool {
	if m != nil && m.Sorted != nil {
		return *m.Sorted
	}
	return false
}

type BasicResp struct {
	ResponseCode     *int32 `protobuf:"varint,1,req,name=response_code" json:"response_code,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Key              []byte `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BasicResp) Reset()         { *m = BasicResp{} }
func (m *BasicResp) String() string { return proto.CompactTextString(m) }
func (*BasicResp) ProtoMessage()    {}

func (m *BasicResp) GetResponseCode() int32 {
	if m != nil && m.ResponseCode != nil {
		return *m.ResponseCode
	}
	return 0
}

func (m *BasicResp) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *BasicResp) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type PagedListReq struct {
	PageNo           *int32 `protobuf:"varint,1,req,name=page_no" json:"page_no,omitempty"`
	PageSize         *int32 `protobuf:"varint,2,req,name=page_size" json:"page_size,omitempty"`
	FromKey          []byte `protobuf:"bytes,3,req,name=from_key" json:"from_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PagedListReq) Reset()         { *m = PagedListReq{} }
func (m *PagedListReq) String() string { return proto.CompactTextString(m) }
func (*PagedListReq) ProtoMessage()    {}

func (m *PagedListReq) GetPageNo() int32 {
	if m != nil && m.PageNo != nil {
		return *m.PageNo
	}
	return 0
}

func (m *PagedListReq) GetPageSize() int32 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return 0
}

func (m *PagedListReq) GetFromKey() []byte {
	if m != nil {
		return m.FromKey
	}
	return nil
}

type PagedListResp struct {
	ResponseCode     *int32       `protobuf:"varint,1,req,name=response_code" json:"response_code,omitempty"`
	PageNo           *int32       `protobuf:"varint,2,opt,name=page_no" json:"page_no,omitempty"`
	PageSize         *int32       `protobuf:"varint,3,opt,name=page_size" json:"page_size,omitempty"`
	ListCnt          *int32       `protobuf:"varint,4,opt,name=list_cnt" json:"list_cnt,omitempty"`
	List             []*BasicResp `protobuf:"bytes,5,rep,name=list" json:"list,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PagedListResp) Reset()         { *m = PagedListResp{} }
func (m *PagedListResp) String() string { return proto.CompactTextString(m) }
func (*PagedListResp) ProtoMessage()    {}

func (m *PagedListResp) GetResponseCode() int32 {
	if m != nil && m.ResponseCode != nil {
		return *m.ResponseCode
	}
	return 0
}

func (m *PagedListResp) GetPageNo() int32 {
	if m != nil && m.PageNo != nil {
		return *m.PageNo
	}
	return 0
}

func (m *PagedListResp) GetPageSize() int32 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return 0
}

func (m *PagedListResp) GetListCnt() int32 {
	if m != nil && m.ListCnt != nil {
		return *m.ListCnt
	}
	return 0
}

func (m *PagedListResp) GetList() []*BasicResp {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
}
