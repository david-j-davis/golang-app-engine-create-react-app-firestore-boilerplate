import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  handleSubmit(formData) {
    console.log(formData)
    // fetch('https://gae.appspot.com/api/user', {
    //   method: 'post',
    //   body: JSON.stringify(formData)
    // })
    // .catch(err => console.log(err))
    // .then(res => res.json())
    // .then(data => {
    //   console.log(data)
    // })
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
        </header>
        <form onSubmit={this.handleSubmit}>
          <input type="text" name="name" required placeholder="Name"></input>
          <input type="email" name="email" required placeholder="Email"></input>
          <button type="Submit">Submit</button>
        </form>
      </div>
    );
  }
}

export default App;
