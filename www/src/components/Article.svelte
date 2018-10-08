<div>
	{ #if article != null }
	<h2>{ article.Title }</h2>
	{@html article.Text}
	<i>
		<h5>{ article.Editor }</h5>
	</i>
	{ /if }
</div>

<script>
  import { fetchArticle } from "../api";

  export default {
    oncreate() {
      const {
        path: { id }
      } = this.get();
      fetchArticle(id)
        .then(article => this.set({ article }))
        .catch(console.error);
    }
  };
</script>