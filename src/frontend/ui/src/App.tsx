import * as React from 'react';
import axios from 'axios';
import './App.css';

interface Props {
}

interface State {
  value: string
  greeting: string
}

export default class App extends React.Component<Props, State> {
  state: State = { value: '', greeting: '' };

  constructor(props: Props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
  }

  handleChange(event: React.ChangeEvent<HTMLInputElement>) {
    const value = event.target.value;
    this.setState({ value })
    axios.get(`/hello/${value}`)
        .then(res => {
          const greeting = res.data;
          this.setState({ greeting: greeting.message });
        });
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <p>
            Lorem ipsum.
          </p>
          <input type="text" value={this.state.value} onChange={this.handleChange} />
          <span>{this.state.greeting}</span>
        </header>
      </div>
    );
  }
}
