<template>
  <div class="main">
    <h2>发布文章</h2>
    <el-input v-model="post.title" placeholder="请输入标题"></el-input>
    <el-divider></el-divider>
    <md-editor-v3 v-model="post.body"></md-editor-v3>
    <el-switch
      v-model="post.is_published"
      size="large"
      active-text="公开"
      inactive-text="未公开"
    />
    <el-divider></el-divider>
    <el-button class="confirm-btn" type="primary" @click="sendNotice"
      >确定</el-button
    >
    <br />
  </div>
</template>

<script setup lang="ts">
import { postApi } from "@/utils/apis";
import { ElMessage } from "element-plus";
let post = $ref({
  title: "",
  body: "",
  is_published: true,
});
async function sendNotice() {
  let res = await postApi(post);
  if (res.success) {
    ElMessage({
      type: "success",
      message: res.message,
    });
  }
}
</script>

<style lang="scss" scoped>
.confirm-btn {
  margin-top: 1rem;
}
</style>
