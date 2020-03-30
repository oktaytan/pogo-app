<template>
  <a-layout-sider
    class="home_layout_sidebar home_layout_sidebar--left animated fadeInRight"
    :style="{ minWidth: '450px', padding: '1rem 2rem' }"
  >
    <a-row>
      <a-input-search
        placeholder="Pogo' da ara..."
        @search="onSearch"
        v-model="searchText"
        size="large"
        id="searchInput"
        autoComplete="off"
      />
      <div class="most_liked_posts">
        <Loader v-if="loading" />
        <right-list
          v-else-if="!loading && searching"
          :posts="postsSearch"
          :search="true"
          @backLiked="cancelSearch"
        />
        <right-list v-else :posts="posts" :search="false" />
      </div>
      <a-layout-footer :style="{ textAlign: 'center' }">
        {{ `&copy; ${new Date().getFullYear()} POGO, Inc.` }}
      </a-layout-footer>
    </a-row>
  </a-layout-sider>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import Loader from "../components/Loader";
import RightList from "./RightList";

export default {
  name: "RightSideBar",
  components: {
    Loader,
    "right-list": RightList
  },
  data() {
    return {
      loading: true,
      searching: false,
      searchText: "",
      posts: [],
      postsSearch: []
    };
  },
  computed: {
    ...mapGetters(["GET_POST_BY_LIKES", "GET_SEARCH_POSTS"])
  },
  mounted() {
    this.FETCH_POST_BY_LIKES().then(res => {
      this.posts = this.GET_POST_BY_LIKES;
      this.loading = false;
    });
  },
  methods: {
    ...mapActions(["FETCH_POST_BY_LIKES", "FETCH_SEARCH_POSTS"]),
    onSearch() {
      if (this.searchText != "") {
        this.searching = true;
        this.FETCH_SEARCH_POSTS({ search: this.searchText }).then(res => {
          this.postsSearch = this.GET_SEARCH_POSTS;
        });
      } else {
        this.searching = false;
      }
    },
    cancelSearch() {
      this.searching = false;
      this.searchText = "";
    }
  }
};
</script>