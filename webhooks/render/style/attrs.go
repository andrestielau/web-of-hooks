package style

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

var sizes = map[string]int{
	"4xs": 360,
	"3xs": 480,
	"2xs": 600,
	"xs":  768,
	"sm":  834,
	"md":  1024,
	"lg":  1280,
	"xl":  1440,
	"2xl": 1600,
	"3xl": 1920,
	"4xl": 2560,
}
var keywords = map[string]string{
	"square":             `aspect-ratio: 1/1;`,
	"video":              `aspect-ratio: 16/9;`,
	"rounded":            `border-radius: 1e9em;`,
	"round":              `border-radius: 50%;`,
	"hidden":             `display: none;`,
	"hide":               `display: none;`,
	"block":              `display: block;`,
	"table":              `display: table;`,
	"flex":               `display: flex;`,
	"grid":               `display: grid;`,
	"contents":           `display: contents;`,
	"inline":             `display: inline;`,
	"inline-block":       `display: inline-block;`,
	"inline-flex":        `display: inline-flex;`,
	"inline-grid":        `display: inline-grid;`,
	"inline-table":       `display: inline-table;`,
	"table-cell":         `display: table-cell;`,
	"table-caption":      `display: table-caption;`,
	"flow-root":          `display: flow-root;`,
	"list-item":          `display: list-item;`,
	"table-row":          `display: table-row;`,
	"table-column":       `display: table-column;`,
	"table-row-group":    `display: table-row-group;`,
	"table-column-group": `display: table-column-group;`,
	"table-header-group": `display: table-header-group;`,
	"table-footer-group": `display: table-footer-group; `,
	"italic":             `font-style: italic;`,
	"oblique":            `font-style: oblique;`,
	"isolate":            `isolation: isolate;`,
	"overflowed":         `overflow: visible;`,
	"untouchable":        `pointer-events: none;`,
	"static":             `position: static;`,
	"fixed":              `position: fixed;`,
	"abs":                `position: absolute;`,
	"rel":                `position: relative;`,
	"sticky":             `position: sticky;`,
	"uppercase":          `text-transform: uppercase;`,
	"lowercase":          `text-transform: lowercase;`,
	"capitalize":         `text-transform: capitalize;`,
	"visible":            `visibility: visible;`,
	"invisible":          `visibility: hidden;`,
	"vw":                 `width: 100vw;`,
	"vh":                 `height: 100vh;`,
	"max-vw":             `max-width: 100vw;`,
	"max-vh":             `max-height: 100vh;`,
	"min-vw":             `min-width: 100vw;`,
	"min-vh":             `min-height: 100vh;`,
	"center-content":     `justify-content: center; align-items: center;`,
	"sr-only": `position: absolute;
				width: 1px;
				height: 1px;
				padding: 0;
				margin: -1px;
				overflow: hidden;
				clip: rect(0,0,0,0);
				white-space: nowrap;
				border-width: 0;`,
	"full":         `width: 100%; height: 100%;`,
	"fit":          `width: fit-content; height: fit-content;`,
	"center":       `left: 0; right: 0; margin-left: auto; margin-right: auto;`,
	"middle":       `top: 0; bottom: 0; margin-top: auto; margin-bottom: auto;`,
	"break-spaces": `white-space: break-spaces;`,
	"break-word":   `overflow-wrap: break-word; overflow: hidden;`,
	"gradient-text": `-webkit-text-fill-color: transparent,
					  -webkit-background-clip: text,
				   	  background-clip: text`,
}
var alias = map[string]string{
	"accent": "accent-color",
	"bg":     "background",
	"fg":     "color",
	"p":      "padding",
	"px":     "padding",
	"py":     "padding",
	"pt":     "padding-top",
	"pb":     "padding-bottom",
	"pl":     "padding-left",
	"pr":     "padding-right",
	"m":      "margin",
	"mt":     "margin-top",
	"mb":     "margin-bottom",
	"ml":     "margin-left",
	"mr":     "margin-right",
	"b":      "border",
	"bt":     "border-top",
	"bb":     "border-bottom",
	"bl":     "border-left",
	"br":     "border-right",
	"r":      "border-radius",
	"w":      "width",
	"h":      "height",
	"z":      "z-index",
}

func C(classes string, args ...any) templ.CSSClass {
	var t strings.Builder
	for _, class := range strings.Split(classes, " ") {
		parts := strings.Split(class, ":")
		key := parts[0]
		if n, ok := keywords[key]; ok {
			t.WriteString(n)
		} else if len(parts) > 1 {
			if n, ok = alias[key]; !ok {
				n = key
			}
			v := parts[1]
			if size, ok := sizes[v]; ok {
				v = strconv.Itoa(size)
			}
			t.WriteString(string(templ.SanitizeCSS(n, v)))
		}
	}
	ret := t.String()
	if len(args) > 0 {
		ret = fmt.Sprintf(ret, args...)
	}
	id := templ.CSSID(`C`, classes)
	return templ.ComponentCSSClass{
		ID:    id,
		Class: templ.SafeCSS(`.` + id + `{` + ret + `}`),
	}
}
