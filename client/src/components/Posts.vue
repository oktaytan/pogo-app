<template>
  <div
    class="posts_container"
    :style="{ height: GET_TOP_BAR_SHOW ? '78vh' : '100vh' }"
  >
    <Loader v-if="loading" />
    <a-list v-else itemLayout="vertical" size="large" :dataSource="posts">
      <a-list-item
        slot="renderItem"
        slot-scope="post"
        ref="postItem"
        key="post.id"
        class="post_item"
        @click="() => detailShow(post.id, post.comments)"
      >
        <a-list-item-meta>
          <div class="post_user_wrap" slot="title">
            <div>
              <span class="post_username">{{ post.author.username }}</span>
              <span class="post_date">{{
                $moment(post.created_at).fromNow()
              }}</span>
            </div>
            <span class="post_title">{{ post.title }}</span>
          </div>
          <a-avatar slot="avatar" class="post_avatar">{{
            post.author.username | userAvatar
          }}</a-avatar>
        </a-list-item-meta>
        {{ post.body.slice(0, 200) }}...
        <a-checkable-tag>devamı</a-checkable-tag>
        <post-actions position="top" :post="post" />
      </a-list-item>
    </a-list>
  </div>
</template>

<script>
import Loader from "./Loader.vue";
import PostActions from "./PostActions";

import { mapGetters, mapActions } from "vuex";

export default {
  name: "Posts",
  components: {
    Loader,
    "post-actions": PostActions
  },
  data() {
    return {
      posts: [],
      loading: true
    };
  },
  mounted() {
    this.loading = true;
    this.FETCH_ALL_POSTS().then(data => {
      this.posts = data;
      this.loading = this.posts && false;
    });
  },
  computed: {
    ...mapGetters(["GET_TOP_BAR_SHOW", "GET_ALL_POSTS", "GET_USER"])
  },
  methods: {
    ...mapActions(["FETCH_ALL_POSTS"]),
    detailShow(id, comments) {
      this.$router.push({ name: "Detaylar", params: { id } });
    }
  }
};
</script>
