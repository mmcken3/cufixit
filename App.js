import React, {Component} from 'react';
import { Image, StyleSheet, AppRegistry } from 'react-native';
import { Container, Header, Content, Form, Body, Item, Input, Button, Label, Title, Text } from 'native-base';
import { StackNavigator } from 'react-navigation';


//var ImagePicker = require('react-native-image-picker');

/*var options = {
  title: 'Select Avatar',
  customButtons: [
    {name: 'fb', title: 'Choose Photo from Facebook'},
  ],
  storageOptions: {
    skipBackup: true,
    path: 'images'
  }
};*/

/*ImagePicker.showImagePicker(options, (response) => {
  console.log('Response = ', response);

  if (response.didCancel) {
    console.log('User cancelled image picker');
  }
  else if (response.error) {
    console.log('ImagePicker Error: ', response.error);
  }
  else if (response.customButton) {
    console.log('User tapped custom button: ', response.customButton);
  }
  else {
    let source = { uri: response.uri };

    // You can also display the image using data:
    // let source = { uri: 'data:image/jpeg;base64,' + response.data };

    this.setState({
      avatarSource: source
    });
  }
});*/


class LoginScreen extends React.Component {
  handleSubmit = () => {
    this.props.navigation.navigate('Report');
  }

  render() {
    return (
      <Container style={styles.container}>
        <Header>
          <Body>
          <Image style={styles.image} source={require('./images/CUfixit.png')}/>
          </Body>
        </Header>
        <Content>
          <Form>
            <Item floatingLabel>
              <Label>Username</Label>
              <Input />
            </Item>
            <Item floatingLabel last>
              <Label>Password</Label>
              <Input secureTextEntry={true}/>
            </Item>
            <Button block
              onPress={this.handleSubmit.bind(this)}>
              <Text>Login</Text>
            </Button>
          </Form>

        </Content>
      </Container>

    );
  }
}

class ReportScreen extends React.Component {
  render() {
    return (
      <Container style={styles.container}>
      <Header>
        <Body>
        <Image source={require('./images/CUfixit.png')} /*style={styles.image}*//>
        </Body>
      </Header>
      </Container>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    //alignItems: 'center',
    //justifyContent: 'center',
  },
  image: {
    flex: 1,
    width: null,
    height: null,
    resizeMode: 'stretch',
}
});

const RootStack = StackNavigator(
  {
    Login: {
      screen: LoginScreen,
    },
    Report: {
      screen: ReportScreen,
    },
  },
  {
    initialRouteName: 'Login',
  }
);

class StartUp extends React.Component {
  render() {
    return <RootStack />;
  }
}

AppRegistry.registerComponent("CUFixit", () => StartUp)
