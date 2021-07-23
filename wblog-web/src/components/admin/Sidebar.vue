<template>
  <el-menu
    theme="dark"
    :default-active="activeName"
    :default-openeds="openNames"
    router
  >
    <el-submenu index="overview">
      <template #title>
        <i class="el-icon-share"></i>
        网站概览
      </template>

      <el-menu-item index="/admin">网站概览</el-menu-item>

      <el-menu-item index="/admin/setting">网站设置</el-menu-item>

      <el-menu-item index="/admin/pushLink">提交链接</el-menu-item>
    </el-submenu>
    <el-submenu index="category">
      <template #title>
        <Icon type="folder"></Icon>
        分类管理
      </template>
      <el-menu-item index="/admin/category/list">分类列表</el-menu-item>
    </el-submenu>
    <el-submenu index="topic">
      <template #title>
        <Icon type="ios-paper"></Icon>
        话题管理
      </template>
      <router-link to="/admin/topic/list">
        <el-menu-item index="/admin/topic/list">全部话题</el-menu-item>
      </router-link>
      <router-link to="/admin/topic/today">
        <el-menu-item index="/admin/topic/today">今日话题</el-menu-item>
      </router-link>
      <router-link to="/admin/topic/yesterday">
        <el-menu-item index="/admin/topic/yesterday">昨日话题</el-menu-item>
      </router-link>
    </el-submenu>
    <el-submenu index="reply">
      <template #title>
        <Icon type="chatboxes"></Icon>
        回复管理
      </template>
      <el-menu-item index="/admin/reply/list">全部回复</el-menu-item>
      <el-menu-item index="/admin/reply/today">今日回复</el-menu-item>
      <el-menu-item index="/admin/reply/yesterday">昨日回复</el-menu-item>
    </el-submenu>
    <el-submenu index="user">
      <template #title>
        <Icon type="ios-people"></Icon>
        用户管理
      </template>
      <el-menu-item index="/admin/user/list">全部用户</el-menu-item>
      <el-menu-item index="/admin/user/today">今日用户</el-menu-item>
      <el-menu-item index="/admin/user/yesterday">昨日用户</el-menu-item>
    </el-submenu>
    <el-submenu index="craler">
      <template #title>
        <Icon type="bug"></Icon>
        爬虫管理
      </template>
      <el-menu-item-group title="账号">
        <el-menu-item index="/admin/crawl/account">爬虫账号</el-menu-item>
      </el-menu-item-group>
      <el-menu-item-group title="网站">
        <el-menu-item index="/admin/crawl/jianshu">简书</el-menu-item>
        <el-menu-item index="/admin/crawl/zhihu">知乎</el-menu-item>
        <el-menu-item index="/admin/crawl/huxiu">虎嗅</el-menu-item>
        <el-menu-item index="/admin/crawl/custom">自定义抓取</el-menu-item>
      </el-menu-item-group>
    </el-submenu>
  </el-menu>
</template>

<script>
export default {
  props: ["activeName"],

  data() {
    return {
      openNames: ["overview", "category", "topic", "reply", "user", "craler"],
    };
  },
};
</script>
