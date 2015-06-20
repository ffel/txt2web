var myApp = angular.module('myApp', ['ngRoute']);

myApp.config(function ($routeProvider) {
    
    $routeProvider
    
    .when('/', {
        templateUrl: 'pages/een-lorem-ipsum-dolor-sit-amet.html',
        controller: 'mctrl'
    })
    
    .when('/een', {
        templateUrl: 'pages/een-lorem-ipsum-dolor-sit-amet.html',
        controller: 'mctrl'
    })
    
    .when('/twee', {
        templateUrl: 'pages/twee-morbi-finibus-rutrum-condimentum..html',
        controller: 'mctrl'
    })

    .when('/drie', {
        templateUrl: 'pages/drie-pellentesque-lobortis-lacus.html',
        controller: 'mctrl'
    })

    .when('/vier', {
        templateUrl: 'pages/dira/vier-nulla-euismod-placerat-nunc-at-mattis.html',
        controller: 'mctrl'
    })

    .when('/vijf', {
        templateUrl: 'pages/dira/vijf-donec-lacus-leo.html',
        controller: 'mctrl'
    })

    .when('/zes', {
        templateUrl: 'pages/dira/zes-fusce-non-aliquet-tortor..html',
        controller: 'mctrl'
    })
    
    .when('/zeven', {
        templateUrl: 'pages/dira/zeven-nulla-ut-faucibus-felis.html',
        controller: 'mctrl'
    })
    
    .when('/acht', {
        templateUrl: 'pages/dirb/acht-pellentesque-lacinia.html',
        controller: 'mctrl'
    })
    
    .when('/negen', {
        templateUrl: 'pages/dirb/negen-vivamus-eget-cursus-erat-in-pharetra-neque.html',
        controller: 'mctrl'
    })
    
    .when('/tien', {
        templateUrl: 'pages/dirb/tien-phasellus-lorem-eros.html',
        controller: 'mctrl'
    })
   
});

myApp.controller('mctrl', ['$scope', '$log', function($scope, $log) {
    
    $scope.name = 'Main';
    
}]);
