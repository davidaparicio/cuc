#!/usr/bin/env bash

#go get -u github.com/faiface/beep@latest
#go get -u github.com/faiface/beep/mp3@latest
#go get -u github.com/faiface/beep/speaker@latest

CGO_ENABLED=1 go run ./main.go --URL "http://neverssl.com/makeSSLgreatAgain" -c 404 -f "assets/mp3/ubuntu_desktop_login.mp3"
CGO_ENABLED=1 go run ./main.go loop -s 15 -c 200 -f "assets/mp3/ubuntu_dialog_info.mp3"