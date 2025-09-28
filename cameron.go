/*
Package cameron implements an avatar generator for Go.
*/
package cameron

import (
	"bytes"
	"crypto/md5"
	"image"
	"image/color"
	"math"
)

// Identicon returns an identicon avatar as an [image.Image] that is visually
// identical to https://github.com/identicons/{login}.png. All geometric rules,
// color calculations, and pixel layouts match the implementation GitHub uses in
// production.
//
// Note that the final image is a square of 6*cell pixels from a 5x5 grid plus a
// half-cell margin on every side.
func Identicon(data []byte, cell int) image.Image {
	digest := md5.Sum(data)
	if cell < 1 {
		cell = 1
	}

	// Split the 16-byte digest into 32 individual 4-bit nibbles.
	var nib [32]byte
	for i := 0; i < 16; i++ {
		nib[2*i] = digest[i] >> 4     // High 4 bits.
		nib[2*i+1] = digest[i] & 0x0f // Low 4 bits.
	}

	// Build the 5x5 symmetry mask.
	//
	// The first 15 nibbles decide the left half of the grid:
	//   - Nibbles 0-4 fill the center column (index 2)
	//   - Nibbles 5-9 fill the column immediately left of center (index 1)
	//   - Nibbles 10-14 fill the leftmost column (index 0)
	//
	// A pixel is set only when its nibble value is even.
	//
	// Once the left half is filled, copy it to columns 3 and 4 to complete
	// the grid and guarantee horizontal symmetry.
	var mask [5][5]bool
	for i := 0; i < 15; i++ {
		if nib[i]%2 == 0 {
			row := i % 5
			col := 2 - i/5
			mask[row][col] = true
		}
	}
	for r := 0; r < 5; r++ {
		mask[r][3] = mask[r][1]
		mask[r][4] = mask[r][0]
	}

	// Derive the foreground color from HSL.
	//
	// The final 7 nibbles are interpreted as HHHSSLL, where
	//   - HHH (12 bits) maps to hue in [0, 360) degrees
	//   - SS (8 bits) maps to saturation in [45, 65] percent
	//   - LL (8 bits) maps to lightness in [55, 75] percent
	var v uint32
	for i := 25; i < 32; i++ {
		v = (v << 4) | uint32(nib[i])
	}
	hueBits := v >> 16
	satBits := (v >> 8) & 0xff
	lgtBits := v & 0xff
	h := float64(hueBits) * 360 / 4095
	s := 65.0 - float64(satBits)*20/255
	l := 75.0 - float64(lgtBits)*20/255
	fg := hslToNRGBA(h, s, l)

	// Use a light gray background as in GitHub's implementation.
	bg := color.NRGBA{R: 240, G: 240, B: 240, A: 255}

	// Allocate the palette-based image and fill it.
	//
	// The bitmap is six logical cells per side with five pattern cells plus
	// a half-cell margin on each edge. Using a palette keeps memory small.
	size := 6 * cell
	img := image.NewPaletted(image.Rect(0, 0, size, size), color.Palette{bg, fg})
	margin := cell / 2                      // Half-cell margin in pixels.
	rowBuf := bytes.Repeat([]byte{1}, cell) // Palette index 1 is fg.
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !mask[r][c] {
				continue
			}
			x := margin + c*cell
			y := margin + r*cell
			for dy := 0; dy < cell; dy++ {
				off := img.PixOffset(x, y+dy)
				copy(img.Pix[off:], rowBuf)
			}
		}
	}
	return img
}

// hslToNRGBA converts HSL values to an opaque [color.NRGBA].
func hslToNRGBA(h, s, l float64) color.NRGBA {
	h /= 360
	s /= 100
	l /= 100

	var q float64
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p := 2*l - q

	r := hueToRGB(p, q, h+1.0/3.0)
	g := hueToRGB(p, q, h)
	b := hueToRGB(p, q, h-1.0/3.0)
	return color.NRGBA{
		R: uint8(math.Round(r * 255)),
		G: uint8(math.Round(g * 255)),
		B: uint8(math.Round(b * 255)),
		A: 255,
	}
}

// hueToRGB converts a hue offset t into a single RGB component.
func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	switch {
	case t < 1.0/6.0:
		return p + (q-p)*6*t
	case t < 1.0/2.0:
		return q
	case t < 2.0/3.0:
		return p + (q-p)*(2.0/3.0-t)*6
	default:
		return p
	}
}
