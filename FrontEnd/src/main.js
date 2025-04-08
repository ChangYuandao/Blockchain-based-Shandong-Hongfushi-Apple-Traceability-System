import { createApp } from 'vue';
import App from './App.vue';
import router from './router';  // 导入路由
import ElementPlus from 'element-plus';  // 导入 ElementPlus
import 'element-plus/dist/index.css';  // 导入 ElementPlus 样式

// 创建 Vue 应用实例
const app = createApp(App);

// 关闭 Vue DevTools（如果你不希望开启）
app.config.devtools = false;

// 使用路由和 ElementPlus 插件
app.use(router); 
app.use(ElementPlus);

// 挂载 Vue 应用
app.mount('#app');
