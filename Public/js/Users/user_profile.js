app.controller('UserProfileCtrl', ['$scope', '$http', 'user', 'user_profile_pic_link', 'user_country', 'can_follow', 'is_following', 'access_token', 'follower_count', 'following_count', function($scope, $http, user, user_profile_pic_link, user_country, can_follow, is_following, access_token, follower_count, following_count) {
	$scope.user = user;
	$scope.user.ProfilePicLink = user_profile_pic_link;
	$scope.user.IsFollowing = is_following;
	$scope.user.Country = user_country;
	$scope.user.CanFollow = can_follow;

	$scope.access_token = access_token;
	
	$scope.follower_count = follower_count;
	$scope.following_count = following_count;

	$scope.isLoading = false;
	$scope.activeTab = null;

	$scope.followers = []; // list of followers
	$scope.followings = []; // list of following users

	$scope.errorAlerts = [];

	var createdDate = new Date(user.Created)
	$scope.user.JoinedDate = monthNames[createdDate.getMonth()] + " " + createdDate.getFullYear();

	$scope.removeAlert = function($index) {
   		$scope.errorAlerts.splice($index, 1);
	}

	$scope.followUser = function($user, $isFollow) {
		$scope.is_following = $isFollow;

		$http.post('/api/follow-user', $.param({
			username: $user.Username,			// user to follow
			is_follow: $isFollow,
			my_token: access_token	// my access token
		})).
		success(function(data, status, headers, config) {
			if (data.Status == 1) {
				$user.IsFollowing = $isFollow;
			} else if (data.Error != null) {
				$scope.errorAlerts.push(data.Error.message);
			}
		}).
			error(function(data, status, headers, config) {
				// TODO: add error handling
		});
	}

	// when reach bottom of page
	$(window).scroll(function(){
		if ($scope.isLoading == false && ($(window).scrollTop() == $(document).height() - $(window).height())) {
			switch ($scope.activeTab) {
				case "#followings-tab":
					if ($scope.followings.length < $scope.following_count) {
						loadFollowings($scope.followings.length);
					}
					break;
				case "#followers-tab":
					if ($scope.followers.length < $scope.follower_count) {
						loadFollowers($scope.followers.length);
					}
					break;
			}
        }
  	}); 

	$("#profile-tab-nav a").on("shown.bs.tab", function(e) {
		 var pattern=/#.+/gi //use regex to get anchor(==selector)
	        var contentID = e.target.toString().match(pattern)[0]; //get anchor         
			$scope.activeTab = contentID;

			switch (contentID) {
				case "#followings-tab":
					$scope.followings.length = 0;
					loadFollowings(0);
					break;
				case "#followers-tab":
					$scope.followers.length = 0;
					loadFollowers(0);
					break;
			};
	       //load content for selected tab
	      // $(contentID).load(baseURL+contentID.replace('#',''), function(){
	       //     $('#myTab').tab(); //reinitialize tabs
	       //});
	});

	// load followings
	function loadFollowings(start) {
		$scope.isLoading = true;

		var username = user.Username;
		var count = 12;

		var url = '/api/get-followings?username=' + username + "&start=" + start + "&count=" + count;
		if ($scope.access_token != "") {
			url += "&token=" + $scope.access_token;
		}

		$http.get(url).success(function(data, status, headers, config) {
			$scope.isLoading = false;

			if (data.Followings) {
				$scope.followings = $scope.followings.concat(data.Followings);
			} else if (data.Error != null) {
				$scope.errorAlerts.push(data.Error.message);
			}
		}).
		error(function(data, status, headers, config) {
			$scope.isLoading = false;
		});
	}

	// load followers
	function loadFollowers(start) {
		$scope.isLoading = true;

		var username = user.Username;
		var count = 12;

		var url = '/api/get-followers?username=' + username + "&start=" + start + "&count=" + count;
		if ($scope.access_token != "") {
			url += "&token=" + $scope.access_token;
		}

		$http.get(url).success(function(data, status, headers, config) {
			$scope.isLoading = false;

			if (data.Followers) {
				$scope.followers = $scope.followers.concat(data.Followers);
			} else if (data.Error != null) {
				$scope.errorAlerts.push(data.Error.message);
			}
		}).
		error(function(data, status, headers, config) {
			$scope.isLoading = false;
		});
	}
}]);