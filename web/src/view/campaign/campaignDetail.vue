<template>
  <el-form :model="formData" ref="vForm" :rules="rules" label-position="left" label-width="120px" size="medium"
    @submit.prevent>
    <el-row>
      <el-col :span="16" class="grid-cell">
        <el-form-item label="Campaign Name" label-width="120px" prop="campaignName">
          <el-input v-model="formData.name" type="text" clearable></el-input>
        </el-form-item>
        <div class="static-content-item">
          <el-divider direction="horizontal"></el-divider>
        </div>
        <el-form-item label="Active" prop="status">
          <el-switch v-model="formData.status"></el-switch>
        </el-form-item>
        <el-form-item label="Start" prop="startDate">
          <el-date-picker is-range v-model="formData.startAt" class="full-width-input" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="End" prop="startDate">
          <el-date-picker is-range v-model="formData.endAt" class="full-width-input" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="Zalo Application" prop="zaloApplicationID">
          <el-select v-model="formData.zaloApplicationID" clearable placeholder="Select">
            <el-option v-for="item in zaloApplicationOptions" :key="item.ID" :label="item.applicationName"
              :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="Decription" prop="description">
          <el-input type="textarea" v-model="formData.description" rows="3"></el-input>
        </el-form-item>
        <el-row>
          <el-col :span="12" class="grid-cell">
            <div class="static-content-item">
              <el-button type="primary" @click="addTrigger">{{ textData.trigger }}</el-button>
              <el-button type="success" class="btn-debug" @click="onTriggerDebugClick">Debug Trigger</el-button>
              
            </div>
          </el-col>
          <el-col :span="12" class="grid-cell">
            <div class="static-content-item">
              <el-button type="primary" @click="addSequence">{{ textData.sequence }}</el-button>
              <el-button type="success" class="btn-debug" @click="onSequenceDebugClick">Debug Sequence</el-button>
            </div>
          </el-col>
        </el-row>
      </el-col>
      <el-col :span="6" class="grid-cell right-panel">
        <div class="static-content-item">
          <el-button type="primary" class="btn-apply" @click="applyClick">Update Campaign</el-button>
          <el-button type="danger" @click="onStartSequenceClick">Start Sequence</el-button>
        </div>
      </el-col>
    </el-row>
    <div class="card-container">
      <!-- <el-card style="{width: 100% !important}" shadow="never">
        <div slot="header" class="clear-fix">
          <span>Statistical</span>
          <i class="float-right el-icon-arrow-down"></i>
        </div>
      </el-card> -->
    </div>
    <div class="tab-container">
      <el-tabs v-model="actionTabs" type="border-card">
        <el-tab-pane name="contactTab" label="Contacts">
          <div class="gva-table-box">
            <el-table style="width: 100%" tooltip-effect="dark" :data="formData.Contacts" row-key="ID"
              @selection-change="handleSelectionChange">
              <el-table-column align="left" label="ID" prop="ID" width="80" />
              <el-table-column align="left" label="Firstname" prop="firstname" />
              <el-table-column align="left" label="Lastname" prop="lastname" />
              <el-table-column align="left" label="Phone" prop="phone" />
              <el-table-column align="left" label="Email" prop="email" />
            </el-table>
            <div class="gva-pagination">
              <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
                :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
                @size-change="handleSizeChange" />
            </div>
          </div>
        </el-tab-pane>
        <el-tab-pane name="triggerTab" label="Trigger">
          <div class="dndflow" @drop="onDrop">
            <VueFlow @dragover="onDragOver" />
            <Sidebar />
          </div>
        </el-tab-pane>
        <el-tab-pane name="sequenceTab" label="Sequence">
          <div class="gva-table-box">
            <el-table style="width: 100%" tooltip-effect="dark" :data="formData.sequences" row-key="ID"
              class="error-table">
              <el-table-column align="left" label="Order" prop="order" width="60" />
              <el-table-column align="left" label="Date" width="180">
                <template #default="scope">{{ formatDate(scope.row.date) }}</template>
              </el-table-column>
              <!-- <el-table-column align="left" label="Time" width="180">
                <template #default="scope">{{ formatShortTime(scope.row.time) }}</template>
              </el-table-column> -->
              <el-table-column align="left" label="Action Name" prop="actionName" />
            </el-table>
          </div>
        </el-tab-pane>
        <el-tab-pane name="errorLogTab" label="Error Log">
          <el-button class="excel-btn" size="small" type="primary" icon="download" @click="handleExcelExportLog()">
            Export</el-button>
          <div class="gva-table-box">
            <el-table style="width: 100%" tooltip-effect="dark" :data="formData.Logs" row-key="ID"
              :row-class-name="logRowClassName" class="error-table">
              <el-table-column align="left" label="Source" prop="source" width="100"/>
              <el-table-column align="left" label="Action" prop="action" />
              <el-table-column align="left" label="Type" prop="type" />
              <el-table-column align="left" label="Message" prop="message" />
              <el-table-column align="left" label="Contact ID" prop="contactID">
                <template #default="scope">
                    <div v-if="getFullNameContactById(scope.row.contactID)">{{getFullNameContactById(scope.row.contactID)}}</div>
                    <div v-else>{{scope.row.contactID}}</div>
                </template>
              </el-table-column>
              <el-table-column align="left" label="Date" width="180">
                <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
              </el-table-column>
            </el-table>
          </div>
        </el-tab-pane>
        <el-tab-pane name="activityTab" label="Activity">
        </el-tab-pane>
      </el-tabs>
    </div>
  </el-form>

