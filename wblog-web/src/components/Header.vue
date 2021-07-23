<template>
  <div class="golang-top-header">
    <div class="golang-top-box">
      <div class="golang-top-header-left">
        <div class="golang-logo-container">
          <a href="/"><img src="/images/logo.png" /></a>
        </div>
        <div class="golang-header-search">
          <form
            @submit.prevent="onSearch"
            action=""
            target="_blank"
            method="get"
            class="golang-top-search"
          >
            <p style="position: relative">
              <el-input
                @focus="onInputFocus"
                @blur="onInputBlur"
                v-model="q"
                type="text"
                class="golang-top-input"
                placeholder="请输入搜索关键词"
                name="topSearch"
              ></el-input>
            </p>
          </form>
        </div>
      </div>
      <div class="golang-top-header-nav">
        <ul>
          <li
            v-for="(item, index) in headRouters"
            :class="item.routeName == $route.name ? `active` : ``"
          >
            <router-link :to="item.to">{{ item.text }}</router-link>
          </li>
        </ul>
      </div>
      <div class="golang-top-header-right">
        <ul>
          <li>
            <a href="https://github.com/yzqdev/mili-vue3" target="_blank"
              >源码</a
            >
          </li>
          <li>
            <a href="https://github.com/yzqdev/mili-vue3/issues" target="_blank"
              >问题反馈</a
            >
          </li>
          <template v-if="user">
            <li class="user-message-wrapbox">
              <el-tooltip
                v-if="userMessages.length"
                trigger="hover"
                title="提示标题"
                placement="bottom"
              >
                <a href="" class="user-message-box"
                  ><Icon class="user-message" type="ios-bell-outline"></Icon
                  ><span class="user-message-tip-count">{{
                    messageCount
                  }}</span></a
                >
                <ul slot="content" class="header-message-list">
                  <li v-for="message in userMessages">
                    <p
                      v-if="message.type === 'messageTypeCommentArticle'"
                      class="header-message-item"
                    >
                      <router-link
                        :to="`/user/${message.fromUser.id}`"
                        target="_blank"
                        class="header-message-user"
                        >{{ message.fromUser.name }}</router-link
                      >&nbsp;回复了你的话题&nbsp;<router-link
                        @click="onReadMessage(message)"
                        :to="`/topic/${message.sourceID}/#reply-${message.commentID}`"
                        target="_blank"
                        class="header-message-content"
                        :style="{ color: message.readed ? '#a0a3a4' : '' }"
                        >{{ message.data.title }}</router-link
                      >
                    </p>
                    <p
                      v-else-if="message.type === 'messageTypeCommentVote'"
                      class="header-message-item"
                    >
                      <router-link
                        :to="`/user/${message.fromUser.id}`"
                        target="_blank"
                        class="header-message-user"
                        >{{ message.fromUser.name }}</router-link
                      >&nbsp;回复了你的投票&nbsp;<a
                        @click="onReadMessage(message)"
                        :href="`/vote/${message.sourceID}/#reply-${message.commentID}`"
                        target="_blank"
                        class="header-message-content"
                        :style="{ color: message.readed ? '#a0a3a4' : '' }"
                        >{{ message.data.title }}</a
                      >
                    </p>
                    <p
                      v-else-if="message.type === 'messageTypeCommentComment'"
                      class="header-message-item"
                    >
                      <a
                        :href="`/user/${message.fromUser.id}`"
                        target="_blank"
                        class="header-message-user"
                        >{{ message.fromUser.name }}</a
                      >&nbsp;回复了你&nbsp;<a
                        @click="onReadMessage(message)"
                        class="header-message-content"
                        :href="
                          message.sourceName === 'article'
                            ? `/topic/${message.sourceID}/#reply-${message.commentID}`
                            : `/vote/${message.sourceID}/#reply-${message.commentID}`
                        "
                        :style="{ color: message.readed ? '#a0a3a4' : '' }"
                        target="_blank"
                        >{{ message.data.commentContent }}</a
                      >
                    </p>
                  </li>
                </ul>
              </el-tooltip>
              <a v-else href="" class="user-message-box"
                ><Icon class="user-message" type="ios-bell-outline"></Icon
              ></a>
            </li>
            <li style="padding-right: 0">
              <el-tooltip
                v-if="user"
                trigger="hover"
                title="提示标题"
                placement="bottom"
              >
                <a :href="`/user/${user.id}`" class="header-usre-box">
                  <span class="header-avatar">
                    <img :src="user.avatarURL" alt="" />
                  </span>
                  <span class="header-user-name">{{ user.name }}</span>
                </a>
                <ul slot="content" class="header-user-box">
                  <li><a :href="`/user/${user.id}`">个人主页</a></li>
                  <li v-if="adminVisible"><a href="/admin">后台管理</a></li>
                  <li><a href="/ac/pwdModify">修改密码</a></li>
                  <li @click="onSignout">退&nbsp&nbsp出</li>
                </ul>
              </el-tooltip>
            </li>
          </template>
          <template v-else>
            <a @click="onSignin"><li style="color: #333">登录</li></a>
            <router-link to="/signup"><li style="color: #333">注册</li></router-link>
          </template>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import UserRole from "@/constant/UserRole";

