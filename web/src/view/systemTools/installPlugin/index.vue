<template>
  <div>
    <el-upload
      class="upload-demo"
      drag
      :action="`${path}/autoCode/installPlugin`"
      :headers="{'x-token':userStore.token}"
      :show-file-list="false"
      :on-success="handleSuccess"
      :on-error="handleSuccess"
      name="plug"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        Drag and drop<em>Click to upload</em>
      </div>
      <template #tip>
        <div class="el-upload__tip">
          Please drag the installed package ZIP to upload here
        </div>
      </template>
    </el-upload>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
import { ElMessage } from 'element-plus'
const userStore = useUserStore()
const path = ref(import.meta.env.VITE_BASE_API)

const handleSuccess = (res) => {
  if (res.code === 0) {
    let msg = ``
    res.data && res.data.forEach((item, index) => {
      msg += `${index + 1}.${item.msg}\n`
    })
    alert(msg)
  } else {
    ElMessage.error(res.msg)
  }
}
</script>
