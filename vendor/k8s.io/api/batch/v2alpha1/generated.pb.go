/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this ***REMOVED***le except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the speci***REMOVED***c language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/api/batch/v2alpha1/generated.proto

/*
	Package v2alpha1 is a generated protocol buffer package.

	It is generated from these ***REMOVED***les:
		k8s.io/kubernetes/vendor/k8s.io/api/batch/v2alpha1/generated.proto

	It has these top-level messages:
		CronJob
		CronJobList
		CronJobSpec
		CronJobStatus
		JobTemplate
		JobTemplateSpec
*/
package v2alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_api_core_v1 "k8s.io/api/core/v1"
import k8s_io_apimachinery_pkg_apis_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated ***REMOVED***le
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *CronJob) Reset()                    { *m = CronJob{} }
func (*CronJob) ProtoMessage()               {}
func (*CronJob) Descriptor() ([]byte, []int) { return ***REMOVED***leDescriptorGenerated, []int{0} }

func (m *CronJobList) Reset()                    { *m = CronJobList{} }
func (*CronJobList) ProtoMessage()               {}
func (*CronJobList) Descriptor() ([]byte, []int) { return ***REMOVED***leDescriptorGenerated, []int{1} }

func (m *CronJobSpec) Reset()                    { *m = CronJobSpec{} }
func (*CronJobSpec) ProtoMessage()               {}
func (*CronJobSpec) Descriptor() ([]byte, []int) { return ***REMOVED***leDescriptorGenerated, []int{2} }

func (m *CronJobStatus) Reset()                    { *m = CronJobStatus{} }
func (*CronJobStatus) ProtoMessage()               {}
func (*CronJobStatus) Descriptor() ([]byte, []int) { return ***REMOVED***leDescriptorGenerated, []int{3} }

func (m *JobTemplate) Reset()                    { *m = JobTemplate{} }
func (*JobTemplate) ProtoMessage()               {}
func (*JobTemplate) Descriptor() ([]byte, []int) { return ***REMOVED***leDescriptorGenerated, []int{4} }

func (m *JobTemplateSpec) Reset()                    { *m = JobTemplateSpec{} }
func (*JobTemplateSpec) ProtoMessage()               {}
func (*JobTemplateSpec) Descriptor() ([]byte, []int) { return ***REMOVED***leDescriptorGenerated, []int{5} }

func init() {
	proto.RegisterType((*CronJob)(nil), "k8s.io.api.batch.v2alpha1.CronJob")
	proto.RegisterType((*CronJobList)(nil), "k8s.io.api.batch.v2alpha1.CronJobList")
	proto.RegisterType((*CronJobSpec)(nil), "k8s.io.api.batch.v2alpha1.CronJobSpec")
	proto.RegisterType((*CronJobStatus)(nil), "k8s.io.api.batch.v2alpha1.CronJobStatus")
	proto.RegisterType((*JobTemplate)(nil), "k8s.io.api.batch.v2alpha1.JobTemplate")
	proto.RegisterType((*JobTemplateSpec)(nil), "k8s.io.api.batch.v2alpha1.JobTemplateSpec")
}
func (m *CronJob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CronJob) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *CronJobList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CronJobList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ListMeta.Size()))
	n4, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *CronJobSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CronJobSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Schedule)))
	i += copy(dAtA[i:], m.Schedule)
	if m.StartingDeadlineSeconds != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.StartingDeadlineSeconds))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ConcurrencyPolicy)))
	i += copy(dAtA[i:], m.ConcurrencyPolicy)
	if m.Suspend != nil {
		dAtA[i] = 0x20
		i++
		if *m.Suspend {
			dAtA[i] = 1
		} ***REMOVED*** {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.JobTemplate.Size()))
	n5, err := m.JobTemplate.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.SuccessfulJobsHistoryLimit != nil {
		dAtA[i] = 0x30
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.SuccessfulJobsHistoryLimit))
	}
	if m.FailedJobsHistoryLimit != nil {
		dAtA[i] = 0x38
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.FailedJobsHistoryLimit))
	}
	return i, nil
}

