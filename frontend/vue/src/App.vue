<script setup lang="ts">
import HeadPage from './components/HeadPage.vue';
import BodyPage from './components/BodyPage.vue';
</script>

<template>
	<div class="layout">
		<div class="container"><HeadPage /></div>
		<BodyPage />
	</div>
	<a href="./temp/开发须知.pdf">点击下载开发须知</a>
</template>

<style scoped>
.layout {
	display: flex;
	flex-direction: column;
}
.container {
		padding: 20px;
}
</style>
