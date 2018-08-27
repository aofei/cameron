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
	jpeg.Encode(h, Identicon([]byte("cameron"), 540, 60), &jpeg.Options{
		Quality: 100,
	})
	assert.Equal(
		t,
		"zkTqoIfVRf76k65VbTogpMKZp7UosRv/uRR430Dcijs=",
		base64.StdEncoding.EncodeToString(h.Sum(nil)),
	)

	h = sha256.New()
	jpeg.Encode(h, Identicon([]byte("cameron"), 540, 65), &jpeg.Options{
		Quality: 100,
	})
	assert.Equal(
		t,
		"4/ID3mdOOmI1eT7gllbTFWNheUv9qAGn98sIqHih9ec=",
		base64.StdEncoding.EncodeToString(h.Sum(nil)),
	)

	h = sha256.New()
	jpeg.Encode(h, Identicon([]byte("cameron"), 540, 1080), &jpeg.Options{
		Quality: 100,
	})
	assert.Equal(
		t,
		"tmW9Svizstxyk0t3C3/ZlGa1ZmyJDSgfAT/Qb/gvQb8=",
		base64.StdEncoding.EncodeToString(h.Sum(nil)),
	)
}
