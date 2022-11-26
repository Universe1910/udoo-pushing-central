<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="Campaign:" prop="campaignID">
          <el-input v-model.number="formData.campaignID" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Action:" prop="action">
          <el-input v-model="formData.action" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Contact:" prop="contactID">
          <el-input v-model.number="formData.contactID" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Message:" prop="message">
          <el-input v-model="formData.message" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Type:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="Please enter" />
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
  name: 'CampaignLog'
}
</script>

<script setup>
import {
  createCampaignLog,
  updateCampaignLog,
  findCampaignLog
} from '@/api/campaignLog'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            campaignID: 0,
            action: '',
            contactID: 0,
            message: '',
            type: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行Search数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCampaignLog({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recampaignLog
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
               res = await createCampaignLog(formData.value)
               break
             case 'update':
               res = await updateCampaignLog(formData.value)
               break
             default:
               res = await createCampaignLog(formData.value)
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
