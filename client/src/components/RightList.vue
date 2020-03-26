<template>
  <a-list :dataSource="posts">
    <div v-if="search" slot="header" :style="{ position: 'relative' }">
      Arama Sonuçları
      <a-icon
        type="close-circle"
        class="right_list_close"
        @click="$emit('backLiked')"
      />
    </div>
    <div v-else slot="header" :style="{ position: 'relative' }">
      En Çok Beğenilenler
    </div>

    <a-list-item slot="renderItem" slot-scope="item" @click="showPostDetail">
      <p class="most_liked_posts_title">{{ item.title }}</p>
      <p class="most_liked_posts_like">{{ item.likes }} beğeni</p>
      <a-dropdown :trigger="['hover']">
        <a-icon class="most_liked_posts_action" type="more" />
        <a-menu slot="overlay" class="most_liked_posts_feedback_menu">
          <a-menu-item key="0">
            <div>
              <a-icon type="trophy" :style="{ marginRight: '0.3rem' }" />
              İlgimi çekti
            </div>
          </a-menu-item>
          <a-menu-item key="1">
            <a-icon type="frown" :style="{ marginRight: '0.3rem' }" />
            İlgi alanım değil
          </a-menu-item>
          <a-menu-item key="3">
            <a-icon type="thunderbolt" :style="{ marginRight: '0.3rem' }" />
            Spam bildirimi
          </a-menu-item>
        </a-menu>
      </a-dropdown>
      <!-- EN çok beğeni alan post detayı -->
      <post-detail v-if="search" ref="postDetailSearched" :postId="item.id" />
      <post-detail v-else ref="postDetailMostLiked" :postId="item.id" />
    </a-list-item>
  </a-list>
</template>

<script>
import PostDetail from "./PostDetail";

export default {
  name: "RightList",
  props: {
    posts: {
      type: Array,
      required: true
    },
    search: {
      type: Boolean,
      required: true
    }
  },
  components: {
    "post-detail": PostDetail
  },
  data() {
    return {};
  },
  methods: {
    showPostDetail() {
      this.$refs.postDetailMostLiked.showModal = true;
    }
  }
};
</script>