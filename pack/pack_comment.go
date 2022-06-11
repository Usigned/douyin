package pack

// tools to pack models to entities
// TODO

import (
	"douyin/dao/mysql"
	"douyin/entity"
)

func UserNames(commentModels []*mysql.Comment) []string {
	if commentModels != nil {
		var ids = make([]string, 0, len(commentModels))
		for _, commentModel := range commentModels {
			ids = append(ids, commentModel.UserName)
		}
		return ids
	}
	return []string{}
}

func Comment(commentModel *mysql.Comment) *entity.Comment {
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

func Comments(commentModels []*mysql.Comment) []*entity.Comment {
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
