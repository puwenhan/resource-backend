package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
)

type Videos struct {
	Model
	Type int   `json:"type"`
	Name string `json:"name"`
	Href string `json:"href"`
}

func (v Videos)Validate() error {
	return validation.ValidateStruct(&v,
		// 名称不得为空,且大小为1-30字
		validation.Field(
			&v.Name,
			validation.Required.Error("名称不得为空"),
			validation.Length(1, 30).Error("名称为1-30字")),
		// 链接不得为空,且为url地址
		validation.Field(&v.Href,
			validation.Required.Error("链接地址不得为空"),
			is.URL.Error("链接必须为URL地址")),
		)
}

// 获取记录总数
func GetVideosTotalCount(maps interface{}) (count int) {
	db.Model(&Videos{}).Where(maps).Count(&count)
	return
}

// 获取数据列表
func GetVideosList(page int, pageSize int, maps interface{}) ([]Videos, error) {
	var videos []Videos

	err := db.Where(maps).Offset((page - 1) * pageSize).Limit(pageSize).Find(&videos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return videos, nil
}

// 添加数据
func AddVideo(name string, url string, typeOfVideo int) error {
	video := Videos{
		Name: name,
		Href: url,
		Type: typeOfVideo,
	}
	// 数据验证
	err := video.Validate()
	if err != nil {
		return err
	}
	if err := db.Create(&video).Error; err != nil {
		return err
	}
	return nil
}

// 根据ID查找数据是否存在
func GetVideoById(id int) (bool, error) {
	var video Videos
	err := db.Select("id").Where("id = ?", id).First(&video).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}

	if video.ID > 0  {
		return true, nil
	}

	return false, nil
}

// 查看数据详情
func GetVideoView(maps interface{}) (video Videos) {
	db.Where(maps).First(&video)
	return
}

// 修改数据
func PutVideoUpdate(id int,video Videos) (err error) {
	// 数据验证
	err = video.Validate()
	if err != nil {
		return err
	}
	if err = db.Model(&Videos{}).Where("id = ?", id).Update(video).Error; err != nil {
		return err
	}
	return nil
}

// 删除数据
func DeleteVideo(id int) (err error) {
	if err := db.Where("id = ?", id).Delete(&Videos{}).Error; err != nil {
		return err
	}
	return nil
}
