package color

import (
	"fmt"
	"image/color"

	"golang.org/x/image/colornames"
)

func Hex(name string) string { return ToHex(colornames.Map[name]) }
func ToHex(c color.RGBA) string {
	if c.A != 0xff {
		return fmt.Sprintf("#%02x%02x%02x%02x", c.R, c.G, c.B, c.A)
	}
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}
