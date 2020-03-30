<template>
  <a-layout id="homeLayout" class="home_layout">
    <!-- Ana navigation -->
    <left-sidebar />
    <a-layout class="home_layout_content animated fadeIn">
      <!-- Mobile navigation -->
      <mobile-nav />
      <!-- Yeni gönderi ekleme -->
      <add-post v-if="GET_TOP_BAR_SHOW" />
      <a-layout-content class="home_layout_content_view">
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
    ...mapGetters(["GET_RIGHT_BAR_SHOW", "GET_TOP_BAR_SHOW"])
  },
  methods: {
    ...mapActions(["USER_REFRESH"])
  }
};
</script>

<style></style>
