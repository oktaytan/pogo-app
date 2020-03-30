<template>
  <a-list>
    <!-- Arama yapıldıysa sağ bar başlığı değiştiriliyor -->
    <div v-if="search" slot="header" :style="{ position: 'relative' }">
      Arama Sonuçları
      <a-icon
        type="close-circle"
        class="right_list_close"
        @click="$emit('backLiked')"
      />
    </div>
    <!-- Arama yapılmadıysa sağ bar başlığı değiştiriliyor -->
    <div v-else slot="header" :style="{ position: 'relative' }">
      En Çok Beğenilenler
    </div>
    <!-- Sağ panelde listelenecek gönderi özetleri -->
    <a-list-item
      v-for="post in posts"
      :key="post.id"
      class="most_liked_posts_item animated fadeIn"
    >
      <!-- Her özete tıkandığında o gönderiye ait detay sayfasına gidilyor -->
      <div class="click_area" @click="() => detailShow(post.id)"></div>
      <!-- Gönderi başlığı -->
      <p class="most_liked_posts_title">{{ post.title }}</p>
      <!-- Gönderi beğeni sayısı -->
      <p class="most_liked_posts_like">{{ post.likes }} beğeni</p>
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
    </a-list-item>
  </a-list>
</template>

<script>
import PostDetail from "./PostDetail";

export default {
  name: "RightList",
  props: ["posts", "search"],
  components: {
    "post-detail": PostDetail
  },
  data() {
    return {};
  },
  methods: {
    // Gönderiye tıklandığında ona ait detay sayfasına gidiliyor
    detailShow(id) {
      this.$router.push("/");
      setTimeout(() => {
        this.$router.push({ name: "Detaylar", params: { id } });
      }, 0);
    }
  }
};
</script>
