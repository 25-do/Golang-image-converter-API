
# Image Conversion API

This is an image conversion API built using Golang and `bimg`, which utilizes `libvips` for efficient image processing. The API allows users to upload images and convert them into different formats.

## Installation

### Prerequisites
This project requires `libvips`, which is a fast image processing library.

#### Install `libvips` on Linux
Run the following command based on your distribution:

```sh
# Ubuntu/Debian
sudo apt update && sudo apt install -y libvips-dev

# Arch Linux
sudo pacman -S vips

# Fedora
sudo dnf install vips-devel
```

#### Install `libvips` on macOS

```sh
brew install vips
```

### Install Go Dependencies

```sh
go mod tidy
```

## Running the API
To start the server, run:

```sh
go run main.go
```

The server will start on `http://localhost:8080`.

## API Usage

### Convert an Image
You can convert images by sending a `POST` request with the image file.

#### Supported Formats
- `jpeg`
- `png`
- `webp`
- `tiff`
- `avif`

### Using cURL

#### Convert a Single Image to JPEG
```sh
curl -X POST -F "file=@/path/to/image.png" "http://localhost:8080/convert?format=jpeg"
```

#### Convert a Single Image to PNG
```sh
curl -X POST -F "file=@/path/to/image.jpg" "http://localhost:8080/convert?format=png"
```

#### Convert a Single Image to WebP
```sh
curl -X POST -F "file=@/path/to/image.jpg" "http://localhost:8080/convert?format=webp"
```

#### Convert Multiple Images
You can upload multiple images by using multiple `-F` flags.

```sh
curl -X POST \
    -F "files=@/path/to/image1.jpg" \
    -F "files=@/path/to/image2.png" \
    -F "files=@/path/to/image3.webp" \
    "http://localhost:8080/convert?format=jpeg"
```

### Response
On success, the API returns a JSON response with the converted file paths:

```json
{
  "converted_files": [
    "output/image1_converted.jpeg",
    "output/image2_converted.jpeg",
    "output/image3_converted.jpeg"
  ]
}
```

## License
This project is open-source and available under the MIT License.

