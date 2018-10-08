<div>
	<Header />
	<div id="router"></div>
</div>

<script>
  import { Store } from "svelte/store";
  import SvelteRouter from "svelte-router";

  import Header from "./Header.svelte";
  import Articles from "./Articles.svelte";
  import Article from "./Article.svelte";

  const router = new SvelteRouter({
    mode: "history",
    routes: {
      "/": {
        Component: Articles,
        data: { articles: [] }
      },
      "/:id": {
        Component: Article,
        data: { article: {} }
      }
    }
  });

  export default {
    oncreate() {
      router.create("#router");
    },
    ondestroy() {
      router.destroy();
    },
    components: {
      Header
    }
  };
</script>