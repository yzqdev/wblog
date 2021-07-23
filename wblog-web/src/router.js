
import Home from "@/components/Home.vue";
import {
  createRouter,
  createWebHistory,
} from "vue-router";
import MainMenu from "@/view/MainMenu.vue";
import signin from "@/view/signin.vue";
import Signup from "@/view/Signup.vue";
import Links from "@/view/Links.vue";
import About from "@/view/About.vue";
import Rank from "@/view/Rank.vue";
import PageIndex from "@/view/PageIndex.vue";
import BookList from "@/view/BookList.vue";
import AdminHome from "@/view/admin/AdminHome.vue";
import AdminIndex from "@/view/admin/AdminIndex.vue";
import Users from "@/view/admin/Users.vue";
import Test from "@/view/Test.vue";

const routes = [
  {
    path: "/",
    redirect: "/main/index",
  },
  {
    path:'/test',
    component: Test
  },
  {
    path: "/admin",
    redirect: "/admin/index",
  },
  {
    path: "/admin",
    name: "admin",
    component: AdminHome,
    children: [
      { path: "index", name: "adminIndex", component: AdminIndex },
      { path: "user/list", name: "users", component: Users,meta:'所有用户' },
    ],
  },
  {
    path: "/main",
    name: "main",
    component: MainMenu,
    children: [
      { path: "about", name: "about", component: About },
      {
        path: "index",
        name: "index",
        component: PageIndex,
        meta: { title: "首页" },
      },
      { path: "rank", name: "rank", component: Rank },
      { path: "book", name: "book", component: BookList },
      {
        path: "links",
        name: "links",
        component: Links,
      },
    ],
  },
  {
    path: "/signin",
    name: "signin",
    component: signin,
  },
  {
    path: "/signup",
    name: "signup",
    component: Signup,
  },

  {
    path: "/home",
    component: Home,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});
export default router;
