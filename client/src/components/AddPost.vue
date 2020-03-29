<template>
  <div>
    <a-layout-header class="home_layout_content_header animated fadeInDown">
      <a-form :form="newPostForm" class="post_form" @submit.prevent="newPost">
        <a-avatar class="header_avatar">{{ userAvatar }}</a-avatar>
        <div class="post_form_items">
          <div class="post_form_items_data">
            <a-form-item>
              <a-input
                class="new_post_input"
                v-model="post.title"
                placeholder="Konu"
                @change="writePost"
              />
            </a-form-item>
            <a-form-item>
              <textarea-autosize
                class="new_post_input"
                v-model="post.body"
                placeholder="Düşüncelerin..."
                maxLength="351"
                :min-height="0"
                :max-height="350"
                @input="writePost"
              />
              <span class="comment_length_home" v-if="post.body ? true : false"
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
import { mapGetters } from "vuex";
import PageHeader from "./PageHeader";

export default {
  name: "AddPost",
  components: {
    "page-header": PageHeader
  },
  data() {
    return {
      newPostForm: this.$form.createForm(this, { name: "newPost" }),
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
    $route(val) {
      if (val.name === "profile") {
        this.post.title = "";
        this.post.body = "";
      }
    }
  },
  computed: {
    ...mapGetters(["GET_USER"]),
    userAvatar() {
      return this.GET_USER && this.GET_USER.username.slice(0, 2).toUpperCase();
    },
    maxLength() {
      return parseInt(this.post.body.split("").length - 1);
    }
  },
  methods: {
    writePost() {
      if (this.post.title !== "" && this.post.body !== "") {
        this.disabled = false;
      } else {
        this.disabled = true;
      }
    },
    newPost() {
      console.log(this.post);
    }
  }
};
</script>
