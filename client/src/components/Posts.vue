<template>
	<div
		class="posts_container"
		:style="{ height: GET_TOP_BAR_SHOW ? '78vh' : '100vh' }"
	>
		<!-- Yeni eklenecek olan gönderi için loader yüklemesi -->
		<loader-skeleton
			v-if="GET_NEW_POST_LOADING"
			class="animated fadeIn"
			:rows="2"
		/>
		<loader v-if="loading" />
		<!-- Tüm gönderilerin listesi -->

		<a-list
			v-else
			itemLayout="vertical"
			size="large"
			:dataSource="GET_ALL_POSTS"
		>
			<a-list-item
				slot="renderItem"
				slot-scope="post"
				ref="postItem"
				key="post.id"
				class="post_item"
				@click="() => detailShow(post.id)"
			>
				<a-list-item-meta>
					<div class="post_user_wrap" slot="title">
						<div>
							<!-- Gönderi yazarının kullanıcı adı -->
							<span class="post_username">{{ post.author.username }}</span>
							<!-- Gönderinin oluşturulma zamanı -->
							<span class="post_date">{{
								$moment(post.created_at).fromNow()
							}}</span>
						</div>
						<!-- Gönderinin başlığı -->
						<span class="post_title">{{ post.title }}</span>
					</div>
					<!-- Gönderi yazarının avatarı -->
					<a-avatar slot="avatar" class="post_avatar">{{
						post.author.username | userAvatar
					}}</a-avatar>
				</a-list-item-meta>
				<!-- Gönderi gövdesi -->
				{{ post.body.slice(0, 200) }}...
				<a-checkable-tag>devamı</a-checkable-tag>
				<!-- Gönderi action alanı -->
				<post-actions
					position="top"
					:post="post"
					:details="false"
					@like="getData"
					:myPost="post.author.username === GET_USER.username ? true : false"
				/>
			</a-list-item>
		</a-list>
	</div>
</template>

<script>
import Loader from './Loader';
import LoaderSkeleton from './LoaderSkeleton';
import PostActions from './PostActions';

import { mapGetters, mapActions } from 'vuex';

export default {
	name: 'Posts',
	components: {
		loader: Loader,
		'loader-skeleton': LoaderSkeleton,
		'post-actions': PostActions
	},
	data() {
		return {
			posts: [],
			loading: true
		};
	},
	mounted() {
		this.loading = true;
		// Tüm gönderi alınıyor
		this.getData();
	},
	computed: {
		...mapGetters([
			'GET_TOP_BAR_SHOW',
			'GET_ALL_POSTS',
			'GET_USER',
			'GET_NEW_POST_LOADING',
			'GET_SOCKET_MESSAGE'
		])
	},
	methods: {
		...mapActions(['FETCH_ALL_POSTS']),
		// Tüm gönderi alacak action tetikleniyor
		async getData() {
			this.FETCH_ALL_POSTS()
				.then((data) => {
					this.loading = this.GET_ALL_POSTS.length > 0 && false;
				})
				.catch((err) =>
					this.$message.error('Bir hata oluştu! Lütfen sayfayı yenileyiniz.')
				);
		},
		// Gönderiye tıklandığında ona ait detay sayfasına gidiyor
		detailShow(id) {
			this.$router.push({ name: 'Detaylar', params: { id } });
		}
	}
};
</script>
