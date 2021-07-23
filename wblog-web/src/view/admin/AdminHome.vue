<template>
  <div class="admin-root">
    <admin-header />
    <el-row class="admin-body">
      <el-col :span="5">
        <admin-sidebar
            :activeName="activeName"
            class="admin-sidebar-container"
        />
      </el-col>
      <el-col class="admin-body-container" :span="18">
        <router-view ></router-view>
      </el-col>
    </el-row>
  </div>
</template>

<script>

import Header from "@/components/admin/Header.vue";
import Sidebar from "@/components/admin/Sidebar.vue";


export default {
  data() {
    return {
      activeName: "",
      siteConfig: this.$store.state.siteConfig,
    };
  },
  head() {
    let siteConfig = this.siteConfig;
    return {
      titleTemplate: "%s - " + siteConfig.title,
      meta: [
        {
          hid: "description",
          name: "description",
          content: siteConfig.description,
        },
        { name: "keywords", content: siteConfig.keywords },
      ],
    };
  },
  components: {
    adminHeader: Header,
    adminSidebar: Sidebar,
  },

  mounted() {
    this.activeName = '';
    // this.activeName = this.$refs.content.$route.path;
  },
};
</script>

<style>
.admin-root {
  min-width: 1200px;
}

.admin-sidebar-container,
.admin-body-container {
  float: left;
}
</style>
