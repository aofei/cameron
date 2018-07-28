package cameron

import (
	"crypto/sha256"
	"encoding/binary"
	"image"
	"image/color"
)

// Identicon returns an identicon avatar based on the data with the length and
// the blockLength. Same parameters, same result.
func Identicon(data []byte, length, blockLength int) image.Image {
	b := sha256.Sum256(data)

	img := image.NewPaletted(
		image.Rect(0, 0, length, length),
		color.Palette{
			color.NRGBA{
				R: b[0],
				G: b[1],
				B: b[2],
				A: 0xff,
			},
			color.NRGBA{
				R: 255 - b[0],
				G: 255 - b[1],
				B: 255 - b[2],
				A: 0xff,
			},
		},
	)

	padding := (length - (length/blockLength)*blockLength) / 2
	if padding == 0 {
		padding = blockLength / 2
	}

	barsCount := (length - 2*padding) / blockLength

	pixels := make([]byte, blockLength)
	for i := 0; i < blockLength; i++ {
		pixels[i] = 1
	}

	v, ri, ci := binary.BigEndian.Uint64(b[:]), 0, 0
	for i := 0; i < barsCount*(barsCount+1)/2; i++ {
		if v>>uint(i)&1 == 1 {
			for i := 0; i < blockLength; i++ {
				x := padding + ri*blockLength
				y := padding + ci*blockLength + i
				copy(img.Pix[img.PixOffset(x, y):], pixels)

				x = padding + (barsCount-1-ri)*blockLength
				copy(img.Pix[img.PixOffset(x, y):], pixels)
			}
		}

		ci++
		if ci == barsCount {
			ci = 0
			ri++
		}
	}

	return img
}
