#include "qmaskgenerator.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMaskGenerator>
#include "_cgo_export.h"

class MyQMaskGenerator: public QMaskGenerator {
public:
};

int QMaskGenerator_Seed(void* ptr){
	return static_cast<QMaskGenerator*>(ptr)->seed();
}

void QMaskGenerator_DestroyQMaskGenerator(void* ptr){
	static_cast<QMaskGenerator*>(ptr)->~QMaskGenerator();
}

