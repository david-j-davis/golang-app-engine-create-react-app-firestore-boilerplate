import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  handleSubmit(e) {
    e.preventDefault()
    
    const email = e.target.elements.email.value
    const name = e.target.elements.name.value
    const formData = {
      "email": email,
      "name": name
    }

    fetch('/api/user', { //In production it would be https://YOUR_PROJECT_ID.appspot.com/api/user
      method: 'post',
      body: JSON.stringify(formData)
    })
    .catch(err => console.log(err))
    .then(res => res.text())
    .then(data => {
      console.log(data)
    })
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
        </header>
        <h2>Example, post form data to Go, and saved to Firestore:</h2>
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
