import * as React from 'react';
import axios from 'axios';
import './App.css';

interface Props {
}

interface State {
  name: string
  greeting: string
  number: string
  squared: number
  count: number
}

export default class App extends React.Component<Props, State> {
  state: State = {
    name: '',
    greeting: '',
    number: '0',
    squared: 0,
    count: 0,
  };

  constructor(props: Props) {
    super(props);
    this.handleNameChange = this.handleNameChange.bind(this);
    this.handleNumberChange = this.handleNumberChange.bind(this);
  }

  handleNameChange(event: React.ChangeEvent<HTMLInputElement>) {
    const name = event.target.value;
    this.setState({ name });
    axios.get(`/hello/${name}`)
      .then(res => {
        const greeting = res.data;
        this.setState({ greeting: greeting.message });
      })
      .catch(reason => {
        console.error(reason);
        this.setState({ greeting: '' });
      });
  }

  handleNumberChange(event: React.ChangeEvent<HTMLInputElement>) {
    const number = event.target.value;
    this.setState({ number });
    axios.get(`/square/${number}`)
      .then(res => {
        const response = res.data;
        this.setState({ squared: response.number });
      })
      .catch(reason => {
        console.error(reason);
        this.setState({ squared: 0 })
      });
  }

  handleButtonClick(event: React.MouseEvent<HTMLInputElement>) {
    axios.get('/count/')
      .then(res => {
        this.setState({ count: 1 })
      })
      .catch(reason => {
        console.error(reason);
        this.setState({ count: 0 })
      })
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <p>
            Lorem ipsum.
          </p>
          <input
            autoFocus={true}
            onChange={this.handleNameChange}
            type="text"
            value={this.state.name}
          />
          <span>{this.state.greeting}</span>
          <input
            autoFocus={true}
            onChange={this.handleNumberChange}
            type="text"
            value={this.state.number}
          />
          <span>{this.state.squared}</span>
          <input
            onClick={this.handleButtonClick}
            type="button"
            value="Click me"
          />
          <span>{this.state.count}</span>
        </header>
      </div>
    );
  }
}
