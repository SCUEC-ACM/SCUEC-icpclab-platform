<script setup lang="ts">
defineOptions({ name: 'BodyPage' });
</script>

<template>
	<hr class="line" />
	<div class="info-row">
		<div class="left">
			<strong>这是实验室的主页</strong>
			<br />
			SCUEC-ICPC现在[2026/03/08]有三个部门：
			<ul>
				<li>训练部</li>
				<li>开发部</li>
				<li>后勤部</li>
			</ul>
		</div>
		<div class="devider"></div>
		<div class="right">
			<strong>开发部当前的目标：</strong>
			<ul>
				<li>绘制主页</li>
				<li>添加“公告”页面</li>
				<li>搭建后端框架</li>
				<li>实现登录功能</li>
			</ul>
		</div>
	</div>
</template>

<style>
.line {
	border: none;
	border-top: 1px solid rgba(0, 0, 0, 0.3);
	margin: 30px;
}
.info-row {
	display: grid;
	grid-template-columns: 1fr 1px 1fr;
	/* 将页面元素分为三块，左右各占权值1，中间的线占用1像素 */
	align-items: center;
	width: 100%;
	column-gap: 16px;
}
.devider {
	height: 95%;
	background: rgba(0, 0, 0, 0.25);
}
.left {
	margin: 5px;
	padding: 15px;
	background-color: rgba(0, 0, 0, 0.05);
	border: 1px solid black;
	border-radius: 10px;
	text-align: start;
	/* 文字按阅读的起始方向对齐 */
	min-width: 0;
	/* 允许元素比内容更窄，就可以对元素换行了 */
	overflow-wrap: break-word;
	/* 长文字自动换行 */
}
.right {
	margin: 5px;
	padding: 15px;
	background-color: rgba(0, 0, 0, 0.05);
	border: 1px solid black;
	border-radius: 10px;
	text-align: start;
	/* 文字按阅读的起始方向对齐 */
	min-width: 0;
	/* 允许元素比内容更窄，就可以对元素换行了 */
	overflow-wrap: break-word;
	/* 长文字自动换行 */
}
</style>
