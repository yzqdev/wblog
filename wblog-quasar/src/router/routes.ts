import {RouteRecordRaw} from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [{path: '', component: () => import('pages/Index.vue')}, {
      path: 'posts',
      name:'posts',
      component: () => import('pages/Posts.vue')
    },{
      path: 'test',
      component: () => import('pages/Test.vue')
    },{
      path: 'post/:id',
      name:'post',
      component: () => import('pages/PostDetail.vue')
    },{
      path: 'links',
      name:'links',
      component: () => import('pages/Links.vue')
    },{
      path: 'about',
      name:'about',
      component: () => import('pages/About.vue')
    },{
      path: 'tags',
      name:'tags',
      component: () => import('pages/Tags.vue')
    }],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/Error404.vue'),
  },
];

export default routes;
