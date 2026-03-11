<script setup lang="ts">
defineOptions({ name: 'HeadPage' });
</script>

<template>
    <div class="header-container">
        <h1 class="title">SCUEC-ICPC <span>活动主页</span></h1>
        <div class="info-tag">该网页由 ICPC 开发部负责</div>
    </div>
</template>

<style scoped>
/* 外部容器：纯白底色，锋利直角，左右排版 */
.header-container {
    display: flex;
    justify-content: space-between; /* 让标题和信息标签分居左右两侧 */
    align-items: center;
    background-color: #ffffff; /* 纯白底色，最经典的浅色 */
    padding: 30px 40px;
    border: 1px solid #e5e5e5; /* 极细的灰色边框，提供锋利感 */
    border-radius: 0; /* 核心：舍弃圆角，使用绝对的直角 */
    border-left: 6px solid #1a1a1a; /* 左侧加入黑色粗实线，增加视觉重心和现代感 */
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.03); /* 极度柔和的微阴影，让它像一张高级卡片浮在页面上 */
}

/* 标题样式：拉大字间距，粗细结合 */
.title {
    font-size: 28px;
    font-weight: 800; /* 主体加粗 */
    color: #111111; /* 使用深灰黑代替纯黑，更柔和高级 */
    margin: 0;
    letter-spacing: 1.5px; /* 适当拉开字间距是产生“高级感”的秘诀 */
}

.title span {
    font-weight: 300; /* 副标题变细，形成强烈的粗细对比 */
    color: #666666;
}

/* 底部/右侧信息标签：极简的线框标签 */
.info-tag {
    font-size: 13px;
    color: #555555;
    background-color: #f9f9f9;
    padding: 8px 16px;
    border: 1px solid #dddddd;
    border-radius: 0; /* 同样使用直角，保持锋利统一 */
    letter-spacing: 1px;
    font-weight: 500;
}
</style>