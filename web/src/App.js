import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import './tachyons.min.css';

class App extends Component {
  render() {
    return (
      <div className="tc">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Welcome to Invoicr</h2>
          <Form/>
        </div>
      </div>
    );
  }
}

class Form extends Component {
  constructor(props) {
    super(props);
    this.state = {
      token: 'Your token will appear here'
    };
    this.handleEmailChange = this.handleEmailChange.bind(this);
    this.handlePasswordChange = this.handlePasswordChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  handleSubmit(e) {
    fetch("/invoice/", {
      method: 'POST',
      body: {}
    })
    .then((res) => {
      res.text().then((text) => {
        this.setState({token: text})
      });
    })
    .catch((err) => {
      console.log(err);
    })
    e.preventDefault();
  }
  handlePasswordChange(e) {
    this.setState({password: e.target.value})
  }
  handleEmailChange(e) {
    this.setState({email: e.target.value})
  }
  render() {
    return (
      <main className="pa4 black-80">
        <form className="measure center" onSubmit={this.handleSubmit}>
          <div className="">
            <input className="b ph3 pv2 input-reset ba b--black bg-transparent grow pointer f6 dib" type="submit" value="Sign up!"/>
          </div>
        </form>
        <p>{this.state.token}</p>
      </main>);
    }
}
export default App;
