<template>
  <div>
    <div class="signup-head">
      <div class="signup-head-content">
        <a href="/">
          <img src="@/assets/images/logo.png" alt="" />
          <span>Golang123</span>
        </a>
      </div>
    </div>
    <div class="signup-box">
      <div class="signup-nav">
        <span class="title">{{ !success ? "账号注册" : "邮箱验证" }}</span>
        <span class="desc"
          >{{
            !success ? "如果您有Golang123账号" : "如果您已经完成验证"
          }}，那么可以<a href="/signin">登录</a></span
        >
      </div>
      <el-form
        v-show="isMounted"
        ref="formCustom"
        :model="formCustom"
        :rules="ruleCustom"
        :label-width="80"
        class="signup-form"
        v-if="!success"
        style="height: 500px"
      >
        <el-form-item label="用户名" prop="username">
          <el-input
            size="large"
            type="text"
            v-model="formCustom.username"
            @on-blur="blur('formCustom.username')"
            class="signup-input"
          ></el-input>
          <span class="signup-label">4-20位可由中文、数字、字母组成</span>
        </el-form-item>
        <el-form-item label="密码" prop="passwd">
          <el-input
            size="large"
            type="password"
            v-model="formCustom.passwd"
            class="signup-input"
          ></el-input>
          <span class="signup-label">密码由6-20个字符组成，区分大小写</span>
        </el-form-item>
        <el-form-item label="确认密码" prop="passwdCheck">
          <el-input
            size="large"
            type="password"
            v-model="formCustom.passwdCheck"
            class="signup-input"
          ></el-input>
          <span class="signup-label">请在此确认您的密码</span>
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input
            size="large"
            @on-blur="blur('formCustom.email')"
            v-model="formCustom.email"
            class="signup-input"
          ></el-input>
          <span class="signup-label">请输入有效的电子邮箱</span>
        </el-form-item>
        <el-button
          type="primary"
          size="large"
          class="signup-button"
          @click="handleSubmit('formCustom')"
          >立即注册</el-button
        >
      </el-form>
      <div v-if="success" class="signup-message-box" style="height: 500px">
        <div>
          <div class="message-mail-icon">
            <img src="@/assets/images/mail.png" alt="" />
          </div>
          <div class="message-mail-right">
            <p class="signup-reminder-text">
              我们发送了一封验证邮件到<span class="signup-resend">{{
                formCustom.email
              }}</span>
            </p>
            <p class="signup-reminder-text">
              请到您的邮箱收信，并点击其中的链接验证您的邮箱
            </p>
            <a
              :href="`http://mail.${
                formCustom.email.split('@')[
                  formCustom.email.split('@').length - 1
                ]
              }`"
              target="_blank"
              ><el-button type="primary">去邮箱验证</el-button></a
            >
            <p class="signup-reminder-text signup-text-bottom">收不到邮件？</p>
            <p class="signup-reminder-small">
              请查看您的垃圾邮件和广告邮件，邮件有可能会被误认为是垃圾邮件或广告邮件
            </p>
            <p class="signup-reminder-small signup-resend click-mouse">
              重新发送
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ErrorCode from "@/constant/ErrorCode";
import { trim, trimBlur } from "@/utils/tool";
import { signUp } from "@/utils/apiConfig";
import { config } from "@/utils/config";
export default {
  name: "Signup",
  data() {
    const validatePass = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入密码"));
      } else {
        if (value.length < 6 || value.length > 20) {
          return callback(new Error("密码必须6-20个字符"));
        }
        if (this.formCustom.passwdCheck !== "") {
          // 对第二个密码框单独验证
          this.$refs.formCustom.validateField("passwdCheck");
        }
        callback();
      }
    };
    const validatePassCheck = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入确认密码"));
      } else if (value !== this.formCustom.passwd) {
        callback(new Error("两次输入密码不一致!"));
      } else {
        callback();
      }
    };

    const validateUserName = (rule, value, callback) => {
      if (value && (value.length < 4 || value.length > 20)) {
        callback(new Error("用户名长度必须4-20位"));
      }
      callback();
    };
    return {
      isMounted: false,
      loading: false,
      formCustom: {
        passwd: "123456",
        passwdCheck: "123456",
        username: "yzqdev",
        email: "yzqdev@outlook.com",
      },
      success: false,
      ruleCustom: {
        passwd: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { validator: validatePass, trigger: "blur" },
        ],
        passwdCheck: [
          { required: true, message: "请输入确认密码", trigger: "blur" },
          { validator: validatePassCheck, trigger: "blur" },
        ],
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          { validator: validateUserName, trigger: "blur" },
        ],
        email: [
          { required: true, message: "请输入邮箱", trigger: "blur" },
          { type: "email", message: "邮箱格式不正确", trigger: "blur" },
        ],
      },
    };
  },
  head() {
    return {
      title: "注册",
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
          signUp({
            name: trim(this.formCustom.username),
            password: trim(this.formCustom.passwd),
            email: trim(this.formCustom.email),
          })
            .then((res) => {
              console.log(res);
              console.log(
                `%c我擦擦擦`,
                `color:red;font-size:16px;background:transparent`
              );
              this.loading = false;
              if (res.errNo === ErrorCode.SUCCESS) {
                this.success = true;
                this.$message.success({
                  duration: config.messageDuration,
                  closable: true,
                  type: "success",
                  message: "提交成功!",
                });
              } else {
                console.log(
                  `%ccourse`,
                  `color:red;font-size:16px;background:transparent`
                );
                this.$message.error(res.msg);
                // this.$message.error({
                //   duration: config.messageDuration,
                //   closable: true, type:'error',
                //   message: res.msg,
                // });
              }
            })
            .catch((err) => {
              this.loading = false;
              this.$message.error({
                duration: config.messageDuration,
                closable: true,
                type: "error",
                message: err.message,
              });
            });
        }
      });
    },
    blur(name) {
      trimBlur(name, this);
    },
  },
  mounted() {
    this.isMounted = true;
  },
};
</script>

<style lang="less" scoped>
.signup-head {
  padding: 10px;
  background-color: #f0f9ff;
  .signup-head-content {
    span {
      font-size: 30px;
      vertical-align: bottom;
      line-height: 60px;
    }
    width: 1000px;
    margin: 0 auto;
    img {
      width: 60px;
    }

    a {
      display: inline-block;
      &:hover {
        text-decoration: none;
      }
    }
  }
}

.signup-box {
  width: 1000px;
  min-height: 594px;
  overflow: hidden;
  background-color: #fff;
  border: 1px solid #e9e9e9;
  font-size: 0;
  margin-left: auto;
  margin-right: auto;
  margin-top: 20px;
  border-radius: 2px;
  position: relative;
  transition: all 0.3s;
}

.signup-box:hover {
  box-shadow: 0 1px 6px rgba(0, 0, 0, 0.2);
  border-color: transparent;
}

.signup-main {
  width: 900px;
  font-size: 0;
  margin: 0 auto;
  background-color: #fff;
}

.signup-nav {
  height: 60px;
  line-height: 60px;
  background: #fff;
  border-bottom: 1px solid #e9e9e9;
  padding: 0 24px;
  border-radius: 2px 2px 0 0;
}

.signup-label {
  margin-left: 20px;
  font-size: 14px;
  color: gray;
}

.signup-nav-sep {
  padding: 0 8px;
}

.signup-nav a {
  color: #80bd01;
}

.signup-nav li {
  float: left;
  line-height: 40px;
}

.signup-nav .title {
  font-size: 18px;
  font-weight: bold;
}

.signup-nav .desc {
  font-size: 12px;
  float: right;
  color: #000;
}

.signup-nav .desc a {
  font-weight: bold;
  color: #000;
}

.signup-nav .desc a:hover {
  text-decoration: underline;
}

.signup-form {
  margin-top: 32px;
  margin-left: 80px;
}
.signup-input {
  width: 400px;
}

.signup-button {
  margin-left: 80px;
  margin-bottom: 20px;
  width: 400px;
}

.message-mail-right {
  display: inline-block;
  vertical-align: top;
  margin-top: 10px;
  margin-left: 20px;
}

.signup-reminder-text {
  font-size: 14px;
  color: #000;
  font-weight: bold;
  margin-bottom: 10px;
}

.signup-reminder-text .signup-resend {
  margin-left: 5px;
}

.signup-text-bottom {
  margin-top: 50px;
}

.signup-reminder-small {
  font-size: 12px;
  color: gray;
  margin-bottom: 10px;
}

.message-mail-icon {
  margin-left: 200px;
  padding: 0 10px;
  display: inline-block;
  border-right: 1px solid;
  border-image: linear-gradient(
      to bottom,
      rgba(158, 158, 158, 0.25) 0%,
      #9e9e9e 50%,
      rgba(158, 158, 158, 0.25) 100%
    )
    30 30;
}

.signup-resend {
  color: #2d8cf0;
}

.click-mouse {
  cursor: pointer;
}

.message-mail-icon img {
  width: 100px;
}
</style>
