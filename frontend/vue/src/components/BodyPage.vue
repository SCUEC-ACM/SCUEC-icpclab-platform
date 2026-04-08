<script setup lang="ts">
import { ref } from 'vue';

defineOptions({ name: 'BodyPage' });

// 使用 activeTab 来记录当前点击了哪个按钮。null 表示全都没点（收起状态）
const activeTab = ref<string | null>(null);

// 切换面板的函数：如果点击的是当前已打开的按钮，就收起；否则打开新面板
const toggleTab = (tabName: string) => {
	if (activeTab.value === tabName) {
		activeTab.value = null;
	} else {
		activeTab.value = tabName;
	}
};
</script>

<template>
	<div class="body-container">
		<div class="button-grid">
			<button
				class="action-btn"
				:class="{ active: activeTab === 'notice' }"
				@click="toggleTab('notice')"
			>
				<span class="btn-text">公告</span>
				<span class="btn-icon">{{ activeTab === 'notice' ? '−' : '+' }}</span>
			</button>

			<button
				class="action-btn"
				:class="{ active: activeTab === 'resource' }"
				@click="toggleTab('resource')"
			>
				<span class="btn-text">资源</span>
				<span class="btn-icon">{{ activeTab === 'resource' ? '−' : '+' }}</span>
			</button>
		</div>

		<div class="expanded-panel" v-show="activeTab !== null">
			<div v-show="activeTab === 'notice'" class="tab-content">
				<p class="notice-alert">
					【2026/03/10】网站前端样式正在全面升级，后端 Gin
					框架准备接入中，请各位成员留意后续更新！
				</p>

				<div class="info-row">
					<div class="panel">
						<h2 class="panel-title">实验室主页</h2>
						<p class="panel-desc">SCUEC-ICPC 现在 [2026/03/08] 有三个部门：</p>
						<ul class="clean-list">
							<li>训练部</li>
							<li>开发部</li>
							<li>后勤部</li>
						</ul>
					</div>

					<div class="devider"></div>

					<div class="panel">
						<h2 class="panel-title">开发部当前目标</h2>
						<ul class="clean-list mt-top">
							<li>绘制主页</li>
							<li>添加“公告”页面</li>
							<li>搭建后端框架</li>
							<li>实现登录功能</li>
						</ul>
					</div>
				</div>
			</div>

			<div v-show="activeTab === 'resource'" class="tab-content">
				<p class="resource-desc">
					这里是内部资料与文件下载区。目前提供《开发须知》文档下载：
				</p>

				<a href="./开发须知.md" download class="download-btn">
					点击下载开发须知.md
					<span class="download-icon">↓</span>
				</a>

				<a href="./核心功能的梳理.md" download class="download-btn">
					点击下载核心功能的梳理.md
					<span class="download-icon">↓</span>
				</a>
			</div>
		</div>
	</div>
</template>

<style scoped>
/* 整体容器 */
.body-container {
	margin-top: 24px;
	display: flex;
	flex-direction: column;
	gap: 16px; /* 按钮和下方内容区的间距 */
}

/* ================= 1. 按钮网格布局 ================= */
.button-grid {
	display: grid;
	/* 核心逻辑：强制划分 4 列，多出来的按钮会自动换到下一行 */
	grid-template-columns: repeat(4, 1fr);
	gap: 16px;
}

/* 按钮基础样式：纯白底，极细边框，锋利直角 */
.action-btn {
	display: flex;
	justify-content: space-between;
	align-items: center;
	background-color: #ffffff;
	border: 1px solid #e5e5e5;
	padding: 16px 20px;
	cursor: pointer;
	border-radius: 0; /* 极致的直角 */
	transition: all 0.2s ease;
	outline: none;
}

/* 按钮悬浮效果：边框加深 */
.action-btn:hover {
	border-color: #111111;
}

/* 按钮激活状态：黑底白字，充满高级感 */
.action-btn.active {
	background-color: #111111;
	border-color: #111111;
	color: #ffffff;
}

.btn-text {
	font-size: 16px;
	font-weight: 600;
	letter-spacing: 2px;
}

.btn-icon {
	font-size: 20px;
	font-weight: 300;
	color: #888888;
}

/* 激活状态下的加减号颜色变成白色 */
.action-btn.active .btn-icon {
	color: #ffffff;
}

/* ================= 2. 底部统一的展开面板 ================= */
.expanded-panel {
	background-color: #ffffff;
	border: 1px solid #e5e5e5;
	border-radius: 0;
	padding: 32px;
}

.tab-content {
	display: flex;
	flex-direction: column;
	gap: 24px;
}

/* 顶部公告提示语 */
.notice-alert {
	margin: 0;
	padding: 12px 16px;
	background-color: #f9f9f9;
	border-left: 4px solid #111111;
	font-size: 14px;
	color: #333333;
}

/* ================= 3. 嵌入的双列布局面板 ================= */
.info-row {
	display: grid;
	grid-template-columns: 1fr 1px 1fr;
	align-items: stretch;
}

.devider {
	background: #e5e5e5;
	width: 1px;
}

.panel {
	padding: 0 24px;
}
.panel:first-child {
	padding-left: 0;
}

.panel-title {
	font-size: 18px;
	font-weight: 700;
	color: #111111;
	margin: 0 0 16px 0;
	letter-spacing: 1px;
}

.panel-desc {
	font-size: 14px;
	color: #666666;
	margin-bottom: 12px;
}

.clean-list {
	list-style: none;
	padding: 0;
	margin: 0;
}
.clean-list.mt-top {
	margin-top: 8px;
}
.clean-list li {
	position: relative;
	padding-left: 16px;
	margin-bottom: 10px;
	color: #333333;
	font-size: 14px;
}
.clean-list li::before {
	content: '';
	position: absolute;
	left: 0;
	top: 6px;
	width: 4px;
	height: 4px;
	background-color: #111111;
}

/* ================= 4. 资源下载按钮 ================= */
.resource-desc {
	font-size: 15px;
	color: #555555;
	margin: 0;
}

/* 重新设计的锋利型下载按钮 */
.download-btn {
	display: inline-flex;
	align-items: center;
	justify-content: space-between;
	width: fit-content;
	gap: 20px;
	padding: 12px 24px;
	background-color: #ffffff;
	border: 1px solid #111111;
	color: #111111;
	text-decoration: none;
	font-weight: 600;
	font-size: 14px;
	border-radius: 0;
	transition: all 0.2s ease;
}

.download-btn:hover {
	background-color: #111111;
	color: #ffffff;
}

.download-icon {
	font-size: 16px;
}
</style>
