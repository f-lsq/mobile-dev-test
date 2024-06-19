# Expo
This repository was created to learn the basics of Expo (19 June 2024).
* [Official Docs](https://docs.expo.dev/)

## Creating a New Project
### 1. Intialise a new Expo app
```
npx create-expo-app <Directory Name>
```

#### Install the minimum required libraries for a new project
```
npx create-expo-app <Directory Name> --template blank
```

#### 2. Move to directory of project
```
cd <Directory Name>
```

### 3. Install dependencies to run the project on the web
```
npx expo install react-dom react-native-web @expo/metro-runtime
```

### 4. Run the app on mobile and web
```
npx expo start
```
Two ways to launch the app
* [Expo Go](https://expo.dev/go) for a physical device
* `w` in the terminal for web app