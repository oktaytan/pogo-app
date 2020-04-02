<template>
  <a-layout-header
    class="haeder home_layout_header"
    :style="{
      background: GET_THEME
        ? GET_COLORS.dark.background
        : GET_COLORS.light.background,
      color: GET_THEME
        ? GET_COLORS.dark.textSecondary
        : GET_COLORS.light.textSecondary,
      borderBottom: GET_THEME
        ? `2px solid ${GET_COLORS.dark.borderColor}`
        : `2px solid ${GET_COLORS.light.borderColor}`
    }"
  >
    <div class="logo">
      <router-link to="/" class="logo_link logo_link--horizontal">
        <img src="../assets/logo.svg" alt="" :style="{ width: '30%' }" />
        <span style="display: inline-block; font-size: 1.4rem;">Pogo</span>
      </router-link>
    </div>
    <!-- Mobil navigsyon menüsü -->
    <a-menu
      theme="dark"
      mode="horizontal"
      :defaultSelectedKeys="['1']"
      :selectedKeys="[key]"
      :style="{
        background: 'transparent'
      }"
    >
      <a-menu-item
        key="1"
        @click="() => changeMenu('/')"
        :style="{
          color: GET_THEME
            ? GET_COLORS.dark.textSecondary
            : GET_COLORS.light.textSecondary
        }"
      >
        <a-icon type="home" />
      </a-menu-item>
      <a-menu-item
        key="2"
        @click="() => changeMenu('/profile')"
        :style="{
          color: GET_THEME
            ? GET_COLORS.dark.textSecondary
            : GET_COLORS.light.textSecondary
        }"
      >
        <a-icon type="user" />
      </a-menu-item>
      <a-menu-item
        key="3"
        @click="() => changeMenu('/settings')"
        :style="{
          color: GET_THEME
            ? GET_COLORS.dark.textSecondary
            : GET_COLORS.light.textSecondary
        }"
      >
        <a-icon type="setting" />
      </a-menu-item>
      <a-menu-item
        key="4"
        @click="$router.push('/logout')"
        :style="{
          color: GET_THEME
            ? GET_COLORS.dark.textSecondary
            : GET_COLORS.light.textSecondary
        }"
      >
        <a-icon type="logout" />
      </a-menu-item>
    </a-menu>
  </a-layout-header>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "MobileNav",
  data() {
    return {
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
