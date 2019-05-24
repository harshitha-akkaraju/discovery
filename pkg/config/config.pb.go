// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Basic struct {
	Username             string                `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             *wrappers.StringValue `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Basic) Reset()         { *m = Basic{} }
func (m *Basic) String() string { return proto.CompactTextString(m) }
func (*Basic) ProtoMessage()    {}
func (*Basic) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ea951a04c6ae584c, []int{0}
}
func (m *Basic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Basic.Unmarshal(m, b)
}
func (m *Basic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Basic.Marshal(b, m, deterministic)
}
func (dst *Basic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Basic.Merge(dst, src)
}
func (m *Basic) XXX_Size() int {
	return xxx_messageInfo_Basic.Size(m)
}
func (m *Basic) XXX_DiscardUnknown() {
	xxx_messageInfo_Basic.DiscardUnknown(m)
}

var xxx_messageInfo_Basic proto.InternalMessageInfo

func (m *Basic) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Basic) GetPassword() *wrappers.StringValue {
	if m != nil {
		return m.Password
	}
	return nil
}

type OAuthToken struct {
	Token                string                `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ApplicationId        *wrappers.StringValue `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OAuthToken) Reset()         { *m = OAuthToken{} }
func (m *OAuthToken) String() string { return proto.CompactTextString(m) }
func (*OAuthToken) ProtoMessage()    {}
func (*OAuthToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ea951a04c6ae584c, []int{1}
}
func (m *OAuthToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthToken.Unmarshal(m, b)
}
func (m *OAuthToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthToken.Marshal(b, m, deterministic)
}
func (dst *OAuthToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthToken.Merge(dst, src)
}
func (m *OAuthToken) XXX_Size() int {
	return xxx_messageInfo_OAuthToken.Size(m)
}
func (m *OAuthToken) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthToken.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthToken proto.InternalMessageInfo

func (m *OAuthToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *OAuthToken) GetApplicationId() *wrappers.StringValue {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

type OAuth2Token struct {
	Token                string                `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	TokenType            *wrappers.StringValue `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	RefreshToken         *wrappers.StringValue `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Expiry               *wrappers.StringValue `protobuf:"bytes,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OAuth2Token) Reset()         { *m = OAuth2Token{} }
func (m *OAuth2Token) String() string { return proto.CompactTextString(m) }
func (*OAuth2Token) ProtoMessage()    {}
func (*OAuth2Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ea951a04c6ae584c, []int{2}
}
func (m *OAuth2Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuth2Token.Unmarshal(m, b)
}
func (m *OAuth2Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuth2Token.Marshal(b, m, deterministic)
}
func (dst *OAuth2Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuth2Token.Merge(dst, src)
}
func (m *OAuth2Token) XXX_Size() int {
	return xxx_messageInfo_OAuth2Token.Size(m)
}
func (m *OAuth2Token) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuth2Token.DiscardUnknown(m)
}

var xxx_messageInfo_OAuth2Token proto.InternalMessageInfo

func (m *OAuth2Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *OAuth2Token) GetTokenType() *wrappers.StringValue {
	if m != nil {
		return m.TokenType
	}
	return nil
}

func (m *OAuth2Token) GetRefreshToken() *wrappers.StringValue {
	if m != nil {
		return m.RefreshToken
	}
	return nil
}

func (m *OAuth2Token) GetExpiry() *wrappers.StringValue {
	if m != nil {
		return m.Expiry
	}
	return nil
}

type Github struct {
	BaseUrl              *wrappers.StringValue `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	UploadUrl            *wrappers.StringValue `protobuf:"bytes,2,opt,name=upload_url,json=uploadUrl,proto3" json:"upload_url,omitempty"`
	Users                []string              `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	Organizations        []string              `protobuf:"bytes,4,rep,name=organizations,proto3" json:"organizations,omitempty"`
	Oauth2               *OAuth2Token          `protobuf:"bytes,10,opt,name=oauth2,proto3" json:"oauth2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Github) Reset()         { *m = Github{} }
func (m *Github) String() string { return proto.CompactTextString(m) }
func (*Github) ProtoMessage()    {}
func (*Github) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ea951a04c6ae584c, []int{3}
}
func (m *Github) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Github.Unmarshal(m, b)
}
func (m *Github) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Github.Marshal(b, m, deterministic)
}
func (dst *Github) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Github.Merge(dst, src)
}
func (m *Github) XXX_Size() int {
	return xxx_messageInfo_Github.Size(m)
}
func (m *Github) XXX_DiscardUnknown() {
	xxx_messageInfo_Github.DiscardUnknown(m)
}

var xxx_messageInfo_Github proto.InternalMessageInfo

func (m *Github) GetBaseUrl() *wrappers.StringValue {
	if m != nil {
		return m.BaseUrl
	}
	return nil
}

func (m *Github) GetUploadUrl() *wrappers.StringValue {
	if m != nil {
		return m.UploadUrl
	}
	return nil
}

func (m *Github) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Github) GetOrganizations() []string {
	if m != nil {
		return m.Organizations
	}
	return nil
}

func (m *Github) GetOauth2() *OAuth2Token {
	if m != nil {
		return m.Oauth2
	}
	return nil
}

type Gitlab struct {
	BaseUrl              *wrappers.StringValue `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	Users                []string              `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	Groups               []string              `protobuf:"bytes,4,rep,name=groups,proto3" json:"groups,omitempty"`
	Private              *OAuthToken           `protobuf:"bytes,10,opt,name=private,proto3" json:"private,omitempty"`
	Oauth                *OAuthToken           `protobuf:"bytes,11,opt,name=oauth,proto3" json:"oauth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Gitlab) Reset()         { *m = Gitlab{} }
func (m *Gitlab) String() string { return proto.CompactTextString(m) }
func (*Gitlab) ProtoMessage()    {}
func (*Gitlab) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ea951a04c6ae584c, []int{4}
}
func (m *Gitlab) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gitlab.Unmarshal(m, b)
}
func (m *Gitlab) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gitlab.Marshal(b, m, deterministic)
}
func (dst *Gitlab) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gitlab.Merge(dst, src)
}
func (m *Gitlab) XXX_Size() int {
	return xxx_messageInfo_Gitlab.Size(m)
}
func (m *Gitlab) XXX_DiscardUnknown() {
	xxx_messageInfo_Gitlab.DiscardUnknown(m)
}

