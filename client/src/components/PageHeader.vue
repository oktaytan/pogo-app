<template>
	<div class="page_header">
		<div class="back_wrap">
			<a-icon
				class="back_icon"
				:type="$route.name === 'Detaylar' ? 'home' : 'left-circle'"
				@click="$router.go(-1)"
			/>
			<span class="back_text">{{ $route.name }}</span>
		</div>
	</div>
</template>

<script>
export default {
	name: 'PageHeader',
	props: {
		title: String
	}
};
</script>
