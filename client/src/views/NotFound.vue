<template>
  <div
    class="not_found_wrap"
    :style="{
      background: GET_THEME
        ? GET_COLORS.dark.backgroundLight
        : GET_COLORS.light.backgroundLight
    }"
  >
    <div
      class="not_found"
      :style="{
        background: GET_THEME
          ? GET_COLORS.dark.background
          : GET_COLORS.light.background
      }"
    >
      <p class="not_found_status">404 Not Found</p>
      <p
        class="not_found_desc"
        :style="{
          color: GET_THEME
            ? GET_COLORS.dark.textPrimary
            : GET_COLORS.light.textPrimary
        }"
      >
        Üzgünüz, aradığınız sayfa bulunamadı.
      </p>
      <a-button size="large" type="primary" @click="$router.push('/')"
        >Anasayfa' ya dön</a-button
      >
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "NotFound",
  computed: {
    ...mapGetters(["GET_THEME", "GET_COLORS"])
  }
};
</script>