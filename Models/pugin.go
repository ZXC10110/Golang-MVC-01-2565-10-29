package Models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zxc10110/mvc_63050096_2565_1/Config"
)

//create feed
func CreateFeed(feedback *Feedback) (err error) {
	if err = Config.DB.Create(&feedback).Error; err != nil {
		return
	}
	return
}

//update
func UpdateFeed(req *Feedback) (err error) {
	if err = Config.DB.Table("Test.Feedback").
		Where("ref_id = ? AND feedback_status = ? OR feedback_status = ?", req.RefId, "open", "escalate").
		Update(map[string]interface{}{
			"time_stamp":      time.Now(),
			"feedback_status": "close",
		}).
		Error; err != nil {
		return
	}
	return
}

//admin update
func AdminUpdate(req *Feedback) (err error) {
	if err = Config.DB.Table("Test.Feedback").
		Where("ref_id = ? AND feedback_status = ?", req.RefId, "open").
		Update(map[string]interface{}{
			"time_stamp":      time.Now(),
			"feedback_status": "escalate",
		}).
		Error; err != nil {
		return
	}
	return
}

//get feedback open escalate
func GetFeedOpen() (feed []Feedback, err error) {
	if err = Config.DB.Where("feedback_status = ?", "open").
		Find(&feed).Error; err != nil {
		return
	}
	return
}

//get feedback open escalate
func GetFeedOpenEscalate() (feed []Feedback, err error) {
	if err = Config.DB.Where("feedback_status = ? or feedback_status = ?", "open", "escalate").
		Order("feedback_status").
		Find(&feed).Error; err != nil {
		return
	}
	return
}

//get feedback close
func GetFeedClose() (feed []Feedback, err error) {
	if err = Config.DB.Where("feedback_status = ?", "close").
		Find(&feed).Error; err != nil {
		return
	}
	return
}
