package xml

//#include "qxmlparseexception.h"
import "C"
import (
	"unsafe"
)

type QXmlParseException struct {
	ptr unsafe.Pointer
}

type QXmlParseException_ITF interface {
	QXmlParseException_PTR() *QXmlParseException
}

func (p *QXmlParseException) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlParseException) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlParseException(ptr QXmlParseException_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlParseException_PTR().Pointer()
	}
	return nil
}

func NewQXmlParseExceptionFromPointer(ptr unsafe.Pointer) *QXmlParseException {
	var n = new(QXmlParseException)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlParseException) QXmlParseException_PTR() *QXmlParseException {
	return ptr
}

func NewQXmlParseException(name string, c int, l int, p string, s string) *QXmlParseException {
	return NewQXmlParseExceptionFromPointer(C.QXmlParseException_NewQXmlParseException(C.CString(name), C.int(c), C.int(l), C.CString(p), C.CString(s)))
}

func NewQXmlParseException2(other QXmlParseException_ITF) *QXmlParseException {
	return NewQXmlParseExceptionFromPointer(C.QXmlParseException_NewQXmlParseException2(PointerFromQXmlParseException(other)))
}

func (ptr *QXmlParseException) ColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QXmlParseException_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlParseException) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QXmlParseException_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlParseException) Message() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlParseException_Message(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlParseException) PublicId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlParseException_PublicId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlParseException) SystemId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlParseException_SystemId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlParseException) DestroyQXmlParseException() {
	if ptr.Pointer() != nil {
		C.QXmlParseException_DestroyQXmlParseException(ptr.Pointer())
	}
}
