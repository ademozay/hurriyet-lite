<div>
	{ #if articles }
	<ul>
		{ #each articles as article }
		<li style="padding: 3px;">
			<RouterLink to='/{ article.Id }'>
				{ article.Title }
			</RouterLink>
		</li>
		{ /each }
	</ul>
	{ /if }
</div>

<script>
  import SvelteRouter from "svelte-router";
  import { fetchArticles } from "../api";

  export default {
    oncreate() {
      fetchArticles()
        .then(articles => this.set({ articles }))
        .catch(console.error);
    },
    components: {
      RouterLink: SvelteRouter.RouterLink
    }
  };
</script>
