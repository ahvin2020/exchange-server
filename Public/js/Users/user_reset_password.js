app.controller('UserResetPasswordCtrl', ['$scope', '$http', 'token', function($scope, $http, token) {
	$scope.token = token;
}]);