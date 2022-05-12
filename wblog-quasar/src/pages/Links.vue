<template>
  <div>

    <q-card>
      <q-card-section>
        友情链接
      </q-card-section>
      <article class="q-mx-md" v-for="item in links">
        <q-btn flat color="primary" @click="goto(item)">{{ item.name }}</q-btn>
      </article>
    </q-card>

  </div>
</template>

<script setup>
import {onMounted} from "vue";
import {getLinks} from "../boot/apis";

let links = $ref([])
onMounted(async () => {
  let data = await getLinks()
  links = data.data
})

function goto(item) {
  window.open(item.url, '_blank')
}
</script>

<style scoped>

</style>
