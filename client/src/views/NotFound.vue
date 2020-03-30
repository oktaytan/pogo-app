<template>
  <div class="not_found">
    <p class="not_found_status">404 Not Found</p>
    <p class="not_found_desc">Üzgünüz, aradığınız sayfa bulunamadı.</p>
    <a-button type="primary" @click="$router.push('/')"
      >Anasayfa' ya dön</a-button
    >
  </div>
</template>

<script>
export default {
  name: "NotFound"
};
</script>