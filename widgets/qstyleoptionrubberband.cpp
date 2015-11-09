#include "qstyleoptionrubberband.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QStyleOptionRubberBand>
#include "_cgo_export.h"

class MyQStyleOptionRubberBand: public QStyleOptionRubberBand {
public:
};

void* QStyleOptionRubberBand_NewQStyleOptionRubberBand(){
	return new QStyleOptionRubberBand();
}

void* QStyleOptionRubberBand_NewQStyleOptionRubberBand2(void* other){
	return new QStyleOptionRubberBand(*static_cast<QStyleOptionRubberBand*>(other));
}

