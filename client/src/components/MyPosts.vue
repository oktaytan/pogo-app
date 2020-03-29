<template>
	<div
		class="posts_container"
		:style="{ height: GET_TOP_BAR_SHOW ? '68vh' : '70vh' }"
	>
		<Loader v-if="loading" />
		<a-list v-else itemLayout="vertical" size="small" :dataSource="listData">
			<a-list-item
				slot="renderItem"
				slot-scope="item"
				ref="postItem"
				key="item.title"
				:style="{ padding: '1.5rem', position: 'relative' }"
				class="post_item"
			>
				<div class="click_area" @click="detailShow"></div>
				<a-list-item-meta :style="{ marginBottom: '0' }">
					<div
						class="post_user_wrap"
						slot="title"
						:style="{ marginBottom: '1.5rem', alignItems: 'center' }"
					>
						<span class="post_title">{{ item.title }}</span>
						<span class="post_date" style="margin-left: 1rem">34 dk önce</span>
						<a-popconfirm
							:style="{ float: 'right', position: 'relative' }"
							placement="topRight"
							title="Bu paylaşımı kaldırmak üzeresiniz?"
							@confirm="deletePost"
							@cancel="cancelDelete"
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
				</a-list-item-meta>
				<div style="margin-top: -1.5rem; width: 85%">
					{{ item.content }}
				</div>
			</a-list-item>
		</a-list>
	</div>
</template>

<script>
import Loader from './Loader.vue';
import PostActions from './PostActions';

import { mapGetters } from 'vuex';

const listData = [];
for (let i = 0; i < 23; i++) {
	listData.push({
		id: i,
		to: '/posts/',
		title: `Post Title ${i}`,
		avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
		description:
			'Ant Design, a design language for background applications, is refined by Ant UED Team.',
		content:
			'We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.'
	});
}

export default {
	name: 'MyPosts',
	components: {
		Loader,
		'post-actions': PostActions
	},
	data() {
		return {
			loading: true,
			listData
		};
	},
	mounted() {
		setTimeout(() => (this.loading = false), 300);
	},
	computed: {
		...mapGetters(['GET_TOP_BAR_SHOW'])
	},
	methods: {
		detailShow(e) {
			this.$refs.postDetail.showModal = true;
		},
		deletePost(e) {
			console.log(e);
			this.$message.success('Paylaşım kaldırıldı!');
		},
		cancelDelete(e) {
			console.log(e);
		}
	}
};
</script>
