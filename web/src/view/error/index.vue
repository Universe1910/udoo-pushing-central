<template>
  <div>
    <div class="big">
      <div class="inner">
        <img src="../../assets/notFound.png">
        <!-- <p>页面被神秘力量吸走了（如果您是开源版请联系我们修复）</p> -->
        <p>The page is absorbed by the mysterious power (if you are an open source version, please contact us to repair it)</p>
        <!-- <p style="font-size:18px;line-height:40px;">常见问题为当前此角色无当前路由，如果Sure要使用本路由，请到角色管理进行分配</p> -->
        <p style="font-size:18px;line-height:40px;">The common problem is that the current role has no current route. If you are determined to use this route, please go to the character management to allocate</p>
        <p>↓</p>
        <img src="../../assets/qm.png" class="leftPic">
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Error'

}
</script>

<style lang="scss">
    .big{
        width: 100%;
        height: calc(100vh - 220px);
        background-color: rgb(244,244,244);
        position: relative;
    }
    .inner{
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }
    .inner p {
        text-align: center;
        font-size: 24px;
    }
    .inner .leftPic{
        width: 60px;
        height: 60px;
        margin-left: 44%;
        margin-top: 20px;
    }
</style>
