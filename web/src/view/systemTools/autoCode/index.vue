<template>
  <div>
    <warning-bar href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
      title="
This function is used for the development environment. It is not recommended to release it to production. Please watch the video for the specific use effect http://www.bilibili.com/video/BV1kv4y1g7nT?p=3" />
    <!-- 从数据库直接获取字段 -->
    <div class="gva-search-box">
      <el-collapse v-model="activeNames" style="margin-bottom:12px">
        <el-collapse-item name="1">
          <template #title>
            <div :style="{ fontSize: '16px', paddingLeft: '20px' }">
              Click here to create code from the existing database
              <el-icon class="header-icon ">
                <pointer />
              </el-icon>
            </div>
          </template>
          <el-form ref="getTableForm" style="margin-top:24px" :inline="true" :model="dbform" label-width="120px">
            <el-form-item label="Select DB Type" prop="selectDBtype">
              <template #label>
                <el-tooltip
                  content="Note: You need to configure the multi-data library yourself in the DB-LIST in advance. If the configuration is not configured, the service party is restarted. (You can choose the corresponding library table here, which can be understood as the choice form from which library)"
                  placement="bottom" effect="light">
                  <div> Business library <el-icon>
                      <QuestionFilled />
                    </el-icon>
                  </div>
                </el-tooltip>
              </template>
              <el-select v-model="dbform.businessDB" clearable style="width:194px"
                placeholder="Choose a business library" @change="getDbFunc">
                <el-option v-for="item in dbList" :key="item.aliasName" :value="item.aliasName" :label="item.aliasName"
                  :disabled="item.disable">
                  <div>
                    <span>{{ item.aliasName }}</span>
                    <span style="float:right;color:#8492a6;font-size:13px">{{ item.dbName }}</span>
                  </div>
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Data Storage Name" prop="structName">
              <el-select v-model="dbform.dbName" clearable filterable placeholder="Please select the database"
                @change="getTableFunc">
                <el-option v-for="item in dbOptions" :key="item.database" :label="item.database"
                  :value="item.database" />
              </el-select>
            </el-form-item>
            <el-form-item label="Table Name" prop="structName">
              <el-select v-model="dbform.tableName" :disabled="!dbform.dbName" filterable
                placeholder="Please select the table">
                <el-option v-for="item in tableOptions" :key="item.tableName" :label="item.tableName"
                  :value="item.tableName" />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button size="small" type="primary" @click="getColumnFunc">Create with this table</el-button>
            </el-form-item>
          </el-form>
        </el-collapse-item>
      </el-collapse>
    </div>
    <div class="gva-search-box">
      <!-- 初始版本自动化代码工具 -->
      <el-form ref="autoCodeForm" :rules="rules" :model="form" label-width="120px" :inline="true">
        <el-form-item label="Struct Name" prop="structName">
          <el-input v-model="form.structName" placeholder="The first letter automatically converts uppercase" />
        </el-form-item>
        <el-form-item label="TableName" prop="tableName">
          <el-input v-model="form.tableName" placeholder="Specify the table name (non -required)" />
        </el-form-item>
        <el-form-item label="Struct abbreviation" prop="abbreviation">
          <el-input v-model="form.abbreviation"
            placeholder="The abbreviation will be used as an entry object name and routing group" />
        </el-form-item>
        <el-form-item label="Struct Description" prop="description">
          <el-input v-model="form.description" placeholder="Description as an automatic API DESCRIBE" />
        </el-form-item>
        <el-form-item label="Package Name" prop="packageName">
          <el-input v-model="form.packageName"
            placeholder="The default name of generating files (recommended in the hump format, the first letters, such as sysxxxxxxx)"
            @blur="toLowerCaseFunc(form, 'packageName')" />
        </el-form-item>
        <el-form-item label="Package" prop="package">
          <el-select v-model="form.package" style="width:194px">
            <el-option v-for="item in pkgs" :key="item.ID" :value="item.packageName" :label="item.packageName" />
          </el-select>
          <el-icon class="auto-icon" @click="getPkgs">
            <refresh />
          </el-icon>
          <el-icon class="auto-icon" @click="goPkgs">
            <document-add />
          </el-icon>
        </el-form-item>
        <el-form-item label="Business library" prop="businessDB">
          <template #label>
            <el-tooltip
              content="Note: You need to configure the multi-data library yourself in the DB-LIST in advance. This item will use GVA to create an automated code (global.gva_db)"
              placement="bottom" effect="light">
              <div>Business library <el-icon>
                  <QuestionFilled />
                </el-icon>
              </div>
            </el-tooltip>
          </template>
          <el-select v-model="form.businessDB" style="width:194px" placeholder="Choose a business library">
            <el-option v-for="item in dbList" :key="item.aliasName" :value="item.aliasName" :label="item.aliasName"
              :disabled="item.disable">
              <div>
                <span>{{ item.aliasName }}</span>
                <span style="float:right;color:#8492a6;font-size:13px">{{ item.dbName }}</span>
              </div>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip
              content="Note: Will automatically add created_by, updated_by, deleted_by to be automatically controlled to facilitate users to control resource permissions"
              placement="bottom" effect="light">
              <div>Create resource logo <el-icon>
                  <QuestionFilled />
                </el-icon>
              </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoCreateResource" />
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip content="Note: Register the automatic API into the database" placement="bottom" effect="light">
              <div>Automatically create API </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoCreateApiToSql" />
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip
              content="Note: The file generated by automatic migration to the corresponding position of the configuration of YAML"
              placement="bottom" effect="light">
              <div> Automatic movement file</div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoMoveFile" />
        </el-form-item>
      </el-form>
    </div>
    <!-- 组件列表 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" @click="editAndAddField()">Add new Field</el-button>
      </div>
      <el-table :data="form.fields">
        <el-table-column align="left" type="index" label="Index" width="80" />
        <el-table-column align="left" prop="fieldName" label="Name" />
        <el-table-column align="left" prop="fieldDesc" label="Description" />
        <el-table-column align="left" prop="require" label="Require">
          <template #default="{ row }">{{ row.require ? "Yes" : "No" }}</template>
        </el-table-column>
        <el-table-column align="left" prop="fieldJson" min-width="120px" label="FieldJson" />
        <el-table-column align="left" prop="fieldType" label="Type" width="130" />
        <el-table-column align="left" prop="dataTypeLong" label="Type Long" width="130" />
        <el-table-column align="left" prop="columnName" label="Column Name" width="130" />
        <el-table-column align="left" prop="comment" label="Describe" width="130" />
        <el-table-column align="left" prop="fieldSearchType" label="Search condition" width="130" />
        <el-table-column align="left" prop="dictType" label="Dict type" width="130" />
        <el-table-column align="left" label="Operate" width="300" fixed="right">
          <template #default="scope">
            <el-button size="small" type="primary" link icon="edit" @click="editAndAddField(scope.row)">Edit</el-button>
            <el-button size="small" type="primary" link :disabled="scope.$index === 0"
              @click="moveUpField(scope.$index)">Move up</el-button>
            <el-button size="small" type="primary" link :disabled="(scope.$index + 1) === form.fields.length"
              @click="moveDownField(scope.$index)">Move down</el-button>
            <el-popover v-model="scope.row.visible" placement="top">
              <p>Action</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="scope.row.visible = false">Cancel</el-button>
                <el-button type="primary" size="small" @click="deleteField(scope.$index)">Sure</el-button>
              </div>
              <template #reference>
                <el-button size="small" type="primary" link icon="delete" @click="scope.row.visible = true">Delete
                </el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <!-- 组件列表 -->
      <div class="gva-btn-list justify-content-flex-end auto-btn-list">
        <el-button size="small" type="primary" @click="enterForm(true)">Preview Code</el-button>
        <el-button size="small" type="primary" @click="enterForm(false)">Generate Code</el-button>
      </div>
    </div>
    <!-- 组件弹窗 -->
    <el-dialog v-model="dialogFlag" width="70%" title="Component Content">
      <FieldDialog v-if="dialogFlag" ref="fieldDialogNode" :dialog-middle="dialogMiddle" />
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Cancel</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Sure</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="previewFlag">
      <template #header>
        <div class="previewCodeTool">
          <p>Operation bar:</p>
          <el-button size="small" type="primary" @click="selectText">Select all</el-button>
          <el-button size="small" type="primary" @click="copy">Copy</el-button>
        </div>
      </template>
      <PreviewCodeDialog v-if="previewFlag" ref="previewNode" :preview-code="preViewCode" />
      <template #footer>
        <div class="dialog-footer" style="padding-top:14px;padding-right:14px">
          <el-button size="small" type="primary" @click="previewFlag = false">Sure</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>

