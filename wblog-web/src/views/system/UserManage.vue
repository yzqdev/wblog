<template>
  <div>用户管理</div>
  <el-table :data="tableData">
    <el-table-column prop="title" label="标题"></el-table-column>
    <el-table-column prop="is_published" label="公开"></el-table-column>
    <el-table-column prop="created_at" label="创建时间">
      <template #default="{ row }">
        <span>{{ formatDate(row.created_at) }}</span>
      </template>
    </el-table-column>
    <el-table-column prop="updated_at" label="更新时间">
      <template #default="{ row }">
        <span>{{ formatDate(row.updated_at) }}</span>
      </template>
    </el-table-column>
    <el-table-column label="操作">
      <template #default="{ row }">
        <el-button type="primary" @click="showDialog(row)">信息</el-button>
        <el-button type="danger" @click="deleteRow(row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  <el-dialog v-model="dialogVisible">
    <template #title> 文章信息 </template>
    <p>{{ cur.id }}</p>
    <p>{{ cur.title }}</p>
    <md-editor-v3 preview-only v-model="cur.body"></md-editor-v3>
  </el-dialog>
</template>

<script setup lang="ts">
import dayjs from "dayjs";
import { ElMessage } from "element-plus";
import {delPostsApi, getPostsApi} from "@/utils/apis";

interface Passage {
  id: string;
  title: string;
is_published:string
  updated_at: string;
  created_at: string;
}

let dialogVisible = $ref<boolean>(false);
let cur = $ref<Passage>();
let tableData = $ref<Passage[]>( );
onMounted(async () => {
 await getAllPosts()

})
async function getAllPosts() {
  let {data}=await getPostsApi()

  tableData=data.posts
}
function formatDate(date: string) {
  return dayjs(date).format("YYYY-MM-DD HH:mm:ss");
}

function showDialog(row: Passage) {
  cur = row;
  dialogVisible = true;
}

async function deleteRow(row: Passage) {
   let data=await  delPostsApi(row.id)
await  getAllPosts()
  ElMessage({ type: "success", message: `删除${row.title}成功!` });
}
</script>

<style scoped></style>
