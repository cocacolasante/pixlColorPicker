package pxcanvas

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/internal/widget"
	"pixl.io/apptype"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent

}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer *PxCanvasRenderer
	PixelData image.Image
	mouseState PxCanvasMouseState
	appState *apptype.State
	reloadImage bool
}

func (pxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(pxCanvas.CanvasOffset.X)
	y0 := int(pxCanvas.CanvasOffset.Y)
	x1 := int(pxCanvas.PxCols * pxCanvas.PxSize + int(pxCanvas.CanvasOffset.X))
	y1 := int(pxCanvas.PxRows * pxCanvas.PxSize + int(pxCanvas.CanvasOffset.Y))

	return image.Rect(x0, y0, x1, y1)


}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) &&
		pos.X < float32(bounds.Max.X) &&
		pos.Y >= float32(bounds.Min.Y) &&
		pos.Y < float32(bounds.Max.Y){
			return true
	}
	return false

}

func NewBlankImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))

	for y := 0; y < rows; y++ {
		for x:= 0; x < cols; x++ {
			img.Set(x, y, c)
		}
	}
	return img
}

func NewPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas{
	pxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		appState: state,
	}
	pxCanvas.PixelData = NewBlankImage(pxCanvas.PxCols, pxCanvas.PxRows, color.Black)
	pxCanvas.ExtendBaseWidget(pxCanvas)

	return pxCanvas
}

func(pxCanvas *PxCanvas) createRenederer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)

	for i :=0; i <len(canvasBorder); i++{
		canvasBorder[i].StrokeColor = color.Gray{}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &PxCanvasRenderer{
		pxCanvas: pxCanvas,
		canvasImage: canvasImage,
		canvasBorder: canvasBorder,
	}

	pxCanvas.renderer = renderer
	return renderer 
}