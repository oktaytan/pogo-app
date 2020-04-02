<template>
  <div class="profile_wrap">
    <!-- Sayfa headerı -->
    <page-header title="Profilim" />
    <!-- Kullanıcı bilgileri -->
    <a-card
      hoverable
      :bordered="false"
      class="profile_card"
      :style="{
        background: GET_THEME
          ? GET_COLORS.dark.background
          : GET_COLORS.light.background,
        color: GET_THEME
          ? GET_COLORS.dark.textPrimary
          : GET_COLORS.light.textPrimary,
        borderBottom: GET_THEME
          ? `10px solid ${GET_COLORS.dark.borderColor}`
          : `10px solid ${GET_COLORS.light.borderColor}`
      }"
    >
      <a-card-meta :description="email">
        <div
          slot="title"
          :style="{
            color: GET_THEME
              ? GET_COLORS.dark.textPrimary
              : GET_COLORS.light.textPrimary
          }"
        >
          {{ username }}
        </div>
        <a-avatar
          class="profile_avatar"
          slot="avatar"
          :style="{
            background: GET_THEME
              ? GET_COLORS.dark.backgroundLight
              : GET_COLORS.light.backgroundLight,
            color: GET_THEME
              ? GET_COLORS.dark.textPrimary
              : GET_COLORS.light.textPrimary,
            fontWeight: 'bold'
          }"
          >{{ username | userAvatar }}</a-avatar
        >
      </a-card-meta>
      <div class="profile_header" v-if="!loadingPosts">
        <!-- Kullanıcının paylaşım sayısı -->
        <a-tag
          color="#1e5181"
          class="post_count"
          :style="{
            background: GET_THEME
              ? GET_COLORS.dark.backgroundLight
              : GET_COLORS.light.backgroundLight,
            color: GET_THEME
              ? GET_COLORS.dark.textPrimary
              : GET_COLORS.light.textPrimary,
            fontWeight: 'bold'
          }"
          >{{ userPostCount }} gönderi</a-tag
        >
        <!-- Kullanıcının beğeni sayısı -->
        <a-tag
          color="#1e5181"
          class="like_count"
          :style="{
            background: GET_THEME
              ? GET_COLORS.dark.backgroundLight
              : GET_COLORS.light.backgroundLight,
            color: GET_THEME
              ? GET_COLORS.dark.textPrimary
              : GET_COLORS.light.textPrimary,
            fontWeight: 'bold'
          }"
          >{{ userLikeCount }} beğeni</a-tag
        >
        <!-- Kullanıcının yaptığı yorum sayısı -->
        <a-tag
          color="#1e5181"
          class="comment_count"
          :style="{
            background: GET_THEME
              ? GET_COLORS.dark.backgroundLight
              : GET_COLORS.light.backgroundLight,
            color: GET_THEME
              ? GET_COLORS.dark.textPrimary
              : GET_COLORS.light.textPrimary,
            fontWeight: 'bold'
          }"
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
      centered
      :visible="newPostModal"
      class="new_post_modal"
      @cancel="cancelPost"
      :footer="false"
      :bodyStyle="{
        color: GET_THEME
          ? GET_COLORS.dark.textPrimary
          : GET_COLORS.light.textPrimary,
        background: GET_THEME
          ? GET_COLORS.dark.background
          : GET_COLORS.light.background,
        paddingBottom: '0',
        borderRadius: '0.5rem'
      }"
    >
      <a-form
        :form="newPostFormFromProfile"
        class="post_form"
        @submit.prevent="newPost"
      >
        <!-- Kullanıcının avatarı -->
        <a-avatar
          class="header_avatar"
          :style="{
            background: GET_THEME
              ? GET_COLORS.dark.backgroundLight
              : GET_COLORS.light.backgroundLight,
            color: GET_THEME
              ? GET_COLORS.dark.textPrimary
              : GET_COLORS.light.textPrimary,
            fontWeight: 'bold'
          }"
          >{{ username | userAvatar }}</a-avatar
        >
        <div class="post_form_items">
          <div class="post_form_items_data">
            <a-form-item>
              <!-- Gönderi başlığı -->
              <a-input
                ref="profile_page_title_input"
                class="new_post_input"
                :style="{
                  marginBottom: '1rem',
                  color: GET_THEME
                    ? GET_COLORS.dark.textPrimary
                    : GET_COLORS.light.textPrimary
                }"
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
                :style="{
                  color: GET_THEME
                    ? GET_COLORS.dark.textPrimary
                    : GET_COLORS.light.textPrimary
                }"
              />
              <!-- Gönderi gövdesinin kalan karakter sayısı -->
              <span
                :style="{
                  color: GET_THEME
                    ? GET_COLORS.dark.textPrimary
                    : GET_COLORS.light.textPrimary
                }"
                class="comment_length_home"
                v-if="post.body ? true : false"
                >{{ maxLength }} / 350</span
              >
            </a-form-item>
          </div>
        </div>
      </a-form>
      <div class="post_form_footer">
        <a-button key="back" @click="cancelPost" type="link" class="cancel_btn"
          >Vazgeç</a-button
        >
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
      </div>
    </a-modal>

    <div
      class="profile_shared"
      :style="{
        borderBottom: GET_THEME
          ? `2px solid ${GET_COLORS.dark.borderColor}`
          : `2px solid ${GET_COLORS.light.borderColor}`
      }"
    >
      <span class="shared_title">Paylaşımlarım</span>
      <!-- Kullanıcının kendi paylaşımlarında arama -->
      <div class="my_post_search">
        <a-icon
          v-if="searchedPosts.length > 0"
          type="close-circle"
          class="cancel_search"
          @click="cancelSearch"
        />
        <div class="search_input_wrap">
          <a-input
            placeholder="Ara..."
            @keyup.enter="onSearch"
            v-model="searchText"
            size="large"
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
      "GET_ALL_COMMENTS",
      "GET_THEME",
      "GET_COLORS"
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
              this.GET_USER_LIKES && this.GET_USER_LIKES.liked_posts
                ? this.GET_USER_LIKES.liked_posts.filter(
                    item => item.liked == 1
                  ).length
                : 0;
          });
          // Kullanıcının yaptığı yorumları getirecek action tetikleniyor
          this.FETCH_ALL_COMMENTS().then(() => {
            const comments =
              this.GET_ALL_COMMENTS &&
              this.GET_ALL_COMMENTS.filter(item => {
                return item.author.username === this.user.username;
              });
            this.userCommentCount = comments ? comments.length : 0;
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
