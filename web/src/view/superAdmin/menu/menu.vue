<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="addMenu('0')">New Root Menu</el-button>
      </div>

      <!-- 由于此处菜单跟左侧列表一一对应所以不需要分页 pageSize默认999 -->
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" min-width="100" prop="ID" />
        <el-table-column align="left" label="Menu" min-width="120" prop="authorityName">
          <template #default="scope">
            <span>{{ scope.row.meta.title }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Icon" min-width="140" prop="authorityName">
          <template #default="scope">
            <div v-if="scope.row.meta.icon" class="icon-column">
              <el-icon>
                <component :is="scope.row.meta.icon" />
              </el-icon>
              <span>{{ scope.row.meta.icon }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Route Name" show-overflow-tooltip min-width="160" prop="name" />
        <el-table-column align="left" label="Route Path" show-overflow-tooltip min-width="160" prop="path" />
        <el-table-column align="left" label="Hide" min-width="100" prop="hidden">
          <template #default="scope">
            <span>{{ scope.row.hidden?"Hide":"Show" }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Parent ID" min-width="90" prop="parentId" />
        <el-table-column align="left" label="Sort" min-width="70" prop="sort" />
        <el-table-column align="left" label="Component" min-width="360" prop="component" />
        <el-table-column align="left" fixed="right" label="Action" width="300">
          <template #default="scope">
            <el-button
              size="small"
              type="primary"
              link
              icon="plus"
              @click="addMenu(scope.row.ID)"
            >Add sub-menu</el-button>
            <el-button
              size="small"
              type="primary"
              link
              icon="edit"
              @click="editMenu(scope.row.ID)"
            >Edit</el-button>
            <el-button
              size="small"
              type="primary"
              link
              icon="delete"
              @click="deleteMenu(scope.row.ID)"
            >Delete</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="handleClose" :title="dialogTitle">
      <warning-bar title="The new menu needs to be configured in the character management to be used to use" />
      <el-form
        v-if="dialogFormVisible"
        ref="menuForm"
        :inline="true"
        :model="form"
        :rules="rules"
        label-position="top"
        label-width="85px"
      >
        <el-form-item label="Route Name" prop="path" style="width:30%">
          <el-input
            v-model="form.name"
            autocomplete="off"
            placeholder="The only English string"
            @change="changeName"
          />
        </el-form-item>
        <el-form-item prop="path" style="width:30%">
          <template #label>
            <div style="display:inline-flex">
              Route Path
              <el-checkbox v-model="checkFlag" style="float:right;margin-left:20px;">Add parameter</el-checkbox>
            </div>
          </template>

          <el-input
            v-model="form.path"
            :disabled="!checkFlag"
            autocomplete="off"
            placeholder="It is recommended that only stitch parameters in the rear"
          />
        </el-form-item>
        <el-form-item label="Hide" style="width:30%">
          <el-select v-model="form.hidden" placeholder="Whether it is hidden in the list">
            <el-option :value="false" label="No" />
            <el-option :value="true" label="Yes" />
          </el-select>
        </el-form-item>
        <el-form-item label="Parent ID" style="width:30%">
          <el-cascader
            v-model="form.parentId"
            style="width:100%"
            :disabled="!isEdit"
            :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="File Path" prop="component" style="width:60%">
          <el-input v-model="form.component" autocomplete="off" placeholder="Page:view/xxx/xx.vue Plug-in:plugin/xx/xx.vue" @blur="fmtComponent" />
          <span style="font-size:12px;margin-right:12px;">If the menu contains the sub-menu, please create Router-View second-level routing page</span><el-button style="margin-top:4px" size="small" @click="form.component = 'view/routerHolder.vue'">点我设置</el-button>
        </el-form-item>
        <el-form-item label="Menu Title" prop="meta.title" style="width:30%">
          <el-input v-model="form.meta.title" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Icon" prop="meta.icon" style="width:30%">
          <icon :meta="form.meta" style="width:100%" />
        </el-form-item>
        <el-form-item label="Sort" prop="sort" style="width:30%">
          <el-input v-model.number="form.sort" autocomplete="off" />
        </el-form-item>
        <el-form-item prop="meta.activeName" style="width:30%">
          <template #label>
            <div>
              <span> Highlight menu</span>
              <el-tooltip content="Note: When you reach this route, the specified left menu specifies that name will be in an active state (light up), which can be empty。" placement="top" effect="light">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </div>
          </template>
          <el-input v-model="form.meta.activeName" :placeholder="form.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="KeepAlive" prop="meta.keepAlive" style="width:30%">
          <el-select v-model="form.meta.keepAlive" style="width:100%" placeholder="Whether the keypalive cache page">
            <el-option :value="false" label="No" />
            <el-option :value="true" label="Yes" />
          </el-select>
        </el-form-item>
        <el-form-item label="CloseTab" prop="meta.closeTab" style="width:30%">
          <el-select v-model="form.meta.closeTab" style="width:100%" placeholder="Auto close Tab">
            <el-option :value="false" label="No" />
            <el-option :value="true" label="Yes" />
          </el-select>
        </el-form-item>
        <el-form-item style="width:30%">
          <template #label>
            <div>
              <span> Whether to be the base page </span>
              <el-tooltip content="This choice is to show the left menu and top information on the left." placement="top" effect="light">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </div>
          </template>

          <el-select v-model="form.meta.defaultMenu" style="width:100%" placeholder="Whether to be the base page">
            <el-option :value="false" label="No" />
            <el-option :value="true" label="Yes" />
          </el-select>
        </el-form-item>
      </el-form>
      <div>
        <el-button
          size="small"
          type="primary"
          icon="edit"
          @click="addParameter(form)"
        >New menu parameter</el-button>
        <el-table :data="form.parameters" style="width: 100%">
          <el-table-column align="left" prop="type" label="Parameter type" width="180">
            <template #default="scope">
              <el-select v-model="scope.row.type" placeholder="Please choose">
                <el-option key="query" value="query" label="query" />
                <el-option key="params" value="params" label="params" />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="key" label="Key" width="180">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.key" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="value" label="Value">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.value" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"
                  size="small"
                  icon="delete"
                  @click="deleteParameter(form.parameters,scope.$index)"
                >Delete</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>

        <el-button
          style="margin-top:12px"
          size="small"
          type="primary"
          icon="edit"
          @click="addBtn(form)"
        >New control button</el-button>
        <el-table :data="form.menuBtn" style="width: 100%">
          <el-table-column align="left" prop="name" label="Button name" width="180">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.name" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="name" label="Description" width="180">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.desc" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"
                  size="small"
                  icon="delete"
                  @click="deleteBtn(form.menuBtn,scope.$index)"
                >Delete</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Cancel</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Sure</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  updateBaseMenu,
  getMenuList,
  addBaseMenu,
  deleteBaseMenu,
  getBaseMenuById
} from '@/api/menu'
import icon from '@/view/superAdmin/menu/icon.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { canRemoveAuthorityBtnApi } from '@/api/authorityBtn'
import { reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const rules = reactive({
  path: [{ required: true, message: '请输入菜单name', trigger: 'blur' }],
  component: [
    { required: true, message: '请输入文件路径', trigger: 'blur' }
  ],
  'meta.title': [
    { required: true, message: '请输入菜单展示名称', trigger: 'blur' }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])
const searchInfo = ref({})
// 查询
const getTableData = async() => {
  const table = await getMenuList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// Add new参数
const addParameter = (form) => {
  if (!form.parameters) {
    form.parameters = []
  }
  form.parameters.push({
    type: 'query',
    key: '',
    value: ''
  })
}

const fmtComponent = () => {
  form.value.component = form.value.component.replace(/\\/g, '/')
}

// Delete参数
const deleteParameter = (parameters, index) => {
  parameters.splice(index, 1)
}

// Add new可控按钮
const addBtn = (form) => {
  if (!form.menuBtn) {
    form.menuBtn = []
  }
  form.menuBtn.push({
    name: '',
    desc: '',
  })
}
// Delete可控按钮
const deleteBtn = async(btns, index) => {
  const btn = btns[index]
  if (btn.ID === 0) {
    btns.splice(index, 1)
    return
  }
  const res = await canRemoveAuthorityBtnApi({ id: btn.ID })
  if (res.code === 0) {
    btns.splice(index, 1)
    return
  }
}

const form = ref({
  ID: 0,
  path: '',
  name: '',
  hidden: false,
  parentId: '',
  component: '',
  meta: {
    activeName: '',
    title: '',
    icon: '',
    defaultMenu: false,
    closeTab: false,
    keepAlive: false
  },
  parameters: [],
  menuBtn: []
})
const changeName = () => {
  form.value.path = form.value.name
}

const handleClose = (done) => {
  initForm()
  done()
}
// Delete菜单
const deleteMenu = (ID) => {
  ElMessageBox.confirm('此操作将永久Delete所有角色下该菜单, 是否继续?', '提示', {
    confirmButtonText: 'Sure',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteBaseMenu({ ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Deleted successfully!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: 'Cancelled'
      })
    })
}
// 初始化弹窗内表格方法
const menuForm = ref(null)
const checkFlag = ref(false)
const initForm = () => {
  checkFlag.value = false
  menuForm.value.resetFields()
  form.value = {
    ID: 0,
    path: '',
    name: '',
    hidden: false,
    parentId: '',
    component: '',
    meta: {
      title: '',
      icon: '',
      defaultMenu: false,
      closeTab: false,
      keepAlive: false
    }
  }
}
// 关闭弹窗

const dialogFormVisible = ref(false)
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}
// 添加menu
const enterDialog = async() => {
  menuForm.value.validate(async valid => {
    if (valid) {
      let res
      if (isEdit.value) {
        res = await updateBaseMenu(form.value)
      } else {
        res = await addBaseMenu(form.value)
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: isEdit.value ? '编辑成功' : 'Added successfully!'
        })
        getTableData()
      }
      initForm()
      dialogFormVisible.value = false
    }
  })
}

const menuOption = ref([
  {
    ID: '0',
    title: '根菜单'
  }
])
const setOptions = () => {
  menuOption.value = [
    {
      ID: '0',
      title: '根目录'
    }
  ]
  setMenuOptions(tableData.value, menuOption.value, false)
}
const setMenuOptions = (menuData, optionsData, disabled) => {
  menuData &&
        menuData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID === form.value.ID,
              children: []
            }
            setMenuOptions(
              item.children,
              option.children,
              disabled || item.ID === form.value.ID
            )
            optionsData.push(option)
          } else {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID === form.value.ID
            }
            optionsData.push(option)
          }
        })
}

// 添加菜单方法，id为 0则为添加根菜单
const isEdit = ref(false)
const dialogTitle = ref('New menu')
const addMenu = (id) => {
   dialogTitle.value = 'New menu'
  form.value.parentId = String(id)
  isEdit.value = false
  setOptions()
  dialogFormVisible.value = true
}
// 修改菜单方法
const editMenu = async(id) => {
  dialogTitle.value = 'Edit menu'
  const res = await getBaseMenuById({ id })
  form.value = res.data.menu
  isEdit.value = true
  setOptions()
  dialogFormVisible.value = true
}

</script>

<script>
export default {
  name: 'Menus',
}
</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
.icon-column{
  display: flex;
  align-items: center;
  .el-icon{
    margin-right: 8px;
  }
}
</style>
