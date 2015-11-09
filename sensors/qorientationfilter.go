package sensors

//#include "qorientationfilter.h"
import "C"
import (
	"unsafe"
)

type QOrientationFilter struct {
	QSensorFilter
}

type QOrientationFilter_ITF interface {
	QSensorFilter_ITF
	QOrientationFilter_PTR() *QOrientationFilter
}

func PointerFromQOrientationFilter(ptr QOrientationFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOrientationFilter_PTR().Pointer()
	}
	return nil
}

func NewQOrientationFilterFromPointer(ptr unsafe.Pointer) *QOrientationFilter {
	var n = new(QOrientationFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOrientationFilter) QOrientationFilter_PTR() *QOrientationFilter {
	return ptr
}

func (ptr *QOrientationFilter) Filter(reading QOrientationReading_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QOrientationFilter_Filter(ptr.Pointer(), PointerFromQOrientationReading(reading)) != 0
	}
	return false
}
