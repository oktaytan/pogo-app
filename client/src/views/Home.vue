<template>
	<div class="home">
		<h1>Home page</h1>
	</div>
</template>

<script>
import { ax } from '../utils/settings';
export default {
	name: 'Home',
	components: {},
	data() {
		return {
			posts: []
		};
	},
	mounted() {
		this.getAllPosts();
	},
	methods: {
		getAllPosts() {
			ax.get('/posts')
				.then((res) => console.log(res.data))
				.catch((err) => console.log(err));
		}
	}
};
</script>
