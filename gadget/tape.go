package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) { //створюємо метод, зсилаємось на (t TapePlayer), назва-Play, і передаємо пісню
	fmt.Println("Plaing ", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stop!")
}

type TapeRecorder struct {
	Mirofones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Plaing ", song)
}
func (t TapeRecorder) Stop() {
	fmt.Println("Stop!")
}
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
