<template>
    <div class="actions">
        <div class="action-item">
            <el-button type="success" class="save padding-8" @click="saveCurrentData">Save Configuration for Order {{
                    currentRow.order
            }}</el-button>
        </div>
    </div>
    <el-row :gutter="24">
        <el-col :span="12">
            <div class="grid-content bg-white padding-18">
                <div class="zalo" v-if="dialogDataActionSendZNSVisible">
                    <el-form :model="znsCurrentTemplateData" label-position="right" ref="elFormRef" :rules="rule"
                        label-width="120px">
                        <el-form-item label="Template:" prop="id">
                            <el-select v-model="znsCurrentTemplateData.id" @change="znsChangeTemplate" clearable
                                placeholder="Select">
                                <el-option v-for="item in ZNSTemplateOptions" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                        <el-form-item label="Variables">
                            <div class="flex justify-space-between mb-4 flex-wrap gap-4">
                                <el-tag class="el-tag-margin-4" v-for="variable in variables" :key="variable"
                                    :type="variable" text bg>
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
                </div>

                <div class="email" v-if="dialogDataActionSendEmailVisible">
                        <el-form :model="emailTemplateSelected" label-position="right" label-width="120px">
                            <el-form-item label=" Select Template:" prop="ID">
                                <el-select v-model="emailTemplateSelected.ID" clearable placeholder="Select">
                                    <el-option v-for="item in emailTemplateOptions" :key="item.ID" :label="item.name"
                                        :value="item.ID" />
                                </el-select>
                            </el-form-item>
                        </el-form>
                </div>
            </div>
        </el-col>
        <el-col :span="12">
            <div class="grid-content ep-bg-purple">
                <el-table :data="sequenceRef" style="width: 100%">
                    <el-table-column label="Order" width="80">
                        <template #default="scope">
                            <el-input type="number" readonly v-model="scope.row.order"></el-input>
                        </template>
                    </el-table-column>
                    <el-table-column label="Date">
                        <template #default="scope">
                            <el-date-picker type="datetime" v-model="scope.row.date" width="120"></el-date-picker>
                        </template>
                    </el-table-column>
                    <!-- <el-table-column label="Time">
                        <template #default="scope">
                            <el-time-picker v-model="scope.row.time" width="120"></el-time-picker>
                        </template>
                    </el-table-column> -->
                    <el-table-column label="Action">
                        <template #default="scope">
                            <el-select class="m-2" placeholder="Select" size="large" v-model="scope.row.actionName"
                                @change="changeActionInSequence(scope.row)">
                                <el-option label="Send Zalo Notificaton Service" value="action-send-zns" />
                                <el-option label="Send Messenger" value="action-send-messenger" />
                                <el-option label="Send Email" value="action-send-email" />
                            </el-select>
                        </template>
                    </el-table-column>
                    <el-table-column fixed="right" label="Action" width="150">
                        <template #default="scope">
                            <el-button link type="primary" size="medium" @click.prevent="onSequenceConfig(scope.row)">
                                Config
                            </el-button>
                            <el-button link type="danger" size="medium"
                                @click.prevent="onSequenceRemove(scope.$index, scope.row)">
                                Remove
                            </el-button>
                        </template>
                    </el-table-column>
                </el-table>
                <el-button class="mt-4" style="width: 100%" @click="onSequenceAdd()">Add Row</el-button>
            </div>
        </el-col>
    </el-row>
</template>

<script setup>

import { nextTick, watch, ref, reactive } from 'vue'


import {
    updateCampaign,
    findCampaign,
} from '@/api/campaign'

import {
    getZaloNotificationTemplate
} from '@/api/zaloApplication'

import { getEmailTemplateList } from '@/api/emailTemplate'

import {
    createSequence,
    deleteSequence,
    deleteSequenceByIds,
    updateSequence,
    findSequence,
    getSequenceByCampaign,
} from '@/api/sequence'

import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'


const router = useRouter()

const route = useRoute()
const campaigndata = ref({})

const currentRow = ref({})

const searchInfo = ref({ campaignID: Number(route.params.id) })



const getCampaign = async () => {
    var id = searchInfo.value.campaignID
    var res = await findCampaign({ ID: id })
    if (res.code === 0) {
        campaigndata.value = res.data.recampaign
        getZnsTemplates()
    }
}

getCampaign()

const sequences = ref([])
const getSequences = async () => {
    var id = searchInfo.value.campaignID
    var res = await getSequenceByCampaign({ campaignId: id })
    if (res.code === 0) {
        var temp = res.data.resequence
        temp.forEach(element => {
            if (element.data) {
                element.data = JSON.parse(element.data)
            } else {
                element.data = {}
            }
        });
        sequenceRef.value = temp
    }
}

getSequences()


const ZNSTemplateOptions = ref([])

const getZnsTemplates = async () => {
    var res = await getZaloNotificationTemplate(campaigndata.value.zaloApplication)
    if (res.code === 0) {
        ZNSTemplateOptions.value = res.data
    }
}