func (m *CronJobStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CronJobStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Active) > 0 {
		for _, msg := range m.Active {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.LastScheduleTime != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.LastScheduleTime.Size()))
		n6, err := m.LastScheduleTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func (m *JobTemplate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JobTemplate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n7, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Template.Size()))
	n8, err := m.Template.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	return i, nil
}

func (m *JobTemplateSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JobTemplateSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n9, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n9
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
	n10, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n10
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CronJob) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *CronJobList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *CronJobSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.Schedule)
	n += 1 + l + sovGenerated(uint64(l))
	if m.StartingDeadlineSeconds != nil {
		n += 1 + sovGenerated(uint64(*m.StartingDeadlineSeconds))
	}
	l = len(m.ConcurrencyPolicy)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Suspend != nil {
		n += 2
	}
	l = m.JobTemplate.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.SuccessfulJobsHistoryLimit != nil {
		n += 1 + sovGenerated(uint64(*m.SuccessfulJobsHistoryLimit))
	}
	if m.FailedJobsHistoryLimit != nil {
		n += 1 + sovGenerated(uint64(*m.FailedJobsHistoryLimit))
	}
	return n
}

func (m *CronJobStatus) Size() (n int) {
	var l int
	_ = l
	if len(m.Active) > 0 {
		for _, e := range m.Active {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.LastScheduleTime != nil {
		l = m.LastScheduleTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *JobTemplate) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Template.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *JobTemplateSpec) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CronJob) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CronJob{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "CronJobSpec", "CronJobSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "CronJobStatus", "CronJobStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CronJobList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CronJobList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "CronJob", "CronJob", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CronJobSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CronJobSpec{`,
		`Schedule:` + fmt.Sprintf("%v", this.Schedule) + `,`,
		`StartingDeadlineSeconds:` + valueToStringGenerated(this.StartingDeadlineSeconds) + `,`,
		`ConcurrencyPolicy:` + fmt.Sprintf("%v", this.ConcurrencyPolicy) + `,`,
		`Suspend:` + valueToStringGenerated(this.Suspend) + `,`,
		`JobTemplate:` + strings.Replace(strings.Replace(this.JobTemplate.String(), "JobTemplateSpec", "JobTemplateSpec", 1), `&`, ``, 1) + `,`,
		`SuccessfulJobsHistoryLimit:` + valueToStringGenerated(this.SuccessfulJobsHistoryLimit) + `,`,
		`FailedJobsHistoryLimit:` + valueToStringGenerated(this.FailedJobsHistoryLimit) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CronJobStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CronJobStatus{`,
		`Active:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Active), "ObjectReference", "k8s_io_api_core_v1.ObjectReference", 1), `&`, ``, 1) + `,`,
		`LastScheduleTime:` + strings.Replace(fmt.Sprintf("%v", this.LastScheduleTime), "Time", "k8s_io_apimachinery_pkg_apis_meta_v1.Time", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *JobTemplate) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JobTemplate{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Template:` + strings.Replace(strings.Replace(this.Template.String(), "JobTemplateSpec", "JobTemplateSpec", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *JobTemplateSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JobTemplateSpec{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "JobSpec", "k8s_io_api_batch_v1.JobSpec", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CronJob) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
		***REMOVED***eldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CronJob: wiretype end group for non-group")
		}
		if ***REMOVED***eldNum <= 0 {
			return fmt.Errorf("proto: CronJob: illegal tag %d (wire type %d)", ***REMOVED***eldNum, wire)
		}
		switch ***REMOVED***eldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *CronJobList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
		***REMOVED***eldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CronJobList: wiretype end group for non-group")
		}
		if ***REMOVED***eldNum <= 0 {
			return fmt.Errorf("proto: CronJobList: illegal tag %d (wire type %d)", ***REMOVED***eldNum, wire)
		}
		switch ***REMOVED***eldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, CronJob{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *CronJobSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
		***REMOVED***eldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CronJobSpec: wiretype end group for non-group")
		}
		if ***REMOVED***eldNum <= 0 {
			return fmt.Errorf("proto: CronJobSpec: illegal tag %d (wire type %d)", ***REMOVED***eldNum, wire)
		}
		switch ***REMOVED***eldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld Schedule", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schedule = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld StartingDeadlineSeconds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StartingDeadlineSeconds = &v
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld ConcurrencyPolicy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConcurrencyPolicy = ConcurrencyPolicy(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld Suspend", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
			b := bool(v != 0)
			m.Suspend = &b
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld JobTemplate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.JobTemplate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld SuccessfulJobsHistoryLimit", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SuccessfulJobsHistoryLimit = &v
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld FailedJobsHistoryLimit", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FailedJobsHistoryLimit = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *CronJobStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
		***REMOVED***eldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CronJobStatus: wiretype end group for non-group")
		}
		if ***REMOVED***eldNum <= 0 {
			return fmt.Errorf("proto: CronJobStatus: illegal tag %d (wire type %d)", ***REMOVED***eldNum, wire)
		}
		switch ***REMOVED***eldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld Active", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Active = append(m.Active, k8s_io_api_core_v1.ObjectReference{})
			if err := m.Active[len(m.Active)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld LastScheduleTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastScheduleTime == nil {
				m.LastScheduleTime = &k8s_io_apimachinery_pkg_apis_meta_v1.Time{}
			}
			if err := m.LastScheduleTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *JobTemplate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
		***REMOVED***eldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JobTemplate: wiretype end group for non-group")
		}
		if ***REMOVED***eldNum <= 0 {
			return fmt.Errorf("proto: JobTemplate: illegal tag %d (wire type %d)", ***REMOVED***eldNum, wire)
		}
		switch ***REMOVED***eldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld Template", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Template.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *JobTemplateSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
		***REMOVED***eldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JobTemplateSpec: wiretype end group for non-group")
		}
		if ***REMOVED***eldNum <= 0 {
			return fmt.Errorf("proto: JobTemplateSpec: illegal tag %d (wire type %d)", ***REMOVED***eldNum, wire)
		}
		switch ***REMOVED***eldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for ***REMOVED***eld Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/batch/v2alpha1/generated.proto", ***REMOVED***leDescriptorGenerated)
}

var ***REMOVED***leDescriptorGenerated = []byte{
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4d, 0x6f, 0xdb, 0x36,
	0x18, 0xc7, 0x2d, 0xc7, 0x6f, 0xa1, 0x97, 0x2d, 0xd1, 0x86, 0xc4, 0xf3, 0x06, 0xd9, 0x50, 0xb0,
	0xc1, 0x18, 0x36, 0x6a, 0x09, 0x86, 0x61, 0xa7, 0x01, 0x53, 0x86, 0x36, 0x4d, 0x53, 0x34, 0x90,
	0x53, 0xa0, 0x28, 0x82, 0xa2, 0x14, 0x45, 0xdb, 0x8c, 0x25, 0x51, 0x10, 0x29, 0x03, 0xbe, 0xf5,
	0xd6, 0x6b, 0x3f, 0x49, 0x2f, 0xed, 0x87, 0x48, 0x7b, 0xca, 0x31, 0x27, 0xa3, 0x51, 0xbf, 0x45,
	0x4f, 0x85, 0x68, 0xf9, 0x25, 0x7e, 0x49, 0xd2, 0x4b, 0x6e, 0xe2, 0xa3, 0xff, 0xff, 0xc7, 0x87,
	0xcf, 0xf3, 0x90, 0xc0, 0xec, 0xfe, 0xc3, 0x21, 0x65, 0x46, 0x37, 0xb2, 0x49, 0xe8, 0x13, 0x41,
	0xb8, 0xd1, 0x23, 0xbe, 0xc3, 0x42, 0x23, 0xfd, 0x81, 0x02, 0x6a, 0xd8, 0x48, 0xe0, 0x8e, 0xd1,
	0xdb, 0x45, 0x6e, 0xd0, 0x41, 0x3b, 0x46, 0x9b, 0xf8, 0x24, 0x44, 0x82, 0x38, 0x30, 0x08, 0x99,
	0x60, 0xea, 0x8f, 0x43, 0x29, 0x44, 0x01, 0x85, 0x52, 0x0a, 0x47, 0xd2, 0xea, 0x1f, 0x6d, 0x2a,
	0x3a, 0x91, 0x0d, 0x31, 0xf3, 0x8c, 0x36, 0x6b, 0x33, 0x43, 0x3a, 0xec, 0xa8, 0x25, 0x57, 0x72,
	0x21, 0xbf, 0x86, 0xa4, 0xea, 0xf6, 0xfc, 0xa6, 0x73, 0xdb, 0x55, 0xf5, 0x29, 0x11, 0x66, 0x21,
	0x59, 0xa4, 0xf9, 0x6b, 0xa2, 0xf1, 0x10, 0xee, 0x50, 0x9f, 0x84, 0x7d, 0x23, 0xe8, 0xb6, 0x93,
	0x00, 0x37, 0x3c, 0x22, 0xd0, 0x22, 0x97, 0xb1, 0xcc, 0x15, 0x46, 0xbe, 0xa0, 0x1e, 0x99, 0x33,
	0xfc, 0x7d, 0x93, 0x81, 0xe3, 0x0e, 0xf1, 0xd0, 0xac, 0x4f, 0x7f, 0x95, 0x05, 0xc5, 0xbd, 0x90,
	0xf9, 0x07, 0xcc, 0x56, 0x5f, 0x80, 0x52, 0x92, 0x8f, 0x83, 0x04, 0xaa, 0x28, 0x75, 0xa5, 0x51,
	0xde, 0xfd, 0x13, 0x4e, 0x0a, 0x3a, 0xc6, 0xc2, 0xa0, 0xdb, 0x4e, 0x02, 0x1c, 0x26, 0x6a, 0xd8,
	0xdb, 0x81, 0x8f, 0xed, 0x53, 0x82, 0xc5, 0x23, 0x22, 0x90, 0xa9, 0x9e, 0x0d, 0x6a, 0x99, 0x78,
	0x50, 0x03, 0x93, 0x98, 0x35, 0xa6, 0xaa, 0xfb, 0x20, 0xc7, 0x03, 0x82, 0x2b, 0x59, 0x49, 0xff,
	0x15, 0x2e, 0x6d, 0x17, 0x4c, 0x73, 0x6a, 0x06, 0x04, 0x9b, 0xdf, 0xa4, 0xcc, 0x5c, 0xb2, 0xb2,
	0x24, 0x41, 0x3d, 0x02, 0x05, 0x2e, 0x90, 0x88, 0x78, 0x65, 0x45, 0xb2, 0x1a, 0xb7, 0x60, 0x49,
	0xbd, 0xf9, 0x6d, 0x4a, 0x2b, 0x0c, 0xd7, 0x56, 0xca, 0xd1, 0xdf, 0x29, 0xa0, 0x9c, 0x2a, 0x0f,
	0x29, 0x17, 0xea, 0xc9, 0x5c, 0x35, 0xe0, 0xed, 0xaa, 0x91, 0xb8, 0x65, 0x2d, 0xd6, 0xd3, 0x9d,
	0x4a, 0xa3, 0xc8, 0x54, 0x25, 0xee, 0x83, 0x3c, 0x15, 0xc4, 0xe3, 0x95, 0x6c, 0x7d, 0xa5, 0x51,
	0xde, 0xd5, 0x6f, 0x4e, 0xdf, 0x5c, 0x4b, 0x71, 0xf9, 0x07, 0x89, 0xd1, 0x1a, 0xfa, 0xf5, 0x37,
	0xb9, 0x71, 0xda, 0x49, 0x79, 0xd4, 0xdf, 0x41, 0x29, 0x69, 0xb5, 0x13, 0xb9, 0x44, 0xa6, 0xbd,
	0x3a, 0x49, 0xa3, 0x99, 0xc6, 0xad, 0xb1, 0x42, 0x7d, 0x02, 0xb6, 0xb8, 0x40, 0xa1, 0xa0, 0x7e,
	0xfb, 0x7f, 0x82, 0x1c, 0x97, 0xfa, 0xa4, 0x49, 0x30, 0xf3, 0x1d, 0x2e, 0x7b, 0xb4, 0x62, 0xfe,
	0x14, 0x0f, 0x6a, 0x5b, 0xcd, 0xc5, 0x12, 0x6b, 0x99, 0x57, 0x3d, 0x01, 0x1b, 0x98, 0xf9, 0x38,
	0x0a, 0x43, 0xe2, 0xe3, 0xfe, 0x11, 0x73, 0x29, 0xee, 0xcb, 0x46, 0xad, 0x9a, 0x30, 0xcd, 0x66,
	0x63, 0x6f, 0x56, 0xf0, 0x79, 0x51, 0xd0, 0x9a, 0x07, 0xa9, 0xbf, 0x80, 0x22, 0x8f, 0x78, 0x40,
	0x7c, 0xa7, 0x92, 0xab, 0x2b, 0x8d, 0x92, 0x59, 0x8e, 0x07, 0xb5, 0x62, 0x73, 0x18, 0xb2, 0x46,
	0xff, 0x54, 0x04, 0xca, 0xa7, 0xcc, 0x3e, 0x26, 0x5e, 0xe0, 0x22, 0x41, 0x2a, 0x79, 0xd9, 0xc3,
	0xdf, 0xae, 0x29, 0xf4, 0xc1, 0x44, 0x2d, 0xe7, 0xee, 0xfb, 0x34, 0xd5, 0xf2, 0xd4, 0x0f, 0x6b,
	0x9a, 0xa9, 0x3e, 0x07, 0x55, 0x1e, 0x61, 0x4c, 0x38, 0x6f, 0x45, 0xee, 0x01, 0xb3, 0xf9, 0x3e,
	0xe5, 0x82, 0x85, 0xfd, 0x43, 0xea, 0x51, 0x51, 0x29, 0xd4, 0x95, 0x46, 0xde, 0xd4, 0xe2, 0x41,
	0xad, 0xda, 0x5c, 0xaa, 0xb2, 0xae, 0x21, 0xa8, 0x16, 0xd8, 0x6c, 0x21, 0xea, 0x12, 0x67, 0x8e,
	0x5d, 0x94, 0xec, 0x6a, 0x3c, 0xa8, 0x6d, 0xde, 0x5b, 0xa8, 0xb0, 0x96, 0x38, 0xf5, 0x0f, 0x0a,
	0x58, 0xbb, 0x72, 0x23, 0xd4, 0x87, 0xa0, 0x80, 0xb0, 0xa0, 0xbd, 0x64, 0x60, 0x92, 0x61, 0xdc,
	0x9e, 0xae, 0x51, 0xf2, 0xae, 0x4d, 0xee, 0xb8, 0x45, 0x5a, 0x24, 0x69, 0x05, 0x99, 0x5c, 0xa3,
	0xff, 0xa4, 0xd5, 0x4a, 0x11, 0xaa, 0x0b, 0xd6, 0x5d, 0xc4, 0xc5, 0x68, 0xd6, 0x8e, 0xa9, 0x47,
	0x64, 0x97, 0xae, 0x96, 0xfe, 0x9a, 0xeb, 0x93, 0x38, 0xcc, 0x1f, 0xe2, 0x41, 0x6d, 0xfd, 0x70,
	0x86, 0x63, 0xcd, 0x91, 0xf5, 0xf7, 0x0a, 0x98, 0xee, 0xce, 0x1d, 0x3c, 0x61, 0x4f, 0x41, 0x49,
	0x8c, 0x46, 0x2a, 0xfb, 0xd5, 0x23, 0x35, 0xbe, 0x8b, 0xe3, 0x79, 0x1a, 0xd3, 0xf4, 0xb7, 0x0a,
	0xf8, 0x6e, 0x46, 0x7f, 0x07, 0xe7, 0xf9, 0xf7, 0xca, 0x93, 0xfc, 0xf3, 0x82, 0xb3, 0xc8, 0x53,
	0x2c, 0x7b, 0x88, 0x4d, 0x78, 0x76, 0xa9, 0x65, 0xce, 0x2f, 0xb5, 0xcc, 0xc5, 0xa5, 0x96, 0x79,
	0x19, 0x6b, 0xca, 0x59, 0xac, 0x29, 0xe7, 0xb1, 0xa6, 0x5c, 0xc4, 0x9a, 0xf2, 0x31, 0xd6, 0x94,
	0xd7, 0x9f, 0xb4, 0xcc, 0xb3, 0xd2, 0xa8, 0x22, 0x5f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x20, 0x1c,
	0xcf, 0x94, 0xe7, 0x07, 0x00, 0x00,
}
