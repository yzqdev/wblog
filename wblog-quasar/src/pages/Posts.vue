<template>
<div class="posts">
 <q-card v-for="item in posts" class="posts-card" @click="gotoPost(item)">
    {{item.title}}
    <q-card-section>

    </q-card-section>


  </q-card>


</div>
</template>

<script setup>
import {onMounted} from "vue";
import {getPosts} from "../boot/apis";
import MdEditorV3 from 'md-editor-v3'
let router=useRouter()
import 'md-editor-v3/lib/style.css';
import {useRouter} from "vue-router";
let posts=$ref([])
let md=$ref('#hsdkfsldsjksdfsdfs')
onMounted(async () => {
let data=await getPosts()
  console.log(`%cdata`,`color:red;font-size:16px;background:transparent`)
  console.log(data)
  posts=data.data.posts
})
function gotoPost(item) {
  router.push({
    name:"post",
    params:{
      id:item.id
    }
  })
}
</script>

<style lang="scss" scoped>
.posts{
  padding: 2rem;
  .posts-card{
    margin: 1rem;
  }
}
</style>
