<template>
  <div
    class="post_actions"
    :class="[position === 'top' ? 'post_actions--top' : 'post_actions--bottom']"
  >
    <!-- Detay sayfasında beğeni gösterme -->
    <div class="detail_action" v-if="details">
      <!-- Detay sayfası fakat kullanıcın kendi gönderisi değilse beğeni yapabilir -->
      <a-tooltip v-if="!myPost">
        <template slot="title">
          <span>beğen</span>
        </template>
        <a-icon
          type="star"
          class="post_actions_icon like"
          theme="outlined"
          @click.stop="liked"
          :style="{
            color: isLiked ? '#f1c40f' : '',
            animationName: likeChange && isLiked ? 'bounce' : ''
          }"
        />
      </a-tooltip>
      <span
        v-if="!myPost"
        :style="{
          animation: likeChange ? 'likeChange 0.4s linear forwards' : 'none',
          marginLeft: '0.4rem'
        }"
        class="like_count"
        >{{ likes }}
      </span>
      <!-- Kullanıcın kendi gönderisi ise beğeni yapamaz, sadece görüntüleyebilir. -->
      <span
        v-if="myPost"
        :style="{
          animation: likeChange ? 'likeChange 0.4s linear forwards' : 'none',
          marginLeft: '0.4rem'
        }"
        class="like_count"
        >{{ likes }} <span style="margin-left: 0.3rem">beğeni</span></span
      >
    </div>
    <!-- Anasayfadaki kullanıcının kendi gönderisi ise beğeniyi sadece görüntüleyebiliyor. -->
    <div class="detail_action" v-if="!details && myPost">
      <span
        :style="{
          animation: likeChange ? 'likeChange 0.4s linear forwards' : 'none',
          marginLeft: '0.4rem'
        }"
        class="like_count"
        >{{ likes }} <span style="margin-left: 0.3rem">beğeni</span></span
      >
    </div>
    <!-- Anasayfadaki kullanıcının kendi gönderisi değil ise beğeni yapabiliyor. -->
    <a-tooltip placement="top" v-if="!details && !myPost">
      <template slot="title">
        <span>{{ likes }} beğeni</span>
      </template>
      <a-icon
        type="star"
        class="post_actions_icon like"
        theme="outlined"
        @click.stop="liked"
        :style="{
          color: isLiked ? '#f1c40f' : '',
          animationName: likeChange && isLiked ? 'bounce' : ''
        }"
      />
    </a-tooltip>
    <!-- Yorum görüntüleme ve ekleyebilme alanını açacak kısım -->
    <a-tooltip placement="top">
      <template slot="title">
        <span>{{ commentCount }}</span>
      </template>
      <!-- Yeni yorum ekleme modal' ı -->
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
          <!-- Gönderinin içeriği -->
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
          <!-- Yeni yorumun yapıldığına ekleneceği alan -->
          <a-timeline-item>
            <a-avatar type="icon" slot="dot" class="posted_avatar">{{
              GET_USER.username | userAvatar
            }}</a-avatar>
            <div
              v-if="!commentAdded"
              class="comment_textarea_wrap"
              style="positon: relative;"
            >
              <!-- Yorum için input -->
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
        theme="outlined"
        @click.stop="openCommentModal"
        :style="{
          color: isCommented == 'filled' ? '#2ecc71' : '',
          animationName: isCommented == 'filled' ? 'bounce' : ''
        }"
      />
    </a-tooltip>
    <!-- Gönderiyi farklı platformlarda paylaşma -->
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

    <!-- Eğer gönderi kullanıcının gönderisi is silebilir. -->
    <a-popconfirm
      v-if="myPost && details"
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
      isCommented: "outlined",
      isShared: false,
      likeChange: false,
      userLikes: [],
      isLiked: false,
      comment: "",
      likes: 0,
      likedStatus: [],
      isTypeComment: false,
      commentModal: false,
      commentEmpty: false,
      commentAdded: false
    };
  },
  mounted() {
    //  Gönderiye ait beğeni sayısı ayarlanıyor
    this.likes = this.post.likes;
    //  Kullanıcı gönderiyi beğenmiş mi beğenmemiş mi kontrolü yapılıyor
    this.currentLiked();
  },
  updated() {
    //  Kullanıcı gönderiyi beğenmiş mi beğenmemiş mi kontrolü yapılıyor
    this.currentLiked();
  },
  computed: {
    ...mapGetters(["GET_USER", "GET_USER_LIKES"]),
    // Yeni eklenecek yorumun uzunluğu ayarlanıyor
    maxLength() {
      return parseInt(this.comment.split("").length - 1);
    },
    // Gönderiye ait yorum sayısı belirleniyor
    commentCount() {
      return this.post && this.post.comments
        ? `${this.post.comments.length} yorum`
        : "yorum yok";
    }
  },
  methods: {
    ...mapActions([
      "ADD_COMMENT",
      "FETCH_POST_BY_ID",
      "POST_LIKE",
      "USER_LIKED",
      "FETCH_USER_LIKES",
      "DELETE_POST"
    ]),
    // Kullanıcın beğendiği gönderilerin beğen ikonu güncelleniyor
    currentLiked() {
      // Kullanıcının tüm beğenileri alınıp her post için beğeni durumu getiriliyor
      this.FETCH_USER_LIKES(this.GET_USER.id).then(res => {
        this.userLikes = this.GET_USER_LIKES.liked_posts;
        let post =
          this.userLikes !== null &&
          this.userLikes.filter(item => {
            return item.post_id === this.post.id;
          });
        this.isLiked = post.length > 0 ? post[0].liked : false;
      });
    },
    // Kullanıcı beğeni yaptığından beğeni sayısı ve beğeni durumu değiştiriliyor
    liked() {
      // Eğer gönderi beğeni almışşsa beğeni azaltılıyor
      if (this.isLiked) {
        // Mevcut beğeni sayısı
        let currentLike = this.likes;
        const payload = {
          id: this.post.id,
          likes: --currentLike
        };

        // Beğeni gönderiliyor
        this.POST_LIKE(payload)
          .then(res => {
            if (res.result.success) {
              const likedPayload = {
                post_id: this.post.id,
                user_id: this.GET_USER.id
              };
              // Yeni beğeni sayısı
              this.likes = res.likes;
              // Beğeni durumuna tekrar bakılıyor
              this.USER_LIKED(likedPayload)
                .then(() => {
                  this.$emit("like");
                  this.likeChange = true;
                  setTimeout(() => (this.likeChange = false), 200);
                })
                .catch(err => this.$message.error(err.message));
            }
          })
          .catch(err => this.$message.error(err.message));
      } else {
        // Eğer gönderi beğeni almamışsa beğeni arttırılıyor

        // Mevcut beğeni sayısı
        let currentLike = this.likes;
        const payload = {
          id: this.post.id,
          likes: ++currentLike
        };

        // Beğeni gönderiliyor
        this.POST_LIKE(payload).then(res => {
          if (res.result.success) {
            const likedPayload = {
              post_id: this.post.id,
              user_id: this.GET_USER.id
            };

            // Yeni beğeni sayısı
            this.likes = res.likes;

            // Beğeni durumuna tekrar bakılıyor
            this.USER_LIKED(likedPayload)
              .then(() => {
                this.$emit("like");
                this.likeChange = true;
                setTimeout(() => (this.likeChange = false), 200);
              })
              .catch(err => this.$message.error(err.message));
          }
        });
      }
    },
    // Yeni yorum eklemek için modal açılıyor
    openCommentModal() {
      this.commentModal = true;
      setTimeout(() => {
        this.$refs.postComment.$el.focus();
      }, 300);
    },
    // Yorum için açılan modal kapanırken verisi resetleniyor
    closeCommentModal() {
      this.commentAdded = false;
      this.commentModal = false;
      this.comment = "";
    },
    // Yeni yorum eklemek için action tetikleniyor
    addComment() {
      if (this.comment == "") {
        this.commentEmpty = true;
        this.isCommented = "outlined";
        setTimeout(() => {
          this.commentEmpty = false;
        }, 2000);
      } else {
        this.loading = true;
        // Gidecek yorum
        const newComment = {
          comment: this.comment,
          post_id: this.post.id,
          user_id: this.GET_USER.id
        };
        // Yorum gönderiliyor
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
    // Göderi silme için action tetikleniyor
    deletePost() {
      if (this.GET_USER.username === this.post.author.username) {
        this.DELETE_POST(this.post.id).then(res => {
          if (res.error) {
            this.$message.error(res.message);
          } else {
            this.$message
              .success("Gönderi kaldırıldı.", 0.3)
              .then(() => this.$router.go(-1));
          }
        });
      }
    },
    // Yorum modalı kapatılıyor
    cancelComment(e) {
      e.stopPropagation();
      this.commentModal = false;
      this.comment = "";
    },
    // Gönderi paylaşımı
    share(platform) {
      this.isShared = true;
      this.$message.success(`${platform}' da paylaşıldı!`);
    }
  }
};
</script>
