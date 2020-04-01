<template>
  <div
    class="posts_container"
    :style="{ height: GET_TOP_BAR_SHOW ? '46vh' : '65vh' }"
  >
    <!-- Yeni gönderi için loader -->
    <loader-skeleton v-if="GET_NEW_POST_LOADING" :rows="0" />
    <!-- Kullanıcının kendi gönderileri -->
    <a-list itemLayout="vertical" size="small">
      <a-list-item
        v-for="post in searchedPosts.length > 0
          ? searchedPosts
          : GET_USERS_POSTS"
        ref="postItem"
        :key="post.id"
        :style="{ padding: '1.5rem', position: 'relative' }"
        class="post_item"
      >
        <div class="click_area" @click.stop="() => detailShow(post.id)"></div>
        <!-- Gönderi başlığı ve oluşturulma zamanı -->
        <a-list-item-meta :style="{ marginBottom: '0' }">
          <div
            class="post_user_wrap"
            slot="title"
            :style="{ marginBottom: '1.5rem', alignItems: 'center' }"
          >
            <span class="post_title">{{ post.title }}</span>
            <span class="post_date" style="margin-left: 1rem">{{
              $moment(post.created_at).fromNow()
            }}</span>
            <!-- Gönderi silme butonu -->
            <a-popconfirm
              :style="{ float: 'right', position: 'relative' }"
              placement="topRight"
              title="Bu paylaşımı kaldırmak üzeresiniz?"
              @confirm="() => deletePost(post)"
              okText="Kaldır"
              cancelText="Vazgeç"
            >
              <a-icon
                type="delete"
                theme="filled"
                class="delete"
                :style="{ top: '1rem', right: '0' }"
              />
            </a-popconfirm>
          </div>
        </a-list-item-meta>
        <!-- Gönderi gövdesi -->
        <div style="margin-top: -1.5rem; width: 85%">
          {{ post.body.slice(0, 200) }}...
          <a-checkable-tag>devamı</a-checkable-tag>
        </div>
      </a-list-item>
    </a-list>
  </div>
</template>

<script>
import LoaderSkeleton from "./LoaderSkeleton";
import PostActions from "./PostActions";

import { mapGetters, mapActions } from "vuex";

export default {
  name: "MyPosts",
  props: ["searchedPosts"],
  components: {
    "loader-skeleton": LoaderSkeleton,
    "post-actions": PostActions
  },
  data() {
    return {
      loading: true,
      user: null
    };
  },
  mounted() {},
  computed: {
    ...mapGetters([
      "GET_TOP_BAR_SHOW",
      "GET_USERS_POSTS",
      "GET_NEW_POST_LOADING"
    ])
  },
  methods: {
    ...mapActions(["DELETE_POST"]),
    // Gönderiye tıklandığında detay sayfasına gidiyor
    detailShow(id) {
      this.$router.push({ name: "Detaylar", params: { id } });
    },
    // Gönderinin silinmesi için action tetikleniyor
    deletePost(post) {
      this.DELETE_POST(post.id).then(res => {
        if (res.error) {
          this.$message.error(res.message);
        } else {
          this.$emit("post-deleted");
          this.$message.success("Gönderi kaldırıldı.");
        }
      });
    }
  }
};
</script>
