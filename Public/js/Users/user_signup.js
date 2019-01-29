app.controller('UserSignupCtrl', ['$scope', '$http', 'country_list', 'user', function($scope, $http, country_list, user) {
	$scope.country_list = country_list;
	$scope.user = user;
}]);