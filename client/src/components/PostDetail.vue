<template>
	<div style="postion: relative">
		<!-- Sayfa headerı -->
		<page-header title="Paylaşım detayları" v-if="!GET_TOP_BAR_SHOW" />

		<Loader v-if="loading" />
		<!-- Gönderi Detayları -->
		<a-card
			v-else
			class="post_detail_wrap"
			:style="{ height: GET_TOP_BAR_SHOW ? '73vh' : '80vh' }"
		>
			<div class="post_detail">
				<div class="post_detail_header">
					<!-- Gönderinin yazarının avatarı -->
					<a-avatar class="post_avatar">{{
						GET_POST_BY_ID.author.username | userAvatar
					}}</a-avatar>
					<div>
						<div>
							<!-- Yazarın kullanıcı adı -->
							<span class="detail_username">{{
								GET_POST_BY_ID.author.username
							}}</span>
							<!-- Gönderinin oluşturulma zamanı -->
							<span class="detail_date">{{
								$moment(GET_POST_BY_ID.created_at).fromNow()
							}}</span>
						</div>
						<!-- Gönderi başlığı -->
						<p class="post_title">{{ GET_POST_BY_ID.title }}</p>
					</div>
				</div>
				<!-- Gönderi gövdesi -->
				<div class="post_detail_content">
					<p>
						{{ GET_POST_BY_ID.body }}
					</p>
					<!-- Gönderiye ait action' lar ( mobile ekranlar için )-->
					<post-actions
						ref="postActions"
						:myPost="
							GET_POST_BY_ID.author.username === GET_USER.username
								? true
								: false
						"
						position="bottom"
						class="mobile_post_actions"
						:details="true"
						:post="GET_POST_BY_ID"
						@addedComment="fetchPost"
						@likedPost="fetchPost"
					/>
				</div>
				<!-- Gönderiye ait action' lar ( geniş ekranlar için )-->
				<post-actions
					ref="postActions"
					:myPost="
						GET_POST_BY_ID.author.username === GET_USER.username ? true : false
					"
					position="top"
					class="desktop_post_actions"
					:details="true"
					:post="GET_POST_BY_ID"
					@addedComment="fetchPost"
					@likedPost="fetchPost"
				/>
				<!-- Gönderiye ait yorumlar -->
				<p class="post_comment_title">{{ commentCount }}</p>
				<div class="post_detail_comments">
					<a-list
						class="comment-list"
						itemLayout="horizontal"
						:dataSource="comments"
					>
						<a-empty
							v-if="comments.length == 0"
							:image="noCommentImg"
							:style="{
								height: '100%',
								marginTop: '1rem'
							}"
							class="empty_wrap"
						>
							<span slot="description" class="empty_desc">
								<span>Henüz yorum yapılmamış</span>
								<!-- Yeni yorum eklemek için moal açılacak -->
								<a @click="openCommentModal">İlk sen yap!</a>
							</span>
						</a-empty>
						<!-- Yorum -->
						<a-list-item
							slot="renderItem"
							slot-scope="item"
							:style="{ position: 'relative' }"
						>
							<a-comment>
								<!-- Yorumu yapanın kullanıcı adı -->
								<span class="comment_author" slot="author">{{
									item.author.username
								}}</span>
								<!-- Yorumu yapanın avatarı -->
								<a-avatar class="comment_avatar" slot="avatar">{{
									item.author.username | userAvatar
								}}</a-avatar>
								<!-- Yorumun içeriği -->
								<p slot="content" class="comment_detail">
									{{ item.comment }}
								</p>
							</a-comment>

							<!-- Yoeum kullanıcın kendi yorumu ise kaldırma ikonu -->
							<a-dropdown
								v-if="item.author.id == GET_USER.id"
								:trigger="['click']"
								placement="bottomRight"
								:style="{
									float: 'right',
									position: 'absolute',
									top: '30%',
									right: '0.5rem',
									transfrom: 'translateY(-50%)'
								}"
							>
								<a-icon type="more" class="delete_comment" />
								<a-menu slot="overlay">
									<a-menu-item
										key="0"
										@click="() => deleteComment(item.author.id, item.id)"
									>
										<a-icon type="delete" theme="filled" />
										<span>Kaldır</span>
									</a-menu-item>
								</a-menu>
							</a-dropdown>
						</a-list-item>
					</a-list>
				</div>
			</div>
		</a-card>
	</div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

import PostActions from './PostActions';
import Loader from './Loader.vue';
import PageHeader from './PageHeader';
import noCommentImg from '../assets/no-comments.png';

export default {
	name: 'PostDetail',
	props: {
		myPost: {
			type: Boolean
		}
	},
	components: {
		Loader,
		'post-actions': PostActions,
		'page-header': PageHeader
	},
	data() {
		return {
			loading: true,
			comments: [],
			commentCount: '',
			noCommentImg
		};
	},
	mounted() {
		// Gönderi alınıyor
		this.fetchPost();
	},
	computed: {
		...mapGetters(['GET_TOP_BAR_SHOW', 'GET_POST_BY_ID', 'GET_USER'])
	},
	methods: {
		...mapActions(['FETCH_POST_BY_ID', 'DELETE_COMMENT']),
		// Yorum yapmak için modalı açılıyor
		openCommentModal() {
			this.$refs.postActions.openCommentModal();
		},
		// Gönderiyi getirecek action tetikleniyor
		fetchPost() {
			this.FETCH_POST_BY_ID(this.$route.params.id)
				.then((data) => {
					this.loading = false;
					// Gönderiye ait yorumlar alınıypr
					this.comments =
						this.GET_POST_BY_ID.comments != null
							? this.GET_POST_BY_ID.comments
							: [];
					// Yorum sayısı belirleniyor
					this.commentCount =
						this.comments.length > 0 ? `${this.comments.length} yorum` : '';
				})
				.catch((err) => {
					this.$message
						.error('Böyle bir gönderi bulunmamaktadır', 1.5)
						.then(() =>
							this.$message.loading("Anasayfa'ya yönlendiriliyorsunuz!", 1.5)
						)
						.then(() => {
							this.$router.push('/');
						});
				});
		},
		// Yorum kullanıcının kendisi ise silmek için action tetikleniyor
		deleteComment(author_id, comment_id) {
			if (author_id === this.GET_USER.id) {
				this.DELETE_COMMENT(comment_id)
					.then((res) => {
						this.$message.success('Yorum kaldırıldı!');
						this.fetchPost();
					})
					.catch((err) => this.$message.error(err.message));
			} else {
				return false;
			}
		}
	}
};
</script>
