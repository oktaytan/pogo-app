<template>
  <div
    class="posts_container"
    :style="{ height: GET_TOP_BAR_SHOW ? '78vh' : '100vh' }"
  >
    <loader-skeleton
      v-if="GET_NEW_POST_LOADING"
      class="animated fadeIn"
      :rows="2"
    />
    <loader v-if="loading" />
    <a-list
      v-else
      itemLayout="vertical"
      size="large"
      :dataSource="GET_ALL_POSTS"
    >
      <a-list-item
        slot="renderItem"
        slot-scope="post"
        ref="postItem"
        key="post.id"
        class="post_item"
        @click="() => detailShow(post.id)"
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
        <post-actions
          position="top"
          :post="post"
          :details="false"
          @like="getData"
          :myPost="post.author.username === GET_USER.username ? true : false"
        />
      </a-list-item>
    </a-list>
  </div>
</template>

<script>
import Loader from "./Loader";
import LoaderSkeleton from "./LoaderSkeleton";
import PostActions from "./PostActions";

import { mapGetters, mapActions } from "vuex";

export default {
  name: "Posts",
  components: {
    loader: Loader,
    "loader-skeleton": LoaderSkeleton,
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
    this.getData();
  },
  computed: {
    ...mapGetters([
      "GET_TOP_BAR_SHOW",
      "GET_ALL_POSTS",
      "GET_USER",
      "GET_NEW_POST_LOADING"
    ])
  },
  methods: {
    ...mapActions(["FETCH_ALL_POSTS"]),
    async getData() {
      this.FETCH_ALL_POSTS()
        .then(data => {
          this.loading = this.GET_ALL_POSTS.length > 0 && false;
        })
        .catch(err => this.$message.error(err.message));
    },
    detailShow(id) {
      this.$router.push({ name: "Detaylar", params: { id } });
    }
  }
};
</script>
