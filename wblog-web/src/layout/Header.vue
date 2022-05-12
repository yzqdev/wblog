<template>
  <header class="header">
    <el-link style="margin-left: 0.5rem" v-html="notice"></el-link>
    <span class="header-right">
      <el-dropdown trigger="click">
        <el-badge v-if="comments.length > 0" :value="comments.length">
          <el-icon :size="20">
            <bell /> </el-icon
        ></el-badge>
        <template #dropdown>
          <section class="notice-tab">
            <article class="notice-head">
            <div class="text-base flex-1">通知</div>
            <div class="text-base flex-1 cursor-pointer text-right  text-primary" @click="readAll">全部已读</div>
            </article>
            <article class="notice-item" v-for="item in comments">
              <div class="m-2">
                <el-avatar
                  size="large"
                  src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
                ></el-avatar>
              </div>
              <div
             class="flex flex-col justify-center"
              >
                <div>
                  <el-link href="hh">{{ item.content }}</el-link>
                </div>
                <div>{{ formatDate(item.created_at) }}</div>
              </div>
            </article>
          </section>
          <section class="border-2 border-gray-100"><el-button class=" w-full" type="primary" text>查看所有</el-button></section>
        </template>
      </el-dropdown>
      <el-icon @click="toggle" :size="20">
        <full-screen />
      </el-icon>
      <el-icon @click="showSettings" :size="20">
        <setting />
      </el-icon>
      <el-dropdown>
        <span class="el-dropdown-link">
          <el-avatar
            :size="30"
            src="https://img-static.mihoyo.com/communityweb/upload/222b847170feb3f2babcc1bd4f0e30dd.png"
          />
          <el-button text class="ml-2"
            >{{ userInfo.nickname ? userInfo.nickname : userInfo.username
            }}<el-icon> <arrow-down /></el-icon
          ></el-button>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>
              <el-link @click="gotoRoute('profile')" :underline="false"
                >个人中心
              </el-link>
            </el-dropdown-item>
            <el-dropdown-item>
              <el-link
                href="https://github.com/yzqdev/slim-admin"
                target="_blank"
                :underline="false"
                >项目地址
              </el-link>
            </el-dropdown-item>
            <el-dropdown-item divided @click="logout"
              >退出登录</el-dropdown-item
            >
          </el-dropdown-menu>
        </template>
      </el-dropdown></span
    >
    <el-drawer
      size="20%"
      v-model="settingDraw"
      title="主题配置"
      direction="rtl"
    >
      <el-row>
        <el-col :span="24">
          <el-switch
            v-model="theme.contentPadding"
            size="large"
            active-text="显示边距"
          />
        </el-col>
        <el-col :span="24">
          <el-switch
            v-model="theme.showFooter"
            size="large"
            active-text="显示footer"
          />
        </el-col>
      </el-row>
    </el-drawer>
  </header>
</template>

<script setup lang="ts">
import { ArrowDown, Bell, FullScreen, Setting } from "@element-plus/icons-vue";
import { useRouter } from "vue-router";
import { onBeforeMount, watch } from "vue";
import { useThemeStore } from "@/store/themeConfig";
import { defaultTheme } from "@/constants/defaultTheme";
import { getCommentsUnreadApi, getUserInfoApi } from "@/utils/apis";
import { useUserStore } from "@/store/user";
import { UserState } from "@/type/storeTypes";
import { formatDate } from "@/utils/util";
let router = useRouter();
let { setThemeConfig } = useThemeStore();
let userStore = useUserStore();
let userInfo: UserState = computed(() => {
  return userStore.$state;
});
function readAll(){

}
let notice = computed(() => {
  return `你好,${
    userStore.nickname ? userStore.nickname : userStore.username
  },今天要炸鱼吗?`;
});
let comments = $ref([]);
let settingDraw = $ref<boolean>(false);
const theme = $(
  useStorage("themeConfig", {
    ...defaultTheme,
  })
);
let noticeTab = $ref("notice");
const { isFullscreen, enter, exit, toggle } = useFullscreen();

function showSettings() {
  settingDraw = true;
}

function gotoRoute(name: string) {
  router.push({
    name: name,
  });
}

function logout() {
  localStorage.clear();
  router.push({ name: "login" });
}

async function getUnread() {
  let data = await getCommentsUnreadApi();
  comments = data.data;
}

async function getUser() {
  let data = await getUserInfoApi();
  userStore.setUserInfo(data.data);
}

onBeforeMount(async () => {
  await getUser();
  await getUnread();
});
watch(
  theme,
  (val, oldVal) => {
    setThemeConfig(val);
  },
  { immediate: true }
);
</script>

<style lang="scss" scoped>
.notice-tab {
  padding: 1rem;
  display: flex;
  flex-direction: column;
  .notice-head{
    display: flex;

  }
  .notice-item {
    width: 20rem;
    display: flex;

    &:hover {
      background-color: rgb(243 244 246);
    }
  }
}
.header {
  height: 3rem;
  padding: 0.5rem 1rem;
  display: flex;
  align-items: center;
  justify-content: space-between;

  .header-right {
    display: flex;
    align-items: center;

    .el-icon {
      margin-left: 1rem;

      &:hover {
        cursor: pointer;
        color: var(--el-color-primary);
      }
    }

    .el-dropdown-link {
      display: flex;
      align-items: center;
      padding: 0 2rem;
      cursor: pointer;
    }
  }
}
</style>
