[![Build Status](https://travis-ci.com/Elgarni/Ascii-Art-with-Docker-and-Go.svg?branch=master)](https://travis-ci.com/Elgarni/Ascii-Art-with-Docker-and-Go)

# Description
This project utilizes [Go-figure](https://github.com/common-nighthawk/go-figure), by making it run in a Docker image.

# Usage
The very simple usage :
```shell
docker run --rm -e text="Hello there" elqarni/hala
```

### Available Options
Options are passed to the image as environment variables. 
* text | t : The text to be displayed on the Terminal
* move | m :  Defines what move will be applied to the text. Available moves are : _**scroll**_ and _**blink**_
```shell
docker run --rm -e text="Hello there" -e move=blink elqarni/hala
```
* hidden | h : A boolean option that lets you hide the text, by passing a base64 encoded text.

For instance, "Hello there" is encoded with base64 to "SGVsbG8gdGhlcmU="

```shell
docker run --rm -e text="SGVsbG8gdGhlcmU=" -e hidden=true elqarni/hala
```