const sequenceRef = ref([])

const emailTemplateSelected = ref(
    {
        ID: 0
    }
)
const changeActionInSequence = async (row) => {
    console.log("changeActionInSequence")
    type.value = "update"
    await enterSequence(row)

}


const znsCurrentTemplateData = ref()
const dialogDataActionSendZNSVisible = ref(false)
const dialogDataActionSendEmailVisible = ref(false)

const openViewSetDataActionSendZNS = (data) => {
    if (typeof data == "undefined" || data == "") {
        znsCurrentTemplateData.value = ZNSTemplateOptions.value[0];
    } else {
        znsCurrentTemplateData.value = data
    }
    dialogDataActionSendZNSVisible.value = true
    dialogDataActionSendEmailVisible.value = false
}

const openViewSetDataActionSendEmail = (data) => {
    if (typeof data == "undefined") {
        emailTemplateSelected.value.ID = 1
    } else {
        emailTemplateSelected.value = data
    }
    dialogDataActionSendEmailVisible.value = true
    dialogDataActionSendZNSVisible.value = false;
}

const saveSendZNSConfig = async () => {
    var id = currentRow.value.ID;
    var newUpdate = getZNSTemplateById(znsCurrentTemplateData.value.id)
    newUpdate.data = znsCurrentTemplateData.value.data

    let sequenceOrder = getSequeneDataById(id)
    sequenceOrder.data = JSON.stringify(newUpdate)
    dialogDataActionSendZNSVisible.value = false;
    type.value = "update";
    await enterSequence(sequenceOrder)
}

const saveEmailConfig = async () => {
    var id = currentRow.value.ID;
    let sequenceOrder = getSequeneDataById(id)
    debugger
    sequenceOrder.data = JSON.stringify(emailTemplateSelected.value)
    dialogDataActionSendEmailVisible.value = false;
    type.value = "update";
    await enterSequence(sequenceOrder)
}

const getZNSTemplateById = (id) => {
    var template = ZNSTemplateOptions.value.filter((e) => e.id == id)
    return template[0]
}

const getSequeneDataById = (id) => {
    var vals = sequenceRef.value.filter((se) => {
        return se.ID == id
    })
    return vals[0]
}

const znsChangeTemplate = (value) => {
    var newTemplate = getZNSTemplateById(value);
    znsCurrentTemplateData.value.data = newTemplate.data
}


const emailTemplateOptions = ref([])
const getEmailTemplate = async () => {
  const res = await getEmailTemplateList({})
  if (res.code === 0) {
    emailTemplateOptions.value = res.data.list
    console.log(emailTemplateOptions.value)
  }
}

getEmailTemplate()


const saveSequence = async () => {
    closeDialog()
}

const onSequenceAdd = async () => {
    var nextOrder = sequenceRef.value.length + 1;
    var newSequence = {
        order: nextOrder - 1,
        date: new Date(),
        actionName: "",
        data: {},
        ID: null
    }
    type.value = "create"
    newSequence.campaignId = searchInfo.value.campaignID
    newSequence.data = '';
    var data = await enterSequence(newSequence)
    newSequence.ID = data.ID
    sequenceRef.value.push(newSequence)
}

const onSequenceRemove = async (index, row) => {
    const res = await deleteSequence({ ID: row.ID })
    if (res.code === 0) {
        sequenceRef.value.splice(index, 1)
        ElMessage({
            type: 'success',
            message: 'successfully deleted'
        })
    }
}


const onSequenceConfig = (row) => {
    currentRow.value = row
    if (row.actionName == "action-send-zns") {
        openViewSetDataActionSendZNS(row.data)
    }
    if (row.actionName == "action-send-email") {
        openViewSetDataActionSendEmail(row.data)
    }
}

const saveCurrentData = () => {
    if (currentRow.value.actionName == "action-send-zns") {
        saveSendZNSConfig()
    }
    if (currentRow.value.actionName == "action-send-email") {
        saveEmailConfig()
    }

}

const type = ref('')

// 更新行
const updateSequenceFunc = async (row) => {
    const res = await findSequence({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.resequence
        dialogFormVisible.value = true
    }
}


// Delete 行
const deleteSequenceFunc = async (row) => {
    const res = await deleteSequence({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: 'successfully deleted'
        })
    }
}

const enterSequence = async (sequence) => {
    let res;
    switch (type.value) {
        case 'create':
            res = await createSequence(sequence)
            break
        case 'update':
            res = await updateSequence(sequence)
            break
        default:
            res = await createSequence(sequence)
            break
    }
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: 'Create /Update successfully'
        })
        return res.data
    }
    return null
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


<style>
.sequence-ref-action {
    display: flex;
    flex-direction: column;
    justify-content: center;
}

.el-tag-margin-4 {
    padding: 8px;
    margin: 4px;
}

.padding-8 {
    padding: 8px;
}

.padding-18 {
    padding: 18px;
}

.save {
    margin-bottom: 12px;
}

.bg-white {
    background-color: #fff;
}
</style>
