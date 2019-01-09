//+build darwin

package pngquant

/*
#cgo CFLAGS: -I../lib -std=c99
#cgo LDFLAGS: -L../lib

#include "libimagequant.h"

#include "libimagequant.c"
#include "pam.c"
#include "mediancut.c"
#include "nearest.c"
#include "blur.c"
#include "kmeans.c"
#include "mempool.c"
*/
import "C"
