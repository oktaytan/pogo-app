<template>
  <div
    class="post_actions"
    :class="[position === 'top' ? 'post_actions--top' : 'post_actions--bottom']"
  >
    <div class="detail_action" v-if="details">
      <a-tooltip>
        <template slot="title">
          <span>beğen</span>
        </template>
        <a-icon
          type="star"
          class="post_actions_icon like"
          :theme="isLiked"
          @click.stop="liked"
          :style="{
            color: isLiked == 'filled' ? '#f1c40f' : '',
            animationName: isLiked == 'filled' ? 'bounce' : ''
          }"
        />
      </a-tooltip>
      <span
        :style="{
          animation: likeChange ? 'likeChange 0.4s linear forwards' : 'none'
        }"
        class="like_count"
        >{{ likes }}</span
      >
    </div>
    <a-tooltip placement="top" v-else>
      <template slot="title">
        <span>34 beğeni</span>
      </template>
      <a-icon
        type="star"
        class="post_actions_icon like"
        :theme="isLiked"
        @click.stop="liked"
        :style="{
          color: isLiked == 'filled' ? '#f1c40f' : '',
          animationName: isLiked == 'filled' ? 'bounce' : ''
        }"
      />
    </a-tooltip>
    <a-tooltip placement="top">
      <template slot="title">
        <span>2 yorum</span>
      </template>
      <div>
        <a-popconfirm
          :placement="details ? 'bottomLeft' : 'bottomRight'"
          okText="Gönder"
          cancelText="Vazgeç"
          @confirm="confirm"
          @cancel="e => cancelComment(e)"
          @visibleChange="visibleChange"
        >
          <span slot="icon"
            ><a-avatar type="icon" class="comment_added_avatar"
              >SD</a-avatar
            ></span
          >
          <template slot="title">
            <div class="comment_textarea_wrap" style="positon: relative;">
              <textarea-autosize
                v-model="comment"
                placeholder="Yorum yap..."
                ref="postComment"
                maxLength="201"
                class="comment_textarea"
                :min-height="0"
                :max-height="110"
              />

              <span class="comment_length" v-if="comment ? true : false"
                >{{ maxLength }} karakter</span
              >
            </div>
          </template>
          <a-icon
            type="highlight"
            class="post_actions_icon comment"
            :theme="isCommented"
            @click.stop="e => e.preventDefault()"
            :style="{
              color: isCommented == 'filled' ? '#2ecc71' : '',
              animationName: isCommented == 'filled' ? 'bounce' : ''
            }"
          />
        </a-popconfirm>
      </div>
    </a-tooltip>
    <a-tooltip placement="top">
      <div>
        <template slot="title">
          <span>paylaş</span>
        </template>
        <a-dropdown
          :trigger="['hover']"
          :placement="details ? 'bottomLeft' : 'bottomRight'"
          id="sharePopconfirm"
        >
          <template slot="overlay">
            <p class="share_pl_wrap">
              <a-icon type="facebook" class="share_pl" @click="share" />
              <a-icon type="twitter" class="share_pl" @click="share" />
              <a-icon type="linkedin" class="share_pl" @click="share" />
              <a-icon type="medium" class="share_pl" @click="share" />
            </p>
          </template>
          <a-icon
            type="share-alt"
            class="post_actions_icon share"
            @click.stop="() => false"
            :style="{
              color: isShared ? '#f5f6fa' : '',
              animationName: isShared ? 'bounce' : ''
            }"
          />
        </a-dropdown>
      </div>
    </a-tooltip>
  </div>
</template>

<script>
export default {
  name: "PostActions",
  props: ["position", "details"],
  data() {
    return {
      isLiked: "outlined",
      isCommented: "outlined",
      isShared: false,
      likes: 34,
      likeChange: false,
      comment: "",
      isTypeComment: false
    };
  },
  computed: {
    maxLength() {
      return parseInt(201 - this.comment.split("").length);
    }
  },
  methods: {
    liked() {
      if (this.isLiked == "outlined") {
        this.isLiked = "filled";
        this.likeChange = true;
        this.likes++;
        setTimeout(() => (this.likeChange = false), 200);
      } else {
        this.isLiked = "outlined";
        this.likeChange = true;
        this.likes = 34;
        setTimeout(() => (this.likeChange = false), 200);
      }
    },
    confirm() {
      this.isCommented = "filled";
    },
    addComment() {
      this.isCommented = "filled";
    },
    visibleChange(show) {
      if (show) {
        setTimeout(() => {
          this.$refs.postComment.$el.focus();
        }, 300);
      } else {
        this.comment = "";
      }
    },
    cancelComment(e) {
      e.stopPropagation();
      this.comment = "";
    },
    share() {
      this.isShared = true;
    }
  }
};
</script>