var xxx_messageInfo_Gitlab proto.InternalMessageInfo

func (m *Gitlab) GetBaseUrl() *wrappers.StringValue {
	if m != nil {
		return m.BaseUrl
	}
	return nil
}

func (m *Gitlab) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Gitlab) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *Gitlab) GetPrivate() *OAuthToken {
	if m != nil {
		return m.Private
	}
	return nil
}

func (m *Gitlab) GetOauth() *OAuthToken {
	if m != nil {
		return m.Oauth
	}
	return nil
}

type Bitbucket struct {
	Users                []string    `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	Teams                []string    `protobuf:"bytes,4,rep,name=teams,proto3" json:"teams,omitempty"`
	Basic                *Basic      `protobuf:"bytes,10,opt,name=basic,proto3" json:"basic,omitempty"`
	Oauth                *OAuthToken `protobuf:"bytes,11,opt,name=oauth,proto3" json:"oauth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Bitbucket) Reset()         { *m = Bitbucket{} }
func (m *Bitbucket) String() string { return proto.CompactTextString(m) }
func (*Bitbucket) ProtoMessage()    {}
func (*Bitbucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ea951a04c6ae584c, []int{5}
}
func (m *Bitbucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bitbucket.Unmarshal(m, b)
}
func (m *Bitbucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bitbucket.Marshal(b, m, deterministic)
}
func (dst *Bitbucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bitbucket.Merge(dst, src)
}
func (m *Bitbucket) XXX_Size() int {
	return xxx_messageInfo_Bitbucket.Size(m)
}
func (m *Bitbucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bitbucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bitbucket proto.InternalMessageInfo

func (m *Bitbucket) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Bitbucket) GetTeams() []string {
	if m != nil {
		return m.Teams
	}
	return nil
}

func (m *Bitbucket) GetBasic() *Basic {
	if m != nil {
		return m.Basic
	}
	return nil
}

func (m *Bitbucket) GetOauth() *OAuthToken {
	if m != nil {
		return m.Oauth
	}
	return nil
}

type Generic struct {
	BaseUrl              string   `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	PerPageParameter     string   `protobuf:"bytes,3,opt,name=per_page_parameter,json=perPageParameter,proto3" json:"per_page_parameter,omitempty"`
	PageParameter        string   `protobuf:"bytes,4,opt,name=page_parameter,json=pageParameter,proto3" json:"page_parameter,omitempty"`
	PageSize             int32    `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Selector             string   `protobuf:"bytes,6,opt,name=selector,proto3" json:"selector,omitempty"`
	Basic                *Basic   `protobuf:"bytes,10,opt,name=basic,proto3" json:"basic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Generic) Reset()         { *m = Generic{} }
func (m *Generic) String() string { return proto.CompactTextString(m) }
func (*Generic) ProtoMessage()    {}
func (*Generic) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ea951a04c6ae584c, []int{6}
}
func (m *Generic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Generic.Unmarshal(m, b)
}
func (m *Generic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Generic.Marshal(b, m, deterministic)
}
func (dst *Generic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Generic.Merge(dst, src)
}
func (m *Generic) XXX_Size() int {
	return xxx_messageInfo_Generic.Size(m)
}
func (m *Generic) XXX_DiscardUnknown() {
	xxx_messageInfo_Generic.DiscardUnknown(m)
}

var xxx_messageInfo_Generic proto.InternalMessageInfo

func (m *Generic) GetBaseUrl() string {
	if m != nil {
		return m.BaseUrl
	}
	return ""
}

func (m *Generic) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Generic) GetPerPageParameter() string {
	if m != nil {
		return m.PerPageParameter
	}
	return ""
}

func (m *Generic) GetPageParameter() string {
	if m != nil {
		return m.PageParameter
	}
	return ""
}

func (m *Generic) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Generic) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *Generic) GetBasic() *Basic {
	if m != nil {
		return m.Basic
	}
	return nil
}

type Static struct {
	RepositoryUrls       []string `protobuf:"bytes,1,rep,name=repository_urls,json=repositoryUrls,proto3" json:"repository_urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Static) Reset()         { *m = Static{} }
func (m *Static) String() string { return proto.CompactTextString(m) }
func (*Static) ProtoMessage()    {}
func (*Static) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ea951a04c6ae584c, []int{7}
}
func (m *Static) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Static.Unmarshal(m, b)
}
func (m *Static) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Static.Marshal(b, m, deterministic)
}
func (dst *Static) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Static.Merge(dst, src)
}
func (m *Static) XXX_Size() int {
	return xxx_messageInfo_Static.Size(m)
}
func (m *Static) XXX_DiscardUnknown() {
	xxx_messageInfo_Static.DiscardUnknown(m)
}

