<template>
  <div class="prepare">
    <p class="prepare_text">Giriş yapılıyor</p>
    <loader :style="{ minHeight: 'initial', height: '60px', margin: '0' }" />
  </div>
</template>

<script>
import Loader from "../components/Loader";

export default {
  name: "Prepare",
  components: {
    loader: Loader
  },
  mounted() {
    if (localStorage.getItem("pogo_reloaded")) {
      localStorage.removeItem("pogo_reloaded");
    } else {
      localStorage.setItem("pogo_reloaded", "1");
      location.reload();
    }
    setTimeout(() => {
      this.$router.push("/");
    }, 300);
  }
};
</script>
