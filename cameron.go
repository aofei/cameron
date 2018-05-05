package cameron

import (
	"crypto/md5"
	"encoding/binary"
	"image"
	"image/color"
)

// NewIdenticon returns an identicon avatar based on the data with the length
// and the squareLength.
func NewIdenticon(data []byte, length, squareLength int) image.Image {
	b := md5.Sum(data)

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

	padding := (length - (length/squareLength)*squareLength) / 2
	if padding == 0 {
		padding = squareLength / 2
	}

	barsCount := (length - 2*padding) / squareLength

	pixels := make([]byte, squareLength)
	for i := 0; i < squareLength; i++ {
		pixels[i] = 1
	}

	v, ri, ci := binary.BigEndian.Uint64(b[:]), 0, 0
	for i := 0; i < barsCount*(barsCount+1)/2; i++ {
		if v>>uint(i)&1 == 1 {
			for i := 0; i < squareLength; i++ {
				x := padding + ri*squareLength
				y := padding + ci*squareLength + i
				copy(img.Pix[img.PixOffset(x, y):], pixels)

				x = padding + (barsCount-1-ri)*squareLength
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
