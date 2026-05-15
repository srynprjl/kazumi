package models

type Speed struct {
	Enabled bool    `json:"enabled"`
	Value   float64 `json:"value"`
}

type Pitch struct {
	Enabled bool    `json:"enabled"`
	Value   float64 `json:"value"`
}

type Reverb struct {
	Enabled bool    `json:"enabled"`
	InGain  float64 `json:"ingain"`
	OutGain float64 `json:"outgain"`
	Decay   float64 `json:"decays"`
	Delay   float64 `json:"delays"`
}

type Options struct {
	Speed  Speed
	Pitch  Pitch
	Reverb Reverb
}

type JSONConfig struct {
	VideoURL string `json:"video"`
	ImageURL string `json:"image"`
	Video    bool   `json:"make_video"`
	Speed    Speed  `json:"speed"`
	Pitch    Pitch  `json:"pitch"`
	Reverb   Reverb `json:"reverb"`
}