import ErrorCode from "@/constant/ErrorCode";
import htmlUtil from "@/utils/html";
import trimHtml from "trim-html";
import {logout, readMessage} from "@/utils/apiConfig";

export default {
  name: "AppHeader",
  data() {
    let user = this.$store.state.user;
    let admins = [
      UserRole.USER_ROLE_ADMIN,
      UserRole.USER_ROLE_SUPER_ADMIN,
      UserRole.USER_ROLE_CRAWLER_ADMIN,
    ];
    let adminVisible = false;
    if (user && admins.indexOf(user.role) >= 0) {
      adminVisible = true;
    }
    return {
      q: "",
      user: user,
      adminVisible: adminVisible,
      isInputFocus: false,
      userMessages: [],
      headRouters: [
        { to: "/", text: "话题", routeName: "index" },
        { to: "/main/book", text: "图书", routeName: "book" },
        { to: "/main/vote", text: "投票", routeName: "vote" },
      ],
      messages: this.$store.state.messages,
      messageCount: this.$store.state.messageCount,
    };
  },
  methods: {
    onSearch() {
      let searchURL =
        "http://zhannei.baidu.com/cse/search?s=2990237584871814305&entry=1&q=" +
        encodeURIComponent(this.q);
      window.open(searchURL);
    },
    onInputFocus() {
      this.isInputFocus = true;
    },
    onInputBlur() {
      this.isInputFocus = false;
    },
    onReadMessage(message) {
       readMessage({
          params: {
            id: message.id,
          },
        })
        .then(() => {
          message.readed = true;
        });
    },
    onSignin() {
      this.$router.push({ name: "signin" });
    },
    onSignout() {
       logout()
        .then((res) => {
          if (res.errNo === ErrorCode.SUCCESS) {
            this.user = null;
            window.location.href = "/signin";
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
  mounted() {
    let messages = this.messages || [];
    let userMessages = messages.slice(0);
    let maxLen = 15;
    for (let i = 0; i < userMessages.length; i++) {
      if (userMessages[i].type === "messageTypeCommentComment") {
        let trimObj = trimHtml(userMessages[i].data.commentContent, {
          limit: maxLen,
          wordBreak: true,
          suffix: "...",
          preserveTags: false,
          moreLink: false,
        });
        let content = trimObj.html;
        content = htmlUtil.trimImg(content);
        userMessages[i].data.commentContent = content;
      }
      let title = userMessages[i].data.title || "";
      if (title.length > maxLen) {
        userMessages[i].data.title = title.substr(0, maxLen) + "...";
      }
      userMessages[i].readed = false;
    }
    this.userMessages = userMessages;
  },
};
</script>

<style></style>
