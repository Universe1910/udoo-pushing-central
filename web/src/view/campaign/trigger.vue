
<script setup>
import { VueFlow, useVueFlow } from '@vue-flow/core'
import { nextTick, watch, ref, reactive } from 'vue'
import Sidebar from '@/components/trigger/sidebar.vue'
import SidebarDetail from '@/components/trigger/sidebar_detail.vue'

import {
  updateCampaign,
  findCampaign,
} from '@/api/campaign'

import {
  createTrigger,
  updateTrigger,
} from '@/api/trigger'


import {
  createZaloApplication,
  deleteZaloApplication,
  deleteZaloApplicationByIds,
  updateZaloApplication,
  findZaloApplication,
  getZaloApplicationList,
  getZaloNotificationTemplate

} from '@/api/zaloApplication'

import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'

import { makeTriggerActions } from '@/utils/trigger.js'


const router = useRouter()

const route = useRoute()
const campaigndata = ref({

})

const searchInfo = ref({ campaignID: Number(route.params.id) })

const getCampaign = async () => {
  var id = searchInfo.value.campaignID
  var res = await findCampaign({ ID: id })
  if (res.code === 0) {
    campaigndata.value = res.data.recampaign
    console.log(campaigndata.value)
    var tf = campaigndata.value.triggerObject.flow;
    var flow = tf.length > 0 ? JSON.parse(tf) : null
    onRestore(flow)
  }
  getZnsTemplates()
}

getCampaign()

const getZnsTemplates = async () => {
  var res = await getZaloNotificationTemplate(campaigndata.value.zaloApplication)
  if (res.code === 0) {
    ZNSTemplateOptions.value = res.data
    console.log(ZNSTemplateOptions.value)
  }
}

const getId = () => { 
  debugger;
  var len = nodes.value.length;
  if (typeof len == "undefined") {
    len = 0;
  }
  var nextId = len  +1
  return `dndnode_${nextId}`
}



const { findNode, onConnect, setNodes, setEdges, dimensions, setTransform, nodes, edges, addEdges, addNodes, viewport, project, vueFlowRef, toObject } = useVueFlow({
  nodes: [
  ],
})

const onRestore = (flow) => {
  if (flow) {
    const [x = 0, y = 0] = flow.position
    nodes.value = flow.nodes
    setNodes(flow.nodes)
    flowSetEvent(flow.nodes)
    setEdges(flow.edges)
    setTransform({ x, y, zoom: flow.zoom || 0 })
  }

}

const flowSetEvent = () => {
  nodes.value.forEach(element => {
    element.events = {
      click: () => {
        currentSelectedNodeId.value = element.id
        if (element.actionName == 'action-send-zns') {
          openDialogSetDataActionSendZNS(element.data)
        }
        if (element.actionName == "utils-wait") {
          openDialogSetDataUtils(element.data)
        }
      },
    }
  });
}

const onDragOver = (event) => {
  event.preventDefault()

  if (event.dataTransfer) {
    event.dataTransfer.dropEffect = 'move'
  }
}

onConnect((params) => addEdges([params]))



const onDrop = (event) => {
  const type = event.dataTransfer?.getData('application/vueflow')
  const label = event.dataTransfer?.getData('nodeLabel')
  const actionName = event.dataTransfer?.getData('actionName')
  const { left, top } = vueFlowRef.value.getBoundingClientRect()

  const position = project({
    x: event.clientX - left,
    y: event.clientY - top,
  })

  var newNodeId = getId()
  const newNode = {
    id: newNodeId,
    type,
    position,
    label: label,
    data: {},
    actionName: actionName,
    events: {
      click: () => {
        currentSelectedNodeId.value = newNodeId
        console.log("currentSelectedNodeId: " + currentSelectedNodeId.value)
        if (actionName == 'action-send-zns') {
          openDialogSetDataActionSendZNS()
        }
        if (actionName == "utils-wait") {
          openDialogSetDataUtils()
        }
      },
    }
  }



  addNodes([newNode])

  // align node position after drop, so it's centered to the mouse
  nextTick(() => {
    const node = findNode(newNode.id)
    const stop = watch(
      () => node.dimensions,
      (dimensions) => {
        if (dimensions.width > 0 && dimensions.height > 0) {
          node.position = { x: node.position.x - node.dimensions.width / 2, y: node.position.y - node.dimensions.height / 2 }
          stop()
        }
      },
      { deep: true, flush: 'post' },
    )
  })
}

