<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="Create Time">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="Start"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="End"></el-date-picker>
        </el-form-item>
        <el-form-item label="Email">
          <el-input v-model="searchInfo.email" placeholder="Email" />

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
          <p>Are you want to delete it?</p>
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

        <!-- <el-table-column align="left" label="uuid" prop="uuid" /> -->
        <el-table-column align="left" label="Firstname" prop="firstname" />
        <el-table-column align="left" label="Lastname" prop="lastname" />
        <el-table-column align="left" label="Email" prop="email" />
        <el-table-column align="left" label="Phone" prop="phone" />
        <el-table-column align="left" label="FB ID" prop="fbid" />
        <el-table-column align="left" label="Zalo ID" prop="zid" />
        <!-- <el-table-column align="left" label="Address" prop="address"   />
        <el-table-column align="left" label="City" prop="city"   />
        <el-table-column align="left" label="State" prop="state"   />
        <el-table-column align="left" label="Zipcode" prop="zipcode"   />
        <el-table-column align="left" label="Country" prop="country"   /> -->
        <el-table-column align="left" label="Date" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Action">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button"
              @click="updateContactFunc(scope.row)">Edit</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Add new Contacts">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <!-- <el-form-item label="uuid:"  prop="uuid" >
          <el-input v-model="formData.uuid" :clearable="false"  placeholder="uuid" />
        </el-form-item> -->
        <el-form-item label="Firstname:" prop="firstname">
          <el-input v-model="formData.firstname" :clearable="true" placeholder=" Please enter" />
        </el-form-item>
        <el-form-item label="Lastname:" prop="lastname">
          <el-input v-model="formData.lastname" :clearable="true" placeholder=" Please enter" />
        </el-form-item>
        <el-form-item label="Email:" prop="email">
          <el-input v-model="formData.email" :clearable="false" placeholder=" Please enter" />
        </el-form-item>
        <el-form-item label="Phone:" prop="phone">
          <el-input v-model="formData.phone" :clearable="false" placeholder=" Please enter" />
        </el-form-item>
        <el-form-item label="Facebook ID:" prop="fbid">
          <el-input v-model="formData.fbid" :clearable="true" placeholder=" Please enter" />
        </el-form-item>
        <el-form-item label="Zalo ID:" prop="zid">
          <el-input v-model="formData.zid" :clearable="true" placeholder=" Please enter" />
        </el-form-item>
        <el-form-item label="Address:" prop="address">
          <el-input v-model="formData.address" :clearable="true" placeholder=" Please enter" />
        </el-form-item>
        <el-form-item label="City:" prop="city">
          <el-input v-model="formData.city" :clearable="true" placeholder=" Please enter" />
        </el-form-item>
        <el-form-item label="State:" prop="state">
          <el-input v-model="formData.state" :clearable="true" placeholder=" Please enter" />
        </el-form-item>
        <el-form-item label="Zipcode:" prop="zipcode">
          <el-input v-model="formData.zipcode" :clearable="true" placeholder=" Please enter" />
        </el-form-item>
        <el-form-item label="Country:" prop="country">
          <el-input v-model="formData.country" :clearable="true" placeholder=" Please enter" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Cancel</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Create</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Contact'
}
</script>

<script setup>
import {
  createContact,
  deleteContact,
  deleteContactByIds,
  updateContact,
  findContact,
  getContactList
} from '@/api/contact'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, getCurrentInstance } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  uuid: '',
  firstname: '',
  lastname: '',
  email: '',
  phone: '',
  fbid: '',
  zid: '',
  address: '',
  city: '',
  state: '',
  zipcode: '',
  country: '',
})

// 验证规则
const rule = reactive({
  email: [{
    required: true,
    message: 'Email can not be empty',
    trigger: ['input', 'blur'],
  }],
  phone: [{
    required: true,
    message: 'Phone can not be empty',
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

// const uuuid = ref({})

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
  const table = await getContactList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    debugger;
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
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
  ElMessageBox.confirm('Are you sure you want to delete?', 'hint', {
    confirmButtonText: 'Delete',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(() => {
    deleteContactFunc(row)
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
      message: '请选择要Delete Please select the data to delete'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteContactByIds({ ids })
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
const updateContactFunc = async (row) => {
  const res = await findContact({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.recontact
    dialogFormVisible.value = true
  }
}


// Delete 行
const deleteContactFunc = async (row) => {
  const res = await deleteContact({ ID: row.ID })
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
    firstname: '',
    lastname: '',
    email: '',
    phone: '',
    fbid: '',
    zid: '',
    address: '',
    city: '',
    state: '',
    zipcode: '',
    country: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  debugger;
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    debugger;
    switch (type.value) {
      case 'create':
        res = await createContact(formData.value)
        break
      case 'update':
        res = await updateContact(formData.value)
        break
      default:
        res = await createContact(formData.value)
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
