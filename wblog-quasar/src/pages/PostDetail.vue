<template>
  <div class="post-detail row wrap justify-start content-stretch">
    <div class="col col-lg-8  col-xs-12    " :class="[$q.screen.gt.xs?'q-pa-md':'']">
    <q-card class="q-pa-md  "><h2>{{ post.title }}</h2>
      <article class="row">
        <div class="col col-lg-4 col-xs-12  text-weight-medium">{{ post.tags }}</div>
        <div class="col col-lg-4 col-xs-12 text-weight-medium">创建时间:{{ formatTime(post.created_at) }}</div>
        <div class="col col-lg-4 col-xs-12  text-weight-medium">更新时间:{{ formatTime(post.updated_at) }}</div>
      </article>
      <q-skeleton v-if="!post.title" style="height: 100vh"></q-skeleton>
      <md-editor v-if="post.title" preview-only style="display: block!important;" v-model="post.body"></md-editor>
      <q-input
        v-model="comment.content"
        filled
        placeholder="欢迎留下您的精彩评论!"
        type="textarea"
      />
      <blank></blank>
      <div class="row">
        <q-input outlined label="用户名" class="col col-lg-4 col-xs-12" v-model="comment.nickname"></q-input>
        <q-input outlined type="email" label="邮箱" class="col col-lg-4 col-xs-12" v-model="comment.email"></q-input>
        <q-input outlined label="网址" class="col col-lg-4 col-xs-12" v-model="comment.link"></q-input>
      </div><blank></blank>
      <div class="row">
        <q-btn color="primary" class="col col-lg-6 col-xs-12 " @click="submitCommit">确定</q-btn>
        <div class="col col-lg-6 col-xs-12  ">


        </div>
      </div>
      <section>
        <q-card style="margin-top: 0.5rem" v-for="item in post.comments">
          <q-card-section>{{item.nickname}}</q-card-section>
          <div class="q-pa-lg">{{item.content}}</div>
        </q-card>
      </section>
    </q-card></div>
    <div class="col col-lg-4 col-xs-12  q-pa-md " v-if="$q.screen.gt.xs">
      <q-card >
        <q-card>
          <q-card-section>
           我
          </q-card-section>

        </q-card>
      </q-card>
    <article style="height: 1rem"></article>
      <q-card >
        <q-card>
          <q-card-section>
            最新文章
          </q-card-section>

        </q-card>
      </q-card>
    </div>
  </div>
</template>

<script setup>
import {onMounted} from "vue";
import {getPostById, postComment} from "../boot/apis";
import {useRoute} from "vue-router";
import MdEditor from 'md-editor-v3'
import Blank from '../components/Blank.vue'
import 'md-editor-v3/lib/style.css'
import {formatTime} from "../utils";

let route = useRoute()
let post = $ref({})
let comment = $ref({
  content: '', nickname: '', email: '', link: ''
})
  function submitCommit(){
     postComment(post.id,comment).then(async () => {
       await getPostDetail()
       alert('success')
     }).catch((err) => {
       alert(err)
     })

}
async function getPostDetail(){
  let data = await getPostById(route.params.id)
  post = data.data
}
onMounted(async () => {
  await getPostDetail()
})
</script>

<style lang="scss" scoped>
.post-detail {
  margin: 0.5rem;
}

</style>
