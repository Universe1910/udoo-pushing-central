<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="Create 时间">
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
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">Delete </el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="OA ID" prop="oaid"  />
        <el-table-column align="left" label="Official Account" prop="officialAccount" />
        <el-table-column align="left" label="Date" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Action">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateZaloOfficalAccountFunc(scope.row)">Edit</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">Delete </el-button>
            </template>
        </el-table-column>
        
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Add new">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="150px">
        <el-form-item label="OA ID:"  prop="oaid" >
          <el-input v-model="formData.oaid" :clearable="false"  placeholder="Please enter" />
        </el-form-item>
        <el-form-item label="Official Account:"  prop="officialAccount" >
          <el-input v-model="formData.officialAccount" :clearable="false"  placeholder="Please enter" />
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
  name: 'ZaloOfficalAccount'
}
</script>

<script setup>
import {
  createZaloOfficalAccount,
  deleteZaloOfficalAccount,
  deleteZaloOfficalAccountByIds,
  updateZaloOfficalAccount,
  findZaloOfficalAccount,
  getZaloOfficalAccountList
} from '@/api/zaloOfficalAccount'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        oaid: '',
        officialAccount: '',
        })

// 验证规则
const rule = reactive({
               oaid : [{
                   required: true,
                   message: 'OA ID can not be empty',
                   trigger: ['input','blur'],
               }],
               officialAccount : [{
                   required: true,
                   message: 'Official Account can not be empty',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// Reset
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

// Search
const getTableData = async() => {
  const table = await getZaloOfficalAccountList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
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
            deleteZaloOfficalAccountFunc(row)
        })
    }


// 批量Delete 控制标记
const deleteVisible = ref(false)

// 多选Delete 
const onDelete = async() => {
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
      const res = await deleteZaloOfficalAccountByIds({ ids })
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
const updateZaloOfficalAccountFunc = async(row) => {
    const res = await findZaloOfficalAccount({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rezaloOfficalAccount
        dialogFormVisible.value = true
    }
}


// Delete 行
const deleteZaloOfficalAccountFunc = async (row) => {
    const res = await deleteZaloOfficalAccount({ ID: row.ID })
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
        oaid: '',
        officialAccount: '',
        }
}
// 弹窗Delete
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createZaloOfficalAccount(formData.value)
                  break
                case 'update':
                  res = await updateZaloOfficalAccount(formData.value)
                  break
                default:
                  res = await createZaloOfficalAccount(formData.value)
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
