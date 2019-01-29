package model

import (
	"github.com/gocraft/dbr"
)

var UserFollowModel UserFollow

type UserFollow struct {
	Id             int64 `json:"Id"`
	UserId         int64 `json:"UserId"`
	FollowerUserId int64 `json:"FollowerUserId"`
}

// func (UserFollowModel UserFollow) GetAll(dbSess *dbr.Session, user_id int64) ([]*UserFollow, error) {
// 	userFollowList := []*UserFollow{}
// 	_, err := dbSess.SelectBySql("SELECT * FROM user_follows WHERE user_id = ?;", user_id).LoadStructs(&userFollowList)

// 	return userFollowList, err
// }

func (UserFollowModel UserFollow) GetFollowers(dbSess *dbr.Session, userId int64, start int64, count int64) ([]*User, error) {
	userList := []*User{}
	_, err := dbSess.SelectBySql("SELECT id, username, first_name, last_name, profile_pic FROM users WHERE id IN "+
		"(SELECT follower_user_id FROM user_follows WHERE user_id=?) LIMIT ?, ?;", userId, start, count).LoadStructs(&userList)

	return userList, err
}

func (UserFollowModel UserFollow) GetFollowings(dbSess *dbr.Session, userId int64, start int64, count int64) ([]*User, error) {
	userList := []*User{}
	_, err := dbSess.SelectBySql("SELECT id, username, first_name, last_name, profile_pic FROM users WHERE id IN "+
		"(SELECT user_id FROM user_follows WHERE follower_user_id=?) LIMIT ?, ?;", userId, start, count).LoadStructs(&userList)

	return userList, err
}

func (UserFollowModel UserFollow) GetFollowerCount(dbSess *dbr.Session, userId int64) (int64, error) {
	followerCount, err := dbSess.SelectBySql("SELECT COUNT(id) FROM user_follows WHERE user_id=?;", userId).ReturnInt64()

	return followerCount, err
}

func (UserFollowModel UserFollow) GetFollowingCount(dbSess *dbr.Session, userId int64) (int64, error) {
	followingCount, err := dbSess.SelectBySql("SELECT COUNT(id) FROM user_follows WHERE follower_user_id=?;", userId).ReturnInt64()

	return followingCount, err
}

func (UserFollowModel UserFollow) GetIsFollowingUser(dbSess *dbr.Session, followerUserId int64, userId int64) (bool, error) {
	userFollowId, err := dbSess.SelectBySql("SELECT id FROM user_follows WHERE user_id = ? AND follower_user_id = ? LIMIT 1;", userId, followerUserId).ReturnInt64()

	var isFollowing bool
	if userFollowId > 0 {
		isFollowing = true
	} else {
		isFollowing = false
	}
	return isFollowing, err
}

func (UserFollowModel UserFollow) GetIsFollowingUsers(dbSess *dbr.Session, followerUserId int64, userIds []int64) (map[int64]bool, error) {
	userFollowList := []*UserFollow{}
	_, err := dbSess.SelectBySql("SELECT user_id FROM user_follows WHERE user_id IN ? AND follower_user_id = ?;", userIds, followerUserId).LoadStructs(&userFollowList)

	// make a list of isFollowing or not
	var isFollowingList map[int64]bool = make(map[int64]bool)
	for _, userId := range userIds {
		isFollowing := false
		for _, userFollow := range userFollowList {
			if userId == userFollow.UserId {
				isFollowing = true
				break
			}
		}

		isFollowingList[userId] = isFollowing
	}

	return isFollowingList, err
}

func (UserFollowModel UserFollow) FollowUser(dbSess *dbr.Session, followerUserId int64, userId int64) error {
	_, err := dbSess.UpdateBySql("INSERT INTO user_follows (user_id, follower_user_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE id=id;", userId, followerUserId).Exec()
	return err
}

func (UserFollowModel UserFollow) UnfollowUser(dbSess *dbr.Session, followerUserId int64, userId int64) error {
	_, err := dbSess.UpdateBySql("DELETE FROM user_follows WHERE user_id=? AND follower_user_id=?;", userId, followerUserId).Exec()
	return err
}
