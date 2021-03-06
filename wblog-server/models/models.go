package models

import (
	"github.com/gookit/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
	"wblog-server/system"
)

// BaseModel I don't need soft delete,so I use customized BaseModel instead gorm.Model
type BaseModel struct {
	Id        string    `json:"id" gorm:"type:varchar(64);primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// table post_tags
type PostTag struct {
	BaseModel
	PostId string // post id
	TagId  string // tag id

}

// query result
type QrArchive struct {
	ArchiveDate time.Time //month
	Total       int       //total
	Year        int       // year
	Month       int       // month
}

type SmmsFile struct {
	BaseModel
	FileName  string `json:"filename"`
	StoreName string `json:"storename"`
	Size      int    `json:"size"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Hash      string `json:"hash"`
	Delete    string `json:"delete"`
	Url       string `json:"url"`
	Path      string `json:"path"`
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	//db, err := gorm.Open(sqlite.Open(system.GetConfiguration().DSN), &gorm.Config{})
	g := system.GetConfiguration()
	color.Redln("链接数据库")
	color.Redln(g.Mysql)
	//dsn := g.Mysql.User + ":" + g.Mysql.Pass + "@tcp(127.0.0.1:3306)/" + g.Mysql.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=localhost user=postgres password=" + g.Pgsql.Pass + " dbname=wblog port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err == nil {
		DB = db
		//db.LogMode(true)
		_ = db.AutoMigrate(&Page{}, &Post{}, &Tag{}, &PostTag{}, &User{}, &Comment{}, &Subscriber{}, &Link{}, &SmmsFile{})
		//db.Model(&PostTag{}).AddUniqueIndex("uk_post_tag", "post_id", "tag_id")
		return db, err
	}
	return nil, err
}

// post_tags
func (pt *PostTag) Insert() error {
	return DB.FirstOrCreate(pt, "post_id = ? and tag_id = ?", pt.PostId, pt.TagId).Error
}

func DeletePostTagByPostId(postId string) error {
	return DB.Delete(&PostTag{}, "post_id = ?", postId).Error
}

/*func GetLinkByUrl(url string) (*Link, error) {
	var link Link
	err := DB.Find(&link, "url = ?", url).Error
	return &link, err
}*/

func (sf SmmsFile) Insert() (err error) {
	err = DB.Create(&sf).Error
	return
}
