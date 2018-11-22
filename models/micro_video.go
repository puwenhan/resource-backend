package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
)

type MicroVideo struct {
	Model
	Url  string `json:"url"`
	View int `json:"view"`
}

func (v MicroVideo) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Url,
			validation.Required.Error("链接地址不得为空"),
			is.URL.Error("视频上传错误")),
	)
}

// 获取数据列表
func MicroVideoList(page int, pageSize int) (microVideos []MicroVideo, count int, err error) {
	err = db.Select("url").Offset((page - 1) * pageSize).Limit(pageSize).Find(&microVideos).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, count, err
	}
	return microVideos, count, nil
}

// 添加数据
func AddMicroVideo(microVideo *MicroVideo) error {
	// 数据验证
	err := microVideo.Validate()
	if err != nil {
		return err
	}
	if err := db.Create(microVideo).Error; err != nil {
		return err
	}
	return nil
}

// 查看数据详情
func GetMicroVideoView(maps interface{}) (MicroVideo MicroVideo) {
	db.Where(maps).First(&MicroVideo)
	return
}

// 删除数据
func DeleteMicroVideo(id int) (err error) {
	if err := db.Where("id = ?", id).Delete(&MicroVideo{}).Error; err != nil {
		return err
	}
	return nil
}
