app.controller('UserSettingsProfileCtrl', ['$scope', '$http', 'country_list', 'gender_list', 'user', function($scope, $http, country_list, gender_list, user) {
	$scope.country_list = country_list;
	$scope.gender_list = gender_list;
	$scope.user = user;
	
	// datepicker
	if ($scope.user.Birthday == "0001-01-01T00:00:00Z") {
		$scope.user.Birthday = "";
	}
	$scope.today = new Date();
	$scope.datePickerIsOpen = false;
	$scope.openDatePicker = function($event) {
		$scope.datePickerIsOpen = true;
	};
}]);