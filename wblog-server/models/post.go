package models

import (
	"database/sql"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
	"strconv"
	"time"
)

// Post

// table posts
type Post struct {
	BaseModel
	Title        string     // title
	Body         string     // body
	View         int        // view count
	IsPublished  bool       // published or not
	Tags         []*Tag     `gorm:"-"` // tags of post
	Comments     []*Comment `gorm:"-"` // comments of post
	CommentTotal int        `gorm:"-"` // count of comment
}

func (post *Post) Insert() error {
	return DB.Create(post).Error
}

func (post *Post) Update() error {
	return DB.Model(post).Updates(map[string]interface{}{
		"title":        post.Title,
		"body":         post.Body,
		"is_published": post.IsPublished,
	}).Error
}

func (post *Post) UpdateView() error {
	return DB.Model(post).Updates(map[string]interface{}{
		"view": post.View,
	}).Error
}

func (post *Post) Delete() error {
	return DB.Delete(post).Error
}

func (post *Post) Excerpt() template.HTML {
	//you can sanitize, cut it down, add images, etc
	policy := bluemonday.StrictPolicy() //remove all html tags
	sanitized := policy.Sanitize(string(blackfriday.MarkdownCommon([]byte(post.Body))))
	runes := []rune(sanitized)
	if len(runes) > 300 {
		sanitized = string(runes[:300])
	}
	excerpt := template.HTML(sanitized + "...")
	return excerpt
}

func ListPublishedPost(tag string, pageIndex, pageSize int) ([]*Post, error) {
	return _listPost(tag, true, pageIndex, pageSize)
}

func ListAllPost(tag string) ([]*Post, error) {
	return _listPost(tag, false, 0, 0)
}

func _listPost(tag string, published bool, pageIndex, pageSize int) ([]*Post, error) {
	var posts []*Post
	var err error
	if len(tag) > 0 {
		tagId, err := strconv.ParseUint(tag, 10, 64)
		if err != nil {
			return nil, err
		}
		var rows *sql.Rows
		if published {
			if pageIndex > 0 {
				rows, err = DB.Raw("select p.* from posts p inner join post_tags pt on p.id = pt.post_id where pt.tag_id = ? and p.is_published = ? order by created_at desc limit ? offset ?", tagId, true, pageSize, (pageIndex-1)*pageSize).Rows()
			} else {
				rows, err = DB.Raw("select p.* from posts p inner join post_tags pt on p.id = pt.post_id where pt.tag_id = ? and p.is_published = ? order by created_at desc", tagId, true).Rows()
			}
		} else {
			rows, err = DB.Raw("select p.* from posts p inner join post_tags pt on p.id = pt.post_id where pt.tag_id = ? order by created_at desc", tagId).Rows()
		}
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var post Post
			DB.ScanRows(rows, &post)
			posts = append(posts, &post)
		}
	} else {
		if published {
			if pageIndex > 0 {
				err = DB.Where("is_published = ?", true).Order("created_at desc").Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&posts).Error
			} else {
				err = DB.Where("is_published = ?", true).Order("created_at desc").Find(&posts).Error
			}
		} else {
			err = DB.Order("created_at desc").Find(&posts).Error
		}
	}
	return posts, err
}

func MustListMaxReadPost() (posts []*Post) {
	posts, _ = ListMaxReadPost()
	return
}

func ListMaxReadPost() (posts []*Post, err error) {
	err = DB.Where("is_published = ?", true).Order("view desc").Limit(5).Find(&posts).Error
	return
}

func MustListMaxCommentPost() (posts []*Post) {
	posts, _ = ListMaxCommentPost()
	return
}

func ListMaxCommentPost() (posts []*Post, err error) {
	var (
		rows *sql.Rows
	)
	rows, err = DB.Raw("select p.*,c.total comment_total from posts p inner join (select post_id,count(*) total from comments group by post_id) c on p.id = c.post_id order by c.total desc limit 5").Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		DB.ScanRows(rows, &post)
		posts = append(posts, &post)
	}
	return
}

func CountPostByTag(tag string) (count int, err error) {
	var (
		tagId uint64
	)
	if len(tag) > 0 {
		tagId, err = strconv.ParseUint(tag, 10, 64)
		if err != nil {
			return
		}
		err = DB.Raw("select count(*) from posts p inner join post_tags pt on p.id = pt.post_id where pt.tag_id = ? and p.is_published = ?", tagId, true).Row().Scan(&count)
	} else {
		err = DB.Raw("select count(*) from posts p where p.is_published = ?", true).Row().Scan(&count)
	}
	return
}

func CountPost() int64 {
	var count int64
	DB.Model(&Post{}).Count(&count)
	return count
}

func GetPostById(id string) (*Post, error) {
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	var post Post
	err = DB.First(&post, "id = ?", pid).Error
	return &post, err
}

func MustListPostArchives() []*QrArchive {
	archives, _ := ListPostArchives()
	return archives
}

func ListPostArchives() ([]*QrArchive, error) {
	var archives []*QrArchive
	//querysql := `select DATE_FORMAT(created_at,'%Y-%m') as month,count(*) as total from posts where is_published = ? group by month order by month desc`
	querysql := `select strftime('%Y-%m',created_at) as month,count(*) as total from posts where is_published = ? group by month order by month desc`
	rows, err := DB.Raw(querysql, true).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var archive QrArchive
		var month string
		rows.Scan(&month, &archive.Total)
		//DB.ScanRows(rows, &archive)
		archive.ArchiveDate, _ = time.Parse("2006-01", month)
		archive.Year = archive.ArchiveDate.Year()
		archive.Month = int(archive.ArchiveDate.Month())
		archives = append(archives, &archive)
	}
	return archives, nil
}

func ListPostByArchive(year, month string, pageIndex, pageSize int) ([]*Post, error) {
	var (
		rows *sql.Rows
		err  error
	)
	if len(month) == 1 {
		month = "0" + month
	}
	condition := fmt.Sprintf("%s-%s", year, month)
	if pageIndex > 0 {
		//querysql := `select * from posts where date_format(created_at,'%Y-%m') = ? and is_published = ? order by created_at desc limit ? offset ?`
		querysql := `select * from posts where strftime('%Y-%m',created_at) = ? and is_published = ? order by created_at desc limit ? offset ?`
		rows, err = DB.Raw(querysql, condition, true, pageSize, (pageIndex-1)*pageSize).Rows()
	} else {
		//querysql := `select * from posts where date_format(created_at,'%Y-%m') = ? and is_published = ? order by created_at desc`
		querysql := `select * from posts where strftime('%Y-%m',created_at) = ? and is_published = ? order by created_at desc`
		rows, err = DB.Raw(querysql, condition, true).Rows()
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := make([]*Post, 0)
	for rows.Next() {
		var post Post
		DB.ScanRows(rows, &post)
		posts = append(posts, &post)
	}
	return posts, nil
}

func CountPostByArchive(year, month string) (count int, err error) {
	if len(month) == 1 {
		month = "0" + month
	}
	condition := fmt.Sprintf("%s-%s", year, month)
	//querysql := `select count(*) from posts where date_format(created_at,'%Y-%m') = ? and is_published = ? order by created_at desc`
	querysql := `select count(*) from posts where strftime('%Y-%m',created_at) = ? and is_published = ?`
	err = DB.Raw(querysql, condition, true).Row().Scan(&count)
	return
}
