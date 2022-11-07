<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="Name:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Start At:" prop="startAt">
          <el-date-picker v-model="formData.startAt" type="date" placeholder="选择Date" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="End At:" prop="endAt">
          <el-date-picker v-model="formData.endAt" type="date" placeholder="选择Date" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="Description:" prop="description">
          <el-input v-model="formData.description" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Parent ID:" prop="parentID">
          <el-input v-model.number="formData.parentID" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Created By:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="false" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Trigger ID:" prop="triggerId">
          <el-input v-model.number="formData.triggerId" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Sequence Id:" prop="sequenceId">
          <el-input v-model.number="formData.sequenceId" :clearable="true" placeholder="Please enter" />
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
  name: 'Campaign'
}
</script>

<script setup>
import {
  createCampaign,
  updateCampaign,
  findCampaign
} from '@/api/campaign'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            name: '',
            startAt: new Date(),
            endAt: new Date(),
            description: '',
            parentID: 0,
            createdBy: 0,
            triggerId: 0,
            sequenceId: 0,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: 'Campaign name can not be empty',
                   trigger: ['input','blur'],
               }],
               createdBy : [{
                   required: true,
                   message: 'Created By can not be empty',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行Search数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCampaign({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recampaign
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
               res = await createCampaign(formData.value)
               break
             case 'update':
               res = await updateCampaign(formData.value)
               break
             default:
               res = await createCampaign(formData.value)
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
