#include "qstyleoptionfocusrect.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionFocusRect>
#include "_cgo_export.h"

class MyQStyleOptionFocusRect: public QStyleOptionFocusRect {
public:
};

void* QStyleOptionFocusRect_NewQStyleOptionFocusRect(){
	return new QStyleOptionFocusRect();
}

void* QStyleOptionFocusRect_NewQStyleOptionFocusRect2(void* other){
	return new QStyleOptionFocusRect(*static_cast<QStyleOptionFocusRect*>(other));
}

