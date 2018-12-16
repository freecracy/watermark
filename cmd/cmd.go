package cmd

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
)

const (
	filePath = "/Users/cn/Desktop/"
	fontPath = "/Library/Fonts/"
	font     = "Arial Italic.ttf"
	fontFile = fontPath + font
)

func CreateTextImage() {

	dx := 200
	dy := 200
	imgfile, _ := os.Create(filePath + "text.png")

	defer imgfile.Close()

	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 0})
		}
	}

	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println("load front fail", err)
	}

	f := freetype.NewContext()
	f.SetDPI(72)
	f.SetFont(font)
	f.SetFontSize(26)
	f.SetClip(img.Bounds())
	f.SetDst(img)
	f.SetSrc(image.NewUniform(color.RGBA{184, 184, 184, 100}))

	//设置字体的位置
	pt := freetype.Pt(40, 40+int(f.PointToFixed(26))>>8)

	_, err = f.DrawString("hello", pt)
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

func MergeImage() {

	imgb, _ := os.Open(filePath + "image.png")
	img, _ := png.Decode(imgb)
	defer imgb.Close()

	wmb, _ := os.Open(filePath + "text.png")
	watermark, _ := png.Decode(wmb)
	defer wmb.Close()

	//把水印写到右下角，并向0坐标各偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	imgw, _ := os.Create(filePath + "new.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})

	defer imgw.Close()
}
