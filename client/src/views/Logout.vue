<template>
  <div
    class="logout_wrap animated fadeIn"
    :style="{
      background: GET_THEME
        ? GET_COLORS.dark.backgroundLight
        : GET_COLORS.light.backgroundLight
    }"
  >
    <div
      class="logout_content"
      :style="{
        background: GET_THEME
          ? GET_COLORS.dark.background
          : GET_COLORS.light.background,
        color: GET_THEME
          ? GET_COLORS.dark.textPrimary
          : GET_COLORS.light.textPrimary
      }"
    >
      <img src="../assets/logo.svg" class="logout_img" />
      <p>Pogo' dan çıkmak üzeresin!</p>
      <span>İstediğin zaman tekrar giriş yapabilirsin.</span>
      <div class="logout_actions">
        <a-button size="large" class="cancel_logout" @click="$router.go(-1)"
          >İptal</a-button
        >
        <a-button size="large" type="primary" @click="logout"
          >Çıkış Yap</a-button
        >
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "Logout",
  computed: {
    ...mapGetters(["GET_THEME", "GET_COLORS"])
  },
  methods: {
    ...mapActions(["USER_LOGOUT"]),
    // Kullanıcıyı logout yapacak action tetikleniyor
    logout() {
      this.USER_LOGOUT()
        .then(res => {
          // Kullanıcı bilgileri token ile birlikte localStorage ekleniyor
          localStorage.removeItem("pogo_user");
          if (localStorage.getItem("pogo_reloaded")) {
            localStorage.removeItem("pogo_reloaded");
          }
          // Kullanıcı login sayfasına yönlendiriliyor.
          this.$router.push("/login");
        })
        .catch(err => {
          this.$message.error("Bir hata oluştu!");
          this.$message.info("Tekrar deneyiniz.");
        });
    }
  }
};
</script>