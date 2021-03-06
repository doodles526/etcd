// Code generated by protoc-gen-go.
// source: append_entries_request.proto
// DO NOT EDIT!

package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type ProtoAppendEntriesRequest struct {
	Term             *uint64                                    `protobuf:"varint,1,req" json:"Term,omitempty"`
	PrevLogIndex     *uint64                                    `protobuf:"varint,2,req" json:"PrevLogIndex,omitempty"`
	PrevLogTerm      *uint64                                    `protobuf:"varint,3,req" json:"PrevLogTerm,omitempty"`
	CommitIndex      *uint64                                    `protobuf:"varint,4,req" json:"CommitIndex,omitempty"`
	LeaderName       *string                                    `protobuf:"bytes,5,req" json:"LeaderName,omitempty"`
	Entries          []*ProtoAppendEntriesRequest_ProtoLogEntry `protobuf:"bytes,6,rep" json:"Entries,omitempty"`
	XXX_unrecognized []byte                                     `json:"-"`
}

func (m *ProtoAppendEntriesRequest) Reset()         { *m = ProtoAppendEntriesRequest{} }
func (m *ProtoAppendEntriesRequest) String() string { return proto.CompactTextString(m) }
func (*ProtoAppendEntriesRequest) ProtoMessage()    {}

func (m *ProtoAppendEntriesRequest) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

func (m *ProtoAppendEntriesRequest) GetPrevLogIndex() uint64 {
	if m != nil && m.PrevLogIndex != nil {
		return *m.PrevLogIndex
	}
	return 0
}

func (m *ProtoAppendEntriesRequest) GetPrevLogTerm() uint64 {
	if m != nil && m.PrevLogTerm != nil {
		return *m.PrevLogTerm
	}
	return 0
}

func (m *ProtoAppendEntriesRequest) GetCommitIndex() uint64 {
	if m != nil && m.CommitIndex != nil {
		return *m.CommitIndex
	}
	return 0
}

func (m *ProtoAppendEntriesRequest) GetLeaderName() string {
	if m != nil && m.LeaderName != nil {
		return *m.LeaderName
	}
	return ""
}

func (m *ProtoAppendEntriesRequest) GetEntries() []*ProtoAppendEntriesRequest_ProtoLogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type ProtoAppendEntriesRequest_ProtoLogEntry struct {
	Index            *uint64 `protobuf:"varint,1,req" json:"Index,omitempty"`
	Term             *uint64 `protobuf:"varint,2,req" json:"Term,omitempty"`
	CommandName      *string `protobuf:"bytes,3,req" json:"CommandName,omitempty"`
	Command          []byte  `protobuf:"bytes,4,opt" json:"Command,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProtoAppendEntriesRequest_ProtoLogEntry) Reset() {
	*m = ProtoAppendEntriesRequest_ProtoLogEntry{}
}
func (m *ProtoAppendEntriesRequest_ProtoLogEntry) String() string { return proto.CompactTextString(m) }
func (*ProtoAppendEntriesRequest_ProtoLogEntry) ProtoMessage()    {}

func (m *ProtoAppendEntriesRequest_ProtoLogEntry) GetIndex() uint64 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *ProtoAppendEntriesRequest_ProtoLogEntry) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

func (m *ProtoAppendEntriesRequest_ProtoLogEntry) GetCommandName() string {
	if m != nil && m.CommandName != nil {
		return *m.CommandName
	}
	return ""
}

func (m *ProtoAppendEntriesRequest_ProtoLogEntry) GetCommand() []byte {
	if m != nil {
		return m.Command
	}
	return nil
}

func init() {
}
