package cameron

import (
	"crypto/md5"
	"encoding/hex"
	"image"
	"image/color"
	"image/png"
	"strconv"
	"testing"
)

func TestIdenticon(t *testing.T) {
	t.Run("GitHub", func(t *testing.T) {
		for _, tt := range []struct {
			name       string
			userID     int
			wantDigest string
		}{
			{"aofei", 5037285, "442b264a5b375f8b79b533199a26ab61"},
			{"github", 9919, "5b768fe2c5e4b47fa819424c555d9787"},
			{"octocat", 583231, "36212fba11a3ded8440d440086ef0290"},
		} {
			t.Run(tt.name, func(t *testing.T) {
				img := Identicon([]byte(strconv.Itoa(tt.userID)), 70)

				h := md5.New()
				if err := png.Encode(h, img); err != nil {
					t.Fatalf("unexpected error %v", err)
				}
				digest := h.Sum(nil)

				got := hex.EncodeToString(digest)
				if got != tt.wantDigest {
					t.Errorf("got %s, want %s", got, tt.wantDigest)
				}
			})
		}
	})

	t.Run("EmptyInputs", func(t *testing.T) {
		img := Identicon(nil, 0)
		got := img.Bounds()
		want := image.Rect(0, 0, 6, 6)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestHSLToNRGBA(t *testing.T) {
	for _, tt := range []struct {
		name    string
		h, s, l float64
		want    color.NRGBA
	}{
		{"Red", 0, 100, 50, color.NRGBA{255, 0, 0, 255}},
		{"Green", 120, 100, 50, color.NRGBA{0, 255, 0, 255}},
		{"Blue", 240, 100, 50, color.NRGBA{0, 0, 255, 255}},
		{"Black", 0, 0, 0, color.NRGBA{0, 0, 0, 255}},
		{"White", 0, 0, 100, color.NRGBA{255, 255, 255, 255}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := hslToNRGBA(tt.h, tt.s, tt.l)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHueToRGB(t *testing.T) {
	const (
		p   = 0.0
		q   = 1.0
		tol = 1e-9
	)
	for _, tt := range []struct {
		name string
		t    float64
		want float64
	}{
		{"DefaultBranchAtZero", 0, p},                       // t == 0
		{"FirstBranch", 1.0 / 12.0, p + (q-p)*6*(1.0/12.0)}, // t < 1/6
		{"SecondBranch", 0.25, q},                           // 1/6 <= t < 1/2
		{"ThirdBranch", 0.6, p + (q-p)*(2.0/3.0-0.6)*6},     // 1/2 <= t < 2/3
		{"WrapNegativeT", -0.2, p},                          // -0.2 wraps to 0.8
		{"WrapTOverOne", 1.2, q},                            // 1.2 wraps to 0.2
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := hueToRGB(p, q, tt.t)
			if diff := got - tt.want; diff < -tol || diff > tol {
				t.Errorf("got %g, want %g", got, tt.want)
			}
		})
	}
}
