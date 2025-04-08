<template>
  <div class="common-layout">
    <el-container class="full-height">
      <!-- 左侧菜单栏 -->
      <el-aside width="200px" class="sidebar">
        <el-menu
          :default-active="activeMenu"
          router
          class="el-menu-vertical-demo"
          background-color="#001529"
          text-color="#fff"
          active-text-color="#ffd04b"
        >
          <el-menu-item index="/submit/input">
            <el-icon><DocumentAdd /></el-icon>
            产品信息录入
          </el-menu-item>
          <el-menu-item index="/submit/query">
            <el-icon><Search /></el-icon>
            溯源信息查询
          </el-menu-item>
        </el-menu>
      </el-aside>

      <!-- 右侧主内容区域 -->
      <el-container>
        <!-- 顶部面包屑 -->
        <el-header class="breadcrumb-header">
          <el-row align="middle">
            <el-col :span="1">
              <el-icon><Position /></el-icon>
            </el-col>
            <el-col :span="22">
              <el-breadcrumb separator="/">
                <el-breadcrumb-item>种植园</el-breadcrumb-item>
                <el-breadcrumb-item>{{ pageTitle }}</el-breadcrumb-item>
              </el-breadcrumb>
            </el-col>
          </el-row>
        </el-header>

        <!-- 主体内容 -->
        <el-main class="main-content">
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import { Position} from '@element-plus/icons-vue'

export default {
  name: 'SubmitLayout',
  components: {
    Position,
  },
  data() {
    return {}
  },
  computed: {
    activeMenu() {
      return this.$route.path
    },
    pageTitle() {
      if (this.$route.path.includes('/input')) return '产品信息录入'
      if (this.$route.path.includes('/query')) return '溯源信息查询'
      return ''
    }
  }
}
</script>

<style scoped>
.common-layout {
  height: 100vh;
}

.full-height {
  height: 100%;
}

.sidebar {
  background-color: #001529;
  color: #fff;
  height: 100vh;
  overflow: auto;
}

.breadcrumb-header {
  padding: 16px;
  background-color: transparent;
  box-shadow: none;
  border-bottom: none;
}

.main-content {
  padding: 20px;
  background-color: #f9f9f9;
  min-height: calc(100vh - 64px); /* 减去 header 高度 */
}

.el-menu {
  border: none;
  border-right-width: 0;
}
</style>
