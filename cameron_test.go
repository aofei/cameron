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
	jpeg.Encode(buf, Identicon([]byte("cameron"), 540, 50), &jpeg.Options{
		Quality: 100,
	})
	assert.Equal(
		t,
		"8c21b5962fb72f62232765eb0b5e6ddd",
		fmt.Sprintf("%x", md5.Sum(buf.Bytes())),
	)
}
