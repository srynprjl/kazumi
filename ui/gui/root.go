package gui

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/srynprjl/kazumi/lib/creation"
	"github.com/srynprjl/kazumi/lib/models"
)

func Gui() {
	jsons := []models.JSONConfig{}
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
	yt_input.OnChanged = func(s string) {
		jsonz.VideoURL = s
	}

	img_input := widget.NewEntry()
	img_input.SetPlaceHolder("Enter image link...")
	img_input.OnChanged = func(s string) {
		jsonz.ImageURL = s
	}

	makeVideoCheck := widget.NewCheck("Make video", func(b bool) {
		jsonz.Video = b
	})

	speed_label := widget.NewLabel("Speed")
	speed_label.TextStyle = fyne.TextStyle{Bold: true}
	speed_enabled := widget.NewCheck("Enable Speed", func(value bool) {
		jsonz.Speed.Enabled = value
	})
	speed_value := widget.NewEntry()
	speed_value.SetPlaceHolder("Speed: 1.25")
	speed_value.OnChanged = func(s string) {
		jsonz.Speed.Value, _ = strconv.ParseFloat(s, 64)
	}

	pitch_label := widget.NewLabel("Pitch")
	pitch_label.TextStyle = fyne.TextStyle{Bold: true}
	pitch_enabled := widget.NewCheck("Enable Pitch", func(value bool) {
		jsonz.Pitch.Enabled = value
	})
	pitch_value := widget.NewEntry()
	pitch_value.SetPlaceHolder("Pitch: 1.33")
	pitch_value.OnChanged = func(s string) {
		jsonz.Pitch.Value, _ = strconv.ParseFloat(s, 64)
	}

	reverb_label := widget.NewLabel("Reverb")
	reverb_label.TextStyle = fyne.TextStyle{Bold: true}
	reverb_enabled := widget.NewCheck("Enable Reverb", func(value bool) {
		jsonz.Reverb.Enabled = value
	})

	ingain_value := widget.NewEntry()
	ingain_value.SetPlaceHolder("InGain")
	ingain_value.OnChanged = func(s string) {
		jsonz.Reverb.InGain, _ = strconv.ParseFloat(s, 64)
	}

	outgain_value := widget.NewEntry()
	outgain_value.SetPlaceHolder("OutGain")
	outgain_value.OnChanged = func(s string) {
		jsonz.Reverb.OutGain, _ = strconv.ParseFloat(s, 64)
	}

	delay_value := widget.NewEntry()
	delay_value.SetPlaceHolder("Delay")
	delay_value.OnChanged = func(s string) {
		jsonz.Reverb.Delay, _ = strconv.ParseFloat(s, 64)
	}

	decay_value := widget.NewEntry()
	decay_value.SetPlaceHolder("Decay")
	decay_value.OnChanged = func(s string) {
		jsonz.Reverb.Decay, _ = strconv.ParseFloat(s, 64)
	}

	entry := container.NewGridWithColumns(2, ingain_value, outgain_value, delay_value, decay_value)

	speed_container := container.NewVBox(speed_label, speed_enabled, speed_value)
	pitch_container := container.NewVBox(pitch_label, pitch_enabled, pitch_value)
	reverb_container := container.NewVBox(reverb_label, reverb_enabled, entry)
	box := container.NewHBox(speed_container, layout.NewSpacer(), pitch_container, layout.NewSpacer(), reverb_container)
	allEntries := []fyne.CanvasObject{yt_input, img_input, makeVideoCheck, speed_value, speed_enabled, pitch_enabled, pitch_value, reverb_enabled, ingain_value, outgain_value, delay_value, decay_value}

	add_btn := widget.NewButton("Add More", func() {
		jsons = append(jsons, jsonz)
		jsonz = models.JSONConfig{}

		// clear all
		for _, field := range allEntries {
			switch w := field.(type) {
			case *widget.Entry:
				w.SetText("")
			case *widget.Check:
				w.SetChecked(false)
			}
		}
	})
	submit_btn := widget.NewButton("Start", func() {
		if jsonz.VideoURL != "" {
			jsons = append(jsons, jsonz)
		}

		if len(jsons) != 0 {
			creation.DownloadUsingJSON(jsons)
			jsons = []models.JSONConfig{}
		}

	})

	btns := container.NewHBox(layout.NewSpacer(), add_btn, submit_btn)
	cont := container.NewVBox(title, yt_input, img_input, makeVideoCheck, box, layout.NewSpacer(), btns)

	myWindow.SetContent(cont)
	myWindow.SetFixedSize(true)
	myWindow.ShowAndRun()

}
