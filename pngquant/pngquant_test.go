package pngquant

import (
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
	assert.NoError(t, err)
	img, err := png.Decode(source)
	assert.NoError(t, err)
	_, err = Compress(img, 60)
	assert.NoError(t, err)
}
