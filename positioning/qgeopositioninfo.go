package positioning

//#include "qgeopositioninfo.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGeoPositionInfo struct {
	ptr unsafe.Pointer
}

type QGeoPositionInfo_ITF interface {
	QGeoPositionInfo_PTR() *QGeoPositionInfo
}

func (p *QGeoPositionInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoPositionInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoPositionInfo(ptr QGeoPositionInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoPositionInfo_PTR().Pointer()
	}
	return nil
}

func NewQGeoPositionInfoFromPointer(ptr unsafe.Pointer) *QGeoPositionInfo {
	var n = new(QGeoPositionInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoPositionInfo) QGeoPositionInfo_PTR() *QGeoPositionInfo {
	return ptr
}

//QGeoPositionInfo::Attribute
type QGeoPositionInfo__Attribute int64

const (
	QGeoPositionInfo__Direction          = QGeoPositionInfo__Attribute(0)
	QGeoPositionInfo__GroundSpeed        = QGeoPositionInfo__Attribute(1)
	QGeoPositionInfo__VerticalSpeed      = QGeoPositionInfo__Attribute(2)
	QGeoPositionInfo__MagneticVariation  = QGeoPositionInfo__Attribute(3)
	QGeoPositionInfo__HorizontalAccuracy = QGeoPositionInfo__Attribute(4)
	QGeoPositionInfo__VerticalAccuracy   = QGeoPositionInfo__Attribute(5)
)

func NewQGeoPositionInfo() *QGeoPositionInfo {
	return NewQGeoPositionInfoFromPointer(C.QGeoPositionInfo_NewQGeoPositionInfo())
}

func NewQGeoPositionInfo2(coordinate QGeoCoordinate_ITF, timestamp core.QDateTime_ITF) *QGeoPositionInfo {
	return NewQGeoPositionInfoFromPointer(C.QGeoPositionInfo_NewQGeoPositionInfo2(PointerFromQGeoCoordinate(coordinate), core.PointerFromQDateTime(timestamp)))
}

func NewQGeoPositionInfo3(other QGeoPositionInfo_ITF) *QGeoPositionInfo {
	return NewQGeoPositionInfoFromPointer(C.QGeoPositionInfo_NewQGeoPositionInfo3(PointerFromQGeoPositionInfo(other)))
}

func (ptr *QGeoPositionInfo) Attribute(attribute QGeoPositionInfo__Attribute) float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGeoPositionInfo_Attribute(ptr.Pointer(), C.int(attribute)))
	}
	return 0
}

func (ptr *QGeoPositionInfo) HasAttribute(attribute QGeoPositionInfo__Attribute) bool {
	if ptr.Pointer() != nil {
		return C.QGeoPositionInfo_HasAttribute(ptr.Pointer(), C.int(attribute)) != 0
	}
	return false
}

func (ptr *QGeoPositionInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QGeoPositionInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoPositionInfo) RemoveAttribute(attribute QGeoPositionInfo__Attribute) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_RemoveAttribute(ptr.Pointer(), C.int(attribute))
	}
}

func (ptr *QGeoPositionInfo) SetAttribute(attribute QGeoPositionInfo__Attribute, value float64) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_SetAttribute(ptr.Pointer(), C.int(attribute), C.double(value))
	}
}

func (ptr *QGeoPositionInfo) SetCoordinate(coordinate QGeoCoordinate_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_SetCoordinate(ptr.Pointer(), PointerFromQGeoCoordinate(coordinate))
	}
}

func (ptr *QGeoPositionInfo) SetTimestamp(timestamp core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_SetTimestamp(ptr.Pointer(), core.PointerFromQDateTime(timestamp))
	}
}

func (ptr *QGeoPositionInfo) Timestamp() *core.QDateTime {
	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QGeoPositionInfo_Timestamp(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGeoPositionInfo) DestroyQGeoPositionInfo() {
	if ptr.Pointer() != nil {
		C.QGeoPositionInfo_DestroyQGeoPositionInfo(ptr.Pointer())
	}
}
