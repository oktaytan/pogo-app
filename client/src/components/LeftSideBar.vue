<template>
  <a-layout-sider
    class="home_layout_sidebar"
    collapsible
    v-model="collapsedLeft"
    :style="{ width: '250px' }"
  >
    <div class="logo">
      <router-link to="/" class="logo_link">
        <img
          src="../assets/logo.svg"
          alt=""
          :style="{ width: !collapsedLeft ? '30%' : '100%' }"
        />
        <span v-if="!collapsedLeft">Pogo</span>
      </router-link>
    </div>
    <a-menu theme="dark" mode="inline" :defaultSelectedKeys="['1']">
      <a-menu-item key="1" @click="$router.push('/')">
        <a-icon type="home" />
        <span class="nav-text">Anasayfa</span>
      </a-menu-item>
      <a-menu-item key="2" @click="$router.push('/profile')">
        <a-icon type="user" />
        <span class="nav-text">Profilim</span>
      </a-menu-item>
      <a-menu-item key="3" @click="$router.push('/settings')">
        <a-icon type="setting" />
        <span class="nav-text">Ayarlar</span>
      </a-menu-item>
      <a-menu-item key="4" @click="logout">
        <a-icon type="logout" />
        <span class="nav-text">Çıkış Yap</span>
      </a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  name: "LeftSideBar",
  data() {
    return {
      collapsedLeft: false
    };
  },
  methods: {
    logout() {
      this.$router.push("/login");
    }
  }
};
</script>