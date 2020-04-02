<template>
  <div>
    <a-layout-header
      class="home_layout_content_header animated fadeInDown"
      :style="{
        background: GET_THEME
          ? GET_COLORS.dark.background
          : GET_COLORS.light.background,
        color: GET_THEME
          ? GET_COLORS.dark.textPrimary
          : GET_COLORS.light.textPrimary,
        fontWeight: 'bold',
        borderBottom: GET_THEME
          ? `10px solid ${GET_COLORS.dark.borderColor}`
          : `10px solid ${GET_COLORS.light.borderColor}`
      }"
    >
      <!-- Yeni gönderi ekleme formu -->
      <a-form :form="newPostForm" class="post_form" @submit.prevent="newPost">
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
          >{{ userAvatar }}</a-avatar
        >
        <div class="post_form_items">
          <div class="post_form_items_data">
            <!-- Yeni gönderi başlığı -->
            <a-form-item>
              <a-input
                class="new_post_input"
                v-model="post.title"
                placeholder="Konu"
                @change="writePost"
                :style="{
                  color: GET_THEME
                    ? GET_COLORS.dark.textPrimary
                    : GET_COLORS.light.textPrimary
                }"
              />
            </a-form-item>
            <!-- Yeni gönderi gövdesi -->
            <a-form-item>
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
              <span
                :style="{
                  color: GET_THEME
                    ? GET_COLORS.dark.textSecondary
                    : GET_COLORS.light.textSecondary
                }"
                class="comment_length_home"
                v-if="post.body ? true : false"
                >{{ maxLength }} / 350</span
              >
            </a-form-item>
          </div>
          <a-form-item :style="{ marginTop: '1rem' }">
            <a-button
              class="new_post_submit_btn"
              type="primary"
              size="large"
              html-type="submit"
              :disabled="disabled"
            >
              Pogo ile paylaş
            </a-button>
          </a-form-item>
        </div>
      </a-form>
    </a-layout-header>

    <page-header title="Paylaşım detayları" v-if="$route.name === 'Detaylar'" />
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";
import PageHeader from "./PageHeader";

export default {
  name: "AddPost",
  components: {
    "page-header": PageHeader
  },
  data() {
    return {
      newPostForm: this.$form.createForm(this, { name: "newPost" }), // Form initialize
      post: {
        title: "",
        body: "",
        created_at: new Date().getTime(),
        updated_at: new Date().getTime(),
        user_id: "",
        likes: 0
      },
      disabled: true
    };
  },
  watch: {
    // Route profil sayfasına gidiyorsa form resetleniyor
    $route(val) {
      if (val.name === "profile") {
        this.post.title = "";
        this.post.body = "";
      }
    }
  },
  computed: {
    ...mapGetters(["GET_USER", "GET_THEME", "GET_COLORS"]),
    // Avatar kullanıcı adının ilk iki harfinden oluşturuluyor
    userAvatar() {
      return this.GET_USER && this.GET_USER.username.slice(0, 2).toUpperCase();
    },
    // Eklenebilecek gönderinin maksimum değeri ayarlanıyor
    maxLength() {
      return parseInt(this.post.body.split("").length - 1);
    }
  },
  methods: {
    ...mapActions(["ADD_NEW_POST", "FETCH_USERS_POSTS"]),
    // Gönderi başlığı ve gövdesi boş değilse paylaş butonu aktif hale geliyor
    writePost() {
      if (this.post.title !== "" && this.post.body !== "") {
        this.disabled = false;
      } else {
        this.disabled = true;
      }
    },
    // Yeni gönderi için aciton tetikleniyor
    newPost() {
      this.post.user_id = this.GET_USER.id;
      this.ADD_NEW_POST(this.post).then(res => {
        if (res.result.success) {
          this.FETCH_USERS_POSTS(this.GET_USER.username).then(() => {
            this.$message.success(res.result.message);
            this.post.title = "";
            this.post.body = "";
          });
        }
      });
    }
  }
};
</script>
