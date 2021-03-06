package models

import "strconv"

// Comment

// table comments
type Comment struct {
	BaseModel
	UserID    uint   `json:"user_id"`                         // 用户id
	Content   string `json:"content"`                         // 内容
	PostID    uint   `json:"post_id"`                         // 文章id
	ReadState bool   `json:"read_state" gorm:"default:false"` // 阅读状态
	//Replies []*Comment // 评论
	Nickname  string `json:"nickname" `
	AvatarUrl string `json:"avatar_url" `
	GithubUrl string `json:"github_url" `
}

func (comment *Comment) Insert() error {
	return DB.Create(comment).Error
}

func (comment *Comment) Update() error {
	return DB.Model(comment).UpdateColumn("read_state", true).Error
}

func SetAllCommentRead() error {
	return DB.Model(&Comment{}).Where("read_state = ?", false).Update("read_state", true).Error
}

func ListUnreadComment() ([]*Comment, error) {
	var comments []*Comment
	err := DB.Where("read_state = ?", false).Order("created_at desc").Find(&comments).Error
	return comments, err
}

func MustListUnreadComment() []*Comment {
	comments, _ := ListUnreadComment()
	return comments
}

func (comment *Comment) Delete() error {
	return DB.Delete(comment, "user_id = ?", comment.UserID).Error
}

func ListCommentByPostID(postId string) ([]*Comment, error) {
	pid, err := strconv.ParseUint(postId, 10, 64)
	if err != nil {
		return nil, err
	}
	var comments []*Comment
	rows, err := DB.Raw("select c.*, nickname, avatar_url from comments c   where c.post_id = ? order by created_at desc", uint(pid)).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment Comment
		DB.ScanRows(rows, &comment)
		comments = append(comments, &comment)
	}
	return comments, err
}

/*func GetComment(id interface{}) (*Comment, error) {
	var comment Comment
	err := DB.First(&comment, id).Error
	return &comment, err
}*/

func CountComment() int64 {
	var count int64
	DB.Model(&Comment{}).Count(&count)
	return count
}
