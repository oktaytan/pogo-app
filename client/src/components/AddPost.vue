<template>
  <a-layout-header class="home_layout_content_header">
    <a-form :form="newPostForm" class="post_form" @submit.prevent="newPost">
      <a-avatar class="header_avatar">CW</a-avatar>
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
              >{{ maxLength }} karakter</span
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
            paylaş
          </a-button>
        </a-form-item>
      </div>
    </a-form>
  </a-layout-header>
</template>

<script>
export default {
  name: "AddPost",
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
  computed: {
    maxLength() {
      return parseInt(351 - this.post.body.split("").length);
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