package core

//#include "qurlquery.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QUrlQuery struct {
	ptr unsafe.Pointer
}

type QUrlQuery_ITF interface {
	QUrlQuery_PTR() *QUrlQuery
}

func (p *QUrlQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QUrlQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQUrlQuery(ptr QUrlQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUrlQuery_PTR().Pointer()
	}
	return nil
}

func NewQUrlQueryFromPointer(ptr unsafe.Pointer) *QUrlQuery {
	var n = new(QUrlQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QUrlQuery) QUrlQuery_PTR() *QUrlQuery {
	return ptr
}

func NewQUrlQuery() *QUrlQuery {
	return NewQUrlQueryFromPointer(C.QUrlQuery_NewQUrlQuery())
}

func NewQUrlQuery3(queryString string) *QUrlQuery {
	return NewQUrlQueryFromPointer(C.QUrlQuery_NewQUrlQuery3(C.CString(queryString)))
}

func NewQUrlQuery2(url QUrl_ITF) *QUrlQuery {
	return NewQUrlQueryFromPointer(C.QUrlQuery_NewQUrlQuery2(PointerFromQUrl(url)))
}

func NewQUrlQuery4(other QUrlQuery_ITF) *QUrlQuery {
	return NewQUrlQueryFromPointer(C.QUrlQuery_NewQUrlQuery4(PointerFromQUrlQuery(other)))
}

func (ptr *QUrlQuery) AddQueryItem(key string, value string) {
	if ptr.Pointer() != nil {
		C.QUrlQuery_AddQueryItem(ptr.Pointer(), C.CString(key), C.CString(value))
	}
}

func (ptr *QUrlQuery) AllQueryItemValues(key string, encoding QUrl__ComponentFormattingOption) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QUrlQuery_AllQueryItemValues(ptr.Pointer(), C.CString(key), C.int(encoding))), "|")
	}
	return make([]string, 0)
}

func (ptr *QUrlQuery) Clear() {
	if ptr.Pointer() != nil {
		C.QUrlQuery_Clear(ptr.Pointer())
	}
}

func (ptr *QUrlQuery) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QUrlQuery_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUrlQuery) Query(encoding QUrl__ComponentFormattingOption) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUrlQuery_Query(ptr.Pointer(), C.int(encoding)))
	}
	return ""
}

func (ptr *QUrlQuery) RemoveAllQueryItems(key string) {
	if ptr.Pointer() != nil {
		C.QUrlQuery_RemoveAllQueryItems(ptr.Pointer(), C.CString(key))
	}
}

func (ptr *QUrlQuery) RemoveQueryItem(key string) {
	if ptr.Pointer() != nil {
		C.QUrlQuery_RemoveQueryItem(ptr.Pointer(), C.CString(key))
	}
}

func (ptr *QUrlQuery) SetQuery(queryString string) {
	if ptr.Pointer() != nil {
		C.QUrlQuery_SetQuery(ptr.Pointer(), C.CString(queryString))
	}
}

func (ptr *QUrlQuery) SetQueryDelimiters(valueDelimiter QChar_ITF, pairDelimiter QChar_ITF) {
	if ptr.Pointer() != nil {
		C.QUrlQuery_SetQueryDelimiters(ptr.Pointer(), PointerFromQChar(valueDelimiter), PointerFromQChar(pairDelimiter))
	}
}

func (ptr *QUrlQuery) Swap(other QUrlQuery_ITF) {
	if ptr.Pointer() != nil {
		C.QUrlQuery_Swap(ptr.Pointer(), PointerFromQUrlQuery(other))
	}
}

func (ptr *QUrlQuery) ToString(encoding QUrl__ComponentFormattingOption) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUrlQuery_ToString(ptr.Pointer(), C.int(encoding)))
	}
	return ""
}

func (ptr *QUrlQuery) DestroyQUrlQuery() {
	if ptr.Pointer() != nil {
		C.QUrlQuery_DestroyQUrlQuery(ptr.Pointer())
	}
}
