package mysql_logic

import (
	"context"
	"errors"
	"phase1/phase1_work/mysql"
	"phase1/phase1_work/mysql/mysql_model"
	"time"
)

func CreatePost(ctx context.Context, post *mysql_model.Posts) (int64, error) {
	if err := mysql.DB.Create(post).Error; err != nil {
		return 0, err
	}

	if post.Id <= 0 {
		return 0, errors.New("create post failed, invalid post id")
	}

	return post.Id, nil
}

func DeletePost(ctx context.Context, postId int64, userId int64) error {
	result := mysql.DB.Model(&mysql_model.Posts{}).Where("id = ? AND user_id = ? AND deleted_at = 0", postId, userId).Update("deleted_at", time.Now().UnixMilli())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("post not found or permission denied")
	}
	return nil
}

func UpdatePost(ctx context.Context, postId int64, userId int64, title, content string) error {
	result := mysql.DB.Model(&mysql_model.Posts{}).Where("id = ? AND user_id = ? AND deleted_at = 0", postId, userId).Updates(map[string]interface{}{
		"title":      title,
		"content":    content,
		"updated_at": time.Now().UnixMilli(),
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("post not found or permission denied")
	}
	return nil
}

func GetPostDetail(ctx context.Context, postId int64) (*mysql_model.Posts, error) {
	var post mysql_model.Posts
	if err := mysql.DB.Where("id = ? AND deleted_at = 0", postId).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func ListPost(ctx context.Context, pageNo, pageSize int32, startTime, endTime int64, postId string) ([]mysql_model.Posts, int64, error) {
	var posts []mysql_model.Posts
	var total int64

	query := mysql.DB.Model(&mysql_model.Posts{}).Where("deleted_at = 0")

	if postId != "" {
		query = query.Where("id = ?", postId)
	}
	if startTime > 0 {
		query = query.Where("updated_at >= ?", startTime)
	}
	if endTime > 0 {
		query = query.Where("updated_at <= ?", endTime)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (pageNo - 1) * pageSize
	if err := query.Order("updated_at DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}
