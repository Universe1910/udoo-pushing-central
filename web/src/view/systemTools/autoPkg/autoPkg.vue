<template>
  <div>
    <warning-bar href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3" title="This function is used for the development environment. It is not recommended to release it to production. Please watch the video for the specific use effect https://www.bilibili.com/video/BV1kv4y1g7nT?p=3" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog('addApi')">Add new</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column align="left" label="ID" width="60" prop="ID" />
        <el-table-column align="left" label="Package Name" width="150" prop="packageName" />
        <el-table-column align="left" label="Display Name" width="150" prop="label" />
        <el-table-column align="left" label="Describe" min-width="150" prop="desc" />

        <el-table-column align="left" label="Action" width="200">
          <template #default="scope">
            <el-button
              icon="delete"
              size="small"
              type="primary"
              link
              @click="deleteApiFunc(scope.row)"
            >Delete</el-button>
          </template>
        </el-table-column>
      </el-table>

    </div>

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Create Package">
      <warning-bar title="Add new pkg is used for automated code use" />
      <el-form ref="pkgForm" :model="form" :rules="rules" label-width="120px">
        <el-form-item label="Package Name" prop="packageName">
          <el-input v-model="form.packageName" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Display name" prop="label">
          <el-input v-model="form.label" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Describe" prop="desc">
          <el-input v-model="form.desc" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Cancel</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Sure</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'AutoPkg',
}
</script>

<script setup>
import {
  createPackageApi,
  getPackageApi,
  deletePackageApi,
} from '@/api/autoCode'
import { ref } from 'vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const form = ref({
  packageName: '',
  label: '',
  desc: '',
})

const validateNum = (rule, value, callback) => {
  if ((/^\d+$/.test(value[0]))) {
    callback(new Error("Can't start with numbers"))
  } else {
    callback()
  }
}

const rules = ref({
  packageName: [
    { required: true, message: 'Please enter the package name', trigger: 'blur' },
    { validator: validateNum, trigger: 'blur' }
  ],
})

const dialogFormVisible = ref(false)
const openDialog = () => {
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  form.value = {
    packageName: '',
    label: '',
    desc: '',
  }
}

const pkgForm = ref(null)
const enterDialog = async() => {
  pkgForm.value.validate(async valid => {
    if (valid) {
      const res = await createPackageApi(form.value)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Add successfully',
          showClose: true
        })
      }
      getTableData()
      closeDialog()
    }
  })
}

const tableData = ref([])
const getTableData = async() => {
  const table = await getPackageApi()
  if (table.code === 0) {
    tableData.value = table.data.pkgs
  }
}

const deleteApiFunc = async(row) => {
  ElMessageBox.confirm('This operation is only PKG storage in the DELETE database. Please delete the same directory structure and keep it consistent with the database!', 'hint', {
    confirmButtonText: 'Sure',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async() => {
      const res = await deletePackageApi(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Deleted successfully!'
        })
        getTableData()
      }
    })
}

getTableData()
</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.warning {
  color: #dc143c;
}
</style>
