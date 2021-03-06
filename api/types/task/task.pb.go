// Code generated by protoc-gen-gogo.
// source: github.com/containerd/containerd/api/types/task/task.proto
// DO NOT EDIT!

/*
	Package task is a generated protocol buffer package.

	It is generated from these files:
		github.com/containerd/containerd/api/types/task/task.proto

	It has these top-level messages:
		Task
		Process
		User
*/
package task

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	StatusUnknown Status = 0
	StatusCreated Status = 1
	StatusRunning Status = 2
	StatusStopped Status = 3
	StatusPaused  Status = 4
)

var Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "CREATED",
	2: "RUNNING",
	3: "STOPPED",
	4: "PAUSED",
}
var Status_value = map[string]int32{
	"UNKNOWN": 0,
	"CREATED": 1,
	"RUNNING": 2,
	"STOPPED": 3,
	"PAUSED":  4,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptorTask, []int{0} }

type Task struct {
	ID          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ContainerID string `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Pid         uint32 `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Status      Status `protobuf:"varint,4,opt,name=status,proto3,enum=containerd.v1.types.Status" json:"status,omitempty"`
	Stdin       string `protobuf:"bytes,5,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout      string `protobuf:"bytes,6,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr      string `protobuf:"bytes,7,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Terminal    bool   `protobuf:"varint,8,opt,name=terminal,proto3" json:"terminal,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptorTask, []int{0} }

type Process struct {
	Pid         uint32                `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Args        []string              `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	Env         []string              `protobuf:"bytes,3,rep,name=env" json:"env,omitempty"`
	User        *User                 `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Cwd         string                `protobuf:"bytes,5,opt,name=cwd,proto3" json:"cwd,omitempty"`
	Terminal    bool                  `protobuf:"varint,6,opt,name=terminal,proto3" json:"terminal,omitempty"`
	ExitStatus  uint32                `protobuf:"varint,7,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	Status      Status                `protobuf:"varint,8,opt,name=status,proto3,enum=containerd.v1.types.Status" json:"status,omitempty"`
	RuntimeData *google_protobuf1.Any `protobuf:"bytes,9,opt,name=runtime_data,json=runtimeData" json:"runtime_data,omitempty"`
	Stdin       string                `protobuf:"bytes,10,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout      string                `protobuf:"bytes,11,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr      string                `protobuf:"bytes,12,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (m *Process) Reset()                    { *m = Process{} }
func (*Process) ProtoMessage()               {}
func (*Process) Descriptor() ([]byte, []int) { return fileDescriptorTask, []int{1} }

type User struct {
	Uid            uint32   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid            uint32   `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"`
	AdditionalGids []uint32 `protobuf:"varint,3,rep,packed,name=additional_gids,json=additionalGids" json:"additional_gids,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorTask, []int{2} }

func init() {
	proto.RegisterType((*Task)(nil), "containerd.v1.types.Task")
	proto.RegisterType((*Process)(nil), "containerd.v1.types.Process")
	proto.RegisterType((*User)(nil), "containerd.v1.types.User")
	proto.RegisterEnum("containerd.v1.types.Status", Status_name, Status_value)
}
func (m *Task) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Task) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.ContainerID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.ContainerID)))
		i += copy(dAtA[i:], m.ContainerID)
	}
	if m.Pid != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.Pid))
	}
	if m.Status != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.Status))
	}
	if len(m.Stdin) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stdin)))
		i += copy(dAtA[i:], m.Stdin)
	}
	if len(m.Stdout) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stdout)))
		i += copy(dAtA[i:], m.Stdout)
	}
	if len(m.Stderr) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stderr)))
		i += copy(dAtA[i:], m.Stderr)
	}
	if m.Terminal {
		dAtA[i] = 0x40
		i++
		if m.Terminal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Process) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Process) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Pid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.Pid))
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Env) > 0 {
		for _, s := range m.Env {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.User != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.User.Size()))
		n1, err := m.User.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Cwd) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.Cwd)))
		i += copy(dAtA[i:], m.Cwd)
	}
	if m.Terminal {
		dAtA[i] = 0x30
		i++
		if m.Terminal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ExitStatus != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.ExitStatus))
	}
	if m.Status != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.Status))
	}
	if m.RuntimeData != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.RuntimeData.Size()))
		n2, err := m.RuntimeData.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Stdin) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stdin)))
		i += copy(dAtA[i:], m.Stdin)
	}
	if len(m.Stdout) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stdout)))
		i += copy(dAtA[i:], m.Stdout)
	}
	if len(m.Stderr) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stderr)))
		i += copy(dAtA[i:], m.Stderr)
	}
	return i, nil
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.Uid))
	}
	if m.Gid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.Gid))
	}
	if len(m.AdditionalGids) > 0 {
		dAtA4 := make([]byte, len(m.AdditionalGids)*10)
		var j3 int
		for _, num := range m.AdditionalGids {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTask(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	return i, nil
}

func encodeFixed64Task(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Task(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTask(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Task) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.ContainerID)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Pid != 0 {
		n += 1 + sovTask(uint64(m.Pid))
	}
	if m.Status != 0 {
		n += 1 + sovTask(uint64(m.Status))
	}
	l = len(m.Stdin)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Stdout)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Stderr)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Terminal {
		n += 2
	}
	return n
}

func (m *Process) Size() (n int) {
	var l int
	_ = l
	if m.Pid != 0 {
		n += 1 + sovTask(uint64(m.Pid))
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			l = len(s)
			n += 1 + l + sovTask(uint64(l))
		}
	}
	if len(m.Env) > 0 {
		for _, s := range m.Env {
			l = len(s)
			n += 1 + l + sovTask(uint64(l))
		}
	}
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Cwd)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Terminal {
		n += 2
	}
	if m.ExitStatus != 0 {
		n += 1 + sovTask(uint64(m.ExitStatus))
	}
	if m.Status != 0 {
		n += 1 + sovTask(uint64(m.Status))
	}
	if m.RuntimeData != nil {
		l = m.RuntimeData.Size()
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Stdin)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Stdout)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Stderr)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	return n
}

func (m *User) Size() (n int) {
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovTask(uint64(m.Uid))
	}
	if m.Gid != 0 {
		n += 1 + sovTask(uint64(m.Gid))
	}
	if len(m.AdditionalGids) > 0 {
		l = 0
		for _, e := range m.AdditionalGids {
			l += sovTask(uint64(e))
		}
		n += 1 + sovTask(uint64(l)) + l
	}
	return n
}

func sovTask(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTask(x uint64) (n int) {
	return sovTask(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Task) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Task{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`ContainerID:` + fmt.Sprintf("%v", this.ContainerID) + `,`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Stdin:` + fmt.Sprintf("%v", this.Stdin) + `,`,
		`Stdout:` + fmt.Sprintf("%v", this.Stdout) + `,`,
		`Stderr:` + fmt.Sprintf("%v", this.Stderr) + `,`,
		`Terminal:` + fmt.Sprintf("%v", this.Terminal) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Process) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Process{`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`Args:` + fmt.Sprintf("%v", this.Args) + `,`,
		`Env:` + fmt.Sprintf("%v", this.Env) + `,`,
		`User:` + strings.Replace(fmt.Sprintf("%v", this.User), "User", "User", 1) + `,`,
		`Cwd:` + fmt.Sprintf("%v", this.Cwd) + `,`,
		`Terminal:` + fmt.Sprintf("%v", this.Terminal) + `,`,
		`ExitStatus:` + fmt.Sprintf("%v", this.ExitStatus) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`RuntimeData:` + strings.Replace(fmt.Sprintf("%v", this.RuntimeData), "Any", "google_protobuf1.Any", 1) + `,`,
		`Stdin:` + fmt.Sprintf("%v", this.Stdin) + `,`,
		`Stdout:` + fmt.Sprintf("%v", this.Stdout) + `,`,
		`Stderr:` + fmt.Sprintf("%v", this.Stderr) + `,`,
		`}`,
	}, "")
	return s
}
func (this *User) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&User{`,
		`Uid:` + fmt.Sprintf("%v", this.Uid) + `,`,
		`Gid:` + fmt.Sprintf("%v", this.Gid) + `,`,
		`AdditionalGids:` + fmt.Sprintf("%v", this.AdditionalGids) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTask(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Task) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Task: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Task: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stderr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stderr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Terminal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Terminal = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Process) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Process: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Process: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Env = append(m.Env, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &User{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cwd", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cwd = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Terminal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Terminal = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitStatus", wireType)
			}
			m.ExitStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExitStatus |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RuntimeData == nil {
				m.RuntimeData = &google_protobuf1.Any{}
			}
			if err := m.RuntimeData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stderr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stderr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gid", wireType)
			}
			m.Gid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTask
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.AdditionalGids = append(m.AdditionalGids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTask
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTask
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTask
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.AdditionalGids = append(m.AdditionalGids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalGids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTask(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTask
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTask
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTask
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTask
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTask
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTask(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTask = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTask   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/types/task/task.proto", fileDescriptorTask)
}

var fileDescriptorTask = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x31, 0x4f, 0x1b, 0x3f,
	0x1c, 0xcd, 0x5d, 0xc2, 0x25, 0x38, 0x09, 0xdc, 0xdf, 0x7f, 0x84, 0x8e, 0xb4, 0x3a, 0x4e, 0x2c,
	0x8d, 0x2a, 0xf5, 0xa2, 0xc2, 0x50, 0xa9, 0x1b, 0x90, 0x08, 0x45, 0x95, 0x42, 0xe4, 0x10, 0x75,
	0x8c, 0x4c, 0xec, 0x5e, 0x2d, 0xc0, 0x8e, 0x6c, 0x1f, 0x94, 0xad, 0x63, 0xc5, 0x77, 0x60, 0x6a,
	0x3f, 0x43, 0x87, 0x7e, 0x02, 0xc6, 0x8e, 0x9d, 0x50, 0xc9, 0x57, 0xe8, 0xd8, 0xa5, 0xb2, 0xef,
	0x48, 0x82, 0x44, 0xa5, 0x2e, 0xa7, 0xe7, 0xf7, 0x9e, 0xce, 0xcf, 0xef, 0xf7, 0x03, 0xaf, 0x13,
	0xa6, 0xdf, 0xa7, 0xc7, 0xf1, 0x58, 0x9c, 0xb5, 0xc6, 0x82, 0x6b, 0xcc, 0x38, 0x95, 0x64, 0x11,
	0xe2, 0x09, 0x6b, 0xe9, 0xcb, 0x09, 0x55, 0x2d, 0x8d, 0xd5, 0x89, 0xfd, 0xc4, 0x13, 0x29, 0xb4,
	0x80, 0xff, 0xcf, 0x5d, 0xf1, 0xf9, 0xcb, 0xd8, 0x9a, 0x1a, 0x6b, 0x89, 0x48, 0x84, 0xd5, 0x5b,
	0x06, 0x65, 0xd6, 0xc6, 0x46, 0x22, 0x44, 0x72, 0x4a, 0x5b, 0xf6, 0x74, 0x9c, 0xbe, 0x6b, 0x61,
	0x7e, 0x99, 0x49, 0x5b, 0xbf, 0x1d, 0x50, 0x3a, 0xc2, 0xea, 0x04, 0xae, 0x03, 0x97, 0x91, 0xc0,
	0x89, 0x9c, 0xe6, 0xf2, 0x9e, 0x37, 0xbd, 0xdd, 0x74, 0xbb, 0x6d, 0xe4, 0x32, 0x02, 0xb7, 0x41,
	0x6d, 0x76, 0xd1, 0x88, 0x91, 0xc0, 0xb5, 0x8e, 0xd5, 0xe9, 0xed, 0x66, 0x75, 0xff, 0x9e, 0xef,
	0xb6, 0x51, 0x75, 0x66, 0xea, 0x12, 0xe8, 0x83, 0xe2, 0x84, 0x91, 0xa0, 0x18, 0x39, 0xcd, 0x3a,
	0x32, 0x10, 0xee, 0x00, 0x4f, 0x69, 0xac, 0x53, 0x15, 0x94, 0x22, 0xa7, 0xb9, 0xb2, 0xfd, 0x24,
	0x7e, 0x24, 0x7d, 0x3c, 0xb0, 0x16, 0x94, 0x5b, 0xe1, 0x1a, 0x58, 0x52, 0x9a, 0x30, 0x1e, 0x2c,
	0x99, 0x3b, 0x51, 0x76, 0x80, 0xeb, 0xe6, 0x57, 0x44, 0xa4, 0x3a, 0xf0, 0x2c, 0x9d, 0x9f, 0x72,
	0x9e, 0x4a, 0x19, 0x94, 0x67, 0x3c, 0x95, 0x12, 0x36, 0x40, 0x45, 0x53, 0x79, 0xc6, 0x38, 0x3e,
	0x0d, 0x2a, 0x91, 0xd3, 0xac, 0xa0, 0xd9, 0x79, 0xeb, 0x97, 0x0b, 0xca, 0x7d, 0x29, 0xc6, 0x54,
	0xa9, 0xfb, 0xd0, 0xce, 0x3c, 0x34, 0x04, 0x25, 0x2c, 0x13, 0x15, 0xb8, 0x51, 0xb1, 0xb9, 0x8c,
	0x2c, 0x36, 0x2e, 0xca, 0xcf, 0x83, 0xa2, 0xa5, 0x0c, 0x84, 0x2f, 0x40, 0x29, 0x55, 0x54, 0xda,
	0x87, 0x55, 0xb7, 0x37, 0x1e, 0x7d, 0xd8, 0x50, 0x51, 0x89, 0xac, 0xcd, 0xfc, 0x60, 0x7c, 0x41,
	0xf2, 0x27, 0x19, 0xf8, 0x20, 0xa0, 0xf7, 0x30, 0x20, 0xdc, 0x04, 0x55, 0xfa, 0x81, 0xe9, 0x51,
	0x5e, 0x5e, 0xd9, 0x86, 0x03, 0x86, 0xca, 0xba, 0x5a, 0x28, 0xb6, 0xf2, 0xef, 0xc5, 0xbe, 0x02,
	0x35, 0x99, 0x72, 0xcd, 0xce, 0xe8, 0x88, 0x60, 0x8d, 0x83, 0x65, 0x1b, 0x7d, 0x2d, 0xce, 0xd6,
	0x24, 0xbe, 0x5f, 0x93, 0x78, 0x97, 0x5f, 0xa2, 0x6a, 0xee, 0x6c, 0x63, 0x8d, 0xe7, 0x13, 0x01,
	0x8f, 0x4f, 0xa4, 0xfa, 0x97, 0x89, 0xd4, 0x16, 0x27, 0xb2, 0x35, 0x00, 0xa5, 0x61, 0x5e, 0x45,
	0x3a, 0x6f, 0x3c, 0x65, 0x76, 0x71, 0x92, 0x7c, 0xc7, 0xea, 0xc8, 0x40, 0xf8, 0x0c, 0xac, 0x62,
	0x42, 0x98, 0x66, 0x82, 0xe3, 0xd3, 0x51, 0xc2, 0x88, 0xb2, 0xdd, 0xd7, 0xd1, 0xca, 0x9c, 0x3e,
	0x60, 0x44, 0x3d, 0xff, 0xea, 0x00, 0x2f, 0xef, 0x24, 0x04, 0xe5, 0x61, 0xef, 0x4d, 0xef, 0xf0,
	0x6d, 0xcf, 0x2f, 0x34, 0xfe, 0xbb, 0xba, 0x8e, 0xea, 0x99, 0x30, 0xe4, 0x27, 0x5c, 0x5c, 0x70,
	0xa3, 0xef, 0xa3, 0xce, 0xee, 0x51, 0xa7, 0xed, 0x3b, 0x8b, 0xfa, 0xbe, 0xa4, 0x58, 0x53, 0x62,
	0x74, 0x34, 0xec, 0xf5, 0xba, 0xbd, 0x03, 0xdf, 0x5d, 0xd4, 0x51, 0xca, 0x39, 0xe3, 0x89, 0xd1,
	0x07, 0x47, 0x87, 0xfd, 0x7e, 0xa7, 0xed, 0x17, 0x17, 0xf5, 0x81, 0x16, 0x93, 0x09, 0x25, 0xf0,
	0x29, 0xf0, 0xfa, 0xbb, 0xc3, 0x41, 0xa7, 0xed, 0x97, 0x1a, 0xfe, 0xd5, 0x75, 0x54, 0xcb, 0xe4,
	0x3e, 0x4e, 0x15, 0x25, 0x8d, 0x95, 0x4f, 0x9f, 0xc3, 0xc2, 0xb7, 0x2f, 0x61, 0x9e, 0x76, 0x2f,
	0xb8, 0xb9, 0x0b, 0x0b, 0x3f, 0xee, 0xc2, 0xc2, 0xc7, 0x69, 0xe8, 0xdc, 0x4c, 0x43, 0xe7, 0xfb,
	0x34, 0x74, 0x7e, 0x4e, 0x43, 0xe7, 0xd8, 0xb3, 0x83, 0xd8, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0xcd, 0x9c, 0x52, 0x26, 0x04, 0x00, 0x00,
}
