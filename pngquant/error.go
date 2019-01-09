package pngquant

/*
#include "libimagequant.h"
*/
import "C"
import "errors"

var (
	ErrQualityTooLow      = errors.New("quality too low")
	ErrValueOutOfRange    = errors.New("value out of range")
	ErrOutOfMemory        = errors.New("out of memory")
	ErrAborted            = errors.New("aborted")
	ErrBitmapNotAvailable = errors.New("bitmap not available")
	ErrBufferTooSmall     = errors.New("buffer too small")
	ErrInvalidPointer     = errors.New("invalid pointer")
	ErrUseAfterFree = errors.New("use after free")
)

func translateError(iqe C.liq_error) error {
	switch iqe {
	case C.LIQ_OK:
		return nil
	case (C.LIQ_QUALITY_TOO_LOW):
		return ErrQualityTooLow
	case (C.LIQ_VALUE_OUT_OF_RANGE):
		return ErrValueOutOfRange
	case (C.LIQ_OUT_OF_MEMORY):
		return ErrOutOfMemory
	case (C.LIQ_ABORTED):
		return ErrAborted
	case (C.LIQ_BITMAP_NOT_AVAILABLE):
		return ErrBitmapNotAvailable
	case (C.LIQ_BUFFER_TOO_SMALL):
		return ErrBufferTooSmall
	case (C.LIQ_INVALID_POINTER):
		return ErrInvalidPointer
	default:
		return errors.New("unknown error")
	}
}
