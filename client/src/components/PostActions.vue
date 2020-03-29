<template>
  <div
    class="post_actions"
    :class="[position === 'top' ? 'post_actions--top' : 'post_actions--bottom']"
  >
    <div class="detail_action" v-if="details">
      <a-tooltip v-if="!myPost">
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
        v-if="myPost"
        :style="{
          animation: likeChange ? 'likeChange 0.4s linear forwards' : 'none',
          marginLeft: '0.4rem'
        }"
        class="like_count"
        >{{ post ? post.likes : "0" }}
        <span style="margin-left: 0.3rem">beğeni</span></span
      >
      <span
        v-else
        :style="{
          animation: likeChange ? 'likeChange 0.4s linear forwards' : 'none'
        }"
        class="like_count"
        >{{ likes }}</span
      >
    </div>
    <a-tooltip placement="top" v-else>
      <template slot="title">
        <span>{{ post ? post.likes : "0" }} beğeni</span>
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
        <span>{{ commentCount }}</span>
      </template>
      <a-modal
        title=""
        centered
        :destroyOnClose="true"
        :visible="commentModal"
        @ok="addComment"
        :okText="commentAdded ? 'Eklendi' : 'Yorumla'"
        :okButtonProps="{ props: { disabled: commentAdded } }"
        :cancelText="commentAdded ? 'Kapat' : 'Vazgeç'"
        class="comment_modal"
        @cancel="closeCommentModal"
      >
        <a-timeline class="comment_timeline">
          <a-timeline-item>
            <a-avatar type="icon" slot="dot" class="posted_avatar">{{
              post.author.username | userAvatar
            }}</a-avatar>
            <div class="post_detail_wrap">
              <div>
                <span class="username">{{ post.author.username }}</span>
                <span class="date">{{
                  $moment(post.created_at).fromNow()
                }}</span>
              </div>
              <p class="title">{{ post.title }}</p>
              <p class="body">
                {{ post.body }}
              </p>
            </div>
          </a-timeline-item>
          <a-timeline-item>
            <a-avatar type="icon" slot="dot" class="posted_avatar">{{
              GET_USER.username | userAvatar
            }}</a-avatar>
            <div
              v-if="!commentAdded"
              class="comment_textarea_wrap"
              style="positon: relative;"
            >
              <textarea-autosize
                v-model="comment"
                placeholder="Ne düşünüyorsun..."
                ref="postComment"
                maxLength="201"
                class="comment_textarea"
                :min-height="100"
                :max-height="130"
              />

              <span class="comment_length" v-if="comment ? true : false"
                >{{ maxLength }} / 200</span
              >
              <span v-if="commentEmpty" class="comment_error animated fadeInUp"
                >Yorum alanı boş bırakılamaz!</span
              >
            </div>
            <div v-else class="post_detail_wrap">
              <Loader v-if="loading" />
              <div v-else>
                <div style="width:100%; position: relative">
                  <span class="username">{{ GET_USER.username }}</span>
                  <span class="date">{{ $moment().fromNow() }}</span>
                </div>
                <p class="body">{{ comment }}</p>
              </div>
            </div>
          </a-timeline-item>
        </a-timeline>
      </a-modal>

      <a-icon
        type="highlight"
        class="post_actions_icon comment"
        :theme="isCommented"
        @click.stop="openCommentModal"
        :style="{
          color: isCommented == 'filled' ? '#2ecc71' : '',
          animationName: isCommented == 'filled' ? 'bounce' : ''
        }"
      />
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
              <a-icon
                type="facebook"
                class="share_pl"
                @click="() => share('Facebook')"
              />
              <a-icon
                type="twitter"
                class="share_pl"
                @click="() => share('Twitter')"
              />
              <a-icon
                type="linkedin"
                class="share_pl"
                @click="() => share('LinkedIn')"
              />
              <a-icon
                type="medium"
                class="share_pl"
                @click="() => share('Medium')"
              />
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

    <a-popconfirm
      v-if="myPost"
      :style="{ float: 'right', position: 'relative', marginLeft: '8px' }"
      placement="top"
      title="Bu paylaşımı kaldırmak üzeresiniz?"
      @confirm="deletePost"
      okText="Kaldır"
      cancelText="Vazgeç"
    >
      <a-icon
        type="delete"
        theme="filled"
        class="delete"
        :style="{ top: '1rem', right: '0' }"
      />
    </a-popconfirm>
  </div>
</template>

<script>
import Loader from "./Loader";
import { mapGetters, mapActions } from "vuex";

export default {
  name: "PostActions",
  props: ["position", "details", "myPost", "post"],
  components: {
    Loader
  },
  data() {
    return {
      loading: false,
      isLiked: "outlined",
      isCommented: "outlined",
      isShared: false,
      likeChange: false,
      comment: "",
      likes: 0,
      isTypeComment: false,
      commentModal: false,
      commentEmpty: false,
      commentAdded: false
    };
  },
  mounted() {
    this.likes = this.post.likes;
  },
  computed: {
    ...mapGetters(["GET_USER"]),
    maxLength() {
      return parseInt(this.comment.split("").length - 1);
    },
    commentCount() {
      return this.post && this.post.comments
        ? `${this.post.comments.length} yorum`
        : "yorum yok";
    }
  },
  methods: {
    ...mapActions(["ADD_COMMENT", "FETCH_POST_BY_ID", "LIKE_POST"]),
    liked() {
      if (this.isLiked == "outlined") {
        let currentLike = this.likes;
        const payload = {
          id: this.post.id,
          likes: ++currentLike
        };

        this.LIKE_POST(payload)
          .then(res => {
            if (res.result.success) {
              this.$emit("like");
              this.likes = res.likes;
              this.isLiked = "filled";
              this.likeChange = true;
              setTimeout(() => (this.likeChange = false), 200);
            }
          })
          .catch(err => this.$message.error(err.message));
      } else {
        let currentLike = this.likes;
        const payload = {
          id: this.post.id,
          likes: --currentLike
        };

        this.LIKE_POST(payload).then(res => {
          if (res.result.success) {
            this.$emit("like");
            this.likes = res.likes;
            this.isLiked = "outlined";
            this.likeChange = true;
            setTimeout(() => (this.likeChange = false), 200);
          }
        });
      }
    },
    openCommentModal() {
      this.commentModal = true;
      setTimeout(() => {
        this.$refs.postComment.$el.focus();
      }, 300);
    },
    closeCommentModal() {
      this.commentAdded = false;
      this.commentModal = false;
      this.comment = "";
    },
    addComment() {
      if (this.comment == "") {
        this.commentEmpty = true;
        this.isCommented = "outlined";
        setTimeout(() => {
          this.commentEmpty = false;
        }, 2000);
      } else {
        this.loading = true;
        // Yorum ekleniyor
        const newComment = {
          comment: this.comment,
          post_id: this.post.id,
          user_id: this.GET_USER.id
        };

        this.ADD_COMMENT(newComment)
          .then(() => {
            this.isCommented = "filled";
            this.commentEmpty = false;
            this.commentAdded = true;
            this.$emit("addedComment");
            this.loading = false;
            this.$message.success("Yorum paylaşıldı.");
          })
          .catch(error => {
            this.$message.error(error.message);
          });
      }
    },
    deletePost(e) {
      console.log(e);
      this.$message.success("Paylaşım kaldırıldı!");
    },
    cancelComment(e) {
      e.stopPropagation();
      this.commentModal = false;
      this.comment = "";
    },
    share(platform) {
      this.isShared = true;
      this.$message.success(`${platform}' da paylaşıldı!`);
    }
  }
};
</script>
