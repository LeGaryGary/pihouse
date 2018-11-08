import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import ReactEcharts, { ObjectMap } from 'echarts-for-react';

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          {/* <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.tsx</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a> */}
          <ReactEcharts
            option={this.getOption()}
            notMerge={true}
            style={
              {
                width:'200px',
                height:'200px'
              }          
            }
          />
        </header>
      </div>
    );
  }

  private getOption = (): ObjectMap => {
    return {
      xAxis: {
          type: 'time'
      },
      yAxis: {
          type: 'value'
      },
      series: [{
          data: [820, 932, 901, 934, 1290, 1330, 1320],
          type: 'line'
      }]
    }
  };

  private getDate = async (): Promise<ObjectMap> => {
    return await fetch("http://stupifybot.com:1337/v1/api/temperature").then
  }
}

export default App;
