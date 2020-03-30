<template>
  <div class="profile_wrap">
    <!-- Sayfa headerı -->
    <page-header title="Profilim" />
    <!-- Kullanıcı bilgileri -->
    <a-card hoverable :bordered="false" class="profile_card">
      <a-card-meta :title="username" :description="email">
        <a-avatar class="profile_avatar" slot="avatar">{{
          username | userAvatar
        }}</a-avatar>
      </a-card-meta>
      <div class="profile_header" v-if="!loadingPosts">
        <!-- Kullanıcının paylaşım sayısı -->
        <a-tag color="#1e5181" class="post_count"
          >{{ userPostCount }} gönderi</a-tag
        >
        <!-- Kullanıcının beğeni sayısı -->
        <a-tag color="#1e5181" class="like_count"
          >{{ userLikeCount }} beğeni</a-tag
        >
        <!-- Kullanıcının yaptığı yorum sayısı -->
        <a-tag color="#1e5181" class="comment_count"
          >{{ userCommentCount }} yorum</a-tag
        >
      </div>
      <!-- Yeni gönderi paylaşımı -->
      <a-button
        v-if="!GET_TOP_BAR_SHOW"
        class="openModalBtn"
        type="primary"
        ghost
        size="large"
        @click="openPostModal"
      >
        Pogo ile paylaş
      </a-button>
    </a-card>
    <!-- yeni gönderi için modal -->
    <a-modal
      title="Pogo ile paylaş"
      centered
      :visible="newPostModal"
      class="new_post_modal"
      @cancel="cancelPost"
    >
      <a-form
        :form="newPostFormFromProfile"
        class="post_form"
        @submit.prevent="newPost"
      >
        <!-- Kullanıcının avatarı -->
        <a-avatar class="header_avatar">{{ username | userAvatar }}</a-avatar>
        <div class="post_form_items">
          <div class="post_form_items_data">
            <a-form-item>
              <!-- Gönderi başlığı -->
              <a-input
                ref="profile_page_title_input"
                class="new_post_input"
                :style="{ marginBottom: '1rem' }"
                v-model="post.title"
                placeholder="Konu"
                @change="writePost"
              />
            </a-form-item>
            <a-form-item>
              <!-- Gönderi gövdesi -->
              <textarea-autosize
                class="new_post_input"
                v-model="post.body"
                placeholder="Düşüncelerin..."
                maxLength="351"
                :min-height="0"
                :max-height="350"
                @input="writePost"
              />
              <!-- Gönderi gövdesinin kalan karakter sayısı -->
              <span class="comment_length_home" v-if="post.body ? true : false"
                >{{ maxLength }} / 350</span
              >
            </a-form-item>
          </div>
        </div>
      </a-form>

      <template slot="footer">
        <a-button key="back" @click="cancelPost">Vazgeç</a-button>
        <a-button
          key="submit"
          :class="disabled && 'disabled_btn'"
          type="primary"
          :loading="loading"
          @click="newPost"
          :disabled="disabled"
        >
          Paylaş
        </a-button>
      </template>
    </a-modal>

    <div class="profile_shared">
      <span class="shared_title">Paylaşımlarım</span>
      <!-- Kullanıcının kendi paylaşımlarında arama -->
      <div class="my_post_search">
        <a-icon
          v-if="searchedPosts.length > 0"
          type="close-circle"
          class="cancel_search"
          @click="cancelSearch"
        />
        <a-input-search
          placeholder="Ara..."
          @search="onSearch"
          v-model="searchText"
          size="large"
          autoComplete="off"
        />
      </div>
    </div>

    <Loader v-if="loadingPosts || searching" />
    <!-- Kullanıcın kendi paylaşımları -->
    <my-posts
      v-else
      @post-deleted="fetchUserPosts"
      :searchedPosts="searchedPosts"
    />
  </div>
</template>
<script>
import { mapGetters, mapActions } from "vuex";
import MyPosts from "./MyPosts";
import PageHeader from "./PageHeader";
import Loader from "./Loader";

