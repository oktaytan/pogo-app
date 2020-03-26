<template>
  <div
    class="posts_container"
    v-infinite-scroll="handleInfiniteOnLoad"
    :infinite-scroll-disabled="busy"
    :infinite-scroll-distance="10"
  >
    <Loader v-if="loading" />
    <a-list v-else itemLayout="vertical" size="large" :dataSource="listData">
      <a-list-item
        slot="renderItem"
        slot-scope="item"
        ref="postItem"
        key="item.title"
        class="post_item"
        @click="detailShow"
      >
        <a-list-item-meta>
          <div class="post_user_wrap" slot="title">
            <div>
              <span class="post_username">cawis</span>
              <span class="post_date">34 dk önce</span>
            </div>
            <span class="post_title">{{ item.title }}</span>
          </div>
          <a-avatar slot="avatar" class="post_avatar">CW</a-avatar>
        </a-list-item-meta>
        {{ item.content }}
        <post-actions position="top" />
        <!-- Gönderi detayları -->
        <post-detail ref="postDetail" :postId="item.id" />
      </a-list-item>
      <div v-if="loading && !busy" class="posts_loading">
        <a-spin />
      </div>
    </a-list>
  </div>
</template>

<script>
import Loader from "./Loader.vue";
import PostDetail from "./PostDetail";
import PostActions from "./PostActions";

const listData = [];
for (let i = 0; i < 23; i++) {
  listData.push({
    id: i,
    to: "/posts/",
    title: `Post Title ${i}`,
    avatar: "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png",
    description:
      "Ant Design, a design language for background applications, is refined by Ant UED Team.",
    content:
      "We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently."
  });
}

export default {
  name: "Posts",
  components: {
    Loader,
    "post-detail": PostDetail,
    "post-actions": PostActions
  },
  data() {
    return {
      loading: true,
      listData,
      busy: false,
      loading: true
    };
  },
  mounted() {
    setTimeout(() => (this.loading = false), 300);
  },
  methods: {
    handleInfiniteOnLoad() {
      // console.log("test");
    },
    detailShow(e) {
      this.$refs.postDetail.showModal = true;
    }
  }
};
</script>
