<template>
	<a-layout id="homeLayout" class="home_layout">
		<a-layout-sider
			class="home_layout_sidebar"
			collapsible
			v-model="collapsedLeft"
			:style="{ width: '250px' }"
		>
			<div class="logo">
				<router-link to="/" class="logo_link">
					<img
						src="../assets/logo.svg"
						alt=""
						:style="{ width: !collapsedLeft ? '30%' : '100%' }"
					/>
					<span v-if="!collapsedLeft">Pogo</span>
				</router-link>
			</div>
			<a-menu theme="dark" mode="inline" :defaultSelectedKeys="['1']">
				<a-menu-item key="1" @click="$router.push('/')">
					<a-icon type="home" />
					<span class="nav-text">Anasayfa</span>
				</a-menu-item>
				<a-menu-item key="2" @click="$router.push('/profile')">
					<a-icon type="user" />
					<span class="nav-text">Profilim</span>
				</a-menu-item>
				<a-menu-item key="3" @click="$router.push('/settings')">
					<a-icon type="setting" />
					<span class="nav-text">Ayarlar</span>
				</a-menu-item>
				<a-menu-item key="4" @click="logout">
					<a-icon type="logout" />
					<span class="nav-text">Çıkış Yap</span>
				</a-menu-item>
			</a-menu>
		</a-layout-sider>
		<a-layout class="home_layout_content">
			<a-layout-header class="haeder home_layout_header">
				<div class="logo">
					<router-link to="/" class="logo_link logo_link--horizontal">
						<img src="../assets/logo.svg" alt="" :style="{ width: '30%' }" />
						<span style="display: inline-block; font-size: 1.4rem;">Pogo</span>
					</router-link>
				</div>
				<a-menu theme="dark" mode="horizontal" :defaultSelectedKeys="['1']">
					<a-menu-item key="1" @click="$router.push('/')">
						<a-tooltip placement="bottom">
							<template slot="title">
								<span>Anasayfa</span>
							</template>
							<a-icon type="home" />
						</a-tooltip>
					</a-menu-item>
					<a-menu-item key="2" @click="$router.push('/profile')">
						<a-tooltip placement="bottom">
							<template slot="title">
								<span>Profilim</span>
							</template>
							<a-icon type="user" />
						</a-tooltip>
					</a-menu-item>
					<a-menu-item key="3" @click="$router.push('/settings')">
						<a-tooltip placement="bottom">
							<template slot="title">
								<span>Ayarlar</span>
							</template>
							<a-icon type="setting" />
						</a-tooltip>
					</a-menu-item>
					<a-menu-item key="4">
						<a-tooltip placement="bottom">
							<template slot="title">
								<span>Çıkış Yap</span>
							</template>
							<a-icon type="logout" />
						</a-tooltip>
					</a-menu-item>
				</a-menu>
			</a-layout-header>
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
									:min-height="0"
									:max-height="350"
									@input="writePost"
								/>
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
			<a-layout-content class="home_layout_content_view">
				<Posts v-if="$route.name == 'home'" />
				<router-view v-else></router-view>
			</a-layout-content>
		</a-layout>
		<a-layout-sider
			class="home_layout_sidebar home_layout_sidebar--left"
			:style="{ minWidth: '450px', padding: '1rem 2rem' }"
		>
			<a-row>
				<a-input-search
					placeholder="Pogo' da ara..."
					@search="onSearch"
					size="large"
					id="searchInput"
				/>
				<div class="most_liked_posts">
					<a-list :dataSource="mostLikedPosts">
						<a-list-item slot="renderItem" slot-scope="item, index">
							<p class="most_liked_posts_title">{{ item.title }}</p>
							<p class="most_liked_posts_like">{{ item.likes }} beğeni</p>
							<a-icon class="most_liked_posts_action" type="more" />
						</a-list-item>
						<div slot="header">En Çok Beğenilenler</div>
					</a-list>
				</div>
				<a-layout-footer :style="{ textAlign: 'center' }">
					{{ `&copy; ${new Date().getFullYear()} POGO, Inc.` }}
				</a-layout-footer>
			</a-row>
		</a-layout-sider>
	</a-layout>
</template>

<script>
import Posts from '../components/Posts';

export default {
	name: 'Home',
	components: {
		Posts
	},
	data() {
		return {
			collapsedLeft: false,
			newPostForm: this.$form.createForm(this, { name: 'newPost' }),
			post: {
				title: '',
				body: '',
				created_at: new Date().getTime(),
				updated_at: new Date().getTime(),
				user_id: '',
				likes: 0
			},
			disabled: true,
			mostLikedPosts: [
				{
					title: 'Title 1',
					likes: 34
				},
				{
					title: 'Title 2',
					likes: 324
				},
				{
					title: 'Title 3',
					likes: 455
				},
				{
					title: 'Title 4',
					likes: 324
				},
				{
					title: 'Title 5',
					likes: 455
				}
			]
		};
	},
	watch: {
		width(val) {
			return val > 992 ? true : false;
		}
	},
	methods: {
		onSearch() {
			console.log('search');
		},
		logout() {
			this.$router.push('/login');
		},
		writePost() {
			if (this.post.title !== '' && this.post.body !== '') {
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

<style></style>
