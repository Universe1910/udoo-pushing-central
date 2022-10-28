<template>
  <div>
    <warning-bar title="To configure the Email configuration file in advance, in order to prevent unnecessary spam, the online experience function does not open this function experience." />
    <div class="gva-form-box">
      <el-form ref="emailForm" label-position="right" label-width="80px" :model="form">
        <el-form-item label="To">
          <el-input v-model="form.to" />
        </el-form-item>
        <el-form-item label="Subject">
          <el-input v-model="form.subject" />
        </el-form-item>
        <el-form-item label="Body">
          <el-input v-model="form.body" type="textarea" />
        </el-form-item>
        <el-form-item>
          <el-button @click="sendTestEmail">Send Test Email</el-button>
          <el-button @click="sendEmail">Send Email</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>

</template>

<script>
export default {
  name: 'Email',
}
</script>

<script setup>
import WarningBar from '@/components/warningBar/warningBar.vue'
import { emailTest } from '@/plugin/email/api/email.js'
import { ElMessage } from 'element-plus'
import { reactive, ref } from 'vue'
const emailForm = ref(null)
const form = reactive({
  to: '',
  subject: '',
  body: '',
})
const sendTestEmail = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage.success('Sent successfully')
  }
}

const sendEmail = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage.success('Sent successfully, please check')
  }
}
</script>

