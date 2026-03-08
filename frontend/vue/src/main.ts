import { createApp } from 'vue'
// 创建整个vue3的仪式
// 从vue中，拿出createAPP这个东西
// 如果不加花括号，就是'默认导出'，即导出默认的东西
import './style.css'
// 引入样式文件
import App from './App.vue'
// 将项目的 根组件 拿进来
// 每个模块有很多命名导出，但是只有一个默认导出

createApp(App).mount('#app')
// 以App这个东西为核心，在#app这个模块构建App.vue的内容