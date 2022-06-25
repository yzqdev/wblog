<template>
<div>{{resp}}</div>
  <el-button @click="updateProfile">点击</el-button>
  <el-button @click="reg">注册</el-button>
</template>

<script setup>
import {getProfileApi, updateProfileApi} from "@/utils/userApi";
import {regApi} from "@/utils/apis";
import {superIndexApi} from "@/utils/superApi";

 let resp=$ref()
 async function updateProfile(){
   let res= await superIndexApi()
   resp=res.data
 }
 async function reg(){
   let res=await  regApi({username:"user1",password:"123456",email:'66@1.com'})
   resp=res.data
}
 onBeforeMount(async () => {
   resp=await  getProfileApi()
 })
</script>

<style scoped>

</style>
