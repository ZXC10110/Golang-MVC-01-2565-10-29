package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zxc10110/mvc_63050096_2565_1/Config"
)

//getPlagiarism
func GetPlagiarism01() (plagiarism PlagiarismResult, err error) {
	if err = Config.DB.Select("p.date_time, t.tid, t.first_name, t.last_name, p.similarity_score").
		Table("Test.PlagiarismResult as p, Test.submissions as s, Test.Tester as t").
		Where("p.sid1 = s.sid AND s.tid = t.tid AND p.similarity_score > 0.35").
		Find(&plagiarism).Error; err != nil {
		return
	}
	return
}

//getPlagiarism
func GetPlagiarism02() (plagiarism PlagiarismResult, err error) {
	if err = Config.DB.Select("p.date_time, t.tid, t.first_name, t.last_name, p.similarity_score").
		Table("Test.PlagiarismResult as p, Test.submissions as s, Test.Tester as t").
		Where("p.sid2 = s.sid AND s.tid = t.tid AND p.similarity_score > 0.35").
		Find(&plagiarism).Error; err != nil {
		return
	}
	return
}
