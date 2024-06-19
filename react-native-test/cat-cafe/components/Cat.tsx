import React, { useState } from 'react';
import {Button, Text, View} from 'react-native';

type CatProps = {
  name: string;
} 

const Cat = (props: CatProps) => {
  const [isHungry, setIsHungry] = useState(true);

  return (
    <View>
      <Text>
        I am a {props.name}, and I am {isHungry ? 'hungry' : 'full'}!
      </Text>
      <Button
        onPress={() => {
          setIsHungry(false);
        }}
        disabled={!isHungry}
        title={isHungry ? 'Give me some food' : 'Thank you!'}
      /> 
    </View>
    
  );
};

export default Cat;