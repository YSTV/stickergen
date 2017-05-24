# StickerGen

A little utility that generates PDF pages of asset tags, designed to be printed on A4 label sheets.

## Paper compatibility
The default output is a layout of 4 columns of 12 rows of stickers, suitable for Avery L4778 sticker paper. There are options to change the layout, but they probably don't work properly... :)

## Usage
```
./stickergen -help
Usage of ./stickergen:
  -drawoutlines
    	Draw the sticker outlines. Useful for testing.
  -h float
    	Sticker height (mm) (default 21.2)
  -numx int
    	Number of stickers wide (default 4)
  -numy int
    	Number of stickers high (default 12)
  -outfile string
    	Path to output file. (default "stickers.pdf")
  -startnum int
    	Asset number to start at (default 1)
  -w float
    	Sticker width (mm) (default 45.7)
```
