<template>
  <a-modal
    title=""
    centered
    v-model="showModal"
    closable
    :footer="null"
    :bodyStyle="{
      minHeight: '400px',
      borderRadius: '10px',
      border: '1px solid #486581'
    }"
    class="post_detail_wrap"
  >
    <div class="post_detail">
      <div class="post_detail_header">
        <a-avatar class="post_avatar">CW</a-avatar>
        <div>
          <div class="">
            <span class="detail_username">cawis {{ postId }}</span>
            <span class="detail_date">34 dk önce</span>
          </div>
          <p class="post_title">Post Title</p>
        </div>
      </div>

      <div class="post_detail_content">
        <p>
          Lorem ipsum dolor sit amet consectetur, adipisicing elit. Sunt labore
          dolore rerum tempore perferendis tempora illum quibusdam nisi,
          reprehenderit deserunt, repudiandae nostrum provident adipisci quidem
          debitis vero, vel neque sint. Lorem ipsum dolor sit amet consectetur,
          adipisicing elit. Sunt labore dolore rerum tempore perferendis tempora
          illum quibusdam nisi, reprehenderit deserunt, repudiandae nostrum
          provident adipisci quidem debitis vero, vel neque sint.
        </p>
        <post-actions position="bottom" :details="true" />
      </div>

      <p class="post_comment_title">{{ `${comments.length} yorum` }}</p>
      <div class="post_detail_comments">
        <a-list
          class="comment-list"
          itemLayout="horizontal"
          :dataSource="comments"
        >
          <a-list-item slot="renderItem" slot-scope="item, index">
            <a-comment>
              <span class="comment_author" slot="author">{{
                item.author
              }}</span>
              <a-avatar class="comment_avatar" slot="avatar">FG</a-avatar>
              <p slot="content">{{ item.content }}</p>
            </a-comment>
          </a-list-item>
        </a-list>
      </div>
    </div>
  </a-modal>
</template>

<script>
import PostAcitons from "./PostActions";

export default {
  name: "PostDetail",
  props: {
    postId: {
      type: Number,
      required: true
    }
  },
  components: {
    "post-actions": PostAcitons
  },
  data() {
    return {
      showModal: false,
      comments: [
        {
          author: "Han Solo",
          content:
            "We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently."
        },
        {
          author: "Han Solo",
          content:
            "We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently."
        }
      ]
    };
  },
  created() {},
  methods: {}
};
</script>
