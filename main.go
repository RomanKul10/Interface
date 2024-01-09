package main

import "gadget/gadget"

type Player interface {
	Play(string)
	Stop()
}

func playlist(device Player, song []string) {
	for _, song := range song {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Track")
	player.Stop()
	recorder := player.(gadget.TapeRecorder)
	recorder.Record()
}

func main() {
	var player Player = gadget.TapePlayer{}
	mixtape := []string{
		"суперпісня",
		"не суперпісня",
		"інша пісня"}
	playlist(player, mixtape)

	player = gadget.TapeRecorder{}
	playlist(player, mixtape)

	TryOut(gadget.TapeRecorder{})
}
