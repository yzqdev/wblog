<template>
  <el-button type="primary" @click="addLinkBtnAction">添加链接</el-button>
  <el-table :data="tableData">
    <el-table-column prop="name" label="名字"></el-table-column>
    <el-table-column prop="url" label="url"></el-table-column>
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
        <el-button type="primary" @click="showDialog(row)">修改</el-button>
        <el-button type="danger" @click="deleteRow(row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  <el-dialog v-model="addLinkVisible" @close="clearDialog">
    <template #title>
      {{ dialogTitle }}
    </template>
    <el-form :model="linkData">
      <el-form-item prop="name" label="名字">
        <el-input v-model="linkData.name"></el-input>
      </el-form-item>
      <el-form-item prop="url" label="链接">
        <el-input v-model="linkData.url"></el-input>
      </el-form-item>
      <el-form-item prop="sort" label="排序">
        <el-input v-model="linkData.sort"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="add">确定</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup lang="ts">
import { formatDate } from "@/utils/util";
import { addLinksApi, delLinksApi, getLinksApi } from "@/utils/apis";
import { ElMessage } from "element-plus";

let tableData = $ref([]);
let addLinkVisible = $ref(false);
let linkData = $ref({ id: "", name: "", url: "", sort: "" });
let dialogTitle = $ref("添加链接");
function add() {
  addLinksApi(linkData).then((res) => {
    if (res.data) {
      getLinks();
      addLinkVisible = false;
      ElMessage({
        type: "success",
        message: "添加成功",
      });
    }
  });
}
function edit() {}
function clearDialog() {
  linkData = {
    id: "",
    name: "",
    url: "",
    sort: "",
  };
}
function showDialog(row) {
  linkData = row;
  dialogTitle = "修改链接";
  addLinkVisible = true;
}

async function deleteRow(row) {
  let data = await delLinksApi(row.id);
  if (data.success) {
    await getLinks();
    ElMessage({
      type: "success",
      message: "删除成功",
    });
  }
}

function addLinkBtnAction() {
  addLinkVisible = true;
}

async function getLinks() {
  let data = await getLinksApi();
  tableData = data.data;
}

onMounted(async () => {
  await getLinks();
});
</script>

<style scoped></style>
