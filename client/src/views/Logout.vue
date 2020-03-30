<template>
  <div class="logout_wrap animated fadeIn">
    <div class="logout_content">
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
import { mapActions } from "vuex";

export default {
  name: "Logout",
  methods: {
    ...mapActions(["USER_LOGOUT"]),
    // Kullanıcıyı logout yapacak action tetikleniyor
    logout() {
      this.USER_LOGOUT()
        .then(res => {
          // Kullanıcı bilgileri token ile birlikte localStorage ekleniyor
          localStorage.removeItem("pogo_user");
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