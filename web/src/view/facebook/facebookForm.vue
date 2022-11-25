<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="Access Token:" prop="accessToken">
          <el-input v-model="formData.accessToken" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="expires In:" prop="expiresIn">
          <el-input v-model.number="formData.expiresIn" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Reauthorize Required In:" prop="reauthorizeRequiredIn">
          <el-input v-model="formData.reauthorizeRequiredIn" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="User ID:" prop="userID">
          <el-input v-model="formData.userID" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Facebook'
}
</script>

<script setup>
import {
  createFacebook,
  updateFacebook,
  findFacebook
} from '@/api/facebook'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            accessToken: '',
            expiresIn: 0,
            reauthorizeRequiredIn: '',
            userID: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行Search数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findFacebook({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.refacebook
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createFacebook(formData.value)
               break
             case 'update':
               res = await updateFacebook(formData.value)
               break
             default:
               res = await createFacebook(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: 'Create /Update successfully'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
