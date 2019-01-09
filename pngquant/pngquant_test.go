package pngquant

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"image/png"
	"os"
	"testing"
)

func TestLibVersion(t *testing.T) {
	assert.Equal(t, LibVersion(), "2.12.2")
}

func TestCompress(t *testing.T) {
	source, err := os.OpenFile("../testdata/1.png", os.O_RDONLY, 0444)
	stat, _ := source.Stat()
	assert.NoError(t, err)
	img, err := png.Decode(source)
	assert.NoError(t, err)

	output1, err := Compress(img, 70, SPEED_DEFAULT)
	assert.NoError(t, err)
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, output1); err != nil {
		assert.NoError(t, err)
	}
	outputSize1 := int64(len(buf.Bytes()))
	assert.True(t, outputSize1 < stat.Size())

	output2, err := Compress(img, 60, SPEED_DEFAULT)
	assert.NoError(t, err)
	buf = new(bytes.Buffer)
	if err := png.Encode(buf, output2); err != nil {
		assert.NoError(t, err)
	}
	outputSize2 := int64(len(buf.Bytes()))
	assert.True(t, outputSize2 < stat.Size())

	assert.True(t, outputSize1 > outputSize2)
}
