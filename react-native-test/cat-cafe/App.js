import React from 'react';
import { Text, ScrollView, View, Image, TextInput} from 'react-native';
import Cat from './components/Cat';
import PizzaTranslator from './components/PizzaTranslator';
import FlatListBasics from './components/FlatListBasics';
import SectionListBasics from './components/SectionListBasics';
import Movies from './components/Movies';

export default function App() {
  return (
    // <ScrollView>
    //   <Text>Some text</Text>
    //   <View>
    //     <Text>Some more text</Text>
    //     <Image source={{
    //       uri: 'https://reactnative.dev/docs/assets/p_cat1.png'
    //     }} style={{
    //       width: 200, 
    //       height: 200
    //       }}></Image>
    //   </View>
    //   <TextInput style={{
    //     height: 40,
    //     borderColor: 'gray',
    //     borderWidth: 1
    //   }} defaultValue='You can type in me'/>
    //   <Cat name="Maru"/>
    //   <Cat name="Jellylorum"/>
    //   <Cat name="Spot"/>
    //   <PizzaTranslator/>
    // </ScrollView>

    // <FlatListBasics/>

    // <SectionListBasics/>

    <Movies/>
  );
}