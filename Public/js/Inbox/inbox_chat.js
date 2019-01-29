app.controller('InboxChatCtrl', ['$scope', '$http', '$location', "listing_id", function($scope, $http, $location, listing_id) {
	var conn = null;
    var chatMessages = null;

    $scope.messages = [];
    $scope.sendChat = function() {
        if (!conn) {
            return false;
        }

        var inputMessage = $scope.inboxChatFormMessage.trim();
        if (inputMessage == "") {
            return false;
        }

        conn.send(inputMessage);
        $scope.inboxChatFormMessage = "";
    };

    function scrollBot() {
        var parentDiv = $("#inbox-chat-messages");
        var lastItem = $("#inbox-chat-messages .media:last-child")
        parentDiv.scrollTop(parentDiv.scrollTop() + lastItem.position().top);
    }

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + $location.host() + ":" + $location.port() + "/ws?listing_id=" + listing_id);
        conn.onclose = function(evt) {
            console.log("left chat room");
        }
        conn.onmessage = function(evt) {
            var message = {
                msg: evt.data
            }
            $scope.messages.push(message);

           $scope.$apply();

           scrollBot();
        }
    } /*else {
        appendLog($("<div><b>Your browser does not support WebSockets.</b></div>"))
    }*/
}]);