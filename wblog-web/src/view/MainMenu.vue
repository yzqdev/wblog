<template>
  <div>
    <AppHeader />
    <div class="common-body">
      <div class="common-body-left">
        <router-view />
      </div>
      <AppSidebar/>
    </div>
    <AppFooter />
<!--    <BackTop></BackTop>-->
  </div>
</template>

<script>
// 课题列表

import AppHeader from "@comp/Header.vue";
import AppSidebar from "@comp/Sidebar.vue";
import AppFooter from "@comp/Footer.vue";
export default {
  components: {AppFooter, AppSidebar, AppHeader}
}
</script>

<style lang="scss" scoped></style>
