<template>
	<div class="intro-screen" @click="closeAnimation" ;>
		<motion.video
			src="./Development.mp4"
			autoplay
			muted
			playsinline
			class="center-image"
			:initial="{ opacity: 0, y: 50 }"
			:animate="{
				opacity: 1,
				y: 0,
				transition: { duration: 1.5, type: 'spring' },
			}"
			@ended="closeAnimation"
		></motion.video>
	</div>
</template>

<script setup>
import { motion } from 'motion-v';

const emit = defineEmits(['close']);

const closeAnimation = () => {
	emit('close');
};

defineOptions({ name: 'Curtain' });
</script>

<style scoped>
.intro-screen {
	position: fixed;
	top: 0;
	left: 0;
	width: 100vw;
	height: 100vh;
	background: radial-gradient(
		circle at center,
		#2a0845 0%,
		#1a0b2e 40%,
		#000000 100%
	);
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
	z-index: 9999;
	cursor: pointer;
}

.center-image {
	width: 400px;
	height: auto;
	transition:
		transform 0.3s ease,
		box-shadow 0.3s ease;
	mix-blend-mode: screen;
}

.center-image:hover {
	transform: scale(1.03);
}
</style>
