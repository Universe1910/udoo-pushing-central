import axios from "axios"; // 引入axios
import { ElMessage, ElMessageBox } from "element-plus";
import { useUserStore } from "@/pinia/modules/user";
import { emitter } from "@/utils/bus.js";
import router from "@/router/index";

const service = axios.create({
  baseURL: import.meta.env.VITE_BASE_API,
  timeout: 99999,
});
let acitveAxios = 0;
let timer;
const showLoading = () => {
  acitveAxios++;
  if (timer) {
    clearTimeout(timer);
  }
  timer = setTimeout(() => {
    if (acitveAxios > 0) {
      emitter.emit("showLoading");
    }
  }, 400);
};

const closeLoading = () => {
  acitveAxios--;
  if (acitveAxios <= 0) {
    clearTimeout(timer);
    emitter.emit("closeLoading");
  }
};
// http request 拦截器
service.interceptors.request.use(
  (config) => {
    if (!config.donNotShowLoading) {
      showLoading();
    }
    const userStore = useUserStore();
    config.headers = {
      "Content-Type": "application/json",
      "x-token": userStore.token,
      "x-user-id": userStore.userInfo.ID,
      ...config.headers,
    };
    return config;
  },
  (error) => {
    closeLoading();
    ElMessage({
      showClose: true,
      message: error,
      type: "error",
    });
    return error;
  }
);

// http response 拦截器
service.interceptors.response.use(
  (response) => {
    const userStore = useUserStore();
    closeLoading();
    if (response.headers["new-token"]) {
      userStore.setToken(response.headers["new-token"]);
    }
    if (response.data.code === 0 || response.headers.success === "true") {
      if (response.headers.msg) {
        response.data.msg = decodeURI(response.headers.msg);
      }
      return response.data;
    } else {
      ElMessage({
        showClose: true,
        message: response.data.msg || decodeURI(response.headers.msg),
        type: "error",
      });
      if (response.data.data && response.data.data.reload) {
        userStore.token = "";
        localStorage.clear();
        router.push({ name: "Login", replace: true });
      }
      return response.data.msg ? response.data : response;
    }
  },
  (error) => {
    closeLoading();

    if (!error.response) {
      ElMessageBox.confirm(
        `
        <p>Detective request errors</p>
        <p>${error}</p>
        `,
        "Request error        ",
        {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: "Repeat later        ",
          cancelButtonText: "Cancel",
        }
      );
      return;
    }

    switch (error.response.status) {
      case 500:
        ElMessageBox.confirm(
          `
        <p>Test interface error        ${error}</p>
        <p>Error code<span style="color:red"> 500 </span>：This kind of error content is common in the background PANIC. Please check the background log. If you affect your normal use</p>
        `,
          "Interface error",
          {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: "Clean up the cache",
            cancelButtonText: "Cancel",
          }
        ).then(() => {
          const userStore = useUserStore();
          userStore.token = "";
          localStorage.clear();
          router.push({ name: "Login", replace: true });
        });
        break;
      case 404:
        ElMessageBox.confirm(
          `
          <p>Test interface error ${error}</p>
          <p>Error code<span style="color:red"> 404 </span>Most of these errors are the interface that is not registered (or no restart) or the request path (method) does not match the API path (method)-if it is an automated code, please check whether there is a space for space.          </p>
          `,
          "Interface error",
          {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: "I got it",
            cancelButtonText: "Cancel",
          }
        );
        break;
    }

    return error;
  }
);
export default service;
