<template>
<div class="post-detail">
  <h2>{{post.title }}</h2>
  <article>

  </article>
  <md-editor preview-only style="display: block!important;" v-model="post.body"></md-editor>
</div>
</template>

<script setup>
 import {onMounted} from "vue";
 import {getPostById} from "../boot/apis";
 import {useRoute} from "vue-router";
 import MdEditor from 'md-editor-v3'
 import 'md-editor-v3/lib/style.css'
let route=useRoute()
 let post=$ref({})
 onMounted(async () => {
   let data=await getPostById(route.params.id)
   post=data.data
 })
</script>

<style lang="scss" scoped>
.post-detail{
  margin: 2rem;
}

</style>
