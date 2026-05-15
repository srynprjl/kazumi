package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/srynprjl/kazumi/lib/models"
)

func Gui() {
	// jsons := []models.JSONConfig{}
	jsonz := models.JSONConfig{}

	myApp := app.New()
	myWindow := myApp.NewWindow("Kazumi")

	// Title
	t := widget.NewLabel("KAZUMI")
	t.TextStyle.Bold = true

	title := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), t, layout.NewSpacer())

	// Input
	yt_input := widget.NewEntry()
	yt_input.SetPlaceHolder("Enter youtube link...")

	img_input := widget.NewEntry()
	img_input.SetPlaceHolder("Enter image link...")

	speed_label := widget.NewLabel("Speed")
	speed_label.TextStyle = fyne.TextStyle{Bold: true}
	speed_enabled := widget.NewCheck("Enable Speed", func(value bool) {
		jsonz.Speed.Enabled = value
	})
	speed_value := widget.NewEntry()
	speed_value.SetPlaceHolder("Speed: 1.25")

	pitch_label := widget.NewLabel("Pitch")
	pitch_label.TextStyle = fyne.TextStyle{Bold: true}
	pitch_enabled := widget.NewCheck("Enable Pitch", func(value bool) {
		jsonz.Pitch.Enabled = value
	})
	pitch_value := widget.NewEntry()
	pitch_value.SetPlaceHolder("Pitch: 1.33")

	reverb_label := widget.NewLabel("Reverb")
	reverb_label.TextStyle = fyne.TextStyle{Bold: true}
	reverb_enabled := widget.NewCheck("Enable Reverb", func(value bool) {
		jsonz.Reverb.Enabled = value
	})

	ingain_value := widget.NewEntry()
	ingain_value.SetPlaceHolder("InGain")

	outgain_value := widget.NewEntry()
	outgain_value.SetPlaceHolder("OutGain")

	delay_value := widget.NewEntry()
	delay_value.SetPlaceHolder("Delay")

	decay_value := widget.NewEntry()
	decay_value.SetPlaceHolder("Decay")

	entry := container.NewGridWithColumns(2, ingain_value, outgain_value, delay_value, decay_value)

	speed_container := container.NewVBox(speed_label, speed_enabled, speed_value)
	pitch_container := container.NewVBox(pitch_label, pitch_enabled, pitch_value)
	reverb_container := container.NewVBox(reverb_label, reverb_enabled, entry)
	box := container.NewHBox(speed_container, layout.NewSpacer(), pitch_container, layout.NewSpacer(), reverb_container)

	add_btn := widget.NewButton("Add More", func() {
	})
	submit_btn := widget.NewButton("Start", func() {

	})

	btns := container.NewHBox(layout.NewSpacer(), add_btn, submit_btn)
	cont := container.NewVBox(title, yt_input, img_input, box, layout.NewSpacer(), btns)

	myWindow.SetContent(cont)
	myWindow.SetFixedSize(true)
	myWindow.ShowAndRun()

}
