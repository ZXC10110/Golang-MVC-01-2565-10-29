package Models

import "time"

func (b *PlagiarismResult) TableName() string {
	return "PlagiarismResult"
}

type PlagiarismResult struct {
	DateTime        time.Time `json:"date_time"`
	Tid             int       `json:"tid"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	SimilarityScore float64   `json:"similarity_score"`
}

type GetPlagiatism struct {
	Student01 PlagiarismResult
	Student02 PlagiarismResult
}
