// Code generated by protoc-gen-go.
// source: request_vote_request.proto
// DO NOT EDIT!

package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type ProtoRequestVoteRequest struct {
	Term             *uint64 `protobuf:"varint,1,req" json:"Term,omitempty"`
	LastLogIndex     *uint64 `protobuf:"varint,2,req" json:"LastLogIndex,omitempty"`
	LastLogTerm      *uint64 `protobuf:"varint,3,req" json:"LastLogTerm,omitempty"`
	CandidateName    *string `protobuf:"bytes,4,req" json:"CandidateName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProtoRequestVoteRequest) Reset()         { *m = ProtoRequestVoteRequest{} }
func (m *ProtoRequestVoteRequest) String() string { return proto.CompactTextString(m) }
func (*ProtoRequestVoteRequest) ProtoMessage()    {}

func (m *ProtoRequestVoteRequest) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

func (m *ProtoRequestVoteRequest) GetLastLogIndex() uint64 {
	if m != nil && m.LastLogIndex != nil {
		return *m.LastLogIndex
	}
	return 0
}

func (m *ProtoRequestVoteRequest) GetLastLogTerm() uint64 {
	if m != nil && m.LastLogTerm != nil {
		return *m.LastLogTerm
	}
	return 0
}

func (m *ProtoRequestVoteRequest) GetCandidateName() string {
	if m != nil && m.CandidateName != nil {
		return *m.CandidateName
	}
	return ""
}

func init() {
}
