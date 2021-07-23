<template>
  <div>
    <div class="about-box">
      <div class="about-main">
        <ul class="about-nav">
          <li>
            <a href="/"><span>首页</span></a>
          </li>
          <li class="about-nav-sep"><span>/</span></li>
          <li><span class="about-label">关于</span></li>
        </ul>
        <div class="about-center">
          <h1 class="about-title">关于golang中文社区</h1>
          <p class="about-text">
            golang中文社区是一个提供<em>分享</em>、<em>互动</em>、<em>学习</em>的平台。我们非常欢迎更多的golang爱好者加入。
          </p>
          <p class="about-text">
            目前, golang在国内越来越受欢迎,
            相关技术发展得也越来越快。golang中文社区希望能够为golang爱好者提供一个<em>开放</em>、<em>学习</em>、<em>交流</em>的平台。同时,
            我们也非常在意大家的使用体验, 对拿不准的功能, 会定期的在<a
              href="/vote"
              >投票专栏</a
            >发起投票, 征求大家的意见。
          </p>
          <p class="about-text">
            golang中文社区的所有代码都是<em>开源</em>的,
            项目名为<em>golang123</em>, github地址为<a
              href="https://github.com/shen100/golang123"
              target="_blank"
              >https://github.com/shen100/golang123</a
            ><br />golang123的技术架构是前后端分离的,
            前端使用vue、iview、node.js、nuxt等技术来开发,
            后端使用go、gin、gorm等技术来开发。golang123的技术选型也是超前的,
            我们大胆得使用nuxt来做<em>前后端同构渲染</em>。
          </p>
          <p class="about-text">
            最后, golang中文社区希望能够和广大地鼠们一起成长!
          </p>
          <div class="about-golang123">
            <pre>
              ___                               _      ___       __
             /\_ \                            /' \   /'___`\   /'__`\
   __     ___\//\ \      __      ___      __ /\_, \ /\_\ /\ \ /\_\L\ \
 /'_ `\  / __`\\ \ \   /'__`\  /' _ `\  /'_ `\/_/\ \\/_/// /__\/_/_\_&lt;_
/\ \L\ \/\ \L\ \\_\ \_/\ \L\.\_/\ \/\ \/\ \L\ \ \ \ \  // /_\ \ /\ \L\ \
\ \____ \ \____//\____\ \__/.\_\ \_\ \_\ \____ \ \ \_\/\______/ \ \____/
 \/___L\ \/___/ \/____/\/__/\/_/\/_/\/_/\/___L\ \ \/_/\/_____/   \/___/
   /\____/                                /\____/
   \_/__/                                 \_/__/
</pre
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "About",
  data() {
    return {};
  },
  created() {
    this.$store.commit("publishTopicVisible", false);
    this.$store.commit("maxBrowseVisible", false);
    this.$store.commit("maxCommentVisible", false);
  },
  head() {
    return {
      title: "关于",
    };
  },
  mounted() {},
};
</script>

<style lang="less" scoped>
body {
  background-color: #e1e1e1;
}

.about-box {
  overflow: hidden;
  font-size: 0;
  margin-left: auto;
  margin-right: auto;
}

.about-main {
  font-size: 0;
  display: inline-block;
  min-height: 666px;
  background: #fff;
}

.about-nav {
  width: 900px;
  height: 40px;
  display: inline-block;
  padding-left: 20px;
  background: #f6f6f6;
}

.about-nav li {
  float: left;
  line-height: 40px;
}

.about-nav-sep {
  padding: 0 8px;
}

.about-nav a {
  color: #80bd01;
}

.about-nav span {
  font-size: 14px;
}

.about-label {
  color: #999;
}

.about-center {
  padding: 20px;
  width: 900px;
}

.about-title {
  font-size: 22px;
  padding-bottom: 20px;
}

.about-text {
  font-size: 16px;
  line-height: 28px;
  margin-bottom: 20px;
}

.about-text a {
  padding: 0 4px;
  text-decoration: underline;
}

.about-golang123 {
  padding-top: 30px;
}

.about-golang123 pre {
  display: block;
  white-space: pre;
  margin: 0;
  word-spacing: initial;
  font-size: 14px;
}
</style>