const convertFlowObjectToFlowData = () => {
  var tf = JSON.stringify(toObject())
  var flow = tf.length > 0 ? JSON.parse(tf) : null
  var dataObject = makeTriggerActions(flow.edges)
  return dataObject;

}

const saveTrigger = async () => {
  console.log('save trigger')
  var triggerObjectData = campaigndata.value.triggerObject.data;
  var triggerObjectFlow = campaigndata.value.triggerObject.flow;

  var dataVal = triggerObjectData.length == 0 ? null : JSON.parse(triggerObjectData)

  var flowVal = triggerObjectFlow.length == 0 ? null : JSON.parse(triggerObjectFlow)

  var flowObject = toObject()

  var dataObject = convertFlowObjectToFlowData()

  var triggerData = {
    flow: JSON.stringify(flowObject),
    campaignId: campaigndata.value.ID,
    data: JSON.stringify(dataObject),
    status: "active"
  }
  let res
  if (!dataVal && !flowVal) {
    console.log('create Trigger');
    res = await createTrigger(triggerData)
  } else {
    var triggerObjectUpdate = campaigndata.value.triggerObject
    triggerObjectUpdate.flow = JSON.stringify(flowObject)
    triggerObjectUpdate.data = JSON.stringify(dataObject)
    res = await updateTrigger(triggerObjectUpdate)
  }
  if (res.code === 0) {
    campaigndata.value.triggerId = res.data.ID
    var updateCampaingRes = await updateCampaign(campaigndata.value)
    if (updateCampaingRes.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Create /Update successfully'
      })
    }
  }

  closeDialog()
}

const type = ref('')

const dialogDataActionSendZNSVisible = ref(false)
const dialogDataUtilsWaitVisible = ref(false)

const actionSendZNSData = ref(
  {
    znsTemplate: ''
  }
)

const openDialogSetDataActionSendZNS = (data) => {
  type.value = 'create'
  if (typeof data == "undefined") {
    znsCurrentTemplateData.value = ZNSTemplateOptions.value[0];
  } else {
    znsCurrentTemplateData.value = data
  }
  dialogDataActionSendZNSVisible.value = true
}

const openDialogSetDataUtils = (data) => {
  if (typeof data == "undefined") {
    waitData.value.time = 0
  } else {
    waitData.value = data
  }
  dialogDataUtilsWaitVisible.value = true
}

const znsCurrentTemplateData = ref()
const waitData = ref({
  time: 0
})

const currentSelectedNodeId = ref()
const znsTemplateSelected = ref()

const ZNSTemplateOptions = ref([])

const saveSendZNSConfig = () => {
  var nodeId = currentSelectedNodeId.value;
  var newUpdate = getZNSTemplateById(znsCurrentTemplateData.value.id)
  newUpdate.data = znsCurrentTemplateData.value.data
  updateNodeData(newUpdate, nodeId)
  dialogDataActionSendZNSVisible.value = false;
}

const saveUtilsWait = () => {
  var nodeId = currentSelectedNodeId.value;
  updateNodeData(waitData.value, nodeId)
  dialogDataUtilsWaitVisible.value = false;
}

const getZNSTemplateById = (id) => {
  var template = ZNSTemplateOptions.value.filter((e) => e.id == id)
  return template[0]
}

const updateNodeData = (data, nodeId) => {
  const node = findNode(nodeId)
  node.data = data
}

