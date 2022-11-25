<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="Create">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="Start"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="End"></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">Search</el-button>
          <el-button size="small" icon="refresh" @click="onReset">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">Add new</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>Are you sure want to delete it?</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">Cancel</el-button>
            <el-button size="small" type="primary" @click="onDelete">Delete</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length"
              @click="deleteVisible = true">Delete </el-button>
          </template>
        </el-popover>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <!-- <el-table-column align="left" label="ID" prop="ID" /> -->
        <el-table-column align="left" label="Application ID" prop="applicationID" />
        <el-table-column align="left" label="Name" prop="applicationName" />
        <el-table-column align="left" label="Development Mode" prop="developmentMode">
          <template #default="scope">{{ formatBoolean(scope.row.developmentMode) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Status" prop="status">
          <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Date">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Action">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button"
              @click="authZaloApplicationFunc(scope.row)">Auth</el-button>
            <el-button type="primary" link icon="edit" size="small" class="table-button"
              @click="updateZaloApplicationFunc(scope.row)">Edit</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">Delete </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Add new">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="150px">
        <el-form-item label="Application Name:" prop="applicationName">
          <el-input v-model="formData.applicationName" :clearable="false" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Application ID:" prop="applicationID">
          <el-input v-model="formData.applicationID" :clearable="false" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Secret Key:" prop="secretKey">
          <el-input type="password" v-model="formData.secretKey" :clearable="false" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Code Verifier:" prop="codeVerifier">
          <el-input type="password" v-model="formData.codeVerifier" :clearable="false" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Callback URL:" prop="callbackURL">
          <el-input v-model="formData.callbackURL" :clearable="false" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Status:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="Active"
            inactive-text="Inactive" clearable></el-switch>
        </el-form-item>
        <el-form-item label="Default App:" prop="defaultApp">
          <el-switch v-model="formData.defaultApp" active-color="#13ce66" inactive-color="#ff4949" active-text="Active"
            inactive-text="Inactive" clearable></el-switch>
        </el-form-item>
        <el-form-item label="Development Mode:" prop="developmentMode">
          <el-switch v-model="formData.developmentMode" active-color="#13ce66" inactive-color="#ff4949" active-text="Active"
            inactive-text="Inactive" clearable></el-switch>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Cancel</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Save</el-button>
        </div>
      </template>
    </el-dialog>
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
  deleteZaloApplication,
  deleteZaloApplicationByIds,
  updateZaloApplication,
  findZaloApplication,
  getZaloApplicationList
} from '@/api/zaloApplication'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

import { useRouter } from 'vue-router'


const router = useRouter()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  applicationName: '',
  applicationID: '',
  secretKey: '',
  status: false,
  callbackURL: '',
  defaultApp: '',
  codeVerifier: '',
  developmentMode: false
})

// 验证规则
const rule = reactive({
  applicationName: [{
    required: true,
    message: 'Application Name can not be empty',
    trigger: ['input', 'blur'],
  }],
  applicationID: [{
    required: true,
    message: 'Application ID can not be empty',
    trigger: ['input', 'blur'],
  }],
  secretKey: [{
    required: true,
    message: 'Secret Key can not be empty',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.status === "") {
    searchInfo.value.status = null
  }
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// Search
const getTableData = async () => {
  const table = await getZaloApplicationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    console.log(tableData.value)
  }
}

getTableData()




// ============== 表格控制部分结束 ===============


// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()



// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}


// Delete 行
const deleteRow = (row) => {
  ElMessageBox.confirm('Are you sure want to delete it?', 'hint', {
    confirmButtonText: 'Delete',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(() => {
    deleteZaloApplicationFunc(row)
  })
}


// 批量Delete 控制标记
const deleteVisible = ref(false)

// 多选Delete 
const onDelete = async () => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: 'Please select the data to delete'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteZaloApplicationByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'successfully deleted'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateZaloApplicationFunc = async (row) => {
  const res = await findZaloApplication({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.rezaloApplication
    dialogFormVisible.value = true
  }
}

const authZaloApplicationFunc = async (row) => {
  router.push({
    name: 'zaloAuth',
    params: {
      id: row.ID
    }
  })
}


// Delete 行
const deleteZaloApplicationFunc = async (row) => {
  const res = await deleteZaloApplication({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'successfully deleted'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    applicationName: '',
    applicationID: '',
    secretKey: '',
    status: false,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    debugger;
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
      closeDialog()
      getTableData()
    }
  })
}
</script>

<style>

</style>
