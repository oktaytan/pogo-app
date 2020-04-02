<template>
  <a-layout-sider
    class="home_layout_sidebar home_layout_sidebar--left animated fadeInRight"
    :style="{
      minWidth: '450px',
      padding: '1rem 2rem',
      background: GET_THEME
        ? GET_COLORS.dark.background
        : GET_COLORS.light.background
    }"
  >
    <a-row>
      <!-- Gönderi arama alanı -->
      <div class="search_input_wrap">
        <a-input
          placeholder="Pogo' da ara..."
          @keyup.enter="onSearch"
          v-model="searchText"
          size="large"
          id="searchInput"
          autoComplete="off"
          :style="{
            background: GET_THEME
              ? GET_COLORS.dark.backgroundLight
              : GET_COLORS.light.backgroundLight,
            color: GET_THEME
              ? GET_COLORS.dark.textPrimary
              : GET_COLORS.light.textPrimary
          }"
        />
        <a-icon class="search_icon" type="search" @click="onSearch" />
      </div>
      <div class="most_liked_posts">
        <Loader v-if="loading" />
        <!-- Arama sonuçlarının listesi -->
        <right-list
          v-else-if="!loading && searching"
          :posts="postsSearch"
          :search="true"
          @backLiked="cancelSearch"
        />
        <!-- En çok beğeni alan gönderiler listesi -->
        <right-list
          v-else
          :posts="posts"
          :search="false"
          @change-limit="lm => fetchMostLikes(lm)"
        />
      </div>
      <a-layout-footer
        :style="{
          textAlign: 'center',
          color: GET_THEME
            ? GET_COLORS.dark.textSecondary
            : GET_COLORS.light.textSecondary
        }"
      >
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
      postsSearch: [],
      limit: 5
    };
  },
  computed: {
    ...mapGetters([
      "GET_POST_BY_LIKES",
      "GET_SEARCH_POSTS",
      "GET_THEME",
      "GET_COLORS"
    ])
  },
  mounted() {
    // Beğeni sayısına göre sıralanmış gönderiler alınıyor
    this.fetchMostLikes(this.limit);
  },
  methods: {
    ...mapActions(["FETCH_POST_BY_LIKES", "FETCH_SEARCH_POSTS"]),
    // En çok beğeni alan gönderileri getirme
    fetchMostLikes(limit) {
      this.FETCH_POST_BY_LIKES(limit).then(res => {
        this.posts = this.GET_POST_BY_LIKES;
        this.loading = false;
      });
    },
    // Arama yapılıyor
    onSearch() {
      if (this.searchText != "") {
        this.searching = true;
        // Arama sonuçalrının getirecek action tetikleniyor
        this.FETCH_SEARCH_POSTS({ search: this.searchText }).then(res => {
          this.postsSearch = this.GET_SEARCH_POSTS;
        });
      } else {
        this.searching = false;
      }
    },
    // Arama sonuçları iptal edilip bir önceki liste getiriliyor
    cancelSearch() {
      this.searching = false;
      this.searchText = "";
    }
  }
};
</script>