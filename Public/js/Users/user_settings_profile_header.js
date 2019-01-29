app.controller('UserSettingsProfileHeaderCtrl', ['$scope', '$http', 'Upload', 'user', 'user_profile_pic_link', function($scope, $http, Upload, user, user_profile_pic_link) {
	$scope.user = user;
	$scope.user_profile_pic_link = user_profile_pic_link;
	
	$scope.uploadPic = function(file) {
		console.log(file);
		console.log($scope.uploadPic);
		$scope.f = file;
		if (file) {
			if (!file.$error) {
				file.upload = Upload.upload({
					url: '/api/upload-profile-pic',
					method : 'POST',
					fields : {
						token : $scope.user.Token
					},
					file : file,
					fileFormDataName : 'profile_pic'
				}).progress(function(evt) {
					file.progress = Math.min(100, parseInt(100.0 * evt.loaded / evt.total));
				}).success(function(data, status, headers, config) {
					if (data.error != null) {
						$scope.errorMsg = data.error.message;
					} else {
						$scope.errorMsg = null;
						
						var random = new Date().getTime();
						$scope.user_profile_pic_link = data.userProfilePicLink + "?rand=" + random;;
						file.result = data;
					}
				}).error(function(data, status, headers, config) {
				})
			} else {
				switch (file.$error) {
				case "maxSize":
					$scope.errorMsg = "File size cannot be bigger than 5MB";
					break;
				}
			}
		}
	}
}]);