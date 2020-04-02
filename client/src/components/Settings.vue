<template>
  <div>
    <!-- Sayfa header'ı -->
    <page-header title="Profilim" />

    <a-card
      class="settings_card"
      :bordered="false"
      :style="{
        background: GET_THEME
          ? GET_COLORS.dark.background
          : GET_COLORS.light.background,
        color: GET_THEME
          ? GET_COLORS.dark.textPrimary
          : GET_COLORS.light.textPrimary
      }"
    >
      <h3 class="settings_name">Kişiselleştirme</h3>
      <div class="settings_wrap">
        <div class="settings_item">
          <!-- Üst barın açılıp kapanması -->
          <span class="label">Üst Bar</span>
          <a-switch
            defaultChecked
            @change="CHANGE_TOP_BAR_SHOW"
            :checked="GET_TOP_BAR_SHOW"
          />
        </div>
        <div class="settings_item">
          <!-- Sağ barın açılıp kapanması -->
          <span class="label">Sağ bar</span>
          <a-switch
            defaultChecked
            @change="CHANGE_RIGHT_BAR_SHOW"
            :checked="GET_RIGHT_BAR_SHOW"
          />
        </div>
        <div class="settings_item">
          <!-- Koyu tema ayarı -->
          <span class="label">Koyu tema</span>
          <a-switch
            defaultChecked
            @change="CHANGE_THEME"
            :checked="GET_THEME"
          />
        </div>
      </div>
    </a-card>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import PageHeader from "./PageHeader";

export default {
  name: "Settings",
  components: {
    "page-header": PageHeader
  },
  computed: {
    ...mapGetters([
      "GET_TOP_BAR_SHOW",
      "GET_RIGHT_BAR_SHOW",
      "GET_THEME",
      "GET_COLORS"
    ])
  },
  methods: {
    ...mapActions([
      "CHANGE_RIGHT_BAR_SHOW",
      "CHANGE_TOP_BAR_SHOW",
      "CHANGE_THEME"
    ])
  }
};
</script>
