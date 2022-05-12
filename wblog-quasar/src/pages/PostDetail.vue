<template>
  <div class="post-detail row wrap justify-start content-stretch">
    <q-card class="col col-lg-8  col-xs-12 q-pa-md   "><h2>{{ post.title }}</h2>
      <article class="row">
<div  class="col col-lg-4 col-xs-12  text-weight-medium">{{post.tags}}</div>
<div  class="col col-lg-4 col-xs-12 text-weight-medium">创建时间:{{formatTime(post.created_at)}}</div>
<div  class="col col-lg-4 col-xs-12  text-weight-medium">更新时间:{{formatTime(post.updated_at)}}</div>
      </article>
      <md-editor preview-only style="display: block!important;" v-model="post.body"></md-editor>
      <q-input
        v-model="comment.content"
        filled
        type="textarea"
      />
      <q-space><div></div></q-space>
      <div class="row">
        <q-input outlined label="用户名" class="col col-lg-4 col-xs-12" v-model="comment.name"></q-input>
        <q-input outlined type="email" label="邮箱" class="col col-lg-4 col-xs-12" v-model="comment.email"></q-input>
        <q-input outlined label="网址" class="col col-lg-4 col-xs-12" v-model="comment.link"></q-input>
      </div>
      <div class="row">
        <q-btn  color="primary"  class="col col-lg-6 col-xs-12 ">确定</q-btn>
        <div class="col col-lg-6 col-xs-12 "> </div>
      </div>
    </q-card>

    <q-card class="col col-lg-4 col-xs-12  q-pa-md " v-if="$q.screen.gt.xs">
      <q-card-section>
        <div>最新文章</div>
      </q-card-section>
    </q-card>
  </div>
</template>

<script setup>
import {onMounted} from "vue";
import {getPostById} from "../boot/apis";
import {useRoute} from "vue-router";
import MdEditor from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import {formatTime} from "../utils";

let route = useRoute()
let post = $ref({})
let comment = $ref({
  content: '', name: '', email: '', link: ''
})
onMounted(async () => {
  let data = await getPostById(route.params.id)
  post = data.data
})
</script>

<style lang="scss" scoped>
.post-detail {
  margin: 1rem;
}

</style>
