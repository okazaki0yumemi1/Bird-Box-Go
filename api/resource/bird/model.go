package bird

import (
	"time"

	"github.com/google/uuid"
)

// type DTO struct {
// 	ID            string `json:"id"`
// 	CommonName    string `json:"name"`
// 	Confidence    string `json:"confidence"`
// 	RecordingDate string `json:"recoding_date"`
// 	InputDevice   string `json:"microphone_name"`
// 	// Description   string `json:"description"`
// }

// type Form struct {
// 	CommonName    string `json:"name"`
// 	Confidence    string `json:"confidence"`
// 	RecordingDate string `json:"recoding_date"`
// 	InputDevice   string `json:"microphone_name"`
// 	// Description   string `json:"description"`
// }

type Bird struct {
	ID            uuid.UUID `gorm:"primarykey"`
	CommonName    string
	Confidence    string
	RecordingDate time.Time
	InputDevice   string
}

type Birds []*Bird
