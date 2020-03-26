<template>
  <a-layout-header class="haeder home_layout_header">
    <div class="logo">
      <router-link to="/" class="logo_link logo_link--horizontal">
        <img src="../assets/logo.svg" alt="" :style="{ width: '30%' }" />
        <span style="display: inline-block; font-size: 1.4rem;">Pogo</span>
      </router-link>
    </div>
    <a-menu theme="dark" mode="horizontal" :defaultSelectedKeys="['1']">
      <a-menu-item key="1" @click="$router.push('/')">
        <a-tooltip placement="bottom">
          <template slot="title">
            <span>Anasayfa</span>
          </template>
          <a-icon type="home" />
        </a-tooltip>
      </a-menu-item>
      <a-menu-item key="2" @click="$router.push('/profile')">
        <a-tooltip placement="bottom">
          <template slot="title">
            <span>Profilim</span>
          </template>
          <a-icon type="user" />
        </a-tooltip>
      </a-menu-item>
      <a-menu-item key="3" @click="$router.push('/settings')">
        <a-tooltip placement="bottom">
          <template slot="title">
            <span>Ayarlar</span>
          </template>
          <a-icon type="setting" />
        </a-tooltip>
      </a-menu-item>
      <a-menu-item key="4">
        <a-tooltip placement="bottom">
          <template slot="title">
            <span>Çıkış Yap</span>
          </template>
          <a-icon type="logout" />
        </a-tooltip>
      </a-menu-item>
    </a-menu>
  </a-layout-header>
</template>

<script>
export default {
  name: "MobileNav",
  methods: {
    logout() {
      this.$router.push("/login");
    }
  }
};
</script>