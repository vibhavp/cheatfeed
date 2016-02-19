all: windows386 windows64 linux386 linuxamd64

windows386:
	GOOS=windows GOARCH=386 go build
	zip windows_386.zip logs.html create.html logs.js cheatfeed.exe

windows64:
	GOOS=windows GOARCH=amd64 go build
	zip windows_amd64.zip logs.html create.html logs.js cheatfeed.exe

linux386:
	GOOS=linux GOARCH=386 go build
	zip linux_386.zip logs.html create.html logs.js cheatfeed

linuxamd64:
	GOOS=linux GOARCH=amd64 go build
	zip linux_amd64.zip logs.html create.html logs.js cheatfeed
