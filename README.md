## How to use ?
___
1) Clone the repo:
```
$ git clone https://github.com/baturalp-kiziltan/static-music-stream-go.git
```
2) Switch to the project's root dir:
```
$ cd static-music-stream-go
```
3) Build an executable using "main.go" file
```
$ go build
```
4) Execute the binary or .exe output which will be created after build process
___
5) Finally, the stream will be available on your local: "http://localhost:4444/stream"
___
* Probably, you will want to change "audio.mp3" file (at the project's root directory) by any type of MPEG supported music file. 
  Just be sure about its naming (must be "audio.mp3"). Optionally, you can edit line:17, colon:50-59 from "main.go" 
  source file to use any naming.