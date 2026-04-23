package mysql_logic

import (
	"context"
	"errors"
	"phase1/phase1_work/mysql"
	"phase1/phase1_work/mysql/mysql_model"
)

func CreateComment(ctx context.Context, comment *mysql_model.Comments) (int64, error) {
	if err := mysql.DB.Create(comment).Error; err != nil {
		return 0, err
	}

	if comment.Id <= 0 {
		return 0, errors.New("create comment failed, invalid comment id")
	}

	return comment.Id, nil
}

func ListCommentByPostId(ctx context.Context, postId int64, pageNo, pageSize int32) ([]mysql_model.Comments, int64, error) {
	var comments []mysql_model.Comments
	var total int64

	query := mysql.DB.Model(&mysql_model.Comments{}).Where("post_id = ? AND deleted_at = 0", postId)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询，按创建时间倒序
	offset := (pageNo - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&comments).Error; err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}
