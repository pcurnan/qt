#ifdef __cplusplus
extern "C" {
#endif

void* QPainter_NewQPainter2(void* device);
int QPainter_Begin(void* ptr, void* device);
void QPainter_DrawArc(void* ptr, void* rectangle, int startAngle, int spanAngle);
void QPainter_DrawChord(void* ptr, void* rectangle, int startAngle, int spanAngle);
void QPainter_DrawConvexPolygon2(void* ptr, void* points, int pointCount);
void QPainter_DrawConvexPolygon(void* ptr, void* points, int pointCount);
void QPainter_DrawEllipse2(void* ptr, void* rectangle);
void QPainter_DrawEllipse(void* ptr, void* rectangle);
void QPainter_DrawGlyphRun(void* ptr, void* position, void* glyphs);
void QPainter_DrawImage3(void* ptr, void* point, void* image);
void QPainter_DrawImage(void* ptr, void* target, void* image, void* source, int flags);
void QPainter_DrawLines2(void* ptr, void* lines, int lineCount);
void QPainter_DrawPicture(void* ptr, void* point, void* picture);
void QPainter_DrawPie(void* ptr, void* rectangle, int startAngle, int spanAngle);
void QPainter_DrawPixmap5(void* ptr, void* point, void* pixmap);
void QPainter_DrawPixmap(void* ptr, void* target, void* pixmap, void* source);
void QPainter_DrawRects2(void* ptr, void* rectangles, int rectCount);
void QPainter_DrawRects(void* ptr, void* rectangles, int rectCount);
void QPainter_DrawText(void* ptr, void* position, char* text);
void QPainter_DrawText5(void* ptr, void* rectangle, int flags, char* text, void* boundingRect);
void QPainter_DrawText8(void* ptr, void* rectangle, char* text, void* option);
void QPainter_DrawText4(void* ptr, void* rectangle, int flags, char* text, void* boundingRect);
void QPainter_DrawTiledPixmap(void* ptr, void* rectangle, void* pixmap, void* position);
void QPainter_EraseRect(void* ptr, void* rectangle);
void QPainter_FillRect5(void* ptr, void* rectangle, void* brush);
void QPainter_FillRect6(void* ptr, void* rectangle, void* color);
void QPainter_FillRect(void* ptr, void* rectangle, void* brush);
void QPainter_FillRect7(void* ptr, void* rectangle, void* color);
void QPainter_Rotate(void* ptr, double angle);
void QPainter_SetBackground(void* ptr, void* brush);
void QPainter_SetBrushOrigin(void* ptr, void* position);
void QPainter_SetClipPath(void* ptr, void* path, int operation);
void QPainter_SetClipRect3(void* ptr, void* rectangle, int operation);
void QPainter_SetClipRect(void* ptr, void* rectangle, int operation);
void QPainter_SetClipRegion(void* ptr, void* region, int operation);
void QPainter_SetViewport(void* ptr, void* rectangle);
void QPainter_SetWindow(void* ptr, void* rectangle);
void* QPainter_NewQPainter();
void* QPainter_Background(void* ptr);
int QPainter_BackgroundMode(void* ptr);
void QPainter_BeginNativePainting(void* ptr);
void* QPainter_Brush(void* ptr);
void* QPainter_ClipRegion(void* ptr);
int QPainter_CompositionMode(void* ptr);
void* QPainter_Device(void* ptr);
void QPainter_DrawArc2(void* ptr, void* rectangle, int startAngle, int spanAngle);
void QPainter_DrawArc3(void* ptr, int x, int y, int width, int height, int startAngle, int spanAngle);
void QPainter_DrawChord2(void* ptr, void* rectangle, int startAngle, int spanAngle);
void QPainter_DrawChord3(void* ptr, int x, int y, int width, int height, int startAngle, int spanAngle);
void QPainter_DrawConvexPolygon4(void* ptr, void* polygon);
void QPainter_DrawConvexPolygon3(void* ptr, void* polygon);
void QPainter_DrawEllipse5(void* ptr, void* center, int rx, int ry);
void QPainter_DrawEllipse4(void* ptr, void* center, double rx, double ry);
void QPainter_DrawEllipse3(void* ptr, int x, int y, int width, int height);
void QPainter_DrawImage4(void* ptr, void* point, void* image);
void QPainter_DrawImage6(void* ptr, void* point, void* image, void* source, int flags);
void QPainter_DrawImage5(void* ptr, void* point, void* image, void* source, int flags);
void QPainter_DrawImage8(void* ptr, void* rectangle, void* image);
void QPainter_DrawImage2(void* ptr, void* target, void* image, void* source, int flags);
void QPainter_DrawImage7(void* ptr, void* rectangle, void* image);
void QPainter_DrawImage9(void* ptr, int x, int y, void* image, int sx, int sy, int sw, int sh, int flags);
void QPainter_DrawLine2(void* ptr, void* line);
void QPainter_DrawLine(void* ptr, void* line);
void QPainter_DrawLine3(void* ptr, void* p1, void* p2);
void QPainter_DrawLine4(void* ptr, void* p1, void* p2);
void QPainter_DrawLine5(void* ptr, int x1, int y1, int x2, int y2);
void QPainter_DrawLines(void* ptr, void* lines, int lineCount);
void QPainter_DrawLines4(void* ptr, void* pointPairs, int lineCount);
void QPainter_DrawLines3(void* ptr, void* pointPairs, int lineCount);
void QPainter_DrawPath(void* ptr, void* path);
void QPainter_DrawPicture2(void* ptr, void* point, void* picture);
void QPainter_DrawPicture3(void* ptr, int x, int y, void* picture);
void QPainter_DrawPie2(void* ptr, void* rectangle, int startAngle, int spanAngle);
void QPainter_DrawPie3(void* ptr, int x, int y, int width, int height, int startAngle, int spanAngle);
void QPainter_DrawPixmap6(void* ptr, void* point, void* pixmap);
void QPainter_DrawPixmap4(void* ptr, void* point, void* pixmap, void* source);
void QPainter_DrawPixmap3(void* ptr, void* point, void* pixmap, void* source);
void QPainter_DrawPixmap8(void* ptr, void* rectangle, void* pixmap);
void QPainter_DrawPixmap2(void* ptr, void* target, void* pixmap, void* source);
void QPainter_DrawPixmap7(void* ptr, int x, int y, void* pixmap);
void QPainter_DrawPixmap11(void* ptr, int x, int y, void* pixmap, int sx, int sy, int sw, int sh);
void QPainter_DrawPixmap10(void* ptr, int x, int y, int w, int h, void* pixmap, int sx, int sy, int sw, int sh);
void QPainter_DrawPixmap9(void* ptr, int x, int y, int width, int height, void* pixmap);
void QPainter_DrawPoint2(void* ptr, void* position);
void QPainter_DrawPoint(void* ptr, void* position);
void QPainter_DrawPoint3(void* ptr, int x, int y);
void QPainter_DrawPoints2(void* ptr, void* points, int pointCount);
void QPainter_DrawPoints(void* ptr, void* points, int pointCount);
void QPainter_DrawPoints4(void* ptr, void* points);
void QPainter_DrawPoints3(void* ptr, void* points);
void QPainter_DrawPolygon2(void* ptr, void* points, int pointCount, int fillRule);
void QPainter_DrawPolygon(void* ptr, void* points, int pointCount, int fillRule);
void QPainter_DrawPolygon4(void* ptr, void* points, int fillRule);
void QPainter_DrawPolygon3(void* ptr, void* points, int fillRule);
void QPainter_DrawPolyline2(void* ptr, void* points, int pointCount);
void QPainter_DrawPolyline(void* ptr, void* points, int pointCount);
void QPainter_DrawPolyline4(void* ptr, void* points);
void QPainter_DrawPolyline3(void* ptr, void* points);
void QPainter_DrawRect2(void* ptr, void* rectangle);
void QPainter_DrawRect(void* ptr, void* rectangle);
void QPainter_DrawRect3(void* ptr, int x, int y, int width, int height);
void QPainter_DrawRoundedRect2(void* ptr, void* rect, double xRadius, double yRadius, int mode);
void QPainter_DrawRoundedRect(void* ptr, void* rect, double xRadius, double yRadius, int mode);
void QPainter_DrawRoundedRect3(void* ptr, int x, int y, int w, int h, double xRadius, double yRadius, int mode);
void QPainter_DrawStaticText2(void* ptr, void* topLeftPosition, void* staticText);
void QPainter_DrawStaticText(void* ptr, void* topLeftPosition, void* staticText);
void QPainter_DrawStaticText3(void* ptr, int left, int top, void* staticText);
void QPainter_DrawText3(void* ptr, void* position, char* text);
void QPainter_DrawText6(void* ptr, int x, int y, char* text);
void QPainter_DrawText7(void* ptr, int x, int y, int width, int height, int flags, char* text, void* boundingRect);
void QPainter_DrawTiledPixmap2(void* ptr, void* rectangle, void* pixmap, void* position);
void QPainter_DrawTiledPixmap3(void* ptr, int x, int y, int width, int height, void* pixmap, int sx, int sy);
int QPainter_End(void* ptr);
void QPainter_EndNativePainting(void* ptr);
void QPainter_EraseRect2(void* ptr, void* rectangle);
void QPainter_EraseRect3(void* ptr, int x, int y, int width, int height);
void QPainter_FillPath(void* ptr, void* path, void* brush);
void QPainter_FillRect3(void* ptr, void* rectangle, int style);
void QPainter_FillRect11(void* ptr, void* rectangle, int color);
void QPainter_FillRect4(void* ptr, void* rectangle, int style);
void QPainter_FillRect12(void* ptr, void* rectangle, int color);
void QPainter_FillRect2(void* ptr, int x, int y, int width, int height, int style);
void QPainter_FillRect10(void* ptr, int x, int y, int width, int height, int color);
void QPainter_FillRect8(void* ptr, int x, int y, int width, int height, void* brush);
void QPainter_FillRect9(void* ptr, int x, int y, int width, int height, void* color);
int QPainter_HasClipping(void* ptr);
int QPainter_IsActive(void* ptr);
int QPainter_LayoutDirection(void* ptr);
double QPainter_Opacity(void* ptr);
void* QPainter_PaintEngine(void* ptr);
int QPainter_RenderHints(void* ptr);
void QPainter_ResetTransform(void* ptr);
void QPainter_Restore(void* ptr);
void QPainter_Save(void* ptr);
void QPainter_Scale(void* ptr, double sx, double sy);
void QPainter_SetBackgroundMode(void* ptr, int mode);
void QPainter_SetBrush2(void* ptr, int style);
void QPainter_SetBrush(void* ptr, void* brush);
void QPainter_SetBrushOrigin2(void* ptr, void* position);
void QPainter_SetBrushOrigin3(void* ptr, int x, int y);
void QPainter_SetClipRect2(void* ptr, int x, int y, int width, int height, int operation);
void QPainter_SetClipping(void* ptr, int enable);
void QPainter_SetCompositionMode(void* ptr, int mode);
void QPainter_SetFont(void* ptr, void* font);
void QPainter_SetLayoutDirection(void* ptr, int direction);
void QPainter_SetOpacity(void* ptr, double opacity);
void QPainter_SetPen3(void* ptr, int style);
void QPainter_SetPen2(void* ptr, void* color);
void QPainter_SetPen(void* ptr, void* pen);
void QPainter_SetRenderHint(void* ptr, int hint, int on);
void QPainter_SetRenderHints(void* ptr, int hints, int on);
void QPainter_SetTransform(void* ptr, void* transform, int combine);
void QPainter_SetViewTransformEnabled(void* ptr, int enable);
void QPainter_SetViewport2(void* ptr, int x, int y, int width, int height);
void QPainter_SetWindow2(void* ptr, int x, int y, int width, int height);
void QPainter_SetWorldMatrixEnabled(void* ptr, int enable);
void QPainter_SetWorldTransform(void* ptr, void* matrix, int combine);
void QPainter_Shear(void* ptr, double sh, double sv);
void QPainter_StrokePath(void* ptr, void* path, void* pen);
int QPainter_TestRenderHint(void* ptr, int hint);
void QPainter_Translate2(void* ptr, void* offset);
void QPainter_Translate(void* ptr, void* offset);
void QPainter_Translate3(void* ptr, double dx, double dy);
int QPainter_ViewTransformEnabled(void* ptr);
int QPainter_WorldMatrixEnabled(void* ptr);
void QPainter_DestroyQPainter(void* ptr);

#ifdef __cplusplus
}
#endif