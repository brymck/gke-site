import * as React from 'react';
import './App.css';

interface Props {
}

interface State {
  value: string
}

class App extends React.Component<Props, State> {
  state: State = { value: '' };

  constructor(props: Props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
  }

  handleChange(event: React.ChangeEvent<HTMLInputElement>) {
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