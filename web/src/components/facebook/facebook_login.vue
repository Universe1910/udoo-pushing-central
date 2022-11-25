<template>
    <button
      @click="logInWithFacebook"
      class="minimal btn-def"
      :class="{
        dark: theme === 'dark',
        light: theme === 'light',
        dark: theme !== 'light',
      }"
      v-if="type === 'minimal'"
    >
      <!-- <img
        :src="theme === 'light' ? facebookColorMin : facebookWhiteMin"
        style="vertical-align: middle; padding: 5px"
        width="45"
        height="45"
      /> -->
    </button>
    <button
      @click="logInWithFacebook"
      v-else-if="type === 'custom'"
      :class="classes"
      :style="styles"
    >
      {{ text || "Login with Facebook" }}
    </button>
    <button
      @click="logInWithFacebook"
      class="normal btn-def"
      :class="{ dark: theme === 'dark', light: theme === 'light' }"
      v-else
    >
      <!-- <img
        style="vertical-align: middle; margin-right: 10px"
        :src="theme === 'dark' ? facebookWhiteFull : facebookColorFull"
        width="35"
        height="35"
      /> -->
      <span>Login with <b>Facebook</b></span>
    </button>
  </template>
  <script>
  
//   import facebookWhiteFull from "../assets/white-full.png";
//   import facebookWhiteMin from "../assets/white-min.png";
//   import facebookColorFull from "../assets/color-full.png";
//   import facebookColorMin from "../assets/color-min.png";
  
  export default {
    name: "facebookLogin",
    props: {
      appId: {
        type: String,
        required: true,
      },
      type: {
        type: String,
        default: "normal",
      },
      text: {
        type: String,
      },
      theme: {
        type: String,
        default: "dark",
      },
      classes: {
        type: String,
      },
      styles: {
        type: String,
      },
    },
    mounted() {
      this.loadFacebookSDK(document, "script", "facebook-jssdk");
      this.initFacebook();
    },
    emits: ["loggedIn", "loginFailed"],
    data() {
      return {
        // facebookWhiteFull,
        // facebookWhiteMin,
        // facebookColorFull,
        // facebookColorMin,
      };
    },
    methods: {
      async logInWithFacebook() {
        await this.loadFacebookSDK(document, "script", "facebook-jssdk");
        await this.initFacebook();
        const globalThis = this;
        window.FB.login(function (response) {
          if (response.authResponse) globalThis.$emit("loggedIn", response);
          else globalThis.$emit("loginFailed");
        });
        return false;
      },
      async initFacebook() {
        const { appId } = this;
        window.fbAsyncInit = function () {
          window.FB.init({
            appId: appId,
            cookie: true,
            xfbml: true,
            version: "v13.0",
          });
          window.FB.AppEvents.logPageView();
        };
      },
      async loadFacebookSDK(d, s, id) {
        var js,
          fjs = d.getElementsByTagName(s)[0];
        if (d.getElementById(id)) {
          return;
        }
        js = d.createElement(s);
        js.id = id;
        js.src = "https://connect.facebook.net/en_US/sdk.js";
        fjs.parentNode.insertBefore(js, fjs);
      },
    },
  };
  </script>
  <style scoped>
  @import url("https://fonts.googleapis.com/css2?family=Domine:wght@700&family=Montserrat:ital,wght@0,200;0,400;1,300&display=swap");
  .btn-def {
    border: none;
    cursor: pointer;
    border-radius: 2px;
  }
  .normal {
    cursor: pointer;
    border: none;
    padding: 10px 15px;
    font-size: 14px;
  }
  .dark {
    background: #3b5998;
    color: #fff;
    box-shadow: 0 0 10px #fff;
  }
  .light {
    color: #3b5998;
    background: #efefef;
    box-shadow: 0 0 10px #3b5998;
  }
  </style>
  