import FieldDialog from '@/view/systemTools/autoCode/component/fieldDialog.vue'
import PreviewCodeDialog from '@/view/systemTools/autoCode/component/previewCodeDialg.vue'
import { toUpperCase, toHump, toSQLLine, toLowerCase } from '@/utils/stringFun'
import { createTemp, getDB, getTable, getColumn, preview, getMeta, getPackageApi } from '@/api/autoCode'
import { getDict } from '@/utils/dictionary'
import { ref, getCurrentInstance, reactive, watch, toRaw } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import WarningBar from '@/components/warningBar/warningBar.vue'


const fieldTemplate = {
  fieldName: '',
  fieldDesc: '',
  fieldType: '',
  dataType: '',
  fieldJson: '',
  columnName: '',
  dataTypeLong: '',
  comment: '',
  require: false,
  errorText: '',
  clearable: true,
  fieldSearchType: '',
  dictType: ''
}
const route = useRoute()
const router = useRouter()
const activeNames = reactive([])
const preViewCode = ref({})
const dbform = ref({
  businessDB: '',
  dbName: '',
  tableName: ''
})
const tableOptions = ref([])
const addFlag = ref('')
const fdMap = ref({})
const form = ref({
  structName: '',
  tableName: '',
  packageName: '',
  package: '',
  abbreviation: '',
  description: '',
  businessDB: '',
  autoCreateApiToSql: true,
  autoMoveFile: true,
  autoCreateResource: false,
  fields: []
})
const rules = ref({
  structName: [
    { required: true, message: 'Please enter the structure name', trigger: 'blur' }
  ],
  abbreviation: [
    { required: true, message: 'Please enter the structure abbreviation', trigger: 'blur' }
  ],
  description: [
    { required: true, message: 'Please enter the structure description', trigger: 'blur' }
  ],
  packageName: [
    {
      required: true,
      message: 'File name: sysxXxxxXxxx',
      trigger: 'blur'
    }
  ],
  package: [
    { required: true, message: 'Please select Package', trigger: 'blur' }
  ]
})
const dialogMiddle = ref({})
const bk = ref({})
const dialogFlag = ref(false)
const previewFlag = ref(false)

