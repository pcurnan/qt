package gui

//#include "qtextframeformat.h"
import "C"
import (
	"unsafe"
)

type QTextFrameFormat struct {
	QTextFormat
}

type QTextFrameFormat_ITF interface {
	QTextFormat_ITF
	QTextFrameFormat_PTR() *QTextFrameFormat
}

func PointerFromQTextFrameFormat(ptr QTextFrameFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFrameFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextFrameFormatFromPointer(ptr unsafe.Pointer) *QTextFrameFormat {
	var n = new(QTextFrameFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextFrameFormat) QTextFrameFormat_PTR() *QTextFrameFormat {
	return ptr
}

//QTextFrameFormat::BorderStyle
type QTextFrameFormat__BorderStyle int64

var (
	QTextFrameFormat__BorderStyle_None       = QTextFrameFormat__BorderStyle(0)
	QTextFrameFormat__BorderStyle_Dotted     = QTextFrameFormat__BorderStyle(1)
	QTextFrameFormat__BorderStyle_Dashed     = QTextFrameFormat__BorderStyle(2)
	QTextFrameFormat__BorderStyle_Solid      = QTextFrameFormat__BorderStyle(3)
	QTextFrameFormat__BorderStyle_Double     = QTextFrameFormat__BorderStyle(4)
	QTextFrameFormat__BorderStyle_DotDash    = QTextFrameFormat__BorderStyle(5)
	QTextFrameFormat__BorderStyle_DotDotDash = QTextFrameFormat__BorderStyle(6)
	QTextFrameFormat__BorderStyle_Groove     = QTextFrameFormat__BorderStyle(7)
	QTextFrameFormat__BorderStyle_Ridge      = QTextFrameFormat__BorderStyle(8)
	QTextFrameFormat__BorderStyle_Inset      = QTextFrameFormat__BorderStyle(9)
	QTextFrameFormat__BorderStyle_Outset     = QTextFrameFormat__BorderStyle(10)
)

//QTextFrameFormat::Position
type QTextFrameFormat__Position int64

const (
	QTextFrameFormat__InFlow     = QTextFrameFormat__Position(0)
	QTextFrameFormat__FloatLeft  = QTextFrameFormat__Position(1)
	QTextFrameFormat__FloatRight = QTextFrameFormat__Position(2)
)

func NewQTextFrameFormat() *QTextFrameFormat {
	return NewQTextFrameFormatFromPointer(C.QTextFrameFormat_NewQTextFrameFormat())
}

func (ptr *QTextFrameFormat) BottomMargin() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextFrameFormat_BottomMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) LeftMargin() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextFrameFormat_LeftMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) RightMargin() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextFrameFormat_RightMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) SetMargin(margin float64) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextFrameFormat) TopMargin() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextFrameFormat_TopMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) Border() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextFrameFormat_Border(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) BorderBrush() *QBrush {
	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QTextFrameFormat_BorderBrush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextFrameFormat) BorderStyle() QTextFrameFormat__BorderStyle {
	if ptr.Pointer() != nil {
		return QTextFrameFormat__BorderStyle(C.QTextFrameFormat_BorderStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QTextFrameFormat_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextFrameFormat) Margin() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextFrameFormat_Margin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) Padding() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextFrameFormat_Padding(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) PageBreakPolicy() QTextFormat__PageBreakFlag {
	if ptr.Pointer() != nil {
		return QTextFormat__PageBreakFlag(C.QTextFrameFormat_PageBreakPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) Position() QTextFrameFormat__Position {
	if ptr.Pointer() != nil {
		return QTextFrameFormat__Position(C.QTextFrameFormat_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFrameFormat) SetBorder(width float64) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetBorder(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QTextFrameFormat) SetBorderBrush(brush QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetBorderBrush(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QTextFrameFormat) SetBorderStyle(style QTextFrameFormat__BorderStyle) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetBorderStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QTextFrameFormat) SetBottomMargin(margin float64) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetBottomMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextFrameFormat) SetHeight(height QTextLength_ITF) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetHeight(ptr.Pointer(), PointerFromQTextLength(height))
	}
}

func (ptr *QTextFrameFormat) SetHeight2(height float64) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetHeight2(ptr.Pointer(), C.double(height))
	}
}

func (ptr *QTextFrameFormat) SetLeftMargin(margin float64) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetLeftMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextFrameFormat) SetPadding(width float64) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetPadding(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QTextFrameFormat) SetPageBreakPolicy(policy QTextFormat__PageBreakFlag) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetPageBreakPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QTextFrameFormat) SetPosition(policy QTextFrameFormat__Position) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetPosition(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QTextFrameFormat) SetRightMargin(margin float64) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetRightMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextFrameFormat) SetTopMargin(margin float64) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetTopMargin(ptr.Pointer(), C.double(margin))
	}
}

func (ptr *QTextFrameFormat) SetWidth(width QTextLength_ITF) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetWidth(ptr.Pointer(), PointerFromQTextLength(width))
	}
}

func (ptr *QTextFrameFormat) SetWidth2(width float64) {
	if ptr.Pointer() != nil {
		C.QTextFrameFormat_SetWidth2(ptr.Pointer(), C.double(width))
	}
}
