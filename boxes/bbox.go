package boxes

import (
	"math"
)

type BOXTYPE uint8

const (
	COCO BOXTYPE = iota
	PASCAL
)

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
	normalize(image_height, image_width float64) error
	make_absolute(image_height, image_width float64) error
	is_normalized() bool
	box_type() BOXTYPE
}

type CocoBoundingBox struct {
	XMin       float64
	YMin       float64
	Height     float64
	Width      float64
	normalized bool
	BoxType    BOXTYPE
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

func (bbox *CocoBoundingBox) normalize(image_height, image_width float64) error {
	if bbox.normalized == true {
		return AlreadyNormalized{}
	}
	bbox.XMin = bbox.XMin / image_width
	bbox.YMin = bbox.YMin / image_height
	bbox.Height = bbox.Height / image_height
	bbox.Width = bbox.Width / image_width
	bbox.normalized = true
	return nil
}

func (bbox *CocoBoundingBox) make_absolute(image_height, image_width float64) error {
	if bbox.normalized == false {
		return AlreadyAbsolute{}
	}
	bbox.XMin = bbox.XMin * image_width
	bbox.YMin = bbox.YMin * image_height
	bbox.Height = bbox.Height * image_height
	bbox.Width = bbox.Width * image_width
	bbox.normalized = false
	return nil
}

func (bbox *CocoBoundingBox) is_normalized() bool {
	return bbox.normalized
}

func (bbox *CocoBoundingBox) box_type() BOXTYPE {
	return COCO
}

func NewCocoBox(xmin, ymin, height, width float64, normalized bool) *CocoBoundingBox {
	return &CocoBoundingBox{
		XMin:       xmin,
		YMin:       ymin,
		Height:     height,
		Width:      width,
		normalized: normalized,
	}
}

type PascalBoundingBox struct {
	XMin       float64
	YMin       float64
	XMax       float64
	YMax       float64
	normalized bool
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

func (bbox *PascalBoundingBox) normalize(image_height, image_width float64) error {
	if bbox.normalized == true {
		return AlreadyNormalized{}
	}
	bbox.XMin = bbox.XMin / image_width
	bbox.YMin = bbox.YMin / image_height
	bbox.XMax = bbox.XMax / image_width
	bbox.YMax = bbox.YMax / image_height
	bbox.normalized = true
	return nil
}

func (bbox *PascalBoundingBox) make_absolute(image_height, image_width float64) error {
	if bbox.normalized == false {
		return AlreadyAbsolute{}
	}
	bbox.XMin = bbox.XMin * image_width
	bbox.YMin = bbox.YMin * image_height
	bbox.XMax = bbox.XMax * image_width
	bbox.YMax = bbox.YMax * image_height
	bbox.normalized = false
	return nil
}

func (bbox *PascalBoundingBox) is_normalized() bool {
	return bbox.normalized
}

func (bbox *PascalBoundingBox) box_type() BOXTYPE {
	return PASCAL
}

func NewPascalBox(xmin, ymin, xmax, ymax float64, normalized bool) *PascalBoundingBox {
	return &PascalBoundingBox{
		XMin:       xmin,
		YMin:       ymin,
		XMax:       xmax,
		YMax:       ymax,
		normalized: normalized,
	}
}

func IoU(box1, box2 boundingbox) (float64, error) {
	if box1.box_type() == box2.box_type() {
		var xA = math.Max(box1.xmin(), box2.xmin())
		var yA = math.Max(box1.ymin(), box2.ymin())
		var xB = math.Min(box1.xmax(), box2.xmax())
		var yB = math.Min(box1.ymax(), box2.ymax())
		var interArea = math.Max((xB-xA), 0.0) * math.Max((yB-yA), 0.0)
		if interArea == 0.0 {
			return 0.0, nil
		}
		var iou = interArea / (box1.area() + box2.area() - interArea)
		return iou, nil
	}
	return -1.0, MismatchedBoxesError{
		type1: box1.box_type().String(),
		type2: box2.box_type().String(),
	}
}

//go:generate stringer -type=BOXTYPE -trimprefix=BoxType
