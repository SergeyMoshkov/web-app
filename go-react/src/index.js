import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import AppFooter from './AppFooter';

class App extends Component {

  render() {
    return (
      <div className='app'>
        <div>
          <h1>Я скоро выучу golang и javascript с react!!!</h1>
        </div>
        <AppFooter />
      </div>
    );
  }
}


ReactDOM.render(<App />, document.getElementById('root'));
