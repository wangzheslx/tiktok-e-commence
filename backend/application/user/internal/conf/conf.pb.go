// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: internal/conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Server        *Server                `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data          *Data                  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	mi := &file_internal_conf_conf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Http          *Server_HTTP           `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc          *Server_GRPC           `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_internal_conf_conf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Database      *Data_Database         `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis         *Data_Redis            `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_internal_conf_conf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

// 授权/认证
type Auth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Jwt           *Auth_JWT              `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	Casdoor       *Auth_Casdoor          `protobuf:"bytes,2,opt,name=casdoor,proto3" json:"casdoor,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Auth) Reset() {
	*x = Auth{}
	mi := &file_internal_conf_conf_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Auth) GetJwt() *Auth_JWT {
	if x != nil {
		return x.Jwt
	}
	return nil
}

func (x *Auth) GetCasdoor() *Auth_Casdoor {
	if x != nil {
		return x.Casdoor
	}
	return nil
}

// 分布式链路追踪
type Trace struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Jaeger        *Trace_Jaeger          `protobuf:"bytes,1,opt,name=jaeger,proto3" json:"jaeger,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Trace) Reset() {
	*x = Trace{}
	mi := &file_internal_conf_conf_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace) ProtoMessage() {}

func (x *Trace) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace.ProtoReflect.Descriptor instead.
func (*Trace) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Trace) GetJaeger() *Trace_Jaeger {
	if x != nil {
		return x.Jaeger
	}
	return nil
}

type Server_HTTP struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	mi := &file_internal_conf_conf_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Server_GRPC struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	mi := &file_internal_conf_conf_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source        string                 `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	mi := &file_internal_conf_conf_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Data_Redis struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	ReadTimeout   *durationpb.Duration   `protobuf:"bytes,3,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout  *durationpb.Duration   `protobuf:"bytes,4,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	mi := &file_internal_conf_conf_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Data_Redis) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

type Auth_JWT struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ServiceKey    string                 `protobuf:"bytes,1,opt,name=service_key,json=serviceKey,proto3" json:"service_key,omitempty"`
	ApiKey        string                 `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Auth_JWT) Reset() {
	*x = Auth_JWT{}
	mi := &file_internal_conf_conf_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth_JWT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth_JWT) ProtoMessage() {}

func (x *Auth_JWT) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth_JWT.ProtoReflect.Descriptor instead.
func (*Auth_JWT) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Auth_JWT) GetServiceKey() string {
	if x != nil {
		return x.ServiceKey
	}
	return ""
}

func (x *Auth_JWT) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

type Auth_Casdoor struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Certificate   string                 `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Server        *Auth_Casdoor_Server   `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Auth_Casdoor) Reset() {
	*x = Auth_Casdoor{}
	mi := &file_internal_conf_conf_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth_Casdoor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth_Casdoor) ProtoMessage() {}

func (x *Auth_Casdoor) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth_Casdoor.ProtoReflect.Descriptor instead.
func (*Auth_Casdoor) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{3, 1}
}

func (x *Auth_Casdoor) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

func (x *Auth_Casdoor) GetServer() *Auth_Casdoor_Server {
	if x != nil {
		return x.Server
	}
	return nil
}

type Auth_Casdoor_Server struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoint      string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	ClientId      string                 `protobuf:"bytes,2,opt,name=client_id,proto3" json:"client_id,omitempty"`
	ClientSecret  string                 `protobuf:"bytes,3,opt,name=client_secret,proto3" json:"client_secret,omitempty"`
	Organization  string                 `protobuf:"bytes,4,opt,name=organization,proto3" json:"organization,omitempty"`
	Application   string                 `protobuf:"bytes,5,opt,name=application,proto3" json:"application,omitempty"`
	FrontendUrl   string                 `protobuf:"bytes,6,opt,name=frontend_url,proto3" json:"frontend_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Auth_Casdoor_Server) Reset() {
	*x = Auth_Casdoor_Server{}
	mi := &file_internal_conf_conf_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth_Casdoor_Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth_Casdoor_Server) ProtoMessage() {}

