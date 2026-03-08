<script setup lang="ts">
defineOptions({ name: 'HeadPage' });
</script>

<template>
	<div class="background">
		<h1>SCUEC-ICPC 活动主页</h1>
		<div class="inf">该网页由ICPC开发部负责</div>
	</div>
</template>

<style scoped>
.background {
	background-color: rgb(160, 178, 194);
	box-shadow: 0 0 3px;
	border-radius: 10px;
	padding: 20px;
}

.background h1 {
	font-size: 30px;
	text-align: left;
}

.inf {
	display: flex;
	box-shadow: 0 0 5px purple;
	color: rgb(208, 1, 1);
	font-size: 20px;
	border-radius: 20px;
	align-items: center;
	justify-content:center;
	/* justify居中 */
}
</style>
