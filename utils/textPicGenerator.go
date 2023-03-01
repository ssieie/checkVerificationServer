package utils

import (
	"bytes"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"io"
	"log"
	"math/rand"
	"os"
)

var fontFace font.Face
var preFontFace font.Face

var (
	width       = 350
	height      = 334
	preWidth    = 150
	preHeight   = 50
	fontSize    = float64(height/2) + 7
	preFontSize = float64(preHeight/2) + 7
)

func PicGenerator() ([]byte, error) {
	dc := gg.NewContext(preWidth, preHeight)
	bgR, bgG, bgB, bgA := getRandColorRange(240, 255)
	dc.SetRGBA255(bgR, bgG, bgB, bgA)
	dc.Clear()
	// 干扰线
	for i := 0; i < 6; i++ {
		x1, y1 := getRandPos(preWidth, preHeight)
		x2, y2 := getRandPos(preWidth, preHeight)
		r, g, b, a := getRandColor(255)
		w := float64(rand.Intn(3) + 1)
		dc.SetRGBA255(r, g, b, a)
		dc.SetLineWidth(w)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}

	dc.SetFontFace(preFontFace)

	text := GetRandomText()
	textLen := len(text)
	for i := 0; i < textLen; i++ {
		r, g, b, _ := getRandColor(100)
		dc.SetRGBA255(r, g, b, 255)
		fontPosX := float64(preWidth/textLen*i) + preFontSize*0.6
		writeText(dc, string(text[i:i+1]), fontPosX, float64(preHeight/2))
	}

	buffer := bytes.NewBuffer(nil)
	err := dc.EncodePNG(buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// 渲染文字
func writeText(dc *gg.Context, text string, x, y float64) {
	xFloat := 5 - rand.Float64()*10 + x
	yFloat := 5 - rand.Float64()*10 + y

	radians := 40 - rand.Float64()*80
	dc.RotateAbout(gg.Radians(radians), x, y)
	dc.DrawStringAnchored(text, xFloat, yFloat, 0.2, 0.5)
	dc.RotateAbout(-1*gg.Radians(radians), x, y)
	dc.Stroke()
}

// 随机坐标
func getRandPos(width, height int) (x float64, y float64) {
	x = rand.Float64() * float64(width)
	y = rand.Float64() * float64(height)
	return x, y
}

// 随机颜色
func getRandColor(maxColor int) (r, g, b, a int) {
	r = int(uint8(rand.Intn(maxColor)))
	g = int(uint8(rand.Intn(maxColor)))
	b = int(uint8(rand.Intn(maxColor)))
	a = int(uint8(rand.Intn(255)))
	return r, g, b, a
}

// 随机颜色范围
func getRandColorRange(miniColor, maxColor int) (r, g, b, a int) {
	if miniColor > maxColor {
		miniColor = 0
		maxColor = 255
	}
	r = int(uint8(rand.Intn(maxColor-miniColor) + miniColor))
	g = int(uint8(rand.Intn(maxColor-miniColor) + miniColor))
	b = int(uint8(rand.Intn(maxColor-miniColor) + miniColor))
	a = int(uint8(rand.Intn(maxColor-miniColor) + miniColor))
	return r, g, b, a
}

func LoadFontFace() {

	ttfFile, err := os.Open(fmt.Sprintf("./font/%s", GetConfig().Font))
	if err != nil {
		log.Fatalf("Failed to open ttfFile: %v", err)
	}
	defer ttfFile.Close()

	all, err := io.ReadAll(ttfFile)
	if err != nil {
		log.Fatalf("ReadAll ttfFile err: %v", err)
	}

	f, err := truetype.Parse(all)
	if err != nil {
		log.Fatalf("truetype Parse err: %v", err)
	}

	preFontFace = truetype.NewFace(f, &truetype.Options{
		Size: preFontSize,
	})
	fontFace = truetype.NewFace(f, &truetype.Options{
		Size: fontSize,
	})
}
