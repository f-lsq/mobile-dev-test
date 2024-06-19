# TypeScript
This repository was created to learn the basics of TypeScript (11 June 2024).
* [Official Docs](https://www.typescriptlang.org/docs/)
* [Personal Notes](https://docs.google.com/document/d/1m89de6EsezM4EBaxZ4uoJIVP7RbKS9jedBJFgSpIWcw/edit#heading=h.q8feypmtuqcr)

## Installation
### Core System Wide Installation
To install TypeScript globally, allowing you to use the `tsc` command anywhere in your terminal, which will execute a TypeScript file.
```
npm install -g typescript
```

To execute a TypeScript file, run 
```
tsc <file name>.ts 
```
Errors in a TypeScript file, does not stop it from transpiling into a JavaScript file.

### Project Config File Installation (For Angular and React)
Since TypeScript is a development tool, it will be installed as a dev dependency.
```
npm install typescript --save-dev
```