func (x *Auth_Casdoor_Server) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth_Casdoor_Server.ProtoReflect.Descriptor instead.
func (*Auth_Casdoor_Server) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{3, 1, 0}
}

func (x *Auth_Casdoor_Server) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Auth_Casdoor_Server) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Auth_Casdoor_Server) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *Auth_Casdoor_Server) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *Auth_Casdoor_Server) GetApplication() string {
	if x != nil {
		return x.Application
	}
	return ""
}

func (x *Auth_Casdoor_Server) GetFrontendUrl() string {
	if x != nil {
		return x.FrontendUrl
	}
	return ""
}

type Trace_Jaeger struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ServiceName   string                 `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Grpc          *Trace_Jaeger_GRPC     `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	Http          *Trace_Jaeger_HTTP     `protobuf:"bytes,3,opt,name=http,proto3" json:"http,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Trace_Jaeger) Reset() {
	*x = Trace_Jaeger{}
	mi := &file_internal_conf_conf_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trace_Jaeger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace_Jaeger) ProtoMessage() {}

func (x *Trace_Jaeger) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace_Jaeger.ProtoReflect.Descriptor instead.
func (*Trace_Jaeger) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Trace_Jaeger) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Trace_Jaeger) GetGrpc() *Trace_Jaeger_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *Trace_Jaeger) GetHttp() *Trace_Jaeger_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

type Trace_Jaeger_GRPC struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoint      string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Trace_Jaeger_GRPC) Reset() {
	*x = Trace_Jaeger_GRPC{}
	mi := &file_internal_conf_conf_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trace_Jaeger_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace_Jaeger_GRPC) ProtoMessage() {}

func (x *Trace_Jaeger_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace_Jaeger_GRPC.ProtoReflect.Descriptor instead.
func (*Trace_Jaeger_GRPC) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{4, 0, 0}
}

func (x *Trace_Jaeger_GRPC) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type Trace_Jaeger_HTTP struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoint      string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Trace_Jaeger_HTTP) Reset() {
	*x = Trace_Jaeger_HTTP{}
	mi := &file_internal_conf_conf_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trace_Jaeger_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace_Jaeger_HTTP) ProtoMessage() {}

func (x *Trace_Jaeger_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace_Jaeger_HTTP.ProtoReflect.Descriptor instead.
func (*Trace_Jaeger_HTTP) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{4, 0, 1}
}

func (x *Trace_Jaeger_HTTP) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

var File_internal_conf_conf_proto protoreflect.FileDescriptor

var file_internal_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb8, 0x02, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x2b, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x2b, 0x0a,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x69, 0x0a, 0x04, 0x48, 0x54,
	0x54, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x69, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x12, 0x18, 0x0a,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x22, 0xdd, 0x02, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x1a, 0x3a,
	0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0xb3, 0x01, 0x0a, 0x05, 0x52,
	0x65, 0x64, 0x69, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x22, 0xdf, 0x03, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x03, 0x6a, 0x77, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x4a, 0x57, 0x54, 0x52, 0x03, 0x6a, 0x77,
	0x74, 0x12, 0x32, 0x0a, 0x07, 0x63, 0x61, 0x73, 0x64, 0x6f, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x61, 0x73, 0x64, 0x6f, 0x6f, 0x72, 0x52, 0x07, 0x63, 0x61,
	0x73, 0x64, 0x6f, 0x6f, 0x72, 0x1a, 0x3f, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a,
	0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x1a, 0xb9, 0x02, 0x0a, 0x07, 0x43, 0x61, 0x73, 0x64, 0x6f,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x61, 0x73, 0x64, 0x6f, 0x6f, 0x72, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0xd2, 0x01,
	0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x75,
	0x72, 0x6c, 0x22, 0x95, 0x02, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06,
	0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e,
	0x4a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x52, 0x06, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x1a, 0xd9,
	0x01, 0x0a, 0x06, 0x4a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x04,
	0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x4a, 0x61,
	0x65, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12,
	0x31, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x2e, 0x4a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74,
	0x74, 0x70, 0x1a, 0x22, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x22, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_conf_conf_proto_rawDescOnce sync.Once
	file_internal_conf_conf_proto_rawDescData = file_internal_conf_conf_proto_rawDesc
)

func file_internal_conf_conf_proto_rawDescGZIP() []byte {
	file_internal_conf_conf_proto_rawDescOnce.Do(func() {
		file_internal_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_conf_conf_proto_rawDescData)
	})
	return file_internal_conf_conf_proto_rawDescData
}

var file_internal_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_internal_conf_conf_proto_goTypes = []any{
	(*Bootstrap)(nil),           // 0: kratos.api.Bootstrap
	(*Server)(nil),              // 1: kratos.api.Server
	(*Data)(nil),                // 2: kratos.api.Data
	(*Auth)(nil),                // 3: kratos.api.Auth
	(*Trace)(nil),               // 4: kratos.api.Trace
	(*Server_HTTP)(nil),         // 5: kratos.api.Server.HTTP
	(*Server_GRPC)(nil),         // 6: kratos.api.Server.GRPC
	(*Data_Database)(nil),       // 7: kratos.api.Data.Database
	(*Data_Redis)(nil),          // 8: kratos.api.Data.Redis
	(*Auth_JWT)(nil),            // 9: kratos.api.Auth.JWT
	(*Auth_Casdoor)(nil),        // 10: kratos.api.Auth.Casdoor
	(*Auth_Casdoor_Server)(nil), // 11: kratos.api.Auth.Casdoor.Server
	(*Trace_Jaeger)(nil),        // 12: kratos.api.Trace.Jaeger
	(*Trace_Jaeger_GRPC)(nil),   // 13: kratos.api.Trace.Jaeger.GRPC
	(*Trace_Jaeger_HTTP)(nil),   // 14: kratos.api.Trace.Jaeger.HTTP
	(*durationpb.Duration)(nil), // 15: google.protobuf.Duration
}
var file_internal_conf_conf_proto_depIdxs = []int32{
	1,  // 0: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
	2,  // 1: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	5,  // 2: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
	6,  // 3: kratos.api.Server.grpc:type_name -> kratos.api.Server.GRPC
	7,  // 4: kratos.api.Data.database:type_name -> kratos.api.Data.Database
	8,  // 5: kratos.api.Data.redis:type_name -> kratos.api.Data.Redis
	9,  // 6: kratos.api.Auth.jwt:type_name -> kratos.api.Auth.JWT
	10, // 7: kratos.api.Auth.casdoor:type_name -> kratos.api.Auth.Casdoor
	12, // 8: kratos.api.Trace.jaeger:type_name -> kratos.api.Trace.Jaeger
	15, // 9: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	15, // 10: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	15, // 11: kratos.api.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	15, // 12: kratos.api.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	11, // 13: kratos.api.Auth.Casdoor.server:type_name -> kratos.api.Auth.Casdoor.Server
	13, // 14: kratos.api.Trace.Jaeger.grpc:type_name -> kratos.api.Trace.Jaeger.GRPC
	14, // 15: kratos.api.Trace.Jaeger.http:type_name -> kratos.api.Trace.Jaeger.HTTP
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_internal_conf_conf_proto_init() }
func file_internal_conf_conf_proto_init() {
	if File_internal_conf_conf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_conf_conf_proto_goTypes,
		DependencyIndexes: file_internal_conf_conf_proto_depIdxs,
		MessageInfos:      file_internal_conf_conf_proto_msgTypes,
	}.Build()
	File_internal_conf_conf_proto = out.File
	file_internal_conf_conf_proto_rawDesc = nil
	file_internal_conf_conf_proto_goTypes = nil
	file_internal_conf_conf_proto_depIdxs = nil
}
