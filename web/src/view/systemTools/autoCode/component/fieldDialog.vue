<template>
  <div>
    <warning-bar title="id , created_at , updated_at , deleted_at Do not create automatically. If the condition is Like only supports string" />
    <el-form
      ref="fieldDialogFrom"
      :model="middleDate"
      label-width="150px"
      label-position="right"
      :rules="rules"
      class="grid-form"
    >
      <el-form-item label="Field Name" prop="fieldName">
        <el-input v-model="middleDate.fieldName" autocomplete="off" style="width:80%" />
        <el-button size="small" style="width:18%;margin-left:2%" @click="autoFill">
          <span style="font-size: 12px">Auto Fill</span>
        </el-button>
      </el-form-item>
      <el-form-item label="Field Desc" prop="fieldDesc">
        <el-input v-model="middleDate.fieldDesc" autocomplete="off" />
      </el-form-item>
      <el-form-item label="FieldJSON" prop="fieldJson">
        <el-input v-model="middleDate.fieldJson" autocomplete="off" />
      </el-form-item>
      <el-form-item label="Database File Name" prop="columnName">
        <el-input v-model="middleDate.columnName" autocomplete="off" />
      </el-form-item>
      <el-form-item label="Describe" prop="comment">
        <el-input v-model="middleDate.comment" autocomplete="off" />
      </el-form-item>
      <el-form-item label="Field Data Type" prop="fieldType">
        <el-select
          v-model="middleDate.fieldType"
          style="width:100%"
          placeholder="Please select Field data type"
          clearable
          @change="clearOther"
        >
          <el-option
            v-for="item in typeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item :label="middleDate.fieldType === 'enum' ? 'Enumeration value' : 'Type length'" prop="dataTypeLong">
        <el-input v-model="middleDate.dataTypeLong" :placeholder="middleDate.fieldType === 'enum'?`Example: 'Beijing', 'Tianjin' `: 'Database type length'" />
      </el-form-item>
      <el-form-item label="Field Search Type" prop="fieldSearchType">
        <el-select
          v-model="middleDate.fieldSearchType"
          style="width:100%"
          placeholder="Please select Field query conditions"
          clearable
        >
          <el-option
            v-for="item in typeSearchOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            :disabled="
              (middleDate.fieldType!=='string'&&item.value==='LIKE')||
                ((middleDate.fieldType!=='int'&&middleDate.fieldType!=='time.Time'&&middleDate.fieldType!=='float64')&&(item.value==='BETWEEN' || item.value==='NOT BETWEEN'))
            "
          />
        </el-select>
      </el-form-item>
      <el-form-item label="Associated dictionary" prop="dictType">
        <el-select
          v-model="middleDate.dictType"
          style="width:100%"
          :disabled="middleDate.fieldType!=='int'"
          placeholder="Please select the dictionary"
          clearable
        >
          <el-option
            v-for="item in dictOptions"
            :key="item.type"
            :label="`${item.type}(${item.name})`"
            :value="item.type"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="Must be filled">
        <el-switch v-model="middleDate.require" />
      </el-form-item>
      <el-form-item label="Can be empty">
        <el-switch v-model="middleDate.clearable" />
      </el-form-item>
      <el-form-item label="Verification failure">
        <el-input v-model="middleDate.errorText" />
      </el-form-item>

    </el-form>
  </div>
</template>

<script setup>
import { toLowerCase, toSQLLine } from '@/utils/stringFun'
import { getSysDictionaryList } from '@/api/sysDictionary'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'

const props = defineProps({
  dialogMiddle: {
    type: Object,
    default: function() {
      return {}
    }
  }
})

const middleDate = ref({})
const dictOptions = ref([])
const typeSearchOptions = ref([
  {
    label: '=',
    value: '='
  },
  {
    label: '<>',
    value: '<>'
  },
  {
    label: '>',
    value: '>'
  },
  {
    label: '<',
    value: '<'
  },
  {
    label: 'LIKE',
    value: 'LIKE'
  },
  {
    label: 'BETWEEN',
    value: 'BETWEEN'
  },
  {
    label: 'NOT BETWEEN',
    value: 'NOT BETWEEN'
  }
])
const typeOptions = ref([
  {
    label: 'String',
    value: 'string'
  },
  {
    label: 'int',
    value: 'int'
  },
  {
    label: 'bool',
    value: 'bool'
  },
  {
    label: 'float64',
    value: 'float64'
  },
  {
    label: 'Time',
    value: 'time.Time'
  },
  {
    label: 'Enum',
    value: 'enum'
  }
])
const rules = ref({
  fieldName: [
    { required: true, message: 'Please enter Field English name', trigger: 'blur' }
  ],
  fieldDesc: [
    { required: true, message: 'Please enter the Chinese name', trigger: 'blur' }
  ],
  fieldJson: [
    { required: true, message: 'Please enter Field formatting JSON', trigger: 'blur' }
  ],
  columnName: [
    { required: true, message: 'Please enter the database field', trigger: 'blur' }
  ],
  fieldType: [
    { required: true, message: 'Please select Field data type', trigger: 'blur' }
  ]
})

const init = async() => {
  middleDate.value = props.dialogMiddle
  const dictRes = await getSysDictionaryList({
    page: 1,
    pageSize: 999999
  })

  dictOptions.value = dictRes.data.list
}
init()

const autoFill = () => {
  middleDate.value.fieldJson = toLowerCase(middleDate.value.fieldName)
  middleDate.value.columnName = toSQLLine(middleDate.value.fieldJson)
}

const clearOther = () => {
  middleDate.value.fieldSearchType = ''
  middleDate.value.dictType = ''
}

const fieldDialogFrom = ref(null)
defineExpose({ fieldDialogFrom })
</script>

<script>

export default {
  name: 'FieldDialog'
}
</script>
<style scoped>
.grid-form{
  display: grid;
  grid-template-columns: 1fr 1fr;
}
.click-text{
  color: #0d84ff;
  font-size: 13px;
  cursor: pointer;
  user-select: none;
}
</style>
