import { createRouter, createWebHistory } from "vue-router";
import LoginPage from "@/pages/LoginPage.vue";
import SubmitPage from "@/pages/Submit/SubmitPage.vue";
import ProcessPage from "@/pages/Process/ProcessPage.vue";
import AuditPage from "@/pages/AuditPage.vue";
import MonitorPage from "@/pages/MonitorPage.vue";

const routes = [
  { path: "/", redirect: "/login" },
  { path: "/login", component: LoginPage },
  {
    path: "/submit",
    component: SubmitPage,
    meta: { requiresRole: "种植园" },
    children: [
      {
        path: "input", // 实际路由为 /submit/input
        component: () => import("@/pages/Submit/SubmitInput.vue"),
      },
      {
        path: "query",
        component: () => import("@/pages/Submit/SubmitQuery.vue"),
      },
      {
        path: "",
        redirect: "/submit/input",
      },
    ],
  },
  {
    path: "/process",
    component: ProcessPage,
    meta: { requiresRole: "物流企业" }, // 需要的角色是 "processing" 或 "logistics"
    children: [
      {
        path: "input", // 实际路由为 /process/input
        component: () => import("@/pages/Process/ProcessInput.vue"),
      },
      {
        path: "query",
        component: () => import("@/pages/Process/ProcessQuery.vue"),
      },
      {
        path: "",
        redirect: "/process/input",
      },
    ],
  },
  { path: "/audit", component: AuditPage, meta: { requiresRole: "管理员" } },
  {
    path: "/monitor",
    component: MonitorPage,
    meta: { requiresRole: "消费者" },
  },
  { path: "/:pathMatch(.*)*", redirect: "/login" }, // 使用新的 catch-all 路由语法
];

const router = createRouter({
  history: createWebHistory(), // 使用 HTML5 History 模式
  routes,
});

// 全局路由守卫
router.beforeEach((to, from, next) => {
  const role = localStorage.getItem("role"); // 假设角色保存在本地存储中

  if (to.meta.requiresRole) {
    const allowedRoles = to.meta.requiresRole.split(",");
    if (allowedRoles.includes(role)) {
      next(); // 如果角色允许访问，继续路由跳转
    } else {
      next("/login"); // 如果角色不允许，跳转到登录页
    }
  } else {
    next(); // 无需角色控制的页面，直接跳转
  }
});

export default router;
