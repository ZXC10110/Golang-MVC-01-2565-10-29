package Controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
	"github.com/zxc10110/mvc_63050096_2565_1/Models"
)

func genShortUUID() (id string) {
	id = shortuuid.New()
	return id
}

func CreateFeedback(c *gin.Context) {
	//call model to request input
	var req Models.Feedback
	c.BindJSON(&req)
	feedback := Models.Feedback{
		RefId:     genShortUUID(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Feedback:  "open",
		TimeStamp: time.Now(),
	}
	//call pugin to create feedback
	er := Models.CreateFeed(&feedback)
	if er != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	} else {
		//print output
		c.JSON(http.StatusOK, feedback)
		return
	}
}

func UpdateFeedback(c *gin.Context) {
	//call model to request input
	var req Models.Feedback
	c.BindJSON(&req)

	//calculate time
	oldTime := req.TimeStamp
	currentTime := time.Now()
	dateEscalate := currentTime.Sub(oldTime)
	diff := dateEscalate.Hours()

	//condition
	if req.Feedback == "close" { //if user update by changing feedback to "close"
		//call pugin to update to database
		er := Models.UpdateFeed(&req)
		if er != nil {
			c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
			return
		}
		//print result
		result := Models.Feedback{
			RefId:     req.RefId,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Feedback:  "close",
			TimeStamp: req.TimeStamp,
		}
		c.JSON(http.StatusOK, result)
		return

	} else if diff > 168 && req.Feedback == "open" { //not modified but date > 7
		//call pugin to update to database
		er := Models.UpdateFeed(&req)
		if er != nil {
			c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
			return
		}
		//print result
		result := Models.Feedback{
			RefId:     req.RefId,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Feedback:  "close",
			TimeStamp: req.TimeStamp,
		}
		c.JSON(http.StatusOK, result)
		return

	} else {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	}
}

func AdminUpdate(c *gin.Context) {
	//call model to request input
	var req Models.Feedback
	c.BindJSON(&req)

	//call pugin to update to database
	er := Models.AdminUpdate(&req)
	if er != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	} else {
		//print result
		result := Models.Feedback{
			RefId:     req.RefId,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Feedback:  "escalate",
			TimeStamp: req.TimeStamp,
		}
		c.JSON(http.StatusOK, result)
		return
	}
}

func GetFeedBack(c *gin.Context) {
	//call pugin
	open, er := Models.GetFeedOpen()
	if er != nil {
		c.JSON(http.StatusNotFound, "ไม่พบข้อมูล")
		return
	}

	//call pugin
	close, err := Models.GetFeedClose()
	if err != nil {
		return
	} else {
		//group output
		openFeed := Models.GetFeedBack{
			Status:   "Open And Escalate",
			FeedBack: open,
		}

		//group output
		closeFeed := Models.GetFeedBack{
			Status:   "Close",
			FeedBack: close,
		}

		//group all output
		groupAll := Models.GetAllFeedBack{
			OpenEscalateFeedback: openFeed,
			CloseFeedback:        closeFeed,
		}

		//print output
		c.JSON(http.StatusOK, groupAll)
		return
	}

}
