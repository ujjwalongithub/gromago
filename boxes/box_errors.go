package boxes

import "fmt"

type MismatchedBoxesError struct {
	type1 string
	type2 string
}

func (z MismatchedBoxesError) Error() string {
	var msg string = fmt.Sprintf("IoU is computed only between boxes of the same type. "+
		"Found %s and %s", z.type1, z.type2)
	return msg
}

type AlreadyNormalized struct {
}

func (z AlreadyNormalized) Error() string {
	return "The boxes are already normalized. They cannot be re-normalized."
}

type AlreadyAbsolute struct {
}

func (z AlreadyAbsolute) Error() string {
	return "The boxes are already in absolute form. They cannot be made absolute again."
}