export default {
  name: "Profile",
  components: {
    Loader,
    "my-posts": MyPosts,
    "page-header": PageHeader
  },
  data() {
    return {
      loading: false,
      loadingPosts: true,
      disabled: true,
      newPostModal: false,
      newPostFormFromProfile: this.$form.createForm(this, { name: "newPost" }),
      post: {
        title: "",
        body: "",
        created_at: new Date().getTime(),
        updated_at: new Date().getTime(),
        user_id: "",
        likes: 0
      },
      searchedPosts: [],
      disabled: true,
      searching: false,
      searchText: "",
      user: null,
      userCommentCount: 0,
      userPostCount: 0,
      userLikeCount: 0
    };
  },
  mounted() {
    // Kullanıcı bilgileri localStorage' dan alınıyor
    this.user = JSON.parse(localStorage.getItem("pogo_user")).user;
    // Kullanıcının kendi gönderileri alınıyor
    this.fetchUserPosts();
  },
  watch: {
    // Kullanıcının kendi yaptığı paylaşımların sayısı izleniyor
    GET_USERS_POSTS(value) {
      this.userPostCount = value !== null && value.length;
    }
  },
  computed: {
    ...mapGetters([
      "GET_TOP_BAR_SHOW",
      "GET_USER",
      "GET_USERS_POSTS",
      "GET_USER_LIKES",
      "GET_ALL_COMMENTS"
    ]),
    // Yeni gönderinin karakter sayısı belirleniyor
    maxLength() {
      return parseInt(this.post.body.split("").length - 1);
    },
    // Kullanıcının kullanıcı adı ayarlanıyor
    username() {
      return JSON.parse(localStorage.getItem("pogo_user")).user.username;
    },
    // Kullanıcının email adresi ayarlanıyor
    email() {
      return JSON.parse(localStorage.getItem("pogo_user")).user.email;
    }
  },
  methods: {
    ...mapActions([
      "FETCH_USERS_POSTS",
      "FETCH_ALL_COMMENTS",
      "FETCH_USER_LIKES",
      "ADD_NEW_POST"
    ]),
    // Kullanıcının kendi gönderilerini getirecek action tetikleniyor
    fetchUserPosts() {
      this.FETCH_USERS_POSTS(this.user.username).then(res => {
        if (res === null) {
          // Hiç gönderi yoksa
          this.$message.info("Hiç paylaşımınız yok!");
          this.loadingPosts = false;
        } else {
          this.userPostCount =
            this.GET_USERS_POSTS && this.GET_USERS_POSTS.length;
          this.loadingPosts = false;
          // Kullanıcının yaptığı beğenileri getirecek action tetikleniyor
          this.FETCH_USER_LIKES(this.user.id).then(() => {
            this.userLikeCount =
              this.GET_USER_LIKES && this.GET_USER_LIKES.liked_posts.length;
          });
          // Kullanıcının yaptığı yorumları getirecek action tetikleniyor
          this.FETCH_ALL_COMMENTS().then(() => {
            const comments = this.GET_ALL_COMMENTS.filter(item => {
              return item.author.username === this.user.username;
            });
            this.userCommentCount = comments.length;
          });
        }
      });
    },
    // Yeni gönderi iin modal açılıyor
    openPostModal() {
      this.newPostModal = true;
      setTimeout(() => {
        this.$refs.profile_page_title_input.$el.focus();
      }, 300);
    },
    // Yeni gönderinin başlığı ve içeriğinin dolu olduğu kontrol edilerek
    // paylaş butonu aktif edilyor
    writePost() {
      if (this.post.title !== "" && this.post.body !== "") {
        this.disabled = false;
      } else {
        this.disabled = true;
      }
    },
    // Yeni gönderi yapılması için action tetikleniyor
    newPost() {
      this.post.user_id = this.GET_USER.id;
      this.ADD_NEW_POST(this.post).then(res => {
        if (res.result.success) {
          // Yeni gönderi eklendikten sonra tüm gönderiler tekrar alınıyor
          this.FETCH_USERS_POSTS(this.GET_USER.username).then(() => {
            this.$message.success(res.result.message);
            // Form alanları resetleniyor
            this.post.title = "";
            this.post.body = "";
          });
        }
      });
    },
    // Yeni gönderi modalı kapatılırken form alanları resetleniyor
    cancelPost() {
      (this.post.title = ""),
        (this.post.body = ""),
        (this.newPostModal = false);
    },
    // Kullanıcının kendi gönderilerinde arama yapılıyor
    onSearch() {
      if (this.searchText != "") {
        this.searching = true;
        this.searchedPosts = this.GET_USERS_POSTS.filter(item => {
          return (
            item.title.toLowerCase().indexOf(this.searchText.toLowerCase()) >
              -1 ||
            item.body.toLowerCase().indexOf(this.searchText.toLowerCase()) > -1
          );
        });
        this.searching = false;
      } else {
        this.searchedPosts = [];
        this.searching = false;
      }
    },
    // Arama iptaledilip mevcut listeye dönülüyor
    cancelSearch() {
      this.searchedPosts = [];
      this.searching = false;
      this.searchText = "";
    }
  }
};
</script>
