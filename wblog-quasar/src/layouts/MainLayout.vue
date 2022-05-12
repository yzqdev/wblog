<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          v-if="$q.screen.lt.sm"
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />
        <template v-if="$q.screen.gt.xs">
          <q-toolbar-title>
            谦谦博客
          </q-toolbar-title>
          <q-toolbar-title v-for="item in linksList" @click="gotoRoute(item)">
           <q-btn flat> {{item.title}}</q-btn>
          </q-toolbar-title>


        </template>
        <div>Quasar v{{ $q.version }}</div>
      </q-toolbar>
    </q-header>

    <q-drawer

      v-model="leftDrawerOpen"

      bordered
    >
      <q-scroll-area style="height: calc(100% - 150px); margin-top: 150px; border-right: 1px solid #ddd">
        <q-list padding>
          <template v-for="item in linksList">
            <q-item clickable :active="curActive==item.link"  @click="gotoRoute(item)">
              <q-item-section avatar>
                <q-icon :name="item.icon"/>
              </q-item-section>

              <q-item-section>
                {{ item.title }}
              </q-item-section>
            </q-item>
          </template>
        </q-list>
      </q-scroll-area>

      <q-img class="absolute-top" src="https://cdn.quasar.dev/img/material.png" style="height: 150px">
        <div class="absolute-bottom bg-transparent">
          <q-avatar size="56px" class="q-mb-sm">
            <img src="https://cdn.quasar.dev/img/boy-avatar.png">
          </q-avatar>
          <div class="text-weight-bold">谦谦</div>
          <div>谦谦说世界</div>
        </div>
      </q-img>
    </q-drawer>

    <q-page-container>
      <router-view/>
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import EssentialLink from 'components/EssentialLink.vue'
import {useRoute, useRouter} from "vue-router";
import {computed, watch} from "vue";

let router = useRouter()
let route=useRoute()
const linksList = $ref([

  {
    title: '博客列表',
    caption: '我的博客列表',
    icon: 'code',
    link: 'posts'
  },
  {
    title: '标签',
    caption: '我的标签',
    icon: 'code',
    link: 'tags'
  },
  {
    title: '关于',
    caption: '关于我',
    icon: 'code',
    link: 'about'
  },

]);
let leftDrawerOpen = $ref(false)
 let curActive=$ref()
watch(route ,(val)  => {
  curActive=val.name
},{ immediate: true } )
function gotoRoute(item) {
  if (item.link.includes('http')) {
    location.assign(item.link)
  } else {
    curActive=route.name
    router.push({name: item.link})
  }
}


function toggleLeftDrawer() {
  leftDrawerOpen = !leftDrawerOpen
}
</script>
