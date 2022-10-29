package Models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zxc10110/mvc_63050096_2565_1/Config"
)

//createPic
func CreateFeed(feedback *Feedback) (err error) {
	if err = Config.DB.Create(&feedback).Error; err != nil {
		return
	}
	return
}

//update
func UpdateFeed(req *Feedback) (err error) {
	if err = Config.DB.Table("Test.Feedback").
		Where("ref_id = ? AND feedback = ? OR feedback = ?", req.RefId, "open", "escalate").
		Update(map[string]interface{}{
			"time_stamp": time.Now(),
			"feedback":   "close",
		}).
		Error; err != nil {
		return
	}
	return
}

//admin update
func AdminUpdate(req *Feedback) (err error) {
	if err = Config.DB.Table("Test.Feedback").
		Where("ref_id = ? AND feedback = ?", req.RefId, "open").
		Update(map[string]interface{}{
			"time_stamp": time.Now(),
			"feedback":   "escalate",
		}).
		Error; err != nil {
		return
	}
	return
}

//get feedback open
func GetFeedOpen() (feed []Feedback, err error) {
	if err = Config.DB.Where("feedback = ? or feedback = ?", "open", "escalate").
		Order("feedback").
		Find(&feed).Error; err != nil {
		return
	}
	return
}

//get feedback close
func GetFeedClose() (feed []Feedback, err error) {
	if err = Config.DB.Where("feedback = ?", "close").
		Find(&feed).Error; err != nil {
		return
	}
	return
}