const toLowerCaseFunc = (form, key) => {
  form[key] = toLowerCase(form[key])
}
const previewNode = ref(null)
const selectText = () => {
  previewNode.value.selectText()
}
const copy = () => {
  previewNode.value.copy()
}
const editAndAddField = (item) => {
  dialogFlag.value = true
  if (item) {
    addFlag.value = 'edit'
    bk.value = JSON.parse(JSON.stringify(item))
    dialogMiddle.value = item
  } else {
    addFlag.value = 'add'
    dialogMiddle.value = JSON.parse(JSON.stringify(fieldTemplate))
  }
}
const moveUpField = (index) => {
  if (index === 0) {
    return
  }
  const oldUpField = form.value.fields[index - 1]
  form.value.fields.splice(index - 1, 1)
  form.value.fields.splice(index, 0, oldUpField)
}
const moveDownField = (index) => {
  const fCount = form.value.fields.length
  if (index === fCount - 1) {
    return
  }
  const oldDownField = form.value.fields[index + 1]
  form.value.fields.splice(index + 1, 1)
  form.value.fields.splice(index, 0, oldDownField)
}

const currentInstance = getCurrentInstance()
const enterDialog = () => {
  currentInstance.refs.fieldDialogNode.fieldDialogFrom.validate(valid => {
    if (valid) {
      dialogMiddle.value.fieldName = toUpperCase(
        dialogMiddle.value.fieldName
      )
      if (addFlag.value === 'add') {
        form.value.fields.push(dialogMiddle.value)
      }
      dialogFlag.value = false
    } else {
      return false
    }
  })
}
const closeDialog = () => {
  if (addFlag.value === 'edit') {
    dialogMiddle.value = bk.value
  }
  dialogFlag.value = false
}
const deleteField = (index) => {
  form.value.fields.splice(index, 1)
}
const autoCodeForm = ref(null)
const enterForm = async (isPreview) => {
  if (form.value.fields.length <= 0) {
    ElMessage({
      type: 'error',
      message: 'Please fill in at least one Field'
    })
    return false
  }
  if (
    form.value.fields.some(item => item.fieldName === form.value.structName)
  ) {
    ElMessage({
      type: 'error',
      message: 'The field of the same name as the structure'
    })
    return false
  }
  autoCodeForm.value.validate(async valid => {
    if (valid) {
      for (const key in form.value) {
        if (typeof form.value[key] === 'string') {
          form.value[key] = form.value[key].trim()
        }
      }
      form.value.structName = toUpperCase(form.value.structName)
      form.value.tableName = form.value.tableName.replace(' ', '')
      if (!form.value.tableName) {
        form.value.tableName = toSQLLine(toLowerCase(form.value.structName))
      }
      if (form.value.structName === form.value.abbreviation) {
        ElMessage({
          type: 'error',
          message: 'Structname and Struct referred to the same as'
        })
        return false
      }
      form.value.humpPackageName = toSQLLine(form.value.packageName)
      if (isPreview) {
        const data = await preview(form.value)
        preViewCode.value = data.data.autoCode
        previewFlag.value = true
      } else {
        const data = await createTemp(form.value)
        if (data.headers?.success === 'false') {
          return
        } else {
          if (form.value.autoMoveFile) {
            ElMessage({
              type: 'success',
              message: 'Automated code creation successfully, automatic movement success'
            })
            return
          }
          ElMessage({
            type: 'success',
            message: 'Automated code creation successfully, downloading'
          })
        }
        const blob = new Blob([data])
        const fileName = 'ginvueadmin.zip'
        if ('download' in document.createElement('a')) {
          // 不是IE浏览器
          const url = window.URL.createObjectURL(blob)
          const link = document.createElement('a')
          link.style.display = 'none'
          link.href = url
          link.setAttribute('download', fileName)
          document.body.appendChild(link)
          link.click()
          document.body.removeChild(link) // 下载完成移除元素
          window.URL.revokeObjectURL(url) // 释放掉blob对象
        } else {
          // IE 10+
          window.navigator.msSaveBlob(blob, fileName)
        }
      }
    } else {
      return false
    }
  })
}

