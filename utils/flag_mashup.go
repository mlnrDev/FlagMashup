package utils

import (
	"bytes"
	"flag-mashup/data"

	"github.com/mlnrDev/mashup"
)

func MashupFlags(src, dst string, maxColors int, codeData *data.CodeData, buf *bytes.Buffer) error {
	srcBody, err := codeData.FetchFlag(src)
	if err != nil {
		return err
	}
	defer srcBody.Close()

	dstBody, err := codeData.FetchFlag(dst)
	if err != nil {
		return err
	}
	defer dstBody.Close()

	return mashup.Mashup(mashup.NewPNGInput(srcBody), mashup.NewPNGInput(dstBody), mashup.NewPNGOutput(buf), maxColors)
}
