let greetings: string = "Hello World";

greetings.toLowerCase();
console.log(greetings);

// number - not numbers!

// not a good practice, should use type inference
// let userId: number = 334455.3; 

let userId = 334455.3; // TypeScript automatically detects that this is a 'number'
userId.toFixed()

// boolean
let isLoggedIn: boolean = false;

// any
//let hero; // currently unsure whether 'hero' is a string, number, etc
let hero: string;

function getHero() {
  return "thor";
}

hero = getHero(); // type is 'any'

// Note: 'any' is not a special type that is being assigned - not a string, boolean, etc
// it is simple a marker to disable typechecking
// Can use 'noImplicitAny' in the TSC config file to flag any implicit 'any' as an error

export {};