package imgopt

import (
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math"

	"github.com/nfnt/resize"
)

func Optimize(r io.Reader, w io.Writer, max int, quality int) error {
	c, format, err := image.DecodeConfig(r)
	if err != nil {
		return err
	}

	if c.Width >= c.Height && c.Width > max {
		c.Width, c.Height = calcSize(c.Width, c.Height, max)
	} else if c.Height > max {
		c.Height, c.Width = calcSize(c.Height, c.Width, max)
	}

	r.(io.Seeker).Seek(0, io.SeekStart)
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	switch format {
	case "jpeg":
		fallthrough
	case "png":
		img = resize.Resize(uint(c.Width), uint(c.Height), img, resize.Lanczos3)
	}

	switch format {
	case "jpeg":
		if err := jpeg.Encode(w, img, &jpeg.Options{
			Quality: quality,
		}); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(w, img); err != nil {
			return err
		}
	}

	return nil
}

func calcSize(w int, h int, max int) (int, int) {
	if w > max {
		return max, int(
			math.Round(
				float64(h) / float64(w) * float64(max),
			),
		)
	}

	return w, h
}
