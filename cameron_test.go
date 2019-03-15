package cameron

import (
	"image/jpeg"
	"testing"

	"github.com/cespare/xxhash/v2"
	"github.com/stretchr/testify/assert"
)

func TestIdenticon(t *testing.T) {
	d := xxhash.New()
	jpeg.Encode(d, Identicon([]byte("cameron"), 540, 60), nil)
	assert.Equal(t, uint64(0x88a6d1f1f6986239), d.Sum64())

	d.Reset()
	jpeg.Encode(d, Identicon([]byte("cameron"), 540, 65), nil)
	assert.Equal(t, uint64(0x4e18c5d23e0bec77), d.Sum64())

	d.Reset()
	jpeg.Encode(d, Identicon([]byte("cameron"), 540, 1080), nil)
	assert.Equal(t, uint64(0x3ec4415b1f794db7), d.Sum64())
}
