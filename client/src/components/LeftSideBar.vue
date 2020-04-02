<template>
  <a-layout-sider
    class="home_layout_sidebar"
    v-model="collapsed"
    :style="{
      width: '250px',
      background: GET_THEME
        ? GET_COLORS.dark.background
        : GET_COLORS.light.background
    }"
  >
    <div class="logo">
      <router-link to="/" class="logo_link">
        <img
          src="../assets/logo.svg"
          alt=""
          :style="{ width: !collapsed ? '30%' : '100%' }"
        />
        <span v-if="!collapsed">Pogo</span>
      </router-link>
    </div>
    <!-- Sol navigasyon menüsü -->
    <a-menu
      :theme="GET_THEME ? 'dark' : 'light'"
      mode="inline"
      :defaultSelectedKeys="['1']"
      :selectedKeys="[key]"
      :style="{
        background: GET_THEME
          ? GET_COLORS.dark.background
          : GET_COLORS.light.background,
        color: GET_THEME
          ? GET_COLORS.dark.textPrimary
          : GET_COLORS.light.textPrimary,
        fontWeight: 'bold'
      }"
    >
      <a-menu-item key="1" @click="() => changeMenu('/')">
        <a-icon type="home" />
        <span class="nav-text">Anasayfa</span>
      </a-menu-item>
      <a-menu-item key="2" @click="() => changeMenu('/profile')">
        <a-icon type="user" />
        <span class="nav-text">Profilim</span>
      </a-menu-item>
      <a-menu-item key="3" @click="() => changeMenu('/settings')">
        <a-icon type="setting" />
        <span class="nav-text">Ayarlar</span>
      </a-menu-item>
      <a-menu-item key="4" @click="$router.push('/logout')">
        <a-icon type="logout" />
        <span class="nav-text">Çıkış</span>
      </a-menu-item>
    </a-menu>

    <a-tooltip placement="right">
      <template slot="title">
        {{ collapsed ? "Göster" : "Gizle" }}
      </template>
      <div
        class="ant-layout-sider-trigger"
        :style="{
          width: !collapsed ? '200px' : '80px',
          background: GET_THEME
            ? GET_COLORS.dark.background
            : GET_COLORS.light.background,
          color: GET_THEME
            ? GET_COLORS.dark.textPrimary
            : GET_COLORS.light.textPrimary
        }"
        @click="() => (collapsed = !collapsed)"
      >
        <a-icon :type="collapsed ? 'right' : 'left'" />
      </div>
    </a-tooltip>
  </a-layout-sider>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "LeftSideBar",
  data() {
    return {
      collapsed: false,
      key: localStorage.getItem("pogo_key") // Seçili menü anahtarı
        ? JSON.parse(localStorage.getItem("pogo_key"))
        : "1"
    };
  },
  created() {
    this.routeChange();
  },
  watch: {
    $route(val) {
      this.routeChange(val.name);
    }
  },
  computed: {
    ...mapGetters(["GET_THEME", "GET_COLORS"])
  },
  methods: {
    // Menu her değiştiğinde "route" 'dan hangi sayfadaysak
    // ona ait anahtar localStorage' a ayarlanıyor
    changeMenu(path) {
      this.$router.push(path);
      this.routeChange();
      localStorage.setItem("pogo_key", JSON.stringify(this.key));
    },
    // Route her değiştiğinde gelinen sayfanın ismine göre anahtar seçiliyor
    routeChange() {
      switch (this.$route.name) {
        case "Anasayfa":
          this.key = "1";
          break;
        case "Profilim":
          this.key = "2";
          break;
        case "Ayarlar":
          this.key = "3";
          break;
        default:
          this.key = "1";
      }
    }
  }
};
</script>
