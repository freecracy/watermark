package cmd

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
)

const (
	filePath = "/Desktop/"
	fontPath = "/Library/Fonts/"
	font     = "Arial Unicode.ttf"
	fontFile = fontPath + font
)

func CreateTextImage() {

	text := flag.String("text", "hello", "水印文字")
	flag.Parse()

	dx := 100
	dy := 50
	homeDir, _ := os.UserHomeDir()
	fmt.Println(homeDir)
	imgfile, _ := os.Create(homeDir + filePath + "text.png")
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
	f.SetSrc(image.NewUniform(color.RGBA{184, 184, 184, 80}))

	//设置字体的位置
	//pt := freetype.Pt(50, 15+int(f.PointToFixed(26))>>8)
	pt := freetype.Pt(10, 25+12)
	_, err = f.DrawString(*text, pt)
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

func MergeImage() {

	filePath := "/Users/libin" + filePath

	srcImage, err := os.Open(filePath + "image.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	srcPng, err := png.Decode(srcImage)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer srcImage.Close()

	textImage, _ := os.Open(filePath + "text.png")
	textPng, _ := png.Decode(textImage)
	defer textImage.Close()

	//把水印写到右下角，并向0坐标各偏移10个像素
	b := srcPng.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, srcPng, image.ZP, draw.Src)
	for i := 10; i < 1000; i = i + 150 {
		for j := 10; j < 500; j = j + 80 {
			offset := image.Pt(srcPng.Bounds().Dx()-textPng.Bounds().Dx()-i, srcPng.Bounds().Dy()-textPng.Bounds().Dy()-j)
			draw.Draw(m, textPng.Bounds().Add(offset), textPng, image.ZP, draw.Over)
		}
	}

	outImage, _ := os.Create(filePath + "new.png")
	png.Encode(outImage, m)
	defer outImage.Close()
}
