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
      posts: [
        {
          id: 1,
          title: "Title 1",
          likes: 34
        },
        {
          id: 2,
          title: "Title 2",
          likes: 324
        },
        {
          id: 3,
          title: "Title 3",
          likes: 455
        },
        {
          id: 4,
          title: "Title 4",
          likes: 324
        },
        {
          id: 5,
          title: "Title 5",
          likes: 455
        }
      ],
      postsSearch: [
        {
          id: 1,
          title: "Search 1",
          likes: 34
        },
        {
          id: 2,
          title: "Search 2",
          likes: 324
        },
        {
          id: 3,
          title: "Search 3",
          likes: 455
        },
        {
          id: 1,
          title: "Search 1",
          likes: 34
        },
        {
          id: 2,
          title: "Search 2",
          likes: 324
        },
        {
          id: 3,
          title: "Search 3",
          likes: 455
        }
      ]
    };
  },
  mounted() {
    setTimeout(() => (this.loading = false), 2000);
  },
  methods: {
    onSearch() {
      if (this.searchText != "") {
        this.searching = true;
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