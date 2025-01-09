package bird

type DTO struct {
	ID            string `json:"id"`
	CommonName    string `json:"name"`
	Confidence    string `json:"confidence"`
	RecordingDate string `json:"recoding_date"`
	InputDevice   string `json:"microphone_name"`
	// Description   string `json:"description"`
}

type Form struct {
	CommonName    string `json:"name"`
	Confidence    string `json:"confidence"`
	RecordingDate string `json:"recoding_date"`
	InputDevice   string `json:"microphone_name"`
	// Description   string `json:"description"`
}
