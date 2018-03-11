import React, {Component} from 'react';
import { Image, StyleSheet } from 'react-native';
import { Container, Header, Content, Form, Body, Item, Input, Button, Label, Title, Text } from 'native-base';
import { StackNavigator } from 'react-navigation';


class LoginScreen extends React.Component {
  handleSubmit = () => {
    this.props.navigation.navigate('Report');
  }

  render() {
    return (
      <Container style={styles.container}>
        <Header>
          <Body>
          <Image style={styles.image} source={require('./CUfixit.png')}/>
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
              <Input />
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
        <Image source={require('./CUfixit.png')} style={styles.image}/>
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
  //  width: null,
    //height: null,
    resizeMode: 'contain',
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

export default class App extends React.Component {
  render() {
    return <RootStack />;
  }
}
