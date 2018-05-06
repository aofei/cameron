package cameron

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"image/jpeg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdenticon(t *testing.T) {
	buf := &bytes.Buffer{}
	jpeg.Encode(buf, Identicon([]byte("cameron"), 10, 1), &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	})
	assert.Equal(
		t,
		fmt.Sprintf("%x", md5.Sum(buf.Bytes())),
		"2d809f43b9ab5e2f80966824a3b6f94b",
	)
}