const znsChangeTemplate = (value) => {
  var newTemplate = getZNSTemplateById(value);
  znsCurrentTemplateData.value.data = newTemplate.data

}

const closeDialog = () => {
  dialogDataActionSendZNSVisible.value = false;
  dialogDataUtilsWaitVisible.value = false;
}


const variables = ref(
  [
    "__CID__", "__LASTNAME__", "__FIRSTNAME__", "__PHONE__", "__EMAIL__",
    "__CAMPAIGN_NAME__", "__CAMPAIGN_START__", "__CAMPAIGN_END__",
    "__ZALO_PHONE__", "__FB_ID__", "__ADDRESS__", "__CITY__", "__STATE__", "__ZIPCODE__", "__COUNTRY__",
    "__DATE__", "__TIME__", "__CREATED_AT__"
  ]
)
</script>

<template>
  <div class="actions">
    <div class="action-item">
      <el-button type="primary" class="save padding-8" @click="saveTrigger">Save Trigger</el-button>
    </div>
  </div>
  <div class="dndflow" @drop="onDrop">

    <VueFlow @dragover="onDragOver" />
    <Sidebar />
  </div>
  <!-- dialog note send ZNS -->
  <el-dialog v-model="dialogDataActionSendZNSVisible" :before-close="closeDialog" title="Pop-up">
    <el-form :model="znsCurrentTemplateData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">

      <el-form-item label="Template:" prop="id">
        <el-select v-model="znsCurrentTemplateData.id" @change="znsChangeTemplate" clearable placeholder="Select">
          <el-option v-for="item in ZNSTemplateOptions" :key="item.id" :label="item.name" :value="item.id" />
        </el-select>
      </el-form-item>
      <el-form-item label="Variables">
        <div class="flex justify-space-between mb-4 flex-wrap gap-4">
          <el-tag class="el-tag-margin-4" v-for="variable in variables" :key="variable" :type="variable" text bg>
            {{ variable }}</el-tag>
        </div>

      </el-form-item>
      <el-form-item label="Template Data">
        <el-table :data="znsCurrentTemplateData.data" style="width: 100%">
          <el-table-column label="Key">
            <template #default="scope">
              <el-input type="text" readonly v-model="scope.row.replaceKey"></el-input>
            </template>
          </el-table-column>
          <el-table-column label="Value">
            <template #default="scope">
              <el-input type="text" v-model="scope.row.replaceVal"></el-input>
            </template>
          </el-table-column>
        </el-table>
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button size="small" @click="closeDialog">Cancel</el-button>
        <el-button size="small" type="primary" @click="saveSendZNSConfig">Save</el-button>
      </div>
    </template>

    //Wait dialog
  </el-dialog>
  <el-dialog v-model="dialogDataUtilsWaitVisible" :before-close="closeDialog" title="Pop-up">
    <el-form :model="waitData" label-position="right" :rules="rule" label-width="120px">
      <el-form-item label="Time to wait:">
        <el-input v-model="waitData.time" type="number" placeholder="Please input">
          <template #append>second</template>
        </el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button size="small" @click="closeDialog">Cancel</el-button>
        <el-button size="small" type="primary" @click="saveUtilsWait">Save</el-button>
      </div>
    </template>
  </el-dialog>
</template>

 
<style>
@import 'https://cdn.jsdelivr.net/npm/@vue-flow/core@1.4.1/dist/style.css';
@import 'https://cdn.jsdelivr.net/npm/@vue-flow/core@1.4.1/dist/theme-default.css';

.vue-flow__minimap {
  transform: scale(75%);
  transform-origin: bottom right;
}

.dndflow {
  flex-direction: column;
  display: flex;
  height: 780px;
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

.el-tag-margin-4 {
  padding: 8px;
  margin: 4px;
}

.padding-8 {
  padding: 8px;
}

.save {
  margin-bottom: 12px;
}
</style>
