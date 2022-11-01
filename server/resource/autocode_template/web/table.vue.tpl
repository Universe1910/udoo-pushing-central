<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="Create At">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="Start"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="End"></el-date-picker>
      </el-form-item>
           {{- range .Fields}}  {{- if .FieldSearchType}} {{- if eq .FieldType "bool" }}
            <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
            <el-select v-model="searchInfo.{{.FieldJson}}" clearable placeholder="Please choose">
                <el-option
                    key="true"
                    label="Yes"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="No"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
           {{- else if .DictType}}
           <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
            <el-select v-model="searchInfo.{{.FieldJson}}" clearable placeholder="Please choose" @clear="()=>{searchInfo.{{.FieldJson}}=undefined}">
              <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            {{- else}}
        <el-form-item label="{{.FieldDesc}}">


        {{- if eq .FieldType "float64" "int"}}
            {{if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
            <el-input v-model.number="searchInfo.start{{.FieldName}}" placeholder="Search conditions (up)" />
            —
            <el-input v-model.number="searchInfo.end{{.FieldName}}" placeholder="Search condition (stop)" />
           {{- else}}
             {{- if .DictType}}
              <el-select v-model="searchInfo.{{.FieldJson}}" placeholder="Please choose" style="width:100%" :clearable="true" >
               <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
             </el-select>
                    {{- else}}
             <el-input v-model.number="searchInfo.{{.FieldJson}}" placeholder="search condition" />
                    {{- end }}
          {{- end}}
        {{- else if eq .FieldType "time.Time"}}
            {{if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
            <el-date-picker v-model="searchInfo.start{{.FieldName}}" type="datetime" placeholder="Search conditions (up)"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.end{{.FieldName}}" type="datetime" placeholder="Search condition (stop)"></el-date-picker>
           {{- else}}
           <el-date-picker v-model="searchInfo.{{.FieldJson}}" type="datetime" placeholder="search condition"></el-date-picker>
          {{- end}}
        {{- else}}
         <el-input v-model="searchInfo.{{.FieldJson}}" placeholder="search condition" />
        {{- end}}

        </el-form-item>{{ end }}{{ end }}{{ end }}
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
        <el-table-column align="left" label="Date" width="180">
            <template #default="scope">{{ "{{ formatDate(scope.row.CreatedAt) }}" }}</template>
        </el-table-column>
        {{- range .Fields}}
        {{- if .DictType}}
        <el-table-column align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
            <template #default="scope">
            {{"{{"}} filterDict(scope.row.{{.FieldJson}},{{.DictType}}Options) {{"}}"}}
            </template>
        </el-table-column>
        {{- else if eq .FieldType "bool" }}
        <el-table-column align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
            <template #default="scope">{{"{{"}} formatBoolean(scope.row.{{.FieldJson}}) {{"}}"}}</template>
        </el-table-column>
         {{- else if eq .FieldType "time.Time" }}
         <el-table-column align="left" label="{{.FieldDesc}}" width="180">
            <template #default="scope">{{"{{"}} formatDate(scope.row.{{.FieldJson}}) {{"}}"}}</template>
         </el-table-column>
        {{- else }}
        <el-table-column align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120" />
        {{- end }}
        {{- end }}
        <el-table-column align="left" label="Action">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="update{{.StructName}}Func(scope.row)">Edit</el-button>
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
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
    {{- range .Fields}}
        <el-form-item label="{{.FieldDesc}}:"  prop="{{.FieldJson}}" >
      {{- if eq .FieldType "bool" }}
          <el-switch v-model="formData.{{.FieldJson}}" active-color="#13ce66" inactive-color="#ff4949" active-text="Yes" inactive-text="No" clearable ></el-switch>
      {{- end }}
      {{- if eq .FieldType "string" }}
          <el-input v-model="formData.{{.FieldJson}}" :clearable="{{.Clearable}}"  placeholder="Please enter" />
      {{- end }}
      {{- if eq .FieldType "int" }}
      {{- if .DictType}}
          <el-select v-model="formData.{{ .FieldJson }}" placeholder="Please choose" style="width:100%" :clearable="{{.Clearable}}" >
            <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
          </el-select>
      {{- else }}
          <el-input v-model.number="formData.{{ .FieldJson }}" :clearable="{{.Clearable}}" placeholder="Please enter" />
      {{- end }}
      {{- end }}
      {{- if eq .FieldType "time.Time" }}
          <el-date-picker v-model="formData.{{ .FieldJson }}" type="date" style="width:100%" placeholder="Selection date" :clearable="{{.Clearable}}"  />
      {{- end }}
      {{- if eq .FieldType "float64" }}
          <el-input-number v-model="formData.{{ .FieldJson }}"  style="width:100%" :precision="2" :clearable="{{.Clearable}}"  />
      {{- end }}
      {{- if eq .FieldType "enum" }}
            <el-select v-model="formData.{{ .FieldJson }}" placeholder="Please choose" style="width:100%" :clearable="{{.Clearable}}" >
               <el-option v-for="item in [{{.DataTypeLong}}]" :key="item" :label="item" :value="item" />
            </el-select>
      {{- end }}
        </el-form-item>
      {{- end }}
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
  name: '{{.StructName}}'
}
</script>

<script setup>
import {
  create{{.StructName}},
  delete{{.StructName}},
  delete{{.StructName}}ByIds,
  update{{.StructName}},
  find{{.StructName}},
  get{{.StructName}}List
} from '@/api/{{.PackageName}}'

// Please keep it on demand for the formatting tool
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// The dictionary generated by automated (may be empty) and field
    {{- range $index, $element := .DictTypes}}
const {{ $element }}Options = ref([])
    {{- end }}
const formData = ref({
        {{- range .Fields}}
        {{- if eq .FieldType "bool" }}
        {{.FieldJson}}: false,
        {{- end }}
        {{- if eq .FieldType "string" }}
        {{.FieldJson}}: '',
        {{- end }}
        {{- if eq .FieldType "int" }}
        {{.FieldJson}}: {{- if .DictType }} undefined{{ else }} 0{{- end }},
        {{- end }}
        {{- if eq .FieldType "time.Time" }}
        {{.FieldJson}}: new Date(),
        {{- end }}
        {{- if eq .FieldType "float64" }}
        {{.FieldJson}}: 0,
        {{- end }}
        {{- end }}
        })

// Verification rules
const rule = reactive({
    {{- range .Fields }}
            {{- if eq .Require true }}
               {{.FieldJson }} : [{
                   required: true,
                   message: '{{ .ErrorText }}',
                   trigger: ['input','blur'],
               }],
            {{- end }}
    {{- end }}
})

const elFormRef = ref()


// =========== The form control part of the form =========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// Reset searchInfo
const onReset = () => {
  searchInfo.value = {}
}

// Search 
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
{{- range .Fields}}{{- if eq .FieldType "bool" }}
  if (searchInfo.value.{{.FieldJson}} === ""){
      searchInfo.value.{{.FieldJson}}=null
  }{{ end }}{{ end }}
  getTableData()
}

// change size
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// Modify the page capacity
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// Get List
const getTableData = async() => {
  const table = await get{{.StructName}}List({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== The end of the form control part =============


// Get the required dictionary may be kept on demand for emptiness
const setOptions = async () =>{
{{- range $index, $element := .DictTypes }}
    {{ $element }}Options.value = await getDictFunc('{{$element}}')
{{- end }}
}

// Get the required dictionary may be kept on demand for emptiness
setOptions()


//Multi -choice data

const multipleSelection = ref([])
// Choice
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// Delete Row
const deleteRow = (row) => {
    ElMessageBox.confirm('Are you sure want to delete it?', 'hint', {
        confirmButtonText: 'Delete',
        cancelButtonText: 'Cancel',
        type: 'warning'
    }).then(() => {
            delete{{.StructName}}Func(row)
        })
    }


// Batch delete control mark
const deleteVisible = ref(false)

// Multi -choice delete 
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
      const res = await delete{{.StructName}}ByIds({ ids })
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

// Behavior control mark (internal pop -up window needs to be added or changed)
const type = ref('')

// Update
const update{{.StructName}}Func = async(row) => {
    const res = await find{{.StructName}}({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.re{{.Abbreviation}}
        dialogFormVisible.value = true
    }
}


// Delete Row
const delete{{.StructName}}Func = async (row) => {
    const res = await delete{{.StructName}}({ ID: row.ID })
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

// pop -up control mark
const dialogFormVisible = ref(false)

// Open the pop -up window
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// Close the pop -up window
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
    {{- range .Fields}}
        {{- if eq .FieldType "bool" }}
        {{.FieldJson}}: false,
        {{- end }}
        {{- if eq .FieldType "string" }}
        {{.FieldJson}}: '',
        {{- end }}
        {{- if eq .FieldType "int" }}
        {{.FieldJson}}: {{- if .DictType }} undefined{{ else }} 0{{- end }},
        {{- end }}
        {{- if eq .FieldType "time.Time" }}
        {{.FieldJson}}: new Date(),
        {{- end }}
        {{- if eq .FieldType "float64" }}
        {{.FieldJson}}: 0,
        {{- end }}
        {{- end }}
        }
}
// Pop -up
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await create{{.StructName}}(formData.value)
                  break
                case 'update':
                  res = await update{{.StructName}}(formData.value)
                  break
                default:
                  res = await create{{.StructName}}(formData.value)
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
