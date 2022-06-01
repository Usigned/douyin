package pack

// tools to pack models to entities
// TODO

import (
	"github.com/Usigned/douyin/dao"
	"github.com/Usigned/douyin/entity"
)

func Comment(commentModel *dao.Comment) *entity.Comment {
	if commentModel != nil {
		return &entity.Comment{
			Id: commentModel.Id,
			User: entity.User{
				Id:   commentModel.Id,
				Name: commentModel.UserName,
			},
			Content:    commentModel.Content,
			CreateDate: commentModel.CreateAt,
		}
	}

	return nil
}

func Comments(commentModels []*dao.Comment) []*entity.Comment {
	if commentModels != nil {
		var comments = make([]*entity.Comment, 0, len(commentModels))
		for _, model := range commentModels {
			comments = append(comments, Comment(model))
		}
		return comments
	}
	return nil
}

func CommentsPtrs(commentPtrs []*entity.Comment) []entity.Comment {
	if commentPtrs != nil {
		var videos = make([]entity.Comment, len(commentPtrs))
		for i, ptr := range commentPtrs {
			videos[i] = *ptr
		}
		return videos
	}
	return []entity.Comment{}
}