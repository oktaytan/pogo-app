<template>
  <a-list
    :style="{
      background: GET_THEME
        ? GET_COLORS.dark.background
        : GET_COLORS.light.background
    }"
  >
    <!-- Arama yapıldıysa sağ bar başlığı değiştiriliyor -->
    <div
      v-if="search"
      slot="header"
      :style="{
        position: 'relative',
        color: GET_THEME
          ? GET_COLORS.dark.textPrimary
          : GET_COLORS.light.textPrimary,
        borderBottom: GET_THEME
          ? `2px solid ${GET_COLORS.dark.borderColor}`
          : `2px solid ${GET_COLORS.light.borderColor}`
      }"
    >
      <div style="padding: 0.7rem 1.5rem; position: relative">
        Arama Sonuçları
        <a-icon
          type="close-circle"
          class="right_list_close"
          @click="$emit('backLiked')"
          :style="{
            color: GET_THEME
              ? GET_COLORS.dark.textPrimary
              : GET_COLORS.light.textPrimary
          }"
        />
      </div>
    </div>
    <!-- Arama yapılmadıysa sağ bar başlığı değiştiriliyor -->
    <div
      v-else
      slot="header"
      :style="{
        position: 'relative',
        color: GET_THEME
          ? GET_COLORS.dark.textPrimary
          : GET_COLORS.light.textPrimary,
        borderBottom: GET_THEME
          ? `2px solid ${GET_COLORS.dark.borderColor}`
          : `2px solid ${GET_COLORS.light.borderColor}`
      }"
    >
      <div style="padding: 0.7rem 1.5rem">
        En Çok Beğenilenler
      </div>
    </div>
    <!-- Sağ panelde listelenecek gönderi özetleri -->
    <a-list-item
      v-for="post in posts"
      :key="post.id"
      class="most_liked_posts_item animated fadeIn"
      :style="{
        borderBottom: GET_THEME
          ? `2px solid ${GET_COLORS.dark.borderColor}`
          : `2px solid ${GET_COLORS.light.borderColor}`,
        color: GET_THEME
          ? GET_COLORS.dark.textPrimary
          : GET_COLORS.light.textPrimary
      }"
    >
      <!-- Her özete tıkandığında o gönderiye ait detay sayfasına gidilyor -->
      <div class="click_area" @click="() => detailShow(post.id)"></div>
      <!-- Gönderi başlığı -->
      <p class="most_liked_posts_title">{{ post.title.slice(0, 40) }}</p>
      <!-- Gönderi beğeni sayısı -->
      <p class="most_liked_posts_like">{{ post.likes }} beğeni</p>
      <a-dropdown :trigger="['hover']">
        <a-icon class="most_liked_posts_action" type="more" />
        <a-menu
          slot="overlay"
          class="most_liked_posts_feedback_menu"
          :style="{
            background: GET_THEME
              ? GET_COLORS.dark.backgroundLight
              : GET_COLORS.light.backgroundLight
          }"
        >
          <a-menu-item
            key="0"
            :style="{
              color: GET_THEME
                ? GET_COLORS.dark.textPrimary
                : GET_COLORS.light.textPrimary
            }"
          >
            <div>
              <a-icon type="trophy" :style="{ marginRight: '0.3rem' }" />
              İlgimi çekti
            </div>
          </a-menu-item>
          <a-menu-item
            key="1"
            :style="{
              color: GET_THEME
                ? GET_COLORS.dark.textPrimary
                : GET_COLORS.light.textPrimary
            }"
          >
            <a-icon type="frown" :style="{ marginRight: '0.3rem' }" />
            İlgi alanım değil
          </a-menu-item>
          <a-menu-item
            key="3"
            :style="{
              color: GET_THEME
                ? GET_COLORS.dark.textPrimary
                : GET_COLORS.light.textPrimary
            }"
          >
            <a-icon type="thunderbolt" :style="{ marginRight: '0.3rem' }" />
            Spam bildirimi
          </a-menu-item>
        </a-menu>
      </a-dropdown>
    </a-list-item>
    <a-list-item
      v-if="!search"
      :style="{ padding: '0.5rem 1.5rem' }"
      @click="changeLimit"
    >
      <a-button type="link">{{
        more ? "Daha az göster" : "Daha fazla göster"
      }}</a-button>
    </a-list-item>
  </a-list>
</template>

<script>
import { mapGetters } from "vuex";
import PostDetail from "./PostDetail";

export default {
  name: "RightList",
  props: ["posts", "search"],
  components: {
    "post-detail": PostDetail
  },
  data() {
    return {
      more: false
    };
  },
  computed: {
    ...mapGetters(["GET_THEME", "GET_COLORS"])
  },
  methods: {
    // Gönderiye tıklandığında ona ait detay sayfasına gidiliyor
    detailShow(id) {
      this.$router.push("/");
      setTimeout(() => {
        this.$router.push({ name: "Detaylar", params: { id } });
      }, 0);
    },
    // En çok beğeni alan gönderilerde
    // kullanıcıya gösterilecek gönderi sayısı arttırılıyor
    changeLimit() {
      if (this.more) {
        this.$emit("change-limit", 5);
        setTimeout(() => {
          this.more = false;
        }, 300);
      } else {
        this.$emit("change-limit", 10);
        setTimeout(() => {
          this.more = true;
        }, 300);
      }
    }
  }
};
</script>
