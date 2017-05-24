package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/jpeg"
	"log"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	var outFile = flag.String("outfile", "stickers.pdf", "Path to output file.")
	var stickersX = flag.Int("numx", 4, "Number of stickers wide")
	var stickersY = flag.Int("numy", 12, "Number of stickers high")
	var stickerW = flag.Float64("w", 45.7, "Sticker width (mm)")
	var stickerH = flag.Float64("h", 21.2, "Sticker height (mm)")
	var startNumber = flag.Int("startnum", 1, "Asset number to start at")
	var drawOutlines = flag.Bool("drawoutlines", false, "Draw the sticker outlines. Useful for testing.")

	flag.Parse()

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(10, 21.3, 10)
	pdf.AddPage()
	pdf.RegisterImageOptions("ystv_logo.png", gofpdf.ImageOptions{"", true})
	for row := 0; row < *stickersY; row++ {
		for col := 0; col < *stickersX; col++ {
			stickerX, stickerY := 10+(float64(col)*(*stickerW+2.6)), 21.3+(float64(row)**stickerH)
			if *drawOutlines {
				// Draw sticker outline
				pdf.MoveTo(stickerX, stickerY)
				pdf.LineTo(stickerX+*stickerW, stickerY)
				pdf.LineTo(stickerX+*stickerW, stickerY+*stickerH)
				pdf.LineTo(stickerX, stickerY+*stickerH)
				pdf.LineTo(stickerX, stickerY)
				pdf.ClosePath()
				pdf.SetLineWidth(0.2)
				pdf.DrawPath("D")
			}
			// Text
			pdf.SetFont("helvetica", "", 7)
			pdf.SetXY(stickerX+17.5, stickerY+14)
			pdf.Cell(20, 3, "York Student Television")
			pdf.SetXY(stickerX+17.5, stickerY+17)
			pdf.Cell(20, 3, "ystv.co.uk")
			// Logo
			pdf.ImageOptions("ystv_logo.png", stickerX+20, stickerY+1.5, 22, 0, false, gofpdf.ImageOptions{"", true}, 0, "")
			// Barcode
			num := col + (row * *stickersX) + *startNumber
			assetBarcode, err := qr.Encode(fmt.Sprintf("YSTV%d", num), qr.L, qr.Auto)
			if err != nil {
				log.Fatal(err)
			}
			assetBarcode, err = barcode.Scale(assetBarcode, 200, 200)
			if err != nil {
				log.Fatal(err)
			}

			var b bytes.Buffer
			if err := jpeg.Encode(&b, assetBarcode, &jpeg.Options{100}); err != nil {
				log.Fatal(err)
			}
			imageName := fmt.Sprintf("barcode-%d", num)
			pdf.RegisterImageOptionsReader(imageName, gofpdf.ImageOptions{"jpg", true}, &b)
			pdf.ImageOptions(imageName, stickerX+1, stickerY+1, 16, 0, false, gofpdf.ImageOptions{"jpg", true}, 0, "")
			pdf.SetXY(stickerX+3, stickerY+17.3)
			pdf.SetFont("courier", "", 9)
			pdf.Cell(20, 3, fmt.Sprintf("%.5d", num))
		}
	}
	if err := pdf.OutputFileAndClose(*outFile); err != nil {
		log.Fatal(err)
	}
}
