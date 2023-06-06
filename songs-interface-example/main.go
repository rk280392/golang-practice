package main

import "github.com/rk280392/golang-practice/my-packages/interfaceGadget"

func playlist(device interfaceGadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := interfaceGadget.TapePlayer{}
	mixtape := []string{"pasoori", "mashooka", "play it all"}
	playlist(player, mixtape)
}
