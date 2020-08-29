# Preparation

## Mandatory

* [x] research the emulation process
  * [x] read [this guide](http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/)
  * [x] [Cowgod's Chip-8 Technical Reference v1.0](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM)
  * [x] [Chip8 Tutorial](http://www.multigesture.net/wp-content/uploads/mirror/goldroad/chip8.shtml)

## Further Reading
* [ ] read [Zenogais' Tutorials](http://www.multigesture.net/wp-content/uploads/mirror/zenogais/Tutorials.htm)
  * [x] [Introduction to Emulation Part 1](http://www.multigesture.net/wp-content/uploads/mirror/zenogais/EmuDoc1.htm)
  * [x] [Introduction to Emulation Part 2](http://www.multigesture.net/wp-content/uploads/mirror/zenogais/EmuDoc2.html) 
  * [ ] [Laying the Ground For An Emulator](http://www.multigesture.net/wp-content/uploads/mirror/zenogais/GroundWork.html) 
  * [ ] [Dynamic Recompiler](http://www.multigesture.net/wp-content/uploads/mirror/zenogais/Dynamic%20Recompiler.html)
  * [x] [Array of Function Pointers](http://www.multigesture.net/wp-content/uploads/mirror/zenogais/FunctionPointers.htm) 
* [x] read [Zilmar's Emubook](http://emubook.emulation64.com/)

# General 
* [x] write Memory
* [x] write CPU
* [x] write Game Loop
* [x] write Graphics
* [ ] handle Input
* [x] handle proper timing for timers (60Hz)
* [ ] write Debugger
* [ ] test roms  
  * [x] BC test rom
  * [ ] PONG(2)  
    Technically works, but the input handling is sluggish
  * [ ] Breakout/Brix  
    Loads graphics, but won't let me play (timing issue?)
  * [ ] Tetris  
    Does not load graphics properly, blocks seem to be stuck at the top and don't move down but continuously spawn? Also the score flashes in and out of existance

# CPU
* [x] implement instructions
* [x] set up registers
* [x] learn how to properly write tests (in go specifically)
* [x] write tests for instructions

# Memory
* [x] implement memory map
* [x] test memory


# Graphics
* [x] research how to implement graphics
* [x] research SDL
* [x] set up screen (64x32 = 2048 px, bw, boolean)
* [ ] read [Lazy Foo's Tutorials](http://lazyfoo.net/tutorials/SDL/)  
  * [x] Lesson 01
  * [x] Lesson 02
  * [x] Lesson 03
  * [x] Lesson 04
  * [x] Lesson 05
  * [x] Lesson 06
  * [x] Lesson 07
  * [x] Lesson 08
  * [x] Lesson 09
  * [x] Lesson 10
  * [x] Lesson 11
  * [x] Lesson 12
  * [x] Lesson 13
  * [ ] Lesson 18
  * [ ] Lesson 19
  * [ ] Lesson 21

# Input
* [x] set up hex based keypad
* [x] set up input.go
* [ ] improve Input handling

# Audio
* [ ] buzzer sound
