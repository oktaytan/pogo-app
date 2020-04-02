<template>
  <a-layout id="homeLayout" class="home_layout">
    <!-- Ana navigation -->
    <left-sidebar />
    <a-layout
      class="home_layout_content animated fadeIn"
      :style="{
        borderLeft: GET_THEME
          ? `2px solid ${GET_COLORS.dark.borderColor}`
          : `2px solid ${GET_COLORS.light.borderColor}`,
        borderRight: GET_THEME
          ? `2px solid ${GET_COLORS.dark.borderColor}`
          : `2px solid ${GET_COLORS.light.borderColor}`
      }"
    >
      <!-- Mobile navigation -->
      <mobile-nav />
      <!-- Yeni gönderi ekleme -->
      <add-post v-if="GET_TOP_BAR_SHOW" />
      <a-layout-content
        class="home_layout_content_view"
        :style="{
          background: GET_THEME
            ? GET_COLORS.dark.background
            : GET_COLORS.light.background,
          color: GET_THEME
            ? GET_COLORS.dark.textPrimary
            : GET_COLORS.light.textPrimary
        }"
      >
        <!-- Tüm gönderiler -->
        <Posts v-if="$route.name == 'Anasayfa'" class="animated fadeIn" />
        <!-- view değişimi -->
        <router-view v-else class="animated fadeIn"></router-view>
      </a-layout-content>
    </a-layout>
    <!-- En çok beğeni alan gönderiler -->
    <right-sidebar v-show="GET_RIGHT_BAR_SHOW" />
  </a-layout>
</template>

<script>
import Posts from "../components/Posts";
import LeftSideBar from "../components/LeftSideBar.vue";
import RightSideBar from "../components/RightSideBar.vue";
import MobileNav from "../components/MobileNav.vue";
import AddPost from "../components/AddPost.vue";

import { mapGetters, mapActions } from "vuex";

export default {
  name: "Home",
  components: {
    "left-sidebar": LeftSideBar,
    "right-sidebar": RightSideBar,
    "mobile-nav": MobileNav,
    "add-post": AddPost,
    Posts
  },
  mounted() {
    // Sayfa yenilenmesi ile kullanıcı bilgilerini state ekleyecek action tetikleniyor
    this.USER_REFRESH();
  },
  computed: {
    ...mapGetters([
      "GET_RIGHT_BAR_SHOW",
      "GET_TOP_BAR_SHOW",
      "GET_THEME",
      "GET_COLORS"
    ])
  },
  methods: {
    ...mapActions(["USER_REFRESH"])
  }
};
</script>

<style></style>