const dbList = ref([])
const dbOptions = ref([])

const getDbFunc = async () => {
  dbform.value.dbName = ''
  dbform.value.tableName = ''
  const res = await getDB({ businessDB: dbform.value.businessDB })
  if (res.code === 0) {
    dbOptions.value = res.data.dbs
    dbList.value = res.data.dbList
  }
}
const getTableFunc = async () => {
  const res = await getTable({ businessDB: dbform.value.businessDB, dbName: dbform.value.dbName })
  if (res.code === 0) {
    tableOptions.value = res.data.tables
  }
  dbform.value.tableName = ''
}

const getColumnFunc = async () => {
  const gormModelList = ['id', 'created_at', 'updated_at', 'deleted_at']
  const res = await getColumn(dbform.value)
  if (res.code === 0) {
    let dbtype = ''
    if (dbform.value.businessDB !== '') {
      const dbtmp = dbList.value.find(item => item.aliasName === dbform.value.businessDB)
      console.log(dbtmp)
      const dbraw = toRaw(dbtmp)
      console.log(dbraw)
      dbtype = dbraw.dbtype
    }
    const tbHump = toHump(dbform.value.tableName)
    form.value.structName = toUpperCase(tbHump)
    form.value.tableName = dbform.value.tableName
    form.value.packageName = tbHump
    form.value.abbreviation = tbHump
    form.value.description = tbHump + 'surface'
    form.value.autoCreateApiToSql = true
    form.value.autoMoveFile = true
    form.value.fields = []
    res.data.columns &&
      res.data.columns.forEach(item => {
        if (!gormModelList.some(gormfd => gormfd === item.columnName)) {
          const fbHump = toHump(item.columnName)
          form.value.fields.push({
            fieldName: toUpperCase(fbHump),
            fieldDesc: item.columnComment || fbHump + 'Field',
            fieldType: fdMap.value[item.dataType],
            dataType: item.dataType,
            fieldJson: fbHump,
            dataTypeLong: item.dataTypeLong && item.dataTypeLong.split(',')[0],
            columnName: dbtype == 'oracle' ? item.columnName.toUpperCase() : item.columnName,
            comment: item.columnComment,
            require: false,
            errorText: '',
            clearable: true,
            fieldSearchType: '',
            dictType: ''
          })
        }
      })
  }
}
const setFdMap = async () => {
  const fdTypes = ['string', 'int', 'bool', 'float64', 'time.Time']
  fdTypes.forEach(async fdtype => {
    const res = await getDict(fdtype)
    res && res.forEach(item => {
      fdMap.value[item.label] = fdtype
    })
  })
}
const getAutoCodeJson = async (id) => {
  const res = await getMeta({ id: Number(id) })
  if (res.code === 0) {
    form.value = JSON.parse(res.data.meta)
  }
}

const pkgs = ref([])
const getPkgs = async () => {
  const res = await getPackageApi()
  if (res.code === 0) {
    pkgs.value = res.data.pkgs
  }
}

const goPkgs = () => {
  router.push({ name: 'autoPkg' })
}

const init = () => {
  getDbFunc()
  setFdMap()
  getPkgs()
  const id = route.params.id
  if (id) {
    getAutoCodeJson(id)
  }
}
init()

watch(() => route.params.id, (id) => {
  if (route.name === 'autoCodeEdit') {
    init()
  }
})

</script>

<script>

export default {
  name: 'AutoCode'
}
</script>

<style scoped lang="scss">
.previewCodeTool {
  display: flex;
  align-items: center;
  padding: 5px 0;
}

.button-box {
  padding: 10px 20px;

  .el-button {
    margin-right: 20px;
    float: right;
  }
}

.auto-btn-list {
  margin-top: 16px;
}

.auto-icon {
  margin-left: 6px;
  color: #666;
  cursor: pointer;
}
</style>
