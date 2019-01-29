app.controller('UserCurrencyCtrl', ['$scope', '$http', '$modal', function($scope, $http, $modal) {
	$scope.get_list = function(idx){
		
		$http.get("/api/usercurrency_list").
			success(function(response) {
				$scope.usercurrency_list = {};
				for (var i=0; i<response.length; i++) {
					$scope.usercurrency_list[response[i].Id] = response[i];
				}
			}).
			error(function(data, status, headers, config) {
				// called asynchronously if an error occurs
				// or server returns response with an error status.
			});
	};

	$scope.get_list();
	
	$scope.delete_usercurrency = function($id) {
		$http.post('/api/usercurrency_delete', $.param({
			id: $id,
		})).
		success(function(data, status, headers, config) {
			delete $scope.usercurrency_list[$id];
		}).
			error(function(data, status, headers, config) {
		});
	};
	
	$scope.show_usercurrency_form = function($id) {
//		var usercurrency = {
//				Id: -1,
//    			UserId: 1,	// hardcoded value, should take from session instead
//		};
//		
//		
//		if ($scope.usercurrency_list[$id] != null) {
//			usercurrency.Id = $scope.usercurrency_list[$id].Id;
//			usercurrency.UserId = $scope.usercurrency_list[$id].UserId;
//			usercurrency.BuyAmount = $scope.usercurrency_list[$id].BuyAmount;
//			usercurrency.BuyCurrency = $scope.usercurrency_list[$id].BuyCurrency;
//			usercurrency.SellAmount = $scope.usercurrency_list[$id].SellAmount;
//			usercurrency.SellCurrency = $scope.usercurrency_list[$id].SellCurrency;
//		}
//		
//        $modal.open({
//            templateUrl: 'user_currency_form.html',
//            controller: "UserCurrencyFormController",
//            inputs: {
//            	usercurrency: usercurrency
//            }
//        }).then(function(modal) {
//            modal.element.modal();
//            modal.close.then(function(modal_result) {
//            	if (modal_result != null) {
//            		$http.post('/api/usercurrency_update', $.param({
//            			id: modal_result.Id,
//	        			user_id: modal_result.UserId,
//	        			buy_amount: modal_result.BuyAmount,
//	        			buy_currency: modal_result.BuyCurrency,
//	        			sell_amount: modal_result.SellAmount,
//	        			sell_currency: modal_result.SellCurrency
//            		})).
//            		success(function(data, status, headers, config) {
//            			modal_result.Id = data; // data is the Id
//            			$scope.usercurrency_list[data] = modal_result;
//            		}).
//            			error(function(data, status, headers, config) {
//            		});
//            	}
//            });
//        });
    };
}]);

//app.controller('UserCurrencyFormController', function($scope, usercurrency, close) {
//	$scope.usercurrency = usercurrency;
//	$scope.close = function(result) {
//		 close(result, 500); // close, but give 500ms for bootstrap to animate
//	};
//});

