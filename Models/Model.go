package Models

import "time"

func (b *Feedback) TableName() string {
	return "Feedback"
}

type Feedback struct {
	RefId     string    `json:"ref_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Feedback  string    `json:"feedback"`
	TimeStamp time.Time `json:"time_stamp"`
}

type GetFeedBack struct {
	Status   string `json:"status"`
	FeedBack []Feedback
}

type GetAllFeedBack struct {
	OpenEscalateFeedback GetFeedBack
	CloseFeedback        GetFeedBack
}
