<script setup lang="ts">
import { ref } from 'vue';
import HeadPage from './components/HeadPage.vue';
import BodyPage from './components/BodyPage.vue';
import Curtain from './components/Raise-the-curtain.vue';

// 控制开场动画的开关，默认开启
const showIntro = ref(true);
</script>

<template>
	<Transition name="fade">
		<Curtain v-if="showIntro" @close="showIntro = false" />
	</Transition>

	<div class="layout">
		<div class="container"><HeadPage /></div>
		<BodyPage />
	</div>
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
