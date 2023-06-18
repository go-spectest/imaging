# gina - Go Image 'N' Assistance
<p align="center">
  <img src="img/question.png" alt="Image">
</p>

The gina command is a CLI tool for image processing. The gina was created to help developers understand 'how to use the image
processing methods provided by the [nao1215/imaging](https://github.com/nao1215/imaging) package'.

The gina aims to be a command with multiple subcommands, focusing on simple functionality rather than being a feature-rich command like [ImageMagick](https://github.com/ImageMagick/ImageMagick).

The gina supports the following image formats:
- JPEG
- GIF
- PNG
- BMP
- TIFF

## How to install / build
### Use go install
If you does not have the golang development environment installed on your system, please install golang from the [golang official website](https://go.dev/doc/install).
```shell
$ go install github.com/nao1215/imaging/cmd/gina@latest
```

### Build gina
```shell
$ git clone git@github.com:nao1215/imaging.git
$ cd imaging
$ make build
```

## How to use
### Available subcommands
```
Available Commands:
  blur        Blur the image according to sigma
  bug-report  Submit a bug report at GitHub
  help        Help about any command
  resize      Resize image
  version     Show imaging command version information
```
### Resize subcommand
resize subcommand resizes the image specified argument and saves it to the file specified by the --output parameter. --output default value is 'output.jpg'.

If you specify either the height or width, the aspect ratio will be maintained during resizing. The file extension specified in the --output parameter can be different from the input image's extension.

```
$ gina resize --width 100 --output resize_awesome.png cmd/gina/img/awesome.png 
save image: resize_awesome.png
```
Original image                     | Resize width=100                            | 
-----------------------------------|----------------------------------------|
![srcImage](img/awesome.png) | ![dstImage](img/resize_awesome.png) |


### Blur subcommand
The blur subcommand outputs an image with blur effect intensity according to the sigma value
```
$ gina blur -s 5.0 --output blur_awesome.png cmd/gina/img/awesome.png 
save image: blur_awesome.png
```
Original image                     | Blur sigma=5.0                           | 
-----------------------------------|----------------------------------------|
![srcImage](img/awesome.png) | ![dstImage](img/blur_awesome.png) |


## LICENSE
### gina command
The gina command is licensed under the MIT License.
- See [LICENSE](./LICENSE) file.
- gina command does not exist original [disintegration/imaging](https://github.com/disintegration/imaging) repository.

### Illustration
This readme use the illustration of Go gopher.
- The Go gopher was designed by Renée French. Illustrations by tottie.
- ©tottie / Renée French 
- Ref. [https://github.com/tottie000/GopherIllustrations](https://github.com/tottie000/GopherIllustrations)