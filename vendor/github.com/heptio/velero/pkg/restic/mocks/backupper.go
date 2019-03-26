// Code generated by mockery v1.0.0
package mocks

import corev1 "k8s.io/api/core/v1"
import logrus "github.com/sirupsen/logrus"
import mock "github.com/stretchr/testify/mock"

import v1 "github.com/heptio/velero/pkg/apis/velero/v1"

// Backupper is an autogenerated mock type for the Backupper type
type Backupper struct {
	mock.Mock
}

// BackupPodVolumes provides a mock function with given fields: backup, pod, log
func (_m *Backupper) BackupPodVolumes(backup *v1.Backup, pod *corev1.Pod, log logrus.FieldLogger) (map[string]string, []error) {
	ret := _m.Called(backup, pod, log)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(*v1.Backup, *corev1.Pod, logrus.FieldLogger) map[string]string); ok {
		r0 = rf(backup, pod, log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 []error
	if rf, ok := ret.Get(1).(func(*v1.Backup, *corev1.Pod, logrus.FieldLogger) []error); ok {
		r1 = rf(backup, pod, log)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error)
		}
	}

	return r0, r1
}