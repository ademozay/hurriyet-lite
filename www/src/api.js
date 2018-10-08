
const URL = "http://localhost:8000";

export const fetchArticles = async () => {
    return new Promise((resolve, reject) => {
        fetch(`${URL}/articles`).
            then(r => r.json().then(resolve)).
            catch(reject)
    });
}

export const fetchArticle = async (id) => {
    return new Promise((resolve, reject) => {
        fetch(`${URL}/article?id=${id}`).
            then(r => r.json().then(resolve)).
            catch(reject)
    });
}