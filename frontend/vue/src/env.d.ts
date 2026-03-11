/// <reference types="vite/client" />

// 1. 解决引入 .vue 文件的红线报错
declare module '*.vue' {
	import type { DefineComponent } from 'vue';
	const component: DefineComponent<{}, {}, any>;
	export default component;
}

// 2. 解决引入视频/图片等静态资源的红线报错
declare module '*.mp4' {
	const src: string;
	export default src;
}
declare module '*.webm' {
	const src: string;
	export default src;
}

// 3. 解决某些没有自带 TypeScript 声明的第三方库报错 (比如 motion-v)
declare module 'motion-v';
