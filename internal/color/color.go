package color

import (
	"os"

	"github.com/devemio/mockio/pkg/color"
)

type Color struct {
	enabled bool
	tc      bool
}

func New(enabled bool) *Color {
	return &Color{enabled: enabled, tc: os.Getenv("COLORTERM") == "truecolor"}
}

func (c *Color) Black(v string) string {
	if !c.enabled {
		return v
	}

	return color.Black(v)
}

func (c *Color) DarkGray(v string) string {
	if !c.enabled {
		return v
	}

	return color.DarkGray(v)
}

func (c *Color) Red(v string) string {
	if !c.enabled {
		return v
	}

	return color.Red(v)
}

func (c *Color) LightRed(v string) string {
	if !c.enabled {
		return v
	}

	return color.LightRed(v)
}

func (c *Color) Green(v string) string {
	if !c.enabled {
		return v
	}

	return color.Green(v)
}

func (c *Color) LightGreen(v string) string {
	if !c.enabled {
		return v
	}

	return color.LightGreen(v)
}

func (c *Color) Brown(v string) string {
	if !c.enabled {
		return v
	}

	return color.Brown(v)
}

func (c *Color) Yellow(v string) string {
	if !c.enabled {
		return v
	}

	return color.Yellow(v)
}

func (c *Color) Blue(v string) string {
	if !c.enabled {
		return v
	}

	return color.Blue(v)
}

func (c *Color) LightBlue(v string) string {
	if !c.enabled {
		return v
	}

	return color.LightBlue(v)
}

func (c *Color) Purple(v string) string {
	if !c.enabled {
		return v
	}

	return color.Purple(v)
}

func (c *Color) LightPurple(v string) string {
	if !c.enabled {
		return v
	}

	return color.LightPurple(v)
}

func (c *Color) Cyan(v string) string {
	if !c.enabled {
		return v
	}

	return color.Cyan(v)
}

func (c *Color) LightCyan(v string) string {
	if !c.enabled {
		return v
	}

	return color.LightCyan(v)
}

func (c *Color) LightGray(v string) string {
	if !c.enabled {
		return v
	}

	if c.tc {
		return "\033[0;90m" + v + "\033[0m"
	}

	return color.LightGray(v)
}

func (c *Color) White(v string) string {
	if !c.enabled {
		return v
	}

	return color.White(v)
}
