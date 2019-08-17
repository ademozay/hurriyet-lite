import React, { Component } from 'react';
import { Link } from 'react-router-dom';

import './App.css'

class App extends Component {

  state = {
    news: []
  }

  componentDidMount() {
    fetch("http://localhost:8090/articles").then((response) => {
      response.json().then((news) => {
        this.setState({ news })
      })
    }).catch((e) => {
      console.error(e)
    })
  }

  render() {
    return (
      <div>
        <strong>
          <a className="heading" href="/">Hürriyet Lite</a>
        </strong>
        <hr />
        <ul>
          {this.renderNews()}
        </ul>
      </div>
    )
  }

  renderNews() {
    return this.state.news.map((news) => {
      return (
        <li key={news.Id}>
          <Link to={news.Id}>
            {news.Title}
          </Link>
        </li>
      )
    })
  }

}

export default App;