</template>


<script>
export default {
  name: "Campaing Detail"
}
</script>
<script setup>
import { reactive, ref, nextTick, watch, } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getDictFunc, formatDate, formatBoolean, filterDict, formatShortTime } from '@/utils/format'
import {
  updateCampaign,
  findCampaign,
  debugCampaign,
  debugSequence,
  startSequence
} from '@/api/campaign'

import {
  getZaloApplicationList
} from '@/api/zaloApplication'

import {
  deleteCampaignLog,
  deleteCampaignLogByIds,
  getCampaignLogList,
  findCampaignLogByCampaign,
} from '@/api/campaignLog'

import exportFromJSON from "export-from-json";


import { VueFlow, useVueFlow } from '@vue-flow/core'

import { useRoute, useRouter } from 'vue-router'


const { findNode, onConnect, setNodes, setEdges, dimensions, setTransform, nodes, edges, addEdges, addNodes, viewport, project, vueFlowRef, toObject } = useVueFlow({
  nodes: [
  ],
})



onConnect((params) => addEdges([params]))

const textData = ref({
  trigger: "Add Trigger",
  sequence: "Add Sequence"
})

const onRestore = (flow) => {
  if (flow) {
    const [x = 0, y = 0] = flow.position
    nodes.value = flow.nodes
    setNodes(flow.nodes)
    // flowSetEvent(flow.nodes)
    setEdges(flow.edges)
    setTransform({ x, y, zoom: flow.zoom || 0 })
    textData.value.trigger = "Edit Trigger"
  }

}

const onDragOver = (event) => {
  event.preventDefault()

  if (event.dataTransfer) {
    event.dataTransfer.dropEffect = 'move'
  }
}

const actionTabs = ref('contactTab')


const router = useRouter()

const route = useRoute()

const searchInfo = ref({ campaignID: Number(route.params.id) })

const getCampaign = async () => {

  var id = searchInfo.value.campaignID
  var res = await findCampaign({ ID: id })
  if (res.code === 0) {
    formData.value = res.data.recampaign
    if (res.code === 0) {
      formData.value = res.data.recampaign
      console.log(formData.value)
      if (formData.value.sequences) {
        textData.value.sequence = "Edit Sequence"
      }
      var tf = formData.value.triggerObject.flow;
      var flow = tf.length > 0 ? JSON.parse(tf) : null
      onRestore(flow)
    }
  }
}

getCampaign()


const campaingLogs = ref([])
const getCampaignAllLog = async () => {
  var id = formData.value.ID
  var res = await findCampaignLogByCampaign({ campaignID: id })
  if (res.code === 0) {
    campaingLogs.value = res.data.recampaignLog
  }
}

const zaloApplicationOptions = ref({})
const getZaloApplications = async () => {
  const res = await getZaloApplicationList({ page: 1, pageSize: 100 })
  if (res.code === 0) {
    zaloApplicationOptions.value = res.data.list
  }
  console.log(zaloApplicationOptions)
}
getZaloApplications()

const getFullNameContactById = (contactID) => {
  var ts = formData.value.Contacts.filter((e) => (e.ID == contactID))
  if (ts.length == 0){
    return false
  }
  return `${ts[0].firstname } ${ts[0].lastname }`
}

const onTriggerDebugClick = async () => {
  var id = searchInfo.value.campaignID
  var res = await debugCampaign({ ID: id })
  console.log(res)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Debug successfully'
    })
  }
}

const onSequenceDebugClick = async () => {
  var id = searchInfo.value.campaignID
  var res = await debugSequence({ ID: id })
  console.log(res)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Debug sequence successfully'
    })
  }
}

const onStartSequenceClick = async () => {
  var id = searchInfo.value.campaignID
  var res = await startSequence({ ID: id })
  console.log(res)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Start sequence successfully'
    })
  }
}



const applyClick = async () => {
  var res = await updateCampaign(formData.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Update successfully'
    })
    // closeDialog()
    // getTableData()
  } else {
    ElMessage({
      type: 'warning',
      message: 'Update failed'
    })
  }
}


