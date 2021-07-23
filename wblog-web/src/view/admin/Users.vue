<template>
  <el-row>
    <user-list
      :list="list"
      :totalCount="totalCount"
      :pageNo="pageNo"
      :pageSize="pageSize"
      :siteTitle="'全部用户'"
      :path="'list'"
      :role="role"
    />
  </el-row>
</template>

<script>
import UserList from "@/components/admin/UserList.vue";
import { getAdminUserList } from "@/utils/adminApiConfig";

export default {
  name: "Users",
  data() {
    return {
      list: "",
      totalCount: "",
      pageNo: "",
      pageSize: "",
      role: "",
    };
  },
  asyncData( ) {

    return getAdminUserList({

        pageNo: query.pageNo || 1,
        role: role,

    }).then((res) => {
      this.list = res.data.users || [];
      this.totalCount = res.data.totalCount;
      this.pageNo = res.data.pageNo;
      this.pageSize = res.data.pageSize;
      this.role = role;
    });
  },
  head() {
    return {
      title: "全部用户",
    };
  },
  mounted() {},
  components: {
    "user-list": UserList,
  },
};
</script>
