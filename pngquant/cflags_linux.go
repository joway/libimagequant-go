//+build linux

package pngquant

/*
#cgo CFLAGS: -O3 -fno-math-errno -fopenmp -funroll-loops -fomit-frame-pointer -Wall -Wno-attributes -std=c99 -DNDEBUG -DUSE_SSE=1 -msse -fexcess-precision=fast
#cgo LDFLAGS: -lm -fopenmp
*/
import "C"
