package cameron

import (
	"crypto/sha256"
	"encoding/base64"
	"image/jpeg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdenticon(t *testing.T) {
	h := sha256.New()
	jpeg.Encode(h, Identicon([]byte("cameron"), 540, 50), &jpeg.Options{
		Quality: 100,
	})
	assert.Equal(
		t,
		"3zDt4JrN67koaZjApe1Be4XYgBvc3IHG58idszxH/9s=",
		base64.StdEncoding.EncodeToString(h.Sum(nil)),
	)
}
