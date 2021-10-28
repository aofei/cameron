package cameron

import (
	"image/jpeg"
	"testing"

	"github.com/cespare/xxhash/v2"
)

func TestIdenticon(t *testing.T) {
	d := xxhash.New()
	jpeg.Encode(d, Identicon([]byte("cameron"), 540, 60), nil)
	digest := d.Sum64()
	if want := uint64(0x88a6d1f1f6986239); digest != want {
		t.Errorf("got %d, want %d", digest, want)
	}

	d.Reset()
	jpeg.Encode(d, Identicon([]byte("cameron"), 540, 65), nil)
	digest = d.Sum64()
	if want := uint64(0x4e18c5d23e0bec77); digest != want {
		t.Errorf("got %d, want %d", digest, want)
	}

	d.Reset()
	jpeg.Encode(d, Identicon([]byte("cameron"), 540, 1080), nil)
	digest = d.Sum64()
	if want := uint64(0x3ec4415b1f794db7); digest != want {
		t.Errorf("got %d, want %d", digest, want)
	}
}
