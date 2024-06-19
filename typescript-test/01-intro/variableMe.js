"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var greetings = "Hello World";
greetings.toLowerCase();
console.log(greetings);
// number - not numbers!
// not a good practice, should use type inference
// let userId: number = 334455.3; 
var userId = 334455.3; // TypeScript automatically detects that this is a 'number'
userId.toFixed();
// boolean
var isLoggedIn = false;
