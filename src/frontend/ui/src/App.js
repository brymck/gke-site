import React, { Component } from 'react';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = { value: '' };

    this.handleChange = this.handleChange.bind(this);
  }

  handleChange(event) {
    const value = event.target.value;
    console.log(value);
    this.setState({ value: event.target.value })
  }


  render() {
    return (
      <div className="App">
        <header className="App-header">
          <p>
            Lorem ipsum.
          </p>
          <input type="text" value={this.state.value} onChange={this.handleChange} />
        </header>
      </div>
    );
  }
}

export default App;
