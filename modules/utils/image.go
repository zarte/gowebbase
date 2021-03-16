package utils

import (
	"errors"
	"fmt"
	"github.com/golang/freetype"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
)


/**
添加水印
 */
func Water(oimgpath string, destpath string,waterstr string,font string) error  {
	if font=="" {
		return  errors.New("no set font")
	}
	//需要加水印的图片
	imgfile, _ := os.Open(oimgpath)
	defer imgfile.Close()
	origin, fm, err := image.Decode(imgfile)
	if err != nil {
		return err
	}
	img := image.NewNRGBA(origin.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, origin.At(x, y))
		}
	}
	//拷贝一个字体文件到运行目录
	fontBytes, err := ioutil.ReadFile(font)
	if err != nil {
		log.Println(err)
	}

	fonttype, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
	}

	f := freetype.NewContext()
	f.SetDPI(72)
	f.SetFont(fonttype)
	f.SetFontSize(12)
	f.SetClip(origin.Bounds())
	f.SetDst(img)
	f.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 0, B: 0, A: 255}))
	pt := freetype.Pt(img.Bounds().Dx()-200, img.Bounds().Dy()-12)
	_, err = f.DrawString(waterstr, pt)

	//draw.Draw(img,jpgimg.Bounds(),jpgimg,image.ZP,draw.Over)

	//保存到新文件中
	newfile, _ := os.Create(destpath)
	defer newfile.Close()
	switch fm {
	case "jpeg":
		err = jpeg.Encode(newfile, img, &jpeg.Options{Quality:100})
		if err != nil {
			fmt.Println(err)
			return err
		}
	case "png":
		err = png.Encode(newfile, img)
		if err != nil {
			fmt.Println(err)
			return err
		}
	case "gif":
		err = gif.Encode(newfile, img, &gif.Options{})
		if err != nil {
			fmt.Println(err)
			return err
		}
	case "bmp":
		err = bmp.Encode(newfile, img)
		if err != nil {
			fmt.Println(err)
			return err
		}
	default:
		return errors.New("ERROR FORMAT")
	}
	return nil
}
/**
裁剪
 */
func Cutimg(in io.Reader, out io.Writer, x0, y0, x1, y1, quality int) error {
	origin, fm, err := image.Decode(in)
	if err != nil {
		return err
	}

	switch fm {
	case "jpeg":
		img := origin.(*image.YCbCr)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)
		return jpeg.Encode(out, subImg, &jpeg.Options{quality})
	case "png":
		switch origin.(type) {
		case *image.NRGBA:
			img := origin.(*image.NRGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
			return png.Encode(out, subImg)
		case *image.RGBA:
			img := origin.(*image.RGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
			return png.Encode(out, subImg)
		}
	case "gif":
		img := origin.(*image.Paletted)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.Paletted)
		return gif.Encode(out, subImg, &gif.Options{})
	case "bmp":
		img := origin.(*image.RGBA)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
		return bmp.Encode(out, subImg)
	default:
		return errors.New("ERROR FORMAT")
	}
	return nil
}
/**
缩略图
-rw-r--r--@  1 wxnacy  staff  64633 Jan 18 18:01 react-app1.png
-rw-r--r--@  1 wxnacy  staff  19857 Jan 19 17:13 react.NearestNeighbor.png
-rw-r--r--   1 wxnacy  staff  23336 Jan 19 17:12 react.Bilinear.png
-rw-r--r--   1 wxnacy  staff  25966 Jan 19 17:12 react.MitchellNetravali.png
-rw-r--r--@  1 wxnacy  staff  27241 Jan 19 17:12 react.Lanczos2.png
-rw-r--r--   1 wxnacy  staff  27356 Jan 19 17:12 react.Bicubic.png
-rw-r--r--@  1 wxnacy  staff  31222 Jan 19 17:09 react.Lanczos3.png
http://wxnacy.com/2019/01/15/go-nfnt-resize/
 */
func Thumbnail(in io.Reader, out io.Writer, width, height, quality int) error {
	origin, fm, err := image.Decode(in)
	if err != nil {
		return err
	}
	if width == 0 || height == 0 {
		width = origin.Bounds().Max.X
		height = origin.Bounds().Max.Y
	}
	if quality == 0 {
		quality = 100
	}
	canvas := resize.Thumbnail(uint(width), uint(height), origin, resize.Lanczos3)

	switch fm {
	case "jpeg":
		return jpeg.Encode(out, canvas, &jpeg.Options{quality})
	case "png":
		return png.Encode(out, canvas)
	case "gif":
		return gif.Encode(out, canvas, &gif.Options{})
	case "bmp":
		return bmp.Encode(out, canvas)
	default:
		return errors.New("ERROR FORMAT")
	}
	return nil
}
