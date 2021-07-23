<template>
  <div class="golang-footer">
    <div class="golang-footer-page-link">
      <a href="https://github.com/shen100/golang123" target="_blank"
        ><div class="golang-footer-github">
          <Icon type="social-github" style="margin-right: 1px"></Icon>
        </div>
        <span style="margin-left: 2px">Github</span></a
      >
      <span>&nbsp;&nbsp;|&nbsp;</span>
      <a href="https://github.com/shen100/golang123/issues" target="_blank"
        >问题反馈</a
      >
      <span>&nbsp;&nbsp;|&nbsp;&nbsp;</span>
      <router-link to="/vote">投票</router-link>
      <span>&nbsp;&nbsp;|&nbsp;&nbsp;</span>
      <router-link to="/rank">积分榜</router-link>
      <span>&nbsp;&nbsp;|&nbsp;&nbsp;</span>
      <!--<a>帮助中心</a>
            <span>&nbsp;&nbsp;|&nbsp;&nbsp;</span>
            <a>服务条款</a>
            <span>&nbsp;&nbsp;|&nbsp;&nbsp;</span> -->
      <router-link to="/main/links">友情链接</router-link>
      <span>&nbsp;&nbsp;|&nbsp;&nbsp;</span>
      <!-- <a>联系我们</a>
            <span>&nbsp;&nbsp;|&nbsp;&nbsp;</span> -->
      <a href="/timeline">成长历史</a>
      <span>&nbsp;&nbsp;|&nbsp;&nbsp;</span>
      <router-link to="/about">关于我们</router-link>
    </div>
    <div class="golang-footer-copyright">
      <span>©&nbsp;2017&nbsp;<a href="/">golang中文社区</a>&nbsp;版权所有</span>
      <span>&nbsp;|&nbsp;</span>
      <span>京ICP备12045181号-2</span>
    </div>
  </div>
</template>
<script>
export default { name: "AppFooter" };
</script>
