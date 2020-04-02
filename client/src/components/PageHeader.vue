<template>
  <div
    class="page_header"
    :style="{
      borderBottom: GET_THEME
        ? `2px solid ${GET_COLORS.dark.borderColor}`
        : `2px solid ${GET_COLORS.light.borderColor}`,
      background: GET_THEME
        ? GET_COLORS.dark.background
        : GET_COLORS.light.background,
      color: GET_THEME
        ? GET_COLORS.dark.textSecondary
        : GET_COLORS.light.textSecondary
    }"
  >
    <div class="back_wrap">
      <!-- Bir önceki sayfaya gitme -->
      <a-icon
        class="back_icon"
        :type="$route.name === 'Detaylar' ? 'home' : 'left-circle'"
        @click="$router.go(-1)"
      />
      <!-- Mevcut sayfanın ismi -->
      <span class="back_text">{{ $route.name }}</span>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "PageHeader",
  props: {
    title: String
  },
  computed: {
    ...mapGetters(["GET_THEME", "GET_COLORS"])
  }
};
</script>
