import React, { Component } from 'react';

class Article extends Component {

    state = {
        article: undefined,
    }

    componentDidMount() {
        const { id } = this.props.match.params;
        const url = `http://localhost:8090/articles/${id}`
        fetch(url).then(response => {
            response.json().then(article => {
                this.setState({ article })
            }).catch(console.error)
        }).catch(console.error);
    }

    render() {
        if (!this.state.article) {
            return <div />;
        }
        const {
            Title: title,
            Text: text,
            Editor: editor,
        } = this.state.article;
        return (
            <div style={{
                maxWidth: "767px", textJustify: "inter-word", textAlign: "justify",
                marginLeft: "auto", marginRight: "auto"
            }}>
                <strong>
                    <a className="heading" href="/">Hürriyet Lite</a>
                </strong>
                <hr />
                <h2>{title}</h2>
                <span dangerouslySetInnerHTML={{ __html: text }} />
                <i><h5>{editor}</h5></i>
            </div>
        );
    }

}

export default Article;