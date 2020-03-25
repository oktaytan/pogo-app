<template>
	<div
		class="posts_container"
		v-infinite-scroll="handleInfiniteOnLoad"
		:infinite-scroll-disabled="busy"
		:infinite-scroll-distance="10"
	>
		<a-list itemLayout="vertical" size="large" :dataSource="listData">
			<a-list-item
				slot="renderItem"
				slot-scope="item"
				key="item.title"
				class="post_item"
			>
				<a-list-item-meta :description="item.description">
					<a slot="title" :href="item.href">{{ item.title }}</a>
					<a-avatar slot="avatar" class="post_avatar">CW</a-avatar>
				</a-list-item-meta>
				{{ item.content }}
				<div class="post_actions">
					<a-tooltip placement="top">
						<template slot="title">
							<span>beğen</span>
						</template>
						<a-icon
							type="star"
							class="post_actions_icon like"
							:theme="isLiked"
							@click="liked"
							:style="{
								color: isLiked == 'filled' ? '#f1c40f' : '',
								animationName: isLiked == 'filled' ? 'bounce' : ''
							}"
						/>
					</a-tooltip>
					<a-tooltip placement="top">
						<template slot="title">
							<span>yorum yap</span>
						</template>
						<a-icon
							type="highlight"
							class="post_actions_icon comment"
							:theme="isCommented"
							@click="addComment"
							:style="{
								color: isCommented == 'filled' ? '#2ecc71' : '',
								animationName: isCommented == 'filled' ? 'bounce' : ''
							}"
						/>
					</a-tooltip>
					<a-tooltip placement="top">
						<template slot="title">
							<span>paylaş</span>
						</template>
						<a-icon
							type="share-alt"
							class="post_actions_icon share"
							@click="share"
							:style="{
								color: isShared ? '#f5f6fa' : '',
								animationName: isShared ? 'bounce' : ''
							}"
						/>
					</a-tooltip>
				</div>
			</a-list-item>
			<div v-if="loading && !busy" class="posts_loading">
				<a-spin />
			</div>
		</a-list>
	</div>
</template>

<script>
const listData = [];
for (let i = 0; i < 23; i++) {
	listData.push({
		href: 'https://www.antdv.com/',
		title: `ant design vue part ${i}`,
		avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
		description:
			'Ant Design, a design language for background applications, is refined by Ant UED Team.',
		content:
			'We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.'
	});
}

export default {
	name: 'Posts',
	data() {
		return {
			listData,
			busy: false,
			loading: false,
			isLiked: 'outlined',
			isCommented: 'outlined',
			isShared: false
		};
	},
	methods: {
		handleInfiniteOnLoad() {
			console.log('test');
		},
		liked() {
			if (this.isLiked == 'outlined') {
				this.isLiked = 'filled';
			} else {
				this.isLiked = 'outlined';
			}
		},
		addComment() {
			this.isCommented = 'filled';
		},
		share() {
			this.isShared = true;
		}
	}
};
</script>
