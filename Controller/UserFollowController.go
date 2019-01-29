package controller

import (
	"encoding/json"
	// "fmt"
	"github.com/gocraft/dbr"
	"github.com/zenazn/goji/web"
	"net/http"
	"strings"
	//	"strconv"

	"exchange.com/exchange/helper"
	. "exchange.com/exchange/helper"
	. "exchange.com/exchange/model"
	"exchange.com/exchange/system"
)

type UserFollowController struct {
	system.Controller
}

// this struct combines user and user_follow, and also stores profile pic link
type UserFollowData struct {
	Id             int64  `json:"Id"`
	Username       string `json:"Username"`
	FirstName      string `json:"FirstName"`
	LastName       string `json:"LastName"`
	ProfilePicLink string `json:"ProfilePicLink"`
	IsFollowing    bool   `json:"IsFollowing"`
	CanFollow      bool   `json:"CanFollow"`
}

func (controller *UserFollowController) FollowUser_Api(c web.C, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	profileUsername := r.FormValue("username") // current profile's username
	isFollow := r.FormValue("is_follow")
	myAccessToken := r.FormValue("my_token") // my access token

	profileUsername = strings.TrimSpace(profileUsername)
	isFollow = strings.TrimSpace(isFollow)
	myAccessToken = strings.TrimSpace(myAccessToken)

	if profileUsername == "" || isFollow == "" || myAccessToken == "" {
		helper.ReturnErr(w, "Invalid request", ERR_SYSTEM)
		return
	}

	isFollowBool, err := helper.StringToBool(isFollow)
	if err != nil {
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	dbSess := controller.GetDb(c).NewSession(nil)

	// check whether user to follow exists or not
	profileUser, err := UserModel.GetUserByUsernameMin(dbSess, profileUsername)
	if err != nil { // there's an error?
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	// get myself
	meUser, err := UserModel.GetUserByAccessTokenMin(dbSess, myAccessToken)
	if err != nil { // there's an error?
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	// follow the user
	if isFollowBool == true {
		err = UserFollowModel.FollowUser(dbSess, meUser.Id, profileUser.Id)
	} else {
		err = UserFollowModel.UnfollowUser(dbSess, meUser.Id, profileUser.Id)
	}

	if err != nil { // there's an error?
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	data := struct {
		Status int `json:"Status"`
	}{
		Status: 1,
	}

	json.NewEncoder(w).Encode(&data)
}

func (controller *UserFollowController) GetFollowers_Api(c web.C, w http.ResponseWriter, r *http.Request) {
	controller.getUserFollowData(c, w, r, "followers")
}

func (controller *UserFollowController) GetFollowings_Api(c web.C, w http.ResponseWriter, r *http.Request) {
	controller.getUserFollowData(c, w, r, "followings")
}

func (controller *UserFollowController) getUserFollowData(c web.C, w http.ResponseWriter, r *http.Request, getType string) {
	profileUsername := r.URL.Query().Get("username") // current profile's username
	start := r.URL.Query().Get("start")
	count := r.URL.Query().Get("count")
	myAccessToken := r.URL.Query().Get("token") // my access token

	profileUsername = strings.TrimSpace(profileUsername)
	start = strings.TrimSpace(start)
	count = strings.TrimSpace(count)

	if profileUsername == "" || start == "" || count == "" {
		helper.ReturnErr(w, "Invalid request", ERR_SYSTEM)
		return
	}

	startInt, err := helper.StringToInt(start)
	if err != nil { // there's an error?
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	countInt, err := helper.StringToInt(count)
	if err != nil { // there's an error?
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	dbSess := controller.GetDb(c).NewSession(nil)

	// get the user
	profileUser, err := UserModel.GetUserByUsernameMin(dbSess, profileUsername)
	if err != nil { // there's an error?
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	// get the list of followers
	var userList []*User

	if getType == "followers" {
		userList, err = UserFollowModel.GetFollowers(dbSess, profileUser.Id, startInt, countInt)
		if err != nil {
			helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
			return
		}
	} else {
		userList, err = UserFollowModel.GetFollowings(dbSess, profileUser.Id, startInt, countInt)
		if err != nil {
			helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
			return
		}
	}

	// get myself
	var meUser User

	if myAccessToken != "" {
		meUser, err = UserModel.GetUserByAccessTokenMin(dbSess, myAccessToken)
		if err != nil { // there's an error?
			helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
			return
		}
	}

	userIds := []int64{}
	userFollowDatas := []UserFollowData{}

	// check whether i am following these followers or not
	for _, profileFollower := range userList {
		userIds = append(userIds, profileFollower.Id)

		userFollowData := UserFollowData{
			Id:             profileFollower.Id,
			Username:       profileFollower.Username,
			FirstName:      profileFollower.FirstName,
			LastName:       profileFollower.LastName,
			ProfilePicLink: helper.GetUserProfilePicLink(profileFollower.ProfilePic),
			IsFollowing:    false,
			CanFollow:      false,
		}
		userFollowDatas = append(userFollowDatas, userFollowData)
	}

	if meUser.Id > 0 && len(userIds) > 0 {
		isFollowingList, err := UserFollowModel.GetIsFollowingUsers(dbSess, meUser.Id, userIds)
		if err != nil && err != dbr.ErrNotFound { // there's an error?
			helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
			return
		}

		// check whether user are following them
		for index, userFollowData := range userFollowDatas {
			userFollowDatas[index].IsFollowing = isFollowingList[userFollowData.Id]

			// only allow to follow if not ownself
			if userFollowData.Id != meUser.Id {
				userFollowDatas[index].CanFollow = true
			}
		}
	}

	if getType == "followers" {
		data := struct {
			Followers []UserFollowData `json:"Followers"`
		}{
			Followers: userFollowDatas,
		}

		json.NewEncoder(w).Encode(&data)
	} else {
		data := struct {
			Followings []UserFollowData `json:"Followings"`
		}{
			Followings: userFollowDatas,
		}

		json.NewEncoder(w).Encode(&data)
	}

}

/*
func (controller *UserFollowController) GetFollowings_Api(c web.C, w http.ResponseWriter, r *http.Request) {
	profileUsername := r.URL.Query().Get("username") // current profile's username
	start := r.URL.Query().Get("start")
	count := r.URL.Query().Get("count")
	myAccessToken := r.URL.Query().Get("token") // my access token

	profileUsername = strings.TrimSpace(profileUsername)
	start = strings.TrimSpace(start)
	count = strings.TrimSpace(count)

	if profileUsername == "" || start == "" || count == "" {
		helper.ReturnErr(w, "Invalid request", ERR_SYSTEM)
		return
	}

	startInt, err := helper.StringToInt(start)
	if err != nil { // there's an error?
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	countInt, err := helper.StringToInt(count)
	if err != nil { // there's an error?
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	dbSess := controller.GetDb(c).NewSession(nil)

	// get the user
	profileUser, err := UserModel.GetUserByUsernameMin(dbSess, profileUsername)
	if err != nil { // there's an error?
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	// get the list of followers
	profileFollowers, err := UserFollowModel.GetFollowings(dbSess, profileUser.Id, startInt, countInt)
	if err != nil {
		helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
		return
	}

	// get myself
	var meUser User

	if myAccessToken != "" {
		meUser, err = UserModel.GetUserByAccessTokenMin(dbSess, myAccessToken)
		if err != nil { // there's an error?
			helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
			return
		}
	}

	profileFollowerIds := []int64{}
	userFollowDatas := []UserFollowData{}

	// check whether i am following these followers or not
	for _, profileFollower := range profileFollowers {
		profileFollowerIds = append(profileFollowerIds, profileFollower.Id)

		userFollowData := UserFollowData{
			Id:             profileFollower.Id,
			Username:       profileFollower.Username,
			FirstName:      profileFollower.FirstName,
			LastName:       profileFollower.LastName,
			ProfilePicLink: helper.GetUserProfilePicLink(profileFollower.ProfilePic),
			IsFollowing:    false,
			CanFollow:      false,
		}
		userFollowDatas = append(userFollowDatas, userFollowData)
	}

	if meUser.Id > 0 && len(profileFollowerIds) > 0 {
		isFollowingList, err := UserFollowModel.GetIsFollowingUsers(dbSess, meUser.Id, profileFollowerIds)
		if err != nil && err != dbr.ErrNotFound { // there's an error?
			helper.ReturnErr(w, err.Error(), ERR_SYSTEM)
			return
		}

		// check whether user are following them
		for index, userFollowData := range userFollowDatas {
			userFollowDatas[index].IsFollowing = isFollowingList[userFollowData.Id]

			// only allow to follow if not ownself
			if userFollowData.Id != meUser.Id {
				userFollowDatas[index].CanFollow = true
			}
		}
	}

	data := struct {
		Followings []UserFollowData `json:"Followings"`
	}{
		Followings: userFollowDatas,
	}

	json.NewEncoder(w).Encode(&data)
}
*/
