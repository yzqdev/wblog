<template>
  <el-select v-model="form" clearable placeholder="创建人">
    <el-option v-for="item in userInfo" :key="item.value" :label="item.label" :value="item.value" /></el-select>
  <i class="el-icon-share"></i>
</template>

<script>
export default {
  name: "Test",
  data(){
    return {
      form: '',
      userInfo: [{label: '全部', value: ''}, {label: '哈哈', value: '1'}, {label: '呵呵', value: '2'}],
    }
  }
}
</script>

<style scoped>

</style>
