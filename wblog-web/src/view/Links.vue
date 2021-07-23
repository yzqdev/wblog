<template>
  <div>
    <div class="links-box">
      <div class="links-main">
        <ul class="links-nav">
          <li>
            <a href="/"><span>首页</span></a>
          </li>
          <li class="links-nav-sep"><span>/</span></li>
          <li><span class="links-label">友情链接</span></li>
        </ul>
        <div class="links-sites">
          <div class="links-site" v-for="site in sites">
            <a :href="site.url" target="_blank">{{ site.name }}</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name:'Links',
  data() {
    return {
      sites: [
        {
          name: "CNode 社区",
          url: "https://cnodejs.org",
        },
        {
          name: "Ruby China",
          url: "https://ruby-china.org/",
        },
      ],
    };
  },
  asyncData(context) {
    return {
      user: context.user,
    };
  },
  head() {
    return {
      title: "友情链接",
    };
  },
};
</script>

<style lang="less" scoped>
body {
  background-color: #e1e1e1;
}

.links-box {
  overflow: hidden;
  font-size: 0;
  margin-left: auto;
  margin-right: auto;
}

.links-main {
  font-size: 0;
  display: inline-block;
  min-height: 666px;
  background: #fff;
}

.links-nav {
  width: 900px;
  height: 40px;
  display: inline-block;
  padding-left: 20px;
  background: #f6f6f6;
}

.links-nav li {
  float: left;
  line-height: 40px;
}

.links-nav-sep {
  padding: 0 8px;
}

.links-nav a {
  color: #80bd01;
}

.links-nav span {
  font-size: 14px;
}

.links-label {
  color: #999;
}

.links-sites {
  width: 900px;
  padding: 20px;
  font-size: 12px;
}

.links-site {
  width: 215px;
  line-height: 32px;
  display: inline-block;
  vertical-align: top;
  border-bottom: 1px dotted #ccc;
}

.links-site a {
  color: #08c;
}

</style>