var xxx_messageInfo_Static proto.InternalMessageInfo

func (m *Static) GetRepositoryUrls() []string {
	if m != nil {
		return m.RepositoryUrls
	}
	return nil
}

type Rds struct {
	Target               string   `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rds) Reset()         { *m = Rds{} }
func (m *Rds) String() string { return proto.CompactTextString(m) }
func (*Rds) ProtoMessage()    {}
func (*Rds) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ea951a04c6ae584c, []int{8}
}
func (m *Rds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rds.Unmarshal(m, b)
}
func (m *Rds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rds.Marshal(b, m, deterministic)
}
func (dst *Rds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rds.Merge(dst, src)
}
func (m *Rds) XXX_Size() int {
	return xxx_messageInfo_Rds.Size(m)
}
func (m *Rds) XXX_DiscardUnknown() {
	xxx_messageInfo_Rds.DiscardUnknown(m)
}

var xxx_messageInfo_Rds proto.InternalMessageInfo

func (m *Rds) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type Account struct {
	Github               *Github    `protobuf:"bytes,1,opt,name=github,proto3" json:"github,omitempty"`
	Gitlab               *Gitlab    `protobuf:"bytes,2,opt,name=gitlab,proto3" json:"gitlab,omitempty"`
	Bitbucket            *Bitbucket `protobuf:"bytes,3,opt,name=bitbucket,proto3" json:"bitbucket,omitempty"`
	Generic              *Generic   `protobuf:"bytes,4,opt,name=generic,proto3" json:"generic,omitempty"`
	Static               *Static    `protobuf:"bytes,5,opt,name=static,proto3" json:"static,omitempty"`
	Rds                  *Rds       `protobuf:"bytes,6,opt,name=rds,proto3" json:"rds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ea951a04c6ae584c, []int{9}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (dst *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(dst, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetGithub() *Github {
	if m != nil {
		return m.Github
	}
	return nil
}

