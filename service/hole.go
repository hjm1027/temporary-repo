package service

import (
	"github.com/lexkong/log"

	"github.com/mental-health/model"
)

func GetCollectionList(userId, limit, page uint32) ([]*model.HoleInfoResponse, error) {
	var response []*model.HoleInfoResponse

	records, err := model.GetHoleCollectionsByUserId(userId, limit, page)
	if err != nil {
		log.Error("GetHoleCollectionsByUserId get records error", err)
		return nil, err
	}

	var holeIds []uint32
	for _, record := range *records {
		holeIds = append(holeIds, record.HoleId)
	}
	//fmt.Println(holeIds)

	for i := 0; i < len(holeIds); i++ {
		//fmt.Println(holeIds[i])
		//get hole
		hole := model.HoleModel{Id: holeIds[i]}
		//fmt.Println(hole)
		if err = hole.GetById(); err != nil {
			log.Error("GetCollectionsByUserId getbyid error", err)
			return nil, err
		}

		//get user
		user, err := model.GetUserInfoById(userId)
		if err != nil {
			log.Error("GetCollectionsByUserId get user error", err)
		}
		userInfo := model.UserHoleResponse{
			Username: user.Username,
			Avatar:   user.Avatar,
		}

		//get like state
		_, isLike := hole.HasLiked(userId)
		_, isFavorite := hole.HasFavorited(userId)

		data := &model.HoleInfoResponse{
			HoleId:      hole.Id,
			Type:        hole.Type,
			Content:     hole.Content,
			LikeNum:     hole.LikeNum,
			ReadNum:     hole.ReadNum,
			FavoriteNum: hole.FavoriteNum,
			IsLike:      isLike,
			IsFavorite:  isFavorite,
			Time:        hole.Time,
			CommentNum:  hole.CommentNum,
			UserInfo:    userInfo,
		}

		response = append(response, data)
	}
	return response, nil
}