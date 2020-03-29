<template>
	<div class="profile_wrap">
		<page-header title="Profilim" />

		<a-card hoverable :bordered="false" class="profile_card">
			<a-card-meta title="cawis" description="zidak@test.com">
				<a-avatar class="profile_avatar" slot="avatar">CW</a-avatar>
			</a-card-meta>
			<div class="profile_header">
				<a-tag color="#1e5181" class="post_count">4 gönderi</a-tag>
				<a-tag color="#1e5181" class="like_count">24 beğeni</a-tag>
				<a-tag color="#1e5181" class="comment_count">32 yorum</a-tag>
			</div>
			<a-button
				class="openModalBtn"
				type="primary"
				ghost
				size="large"
				@click="openPostModal"
			>
				Pogo ile paylaş
			</a-button>
		</a-card>
		<a-modal
			title="Pogo ile paylaş"
			centered
			:visible="newPostModal"
			class="new_post_modal"
			@cancel="cancelPost"
		>
			<a-form
				:form="newPostFormFromProfile"
				class="post_form"
				@submit.prevent="addPost"
			>
				<a-avatar class="header_avatar">CW</a-avatar>
				<div class="post_form_items">
					<div class="post_form_items_data">
						<a-form-item>
							<a-input
								ref="profile_page_title_input"
								class="new_post_input"
								:style="{ marginBottom: '1rem' }"
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
				</div>
			</a-form>

			<template slot="footer">
				<a-button key="back" @click="cancelPost">Vazgeç</a-button>
				<a-button
					key="submit"
					:class="disabled && 'disabled_btn'"
					type="primary"
					:loading="loading"
					@click="addPost"
					:disabled="disabled"
				>
					Paylaş
				</a-button>
			</template>
		</a-modal>
		<div class="profile_shared">
			<span class="shared_title">Paylaşımlarım</span>

			<a-input-search
				class="my_post_search"
				placeholder="Ara..."
				@search="onSearch"
				v-model="searchText"
				size="large"
				autoComplete="off"
			/>
		</div>
		<my-posts />
	</div>
</template>
<script>
import { mapGetters } from 'vuex';
import MyPosts from './MyPosts';
import PageHeader from './PageHeader';

export default {
	name: 'Profile',
	components: {
		'my-posts': MyPosts,
		'page-header': PageHeader
	},
	data() {
		return {
			loading: false,
			disabled: true,
			newPostModal: false,
			newPostFormFromProfile: this.$form.createForm(this, { name: 'newPost' }),
			post: {
				title: '',
				body: '',
				created_at: new Date().getTime(),
				updated_at: new Date().getTime(),
				user_id: '',
				likes: 0
			},
			disabled: true,
			searching: false,
			searchText: ''
		};
	},
	computed: {
		...mapGetters(['GET_TOP_BAR_SHOW']),
		maxLength() {
			return parseInt(this.post.body.split('').length - 1);
		}
	},
	methods: {
		openPostModal() {
			this.newPostModal = true;
			setTimeout(() => {
				this.$refs.profile_page_title_input.$el.focus();
			}, 300);
		},
		writePost() {
			if (this.post.title !== '' && this.post.body !== '') {
				this.disabled = false;
			} else {
				this.disabled = true;
			}
		},
		addPost() {
			this.loading = true;
			setTimeout(() => {
				this.loading = false;
				console.log(this.post);
			}, 1000);
		},
		cancelPost() {
			(this.post.title = ''),
				(this.post.body = ''),
				(this.newPostModal = false);
		},
		onSearch() {
			if (this.searchText != '') {
				this.searching = true;
			} else {
				this.searching = false;
			}
		},
		cancelSearch() {
			this.searching = false;
			this.searchText = '';
		}
	}
};
</script>
