package creation

type Speed struct {
	Enabled bool
	Value   float64
}

type Pitch struct {
	Enabled bool
	Value   float64
}

type Reverb struct {
	Enabled bool
	InGain  float32
	OutGain float32
	Decay   float32
	Delay   float32
}

type Options struct {
	Speed  Speed
	Pitch  Pitch
	Reverb Reverb
}
