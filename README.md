# chip-infinite
Yet another Chip-8 emulator made purely for educational purposes

## Build Instructions

### Install Dependencies
SDL2 is required for this program.

Then clone this repository and build the executable:
```
git clone https://github.com/OutOfCache/chip-infinite.git
cd chip-infinite
go build
```

## Usage
```
./chip-infinite /path/to/rom
```

## Future Features
* [ ] debugger
* [ ] support for SUPER-CHIP
* [ ] command line options  
  * [ ] screen size
  * [ ] configure button layout
* [ ] GUI

## Resources
### Chip-8 and its emulation
* [Wikipedia](http://en.wikipedia.org/wiki/CHIP-8)
* [How to write an emulator (CHIP-8 interpreter)](http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/)
* [Cowgod's Chip-8 Technical Reference](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#00E0)
* [Zenogais' Emulation Tutorials](http://www.multigesture.net/wp-content/uploads/mirror/zenogais/Tutorials.htm)
* [Zilmar's EmuBook](http://emubook.emulation64.com/)
* [metteo's test rom](https://github.com/metteo/chip8-test-rom)
* [corax89's test rom](https://github.com/corax89/chip8-test-rom)

### SDL
* [Lazy Foo' Productions' SDL Tutorials](https://lazyfoo.net/tutorials/SDL/index.php)
* [go-sdl2's godoc page](https://godoc.org/github.com/veandco/go-sdl2/sdl#PauseAudio)
* [go-sdl2's examples](https://github.com/veandco/go-sdl2-examples/tree/e79e66a8c075ffd2bd0511f9f2f6f7f7047d4c4c/examples)
* [SDL Wiki](https://wiki.libsdl.org/FrontPage)
* [Audio Programming Tutorial](https://www.youtube.com/playlist?list=PLEETnX-uPtBVpZvp-R2daNfy9k3-L-Q3u)
* [golang's cgo (for SDL Audio)](https://github.com/golang/go/wiki/cgo#function-variables)


### Helpful github projects
* https://github.com/jamesmcm/chip8go
* https://github.com/fogleman/nes

