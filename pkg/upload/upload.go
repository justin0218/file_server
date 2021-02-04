package upload

import (
	"bytes"
	"file_server/api/proto"
	"fmt"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func Local(fileName string, src []byte) (ret *proto.UploadLocalRes, err error) {
	ret = new(proto.UploadLocalRes)
	if fileName == "" {
		fileName = uuid.New().String()
	}
	ret.Format = GetFileType(src)
	//开启压缩
	src, err = CompressedImage(src, ret.Format)
	if err != nil {
		return
	}
	fullFileName := fmt.Sprintf("/www/resources/%s.%s", fileName, ret.Format)
	file, e := os.Create(fullFileName)
	if e != nil {
		err = e
		return
	}
	n, e := file.Write(src)
	if e != nil {
		err = e
		return
	}
	ret.Size = int64(n)
	ret.Name = fileName
	ret.Url = fmt.Sprintf("http://momoman.cn/resources/%s.%s", fileName, ret.Format)
	return
}

func CompressedImage(src []byte, format string) (after []byte, err error) {
	var img image.Image
	if format == "jpg" || format == "jpeg" {
		img, err = jpeg.Decode(bytes.NewBuffer(src))
		if err != nil {
			return
		}
		buf := new(bytes.Buffer)
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 10})
		if err != nil {
			return
		}
		after = buf.Bytes()
		return
	} else if format == "png" {
		img, err = png.Decode(bytes.NewBuffer(src))
		img = resize.Thumbnail(100, 100, img, resize.Lanczos3)
		buf := new(bytes.Buffer)
		err = png.Encode(buf, img)
		after = buf.Bytes()
		return
	} else {
		after = src
		return
	}
	return
}
