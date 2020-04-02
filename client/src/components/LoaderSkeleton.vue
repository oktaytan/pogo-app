<template>
  <div class="skeleton">
    <loader class="skeleton_loader" />
    <a-skeleton avatar :paragraph="{ rows: rows }" active />
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import Loader from "./Loader";

export default {
  name: "LoaderSkeleton",
  props: ["rows"],
  components: {
    loader: Loader
  }
};
</script>