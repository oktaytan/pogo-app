<template>
  <a-layout-header class="haeder home_layout_header">
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
    >
      <a-menu-item key="1" @click="() => changeMenu('/')">
        <a-icon type="home" />
      </a-menu-item>
      <a-menu-item key="2" @click="() => changeMenu('/profile')">
        <a-icon type="user" />
      </a-menu-item>
      <a-menu-item key="3" @click="() => changeMenu('/settings')">
        <a-icon type="setting" />
      </a-menu-item>
      <a-menu-item key="4" @click="$router.push('/logout')">
        <a-icon type="logout" />
      </a-menu-item>
    </a-menu>
  </a-layout-header>
</template>

<script>
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
