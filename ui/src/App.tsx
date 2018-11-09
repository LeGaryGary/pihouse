import React, { Component } from 'react';
import './App.css';
import ReactEcharts, { ObjectMap } from 'echarts-for-react';

interface IProps{

}

interface IState{
  data: []
}

class App extends Component<IProps, IState> {
  constructor(props: any){
    super(props);
    this.state = {data: []}
  }

  async componentDidMount() {
    const data = (await fetch("http://stupifybot.com:1337/v1/api/temperature")
    .then(v => {console.log(v); return v.json()})
    .then(json => json.map((reading: TemperatureReading) => [reading.CreatedAt, reading.Value]))
    );
    this.setState({data: data});
  }

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
                width:window.innerWidth,
                height:window.innerHeight
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
          data: this.state.data,
          type: 'scatter'
      }]
    }
  };
}

interface TemperatureReading{
  CreatedAt: string,
  Value: string
}

export default App;
