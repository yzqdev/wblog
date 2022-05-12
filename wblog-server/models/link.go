package models

import "time"

// Link
// table link
type Link struct {
	BaseModel
	DeletedAt time.Time `json:"deleted_at"`
	Name      string    `json:"name"`                  //名称
	Url       string    `json:"url"`                   //地址
	Sort      int       `json:"sort" gorm:"default:0"` //排序
	View      int       `json:"view"`                  //访问次数
}

func (link *Link) Insert() error {
	return DB.FirstOrCreate(link, "url = ?", link.Url).Error
}

func (link *Link) Update() error {
	return DB.Save(link).Error
}

func (link *Link) Delete() error {
	return DB.Delete(link).Error
}

func ListLinks() ([]*Link, error) {
	var links []*Link
	err := DB.Order("sort asc").Find(&links).Error
	return links, err
}

func MustListLinks() []*Link {
	links, _ := ListLinks()
	return links
}

func GetLinkById(id uint) (*Link, error) {
	var link Link
	err := DB.FirstOrCreate(&link, "id = ?", id).Error
	return &link, err
}
