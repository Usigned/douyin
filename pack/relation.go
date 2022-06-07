package pack

import "douyin/entity"

func RelationsPtrs(relationsPtrs []*entity.User) []entity.User {
	if relationsPtrs != nil {
		var relation = make([]entity.User, len(relationsPtrs))
		for i, ptr := range relationsPtrs {
			relation[i] = *ptr
		}
		return relation
	}
	return []entity.User{}
}