func (m *Account) GetGitlab() *Gitlab {
	if m != nil {
		return m.Gitlab
	}
	return nil
}

func (m *Account) GetBitbucket() *Bitbucket {
	if m != nil {
		return m.Bitbucket
	}
	return nil
}

func (m *Account) GetGeneric() *Generic {
	if m != nil {
		return m.Generic
	}
	return nil
}

func (m *Account) GetStatic() *Static {
	if m != nil {
		return m.Static
	}
	return nil
}

func (m *Account) GetRds() *Rds {
	if m != nil {
		return m.Rds
	}
	return nil
}

type Configuration struct {
	Accounts             []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ea951a04c6ae584c, []int{10}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (dst *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(dst, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Basic)(nil), "cloud.deps.rds.config.Basic")
	proto.RegisterType((*OAuthToken)(nil), "cloud.deps.rds.config.OAuthToken")
	proto.RegisterType((*OAuth2Token)(nil), "cloud.deps.rds.config.OAuth2Token")
	proto.RegisterType((*Github)(nil), "cloud.deps.rds.config.Github")
	proto.RegisterType((*Gitlab)(nil), "cloud.deps.rds.config.Gitlab")
	proto.RegisterType((*Bitbucket)(nil), "cloud.deps.rds.config.Bitbucket")
	proto.RegisterType((*Generic)(nil), "cloud.deps.rds.config.Generic")
	proto.RegisterType((*Static)(nil), "cloud.deps.rds.config.Static")
	proto.RegisterType((*Rds)(nil), "cloud.deps.rds.config.Rds")
	proto.RegisterType((*Account)(nil), "cloud.deps.rds.config.Account")
	proto.RegisterType((*Configuration)(nil), "cloud.deps.rds.config.Configuration")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_config_ea951a04c6ae584c) }

