<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="Created At">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="Start time"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="End time"></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">Search</el-button>
          <el-button size="small" icon="refresh" @click="onReset">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">Add New</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>You sure you want to delete it?</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">Cancel</el-button>
            <el-button size="small" type="primary" @click="onDelete">Sure</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length"
              @click="deleteVisible = true">Delete</el-button>
          </template>
        </el-popover>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="date" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Name" prop="name" width="120" />
        <!-- <el-table-column align="left" label="Content" prop="content" width="120" /> -->
        <el-table-column align="left" label="Subject" prop="subject" width="120" />
        <el-table-column align="left" label="Created" prop="createdObject.nickName" width="120" />
        <el-table-column align="left" label="Action">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button"
              @click="toEmailBuilder(scope.row)">Builder</el-button>
            <el-button type="primary" link icon="edit" size="small" class="table-button"
              @click="updateEmailTemplateFunc(scope.row)">Edit</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">Delete</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Pop-up">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
        <el-form-item label="Name:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <!-- <el-form-item label="Content HTML:"  prop="content" >
          <el-input v-model="formData.content" :clearable="true"  placeholder="Please enter" />
        </el-form-item> -->
        <el-form-item label="Subject:" prop="subject">
          <el-input v-model="formData.subject" :clearable="true" placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Created by:" prop="createdBy">
          <!-- <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="Plse enter" /> -->
          <el-select v-model="selectedUser" clearable placeholder="Select">
            <el-option v-for="item in options" :key="item.ID" :label="item.nickName" :value="item.ID" />
          </el-select>
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
  name: 'EmailTemplate'
}
</script>

<script setup>
import {
  createEmailTemplate,
  deleteEmailTemplate,
  deleteEmailTemplateByIds,
  updateEmailTemplate,
  findEmailTemplate,
  getEmailTemplateList
} from '@/api/emailTemplate'

import {
  getUserList,
} from '@/api/user'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'


const router = useRouter()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: '',
  // content: '',
  subject: '',
  createdBy: 1,
})

const selectedUser = ref(0);
const options = ref([]);


const getUsers = async () => {
  
    const table = await getUserList({ page: 1, pageSize: 100 })
    if(table.code === 0){
      options.value = table.data.list

    }    
    console.log(options.value);

}

// 验证规则
const rule = reactive({
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

// 查询
const getTableData = async () => {
  
  const table = await getEmailTemplateList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    // debugger;
    //   console.log(tableData.value)
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getUsers();
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

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('You sure you want to delete it?', 'hint', {
    confirmButtonText: 'Sure',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(() => {
    deleteEmailTemplateFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: 'Please select the data to be deleted'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteEmailTemplateByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Deleted successfully'
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
const updateEmailTemplateFunc = async (row) => {
  const res = await findEmailTemplate({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reemail_template
    dialogFormVisible.value = true
  }
}


const toEmailBuilder = (row) => {
  router.push({
    name: 'emailBuilder',
    params: {
      id: row.ID
    }
  })
}

// 删除行
const deleteEmailTemplateFunc = async (row) => {
  const res = await deleteEmailTemplate({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Delete successfully'
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
    name: '',
    content: '',
    subject: '',
    createdBy: 1,
  }
}
// 弹窗确定
const enterDialog = async () => {

  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    debugger;
    formData.value.createdBy = selectedUser.value;
    switch (type.value) {
      case 'create':
        res = await createEmailTemplate(formData.value)
        break
      case 'update':
        res = await updateEmailTemplate(formData.value)
        break
      default:
        res = await createEmailTemplate(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Create/change success'
      })
      closeDialog()
      getTableData()
    }
  })
}
</script>

<style>

</style>
