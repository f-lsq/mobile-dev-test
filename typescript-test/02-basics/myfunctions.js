"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
function addTwo(num) {
    // function addTwo(num) {
    // }
    // num is of type 'any', meaning that there will be no type checking
    // num.toUpperCase(); -> this function will be allowed even though it shouldnt be
    return num + 2;
}
function getUpper(val) {
    return val.toUpperCase();
}
function signUpUser(name, email, isPaid) {
}
var loginUser = function (name, email, isPaid) {
    // isPaid will have a default value of 'false'
    if (isPaid === void 0) { isPaid = false; }
};
addTwo(5);
getUpper("hello");
signUpUser("Hello", "hello123@world.com", false);
loginUser("h", "h@h.com");
