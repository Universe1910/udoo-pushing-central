<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="Application Name:" prop="applicationName">
          <el-input v-model="formData.applicationName" :clearable="false" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Application ID:" prop="applicationID">
          <el-input v-model="formData.applicationID" :clearable="false" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Secret Key:" prop="secretKey">
          <el-input v-model="formData.secretKey" :clearable="false" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Status:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  name: 'ZaloApplication'
}
</script>

<script setup>
import {
  createZaloApplication,
  updateZaloApplication,
  findZaloApplication
} from '@/api/zaloApplication'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            applicationName: '',
            applicationID: '',
            secretKey: '',
            status: false,
        })
// 验证规则
const rule = reactive({
               applicationName : [{
                   required: true,
                   message: 'Application Name can not be empty',
                   trigger: ['input','blur'],
               }],
               applicationID : [{
                   required: true,
                   message: 'Application ID can not be empty',
                   trigger: ['input','blur'],
               }],
               secretKey : [{
                   required: true,
                   message: 'Secret Key can not be empty',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行Search数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findZaloApplication({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rezaloApplication
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
               res = await createZaloApplication(formData.value)
               break
             case 'update':
               res = await updateZaloApplication(formData.value)
               break
             default:
               res = await createZaloApplication(formData.value)
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
