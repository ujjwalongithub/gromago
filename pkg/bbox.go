package pkg

import "math"

type boundingbox interface {
	xmin() float64
	ymin() float64
	xmax() float64
	ymax() float64
	height() float64
	width() float64
	aspect_ratio() float64
	center() (float64, float64)
	area() float64
}

type CocoBoundingBox struct {
	XMin   float64
	YMin   float64
	Height float64
	Width  float64
}

func (bbox *CocoBoundingBox) xmin() float64 {
	return bbox.XMin
}

func (bbox *CocoBoundingBox) ymin() float64 {
	return bbox.YMin
}

func (bbox *CocoBoundingBox) xmax() (xmax float64) {
	xmax = bbox.XMin + bbox.Width
	return
}

func (bbox *CocoBoundingBox) ymax() (ymax float64) {
	ymax = bbox.YMin + bbox.Height
	return
}

func (bbox *CocoBoundingBox) height() float64 {
	return bbox.Height
}

func (bbox *CocoBoundingBox) width() float64 {
	return bbox.Width
}

func (bbox *CocoBoundingBox) aspect_ratio() float64 {
	return bbox.Height / bbox.Width
}

func (bbox *CocoBoundingBox) center() (x, y float64) {
	x = bbox.XMin + bbox.Width/2
	y = bbox.YMin + bbox.Height/2
	return
}

func (bbox *CocoBoundingBox) area() float64 {
	return bbox.Height * bbox.Width
}

func NewCocoBox(xmin, ymin, height, width float64) *CocoBoundingBox {
	return &CocoBoundingBox{
		XMin:   xmin,
		YMin:   ymin,
		Height: height,
		Width:  width,
	}
}

type PascalBoundingBox struct {
	XMin float64
	YMin float64
	XMax float64
	YMax float64
}

func (bbox *PascalBoundingBox) xmin() float64 {
	return bbox.XMin
}

func (bbox *PascalBoundingBox) ymin() float64 {
	return bbox.YMin
}

func (bbox *PascalBoundingBox) xmax() float64 {
	return bbox.XMax
}

func (bbox *PascalBoundingBox) ymax() float64 {
	return bbox.YMax
}

func (bbox *PascalBoundingBox) height() (height float64) {
	height = bbox.YMax - bbox.YMin
	return
}

func (bbox *PascalBoundingBox) width() (width float64) {
	width = bbox.XMax - bbox.XMin
	return
}

func (bbox *PascalBoundingBox) aspect_ratio() float64 {
	return bbox.height() / bbox.width()
}

func (bbox *PascalBoundingBox) center() (x, y float64) {
	x = (bbox.XMax - bbox.XMin) / 2
	y = (bbox.YMax - bbox.YMin) / 2
	return
}

func (bbox *PascalBoundingBox) area() float64 {
	return bbox.height() * bbox.width()
}

func NewPascalBox(xmin, ymin, xmax, ymax float64) *PascalBoundingBox {
	return &PascalBoundingBox{
		XMin: xmin,
		YMin: ymin,
		XMax: xmax,
		YMax: ymax,
	}
}

func IoU(box1, box2 boundingbox) float64 {
	var xA = math.Max(box1.xmin(), box2.xmin())
	var yA = math.Max(box1.ymin(), box2.ymin())
	var xB = math.Min(box1.xmax(), box2.xmax())
	var yB = math.Min(box1.ymax(), box2.ymax())
	var interArea = math.Max((xB-xA), 0.0) * math.Max((yB-yA), 0.0)
	if interArea == 0.0 {
		return 0.0
	}
	var iou = interArea / (box1.area() + box2.area() - interArea)
	return iou
}
