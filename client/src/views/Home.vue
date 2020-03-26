<template>
  <a-layout id="homeLayout" class="home_layout">
    <!-- Ana navigation -->
    <left-sidebar />
    <a-layout class="home_layout_content">
      <!-- Mobile navigation -->
      <mobile-nav />
      <!-- Yeni gönderi ekleme -->
      <add-post />
      <a-layout-content class="home_layout_content_view">
        <!-- Tüm gönderiler -->
        <Posts v-if="$route.name == 'home'" />
        <!-- view değişimi -->
        <router-view v-else></router-view>
      </a-layout-content>
    </a-layout>
    <!-- En çok beğeni alan gönderiler -->
    <right-sidebar />
  </a-layout>
</template>

<script>
import Posts from "../components/Posts";

import LeftSideBar from "../components/LeftSideBar.vue";
import RightSideBar from "../components/RightSideBar.vue";
import MobileNav from "../components/MobileNav.vue";
import AddPost from "../components/AddPost.vue";

export default {
  name: "Home",
  components: {
    "left-sidebar": LeftSideBar,
    "right-sidebar": RightSideBar,
    "mobile-nav": MobileNav,
    "add-post": AddPost,
    Posts
  }
};
</script>

<style></style>
