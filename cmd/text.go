package cmd

// import (
// 	"fmt"

// 	"github.com/disintegration/imaging"
// 	"github.com/golang/freetype"
// 	"github.com/op/go-logging"
// )

// func init() {

// }

// func HandleUserImage(string, error) {
// 	if m, err := imaging.Open("target.png"); err != nil {
// 		fmt.Printf("open target image failed")
// 	}
// 	if bm, err := imaging.Open("bg.png"); err != nil {
// 		fmt.Printf("open gb image failed")
// 	}
// }

// func HandleUserImage() (string, error) {
// 	m, err := imaging.Open("target.png")
// 	if err != nil {
// 		fmt.Printf("open file failed")
// 	}

// 	bm, err := imaging.Open("bg.jpg")
// 	if err != nil {
// 		fmt.Printf("open file failed")
// 	}

// 	// 图片按比例缩放
// 	dst := imaging.Resize(m, 200, 200, imaging.Lanczos)
// 	// 将图片粘贴到背景图的固定位置
// 	result := imaging.Overlay(bm, dst, image.Pt(120, 140), 1)
// 	writeOnImage(result)

// 	fileName := fmt.Sprintf("%d.jpg", 1234)
// 	err = imaging.Save(result, fileName)
// 	if err != nil {
// 		return "", err
// 	}

// 	return fileName, nil
// }
