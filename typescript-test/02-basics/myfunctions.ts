function addTwo(num: number): number {
  // function addTwo(num) {
  // }
  // num is of type 'any', meaning that there will be no type checking
  // num.toUpperCase(); -> this function will be allowed even though it shouldnt be
  
  return num + 2;
  // return "hello"
}

function getUpper(val: string) {
  return val.toUpperCase();
}

function signUpUser(name: string, email: string , isPaid: boolean) {

}

let loginUser = (name: string, email: string, isPaid: boolean = false) => {
  // isPaid will have a default value of 'false'

}

let myValue = addTwo(5);
getUpper("hello");
signUpUser("Hello", "hello123@world.com" , false);
loginUser("h", "h@h.com")

// There are cases where more than one type can be returned
// function getValue(myVal: number) {
//   if (myVal > 5) {
//     return true
//   }
//   return "200 OK"
// }

const getHello = (s: string): string => {
  return ""
}

const heros = ["thor", "spiderman", "ironman"]
// const heros = [1, 2, 3]
heros.map((hero):string => {
  return `hero is ${hero}`
})

function consoleError(errMsg: string): void {
  console.log(errMsg);
}

function handleError(errMsg: string): never {
  throw new Error(errMsg);
}

export {};