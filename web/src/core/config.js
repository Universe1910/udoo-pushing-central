/**
 * 网站配置文件
 */

const config = {
  appName: 'Puhsing Central',
  appLogo: 'https://dl.dropboxusercontent.com/s/1k19249eifen3e8/Udoo%20Square.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> Wellcome to Udoo`
      )
    )
    console.log(
      chalk.green(
        `> Current Version:v2.5.4`
      )
    )
    // console.log(
    //   chalk.green(
    //     `> 加群方式:微信：shouzi_1994 QQ群：622360840`
    //   )
    // )
    // console.log(
    //   chalk.green(
    //     `> GVA讨论社区：https://support.qq.com/products/371961`
    //   )
    // )
    // console.log(
    //   chalk.green(
    //     `> 插件市场:https://plugin.gin-vue-admin.com`
    //   )
    // )
    // console.log(
    //   chalk.green(
    //     `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
    //   )
    // )
    // console.log(
    //   chalk.green(
    //     `> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`
    //   )
    // )
    // console.log(
    //   chalk.green(
    //     `> 如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/coffee/index.html`
    //   )
    // )
    // console.log('\n')
  }
}

export default config
