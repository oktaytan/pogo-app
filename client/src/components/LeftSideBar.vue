<template>
	<a-layout-sider
		class="home_layout_sidebar"
		collapsible
		v-model="collapsed"
		:style="{ width: '250px' }"
	>
		<div class="logo">
			<router-link to="/" class="logo_link">
				<img
					src="../assets/logo.svg"
					alt=""
					:style="{ width: !collapsed ? '30%' : '100%' }"
				/>
				<span v-if="!collapsed">Pogo</span>
			</router-link>
		</div>
		<a-menu
			theme="dark"
			mode="inline"
			:defaultSelectedKeys="['1']"
			:selectedKeys="[key]"
		>
			<a-menu-item key="1" @click="() => changeMenu('/')">
				<a-icon type="home" />
				<span class="nav-text">Anasayfa</span>
			</a-menu-item>
			<a-menu-item key="2" @click="() => changeMenu('/profile')">
				<a-icon type="user" />
				<span class="nav-text">Profilim</span>
			</a-menu-item>
			<a-menu-item key="3" @click="() => changeMenu('/settings')">
				<a-icon type="setting" />
				<span class="nav-text">Ayarlar</span>
			</a-menu-item>
			<a-menu-item key="4" @click="$router.push('/logout')">
				<a-icon type="logout" />
				<span class="nav-text">Çıkış Yap</span>
			</a-menu-item>
		</a-menu>
	</a-layout-sider>
</template>

<script>
export default {
	name: 'LeftSideBar',
	data() {
		return {
			collapsed: false,
			key: localStorage.getItem('pogo_key')
				? JSON.parse(localStorage.getItem('pogo_key'))
				: '1'
		};
	},
	created() {
		this.routeChange();
	},
	watch: {
		$route(val) {
			this.routeChange(val.name);
		}
	},
	methods: {
		changeMenu(path) {
			this.$router.push(path);
			this.routeChange();
			localStorage.setItem('pogo_key', JSON.stringify(this.key));
		},
		routeChange() {
			switch (this.$route.name) {
				case 'Anasayfa':
					this.key = '1';
					break;
				case 'Profilim':
					this.key = '2';
					break;
				case 'Ayarlar':
					this.key = '3';
					break;
				default:
					this.key = '1';
			}
		}
	}
};
</script>
