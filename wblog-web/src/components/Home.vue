<template>
  <el-button type="warning" @click="showSiteInfo">home</el-button>
  <article>{{ msg }}</article>
</template>

<script>
// 课题列表
import {ref} from 'vue'
import {getSiteInfo} from "@/utils/apiConfig";

export default {
  setup() {
    let msg = ref('')
    const showSiteInfo = () => {
      getSiteInfo().then((res) => {
        console.log(res)
        msg.value=res.data
      })
    }
    return {
      msg,
      showSiteInfo
    }
  }
}
</script>

<style lang="scss" scoped></style>