var fileDescriptor_config_ea951a04c6ae584c = []byte{
	// 752 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdf, 0x6e, 0xfb, 0x34,
	0x14, 0x56, 0x7e, 0x6d, 0xd3, 0xe6, 0x74, 0x1d, 0xc8, 0x1a, 0x28, 0x8c, 0x0d, 0x95, 0x08, 0xc4,
	0x2e, 0xa6, 0x54, 0x14, 0xd0, 0xc6, 0x26, 0x21, 0x6d, 0xbb, 0x98, 0x10, 0x17, 0x4c, 0xd9, 0xc6,
	0x05, 0x12, 0xaa, 0x9c, 0xc4, 0x4b, 0xad, 0xa5, 0xb1, 0x65, 0x3b, 0x8c, 0xee, 0x79, 0x78, 0x26,
	0x24, 0x1e, 0x80, 0x57, 0xe0, 0x82, 0x2b, 0xe4, 0x3f, 0x59, 0xb7, 0x69, 0x19, 0xd5, 0xb8, 0xf3,
	0xb1, 0xbf, 0xcf, 0xe7, 0xf3, 0xf9, 0x8e, 0x0f, 0x6c, 0x64, 0xac, 0xba, 0xa1, 0x45, 0xcc, 0x05,
	0x53, 0x0c, 0x7d, 0x90, 0x95, 0xac, 0xce, 0xe3, 0x9c, 0x70, 0x19, 0x8b, 0x5c, 0xc6, 0xf6, 0x70,
	0xfb, 0xdb, 0x82, 0xaa, 0x79, 0x9d, 0xc6, 0x19, 0x5b, 0x4c, 0x0a, 0x56, 0xe2, 0xaa, 0x98, 0x18,
	0x7c, 0x5a, 0xdf, 0x4c, 0xb8, 0x5a, 0x72, 0x22, 0x27, 0x77, 0x02, 0x73, 0x4e, 0xc4, 0x6a, 0x61,
	0x6f, 0x8c, 0x7e, 0x81, 0xde, 0x29, 0x96, 0x34, 0x43, 0xdb, 0x30, 0xa8, 0x25, 0x11, 0x15, 0x5e,
	0x90, 0xd0, 0x1b, 0x7b, 0x7b, 0x41, 0xf2, 0x10, 0xa3, 0x43, 0x18, 0x70, 0x2c, 0xe5, 0x1d, 0x13,
	0x79, 0xf8, 0x6e, 0xec, 0xed, 0x0d, 0xa7, 0x3b, 0x71, 0xc1, 0x58, 0x51, 0x92, 0xb8, 0xc9, 0x13,
	0x5f, 0x2a, 0x41, 0xab, 0xe2, 0x27, 0x5c, 0xd6, 0x24, 0x79, 0x40, 0x47, 0x05, 0xc0, 0x8f, 0x27,
	0xb5, 0x9a, 0x5f, 0xb1, 0x5b, 0x52, 0xa1, 0x2d, 0xe8, 0x29, 0xbd, 0x70, 0x09, 0x6c, 0x80, 0xce,
	0x60, 0x13, 0x73, 0x5e, 0xd2, 0x0c, 0x2b, 0xca, 0xaa, 0x19, 0x5d, 0x2f, 0xc7, 0xe8, 0x11, 0xe7,
	0xfb, 0x3c, 0xfa, 0xd3, 0x83, 0xa1, 0xc9, 0x34, 0x7d, 0x2d, 0xd5, 0x31, 0x80, 0x59, 0xcc, 0x74,
	0x51, 0xd6, 0x4a, 0x13, 0x18, 0xfc, 0xd5, 0x92, 0x13, 0x74, 0x02, 0x23, 0x41, 0x6e, 0x04, 0x91,
	0xf3, 0x99, 0xbd, 0xba, 0xb3, 0x06, 0x7f, 0xc3, 0x51, 0xac, 0xaa, 0xaf, 0xc1, 0x27, 0xbf, 0x71,
	0x2a, 0x96, 0x61, 0x77, 0x0d, 0xae, 0xc3, 0x46, 0x7f, 0x7b, 0xe0, 0x9f, 0x1b, 0x87, 0xd1, 0x01,
	0x0c, 0x52, 0x2c, 0xc9, 0xac, 0x16, 0xa5, 0x79, 0xd9, 0x7f, 0x5d, 0xd1, 0xd7, 0xe8, 0x6b, 0x51,
	0xea, 0x97, 0xd7, 0xbc, 0x64, 0x38, 0x37, 0xd4, 0xb5, 0x5e, 0x6e, 0xf1, 0x9a, 0xbc, 0x05, 0x3d,
	0xdd, 0x0b, 0x32, 0xec, 0x8c, 0x3b, 0xba, 0x98, 0x26, 0x40, 0x9f, 0xc1, 0x88, 0x89, 0x02, 0x57,
	0xf4, 0xde, 0x98, 0x20, 0xc3, 0xae, 0x39, 0x7d, 0xba, 0x89, 0x8e, 0xc0, 0x67, 0x58, 0xfb, 0x12,
	0x82, 0x49, 0x1a, 0xc5, 0x2f, 0xf6, 0x70, 0xfc, 0xc8, 0xbc, 0xc4, 0x31, 0xa2, 0xbf, 0xec, 0xc3,
	0x4b, 0xfc, 0x3f, 0x1e, 0xfe, 0xb2, 0xf6, 0x0f, 0xc1, 0x2f, 0x04, 0xab, 0x79, 0x23, 0xda, 0x45,
	0xe8, 0x18, 0xfa, 0x5c, 0xd0, 0x5f, 0xb1, 0x22, 0x4e, 0xee, 0xa7, 0xaf, 0xc9, 0xb5, 0x6a, 0x1b,
	0x06, 0x3a, 0x80, 0x9e, 0x11, 0x1e, 0x0e, 0xd7, 0xa5, 0x5a, 0x7c, 0xf4, 0xbb, 0x07, 0xc1, 0x29,
	0x55, 0x69, 0x9d, 0xdd, 0x12, 0xd5, 0xa2, 0x58, 0x37, 0x34, 0xc1, 0x8b, 0x46, 0xb0, 0x0d, 0xd0,
	0x14, 0x7a, 0xa9, 0xfe, 0xbe, 0x4e, 0xed, 0x4e, 0x4b, 0x4a, 0xf3, 0xc5, 0x13, 0x0b, 0x7d, 0xbb,
	0xcc, 0x7f, 0x3c, 0xe8, 0x9f, 0x93, 0x8a, 0x08, 0x9a, 0xa1, 0x8f, 0x9e, 0xf9, 0x11, 0xac, 0x2a,
	0x8e, 0xa0, 0xcb, 0xb1, 0x9a, 0x9b, 0x26, 0x0b, 0x12, 0xb3, 0x46, 0xfb, 0x80, 0x38, 0x11, 0x33,
	0x8e, 0x0b, 0x32, 0xe3, 0x58, 0xe0, 0x05, 0x51, 0x44, 0x98, 0x0f, 0x14, 0x24, 0xef, 0x73, 0x22,
	0x2e, 0x70, 0x41, 0x2e, 0x9a, 0x7d, 0xf4, 0x39, 0x6c, 0x3e, 0x43, 0x76, 0x0d, 0x72, 0xc4, 0x9f,
	0xc0, 0x3e, 0x86, 0xc0, 0xc0, 0x24, 0xbd, 0x27, 0x61, 0x6f, 0xec, 0xed, 0xf5, 0xf4, 0xe4, 0x29,
	0xc8, 0x25, 0xbd, 0x27, 0x7a, 0x9e, 0x49, 0x52, 0x92, 0x4c, 0x31, 0x11, 0xfa, 0x76, 0x9e, 0x35,
	0xf1, 0x5b, 0xaa, 0x16, 0x7d, 0x09, 0xfe, 0xa5, 0xc2, 0x8a, 0x66, 0xe8, 0x0b, 0x78, 0x4f, 0x10,
	0xce, 0x24, 0x55, 0x4c, 0x2c, 0x75, 0x01, 0x64, 0xe8, 0x19, 0x4f, 0x36, 0x57, 0xdb, 0xd7, 0xa2,
	0x94, 0xd1, 0x2e, 0x74, 0x92, 0xdc, 0xf4, 0x9a, 0xc2, 0xa2, 0x20, 0xca, 0x15, 0xca, 0x45, 0xd1,
	0x1f, 0xef, 0xa0, 0x7f, 0x92, 0x65, 0xac, 0xae, 0x14, 0xfa, 0x06, 0x7c, 0x3b, 0xc3, 0x5d, 0x73,
	0xef, 0xb6, 0x48, 0xb2, 0x63, 0x20, 0x71, 0x60, 0x47, 0x2b, 0x71, 0xea, 0x7e, 0xf4, 0x2b, 0xb4,
	0x12, 0x5b, 0x9a, 0xfe, 0x4c, 0xdf, 0x41, 0x90, 0x36, 0xed, 0xe6, 0xa6, 0xd8, 0xb8, 0xad, 0x06,
	0x0d, 0x2e, 0x59, 0x51, 0xd0, 0x21, 0xf4, 0x0b, 0xdb, 0x07, 0x6e, 0x8e, 0x7d, 0xd2, 0x96, 0xd7,
	0xa2, 0x92, 0x06, 0xae, 0x05, 0x4b, 0x53, 0x45, 0xe3, 0x57, 0xbb, 0x60, 0x5b, 0xea, 0xc4, 0x81,
	0xd1, 0x3e, 0x74, 0x44, 0x2e, 0x8d, 0x8f, 0xc3, 0xe9, 0x76, 0x0b, 0x27, 0xc9, 0x65, 0xa2, 0x61,
	0xd1, 0x0f, 0x30, 0x3a, 0x33, 0x5b, 0xb5, 0x30, 0x43, 0x08, 0x1d, 0xc1, 0x00, 0xdb, 0x42, 0x5b,
	0xab, 0xda, 0x05, 0x3b, 0x3f, 0x92, 0x07, 0xfc, 0xe9, 0xe0, 0x67, 0xdf, 0x9e, 0xa5, 0xbe, 0x19,
	0x34, 0x5f, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x6f, 0x35, 0x0c, 0x93, 0x07, 0x00, 0x00,
}
