
var myApp = angular.module('myApp', ['ngRoute']);

myApp.config(function ($routeProvider) {
    
    $routeProvider

    .when('/', {
        templateUrl: 'pages/index.html',
        controller: 'mctrl'
    })
    
    
    
	.when('/een-lorem-ipsum-dolor-sit-amet', {
	    templateUrl: 'pages/een-lorem-ipsum-dolor-sit-amet.html',
	    controller: 'mctrl'
	})
    
	.when('/twee-morbi-finibus-rutrum-condimentum.', {
	    templateUrl: 'pages/twee-morbi-finibus-rutrum-condimentum..html',
	    controller: 'mctrl'
	})
    
	.when('/images', {
	    templateUrl: 'pages/images.html',
	    controller: 'mctrl'
	})
    
	.when('/drie-pellentesque-lobortis-lacus', {
	    templateUrl: 'pages/drie-pellentesque-lobortis-lacus.html',
	    controller: 'mctrl'
	})
    
	.when('/dira/vier-nulla-euismod-placerat-nunc-at-mattis', {
	    templateUrl: 'pages/dira/vier-nulla-euismod-placerat-nunc-at-mattis.html',
	    controller: 'mctrl'
	})
    
	.when('/dira/vijf-donec-lacus-leo', {
	    templateUrl: 'pages/dira/vijf-donec-lacus-leo.html',
	    controller: 'mctrl'
	})
    
	.when('/dira/zes-fusce-non-aliquet-tortor.', {
	    templateUrl: 'pages/dira/zes-fusce-non-aliquet-tortor..html',
	    controller: 'mctrl'
	})
    
	.when('/dira/zeven-nulla-ut-faucibus-felis', {
	    templateUrl: 'pages/dira/zeven-nulla-ut-faucibus-felis.html',
	    controller: 'mctrl'
	})
    
	.when('/dirb/acht-pellentesque-lacinia', {
	    templateUrl: 'pages/dirb/acht-pellentesque-lacinia.html',
	    controller: 'mctrl'
	})
    
	.when('/dirb/negen-vivamus-eget-cursus-erat-in-pharetra-neque', {
	    templateUrl: 'pages/dirb/negen-vivamus-eget-cursus-erat-in-pharetra-neque.html',
	    controller: 'mctrl'
	})
    
	.when('/dirb/tien-phasellus-lorem-eros', {
	    templateUrl: 'pages/dirb/tien-phasellus-lorem-eros.html',
	    controller: 'mctrl'
	})
    
	.when('/index', {
	    templateUrl: 'pages/index.html',
	    controller: 'mctrl'
	})
    
	.when('/dirb/index', {
	    templateUrl: 'pages/dirb/index.html',
	    controller: 'mctrl'
	})
    
	.when('/dira/index', {
	    templateUrl: 'pages/dira/index.html',
	    controller: 'mctrl'
	})
    
    
   
});

myApp.controller('mctrl', ['$scope', '$log', function($scope, $log) {
    
    $scope.name = 'Main';
    
}]);
