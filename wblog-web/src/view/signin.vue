<template>
  <el-row
      type="flex"
      align="middle"
      justify="center"
      class="golang-signin-container"
  >
    <el-col :xs="24" :lg="6" :md="14">
      <a class="golang-signin-title" href="/">GOLANG123</a>
      <p class="golang-signin-desc">和地鼠们分享你的知识、经验和见解</p>
      <el-form
          ref="formCustom"
          :model="formCustom"
          :rules="ruleCustom"
          class="signup-form"
      >
        <el-form-item prop="username">
          <el-input
              size="large"
              v-model="formCustom.username"
              @on-blur="blur('formCustom.username')"
              placeholder="用户名 / 邮箱"
          ></el-input>
        </el-form-item>
        <el-form-item prop="passwd">
          <el-input
              size="large"
              type="password"
              v-model="formCustom.passwd"
              placeholder="密码"
              @keydown.native="handleKeyUp"
          ></el-input>
        </el-form-item>
        <div v-if="luosimaoSiteKey" style="min-height: 44px">
          <div
              class="l-captcha"
              data-width="100%"
              :data-site-key="luosimaoSiteKey"
              data-callback="luosimaoCallback"
          ></div>
        </div>
        <p style="text-align: right; padding-right: 2px; margin-top: 10px">
          <a
              href="/signup"
              class="golang-common-link"
              style="margin-right: 12px"
          >立即注册</a
          >
          <a href="/ac/pwdReset" class="golang-common-link">忘记密码</a>
        </p>
        <el-form-item style="margin-top: 10px">
          <el-button
              size="large"
              type="primary"
              @click="handleSubmit('formCustom')"
              style="width: 100%"
          >登&nbsp;&nbsp;录
          </el-button>
        </el-form-item>
      </el-form>
    </el-col>
    <!--    <script-->
    <!--        type="text/javascript"-->
    <!--        color="51,133,255"-->
    <!--        opacity="0.7"-->
    <!--        zIndex="1"-->
    <!--        count="80"-->
    <!--        src="/javascripts/canvasnest/canvas-nest.min.js"-->
    <!--    ></script>-->
  </el-row>
</template>

<script>
import ErrorCode from "@/constant/ErrorCode";
import {trim, trimBlur} from "@/utils/tool";
import {signIn} from "@/utils/apiConfig";

let config = {luosimaoSiteKey: ""};
export default {
  name: "Signin",
  data() {
    return {
      luosimaoRes: "",
      luosimaoSiteKey: config.luosimaoSiteKey,
      loading: false,
      formCustom: {
        passwd: "123456",
        username: "yzqdev",
      },
      success: false,
      ruleCustom: {
        passwd: [{required: true, message: "请输入密码", trigger: "blur"}],
        username: [
          {required: true, message: "请输入用户名", trigger: "blur"},
        ],
      },
    };
  },
  asyncData(context) {
    let user = context.user;
    let redirectURL;

    let myURL = url.parse(context.req.url, true);
    if (myURL.query && myURL.query.ref) {
      redirectURL = decodeURIComponent(myURL.query.ref);
      let redirectObj = url.parse(redirectURL, true);
      let pathname = redirectObj.pathname;
      // 由重置密码或激活账号跳过来的，登录后直接跳到首页
      if (pathname.match(/\/reset\/.+/) || pathname.match(/\/active\/.+/)) {
        redirectURL = "/";
      }
    } else {
      redirectURL = "/";
    }
    if (user) {
      context.redirect(redirectURL);
      return;
    }
    return {
      user: user,
      ref: myURL.query.ref,
      redirectURL: redirectURL,
    };
  },

  head() {
    return {
      title: "登录",
      script: [{src: "//captcha.luosimao.com/static/js/api.js"}],
    };
  },
  methods: {
    handleSubmit(name) {
      this.$refs[name].validate((valid) => {
        if (valid) {
          if (this.loading) {
            return;
          }
          this.loading = true;
          signIn(
              this.formCustom.username.indexOf("@") > 0 ? "email" : "username",

              {
                signinInput: trim(this.formCustom.username),
                password: trim(this.formCustom.passwd),
                luosimaoRes: this.luosimaoRes,
              }
          )
              .then((res) => {
                this.loading = false;
                console.log(res);

                if (res.errNo === ErrorCode.SUCCESS) {
                  localStorage.setItem("token", res.data.token);
                  this.$store.commit("user", res.data.user);

                  this.$router.push({name: "main"});
                } else if (res.errNo === ErrorCode.IN_ACTIVE) {
                  window.location.href =
                      "/verify/mail?e=" + encodeURIComponent(res.data.email);
                } else {
                  // 没有配置luosimaoSiteKey的话，就没有验证码功能
                  this.luosimaoSiteKey && window.LUOCAPTCHA.reset();
                  this.$message.error({
                    duration: config.messageDuration,
                    closable: true,
                    content: res.msg,
                  });
                }
              })
              .catch((err) => {
                this.loading = false;
                this.$Message.error({
                  duration: config.messageDuration,
                  closable: true,
                  content: err.message || err.msg,
                });
              });
        }
      });
    },
    handleKeyUp(e) {
      if (e.keyCode === 13) {
        return this.handleSubmit("formCustom");
      }
    },
    blur(name) {
      trimBlur(name, this);
    },
  },
  mounted() {
    localStorage.clear();
    window.luosimaoCallback = (response) => {
      this.luosimaoRes = response;
    };
  },
};
</script>

<style scoped>
html, body {
  background-color: #f7fafc;
}


.golang-signin-container {
  height: 100%;
}

.golang-signin-title {
  display: block;
  color: #2d8cf0;
  font-size: 36px;
  font-style: italic;
  letter-spacing: 3px;
  text-align: center;
}

.golang-signin-title:hover {
  color: #2d8cf0;
  text-decoration: none;
}

.golang-signin-desc {
  color: #959595;
  margin: 10px 0 20px 0;
  text-align: center;
  letter-spacing: 3px;
  font-size: 14px;
}

a:hover {
  text-decoration: underline;
}

</style>
