/*
 Copyright (c) 2017, CodeHat
 https://codehat.de
 All rights reserved.
 */

$(document).ready(function () {

    // Get current SignColors version and update date
    $.getJSON("https://api.codehat.de/plugin/1", function (data) {
        var updated_at = moment(data.updated_at);
        $('#signcolors').text('v' + data.version);
        $('#signcolors_updatetime').text(updated_at.format('Do MMM YYYY, HH:mm'));
    });
    // Get current PlayerSupport version and update date
    $.getJSON("https://api.codehat.de/plugin/2", function (data) {
        var updated_at = moment(data.updated_at);
        $('#playersupport').text('v' + data.version);
        $('#playersupport_updatetime').text(updated_at.format('Do MMM YYYY, HH:mm'));
    });

});