const formData = ref(
  {
    name: "",
    startAt: null,
    endAt: null,
    description: "",
    status: false,
    sequenceId: 0,
    triggerId: 0,
    zaloApplicationID: 0,
    contacts: []
  }
)

const rule = reactive({})


const formRef = reactive({})

const addTrigger = () => {
  router.push({
    name: 'triggerBuilder',
    params: {
      id: searchInfo.value.campaignID
    }
  })
}
const addSequence = () => {
  router.push({
    name: 'sequenceBuilder',
    params: {
      id: searchInfo.value.campaignID
    }
  })
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
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

const handleExcelExportLog = async () => {
  const fileName = "TriggerLog.xlsx"
  if (!fileName || typeof fileName !== 'string') {
    fileName = 'ExcelExport.xlsx'
  }

  await getCampaignAllLog()
  const data = campaingLogs.value;
  const exportType = exportFromJSON.types.csv;
  if (data) exportFromJSON({ data, fileName, exportType });
}

const logRowClassName = (row, rowIndex) => {
  let rowClass = `${row.row.type}-row`;
  return rowClass.toLowerCase();
}

</script>

<style lang="scss">
.error-table .error-row {
  // --el-table-tr-bg-color: var(--el-color-danger-light-9) !important;
  background-color: var(--el-color-danger-light-9) !important;
}

.error-table .success-row {
  // --el-table-tr-bg-color: var(--el-color-success-light-9) !important;
  background-color: var(--el-color-success-light-9) !important;
}

.error-table .verbose-row {
  background-color: var(--el-color-warning-light-9) !important;
}

.right-panel {
  margin-left: 24px;
}

.el-input-number.full-width-input,
.el-cascader.full-width-input {
  width: 100% !important;
}

.el-form-item--medium {
  .el-radio {
    line-height: 36px !important;
  }

  .el-rate {
    margin-top: 8px;
  }
}

.el-form-item--small {
  .el-radio {
    line-height: 32px !important;
  }

  .el-rate {
    margin-top: 6px;
  }
}

.el-form-item--mini {
  .el-radio {
    line-height: 28px !important;
  }

  .el-rate {
    margin-top: 4px;
  }
}

.clear-fix:before,
.clear-fix:after {
  display: table;
  content: "";
}

.clear-fix:after {
  clear: both;
}

.float-right {
  float: right;
}
</style>

<style lang="scss" scoped>
div.table-container {
  table.table-layout {
    width: 100%;
    table-layout: fixed;
    border-collapse: collapse;

    td.table-cell {
      display: table-cell;
      height: 36px;
      border: 1px solid #e1e2e3;
    }
  }
}

div.tab-container {}

.label-left-align :deep(.el-form-item__label) {
  text-align: left;
}

.label-center-align :deep(.el-form-item__label) {
  text-align: center;
}

.label-right-align :deep(.el-form-item__label) {
  text-align: right;
}

.custom-label {}

.static-content-item {
  min-height: 20px;
  display: flex;
  align-items: center;

  :deep(.el-divider--horizontal) {
    margin: 0;
  }
}

@import 'https://cdn.jsdelivr.net/npm/@vue-flow/core@1.4.1/dist/style.css';
@import 'https://cdn.jsdelivr.net/npm/@vue-flow/core@1.4.1/dist/theme-default.css';

.vue-flow__minimap {
  transform: scale(75%);
  transform-origin: bottom right;
}

.dndflow {
  flex-direction: column;
  display: flex;
  height: 440px;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background: #EFF5F5;
  background-image: linear-gradient(rgba(148, 148, 148, 0.7) .1em, transparent .1em), linear-gradient(90deg, rgba(148, 148, 148, 0.7) .1em, transparent .1em);
  background-size: 3em 3em;
}

.dndflow aside {
  color: #fff;
  font-weight: 700;
  border-right: 1px solid #eee;
  padding: 15px 10px;
  font-size: 12px;
  background: rgba(16, 185, 129, .75);
  -webkit-box-shadow: 0px 5px 10px 0px rgba(0, 0, 0, .3);
  box-shadow: 0 5px 10px #0000004d
}

.dndflow aside .nodes>* {
  margin-bottom: 10px;
  cursor: grab;
  font-weight: 500;
  -webkit-box-shadow: 5px 5px 10px 2px rgba(0, 0, 0, .25);
  box-shadow: 5px 5px 10px 2px #00000040
}

.dndflow aside .description {
  margin-bottom: 10px
}

.dndflow .vue-flow-wrapper {
  flex-grow: 1;
  height: 100%
}

@media screen and (min-width: 640px) {
  .dndflow {
    flex-direction: row
  }

  .dndflow aside {
    min-width: 25%
  }
}

@media screen and (max-width: 639px) {
  .dndflow aside .nodes {
    display: flex;
    flex-direction: row;
    gap: 5px
  }
}
</style>