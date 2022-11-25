
<template>
  <el-form :model="zaloBuilderForm" ref="vForm" :rules="rules" label-position="left" label-width="80px" size="medium"
    @submit.prevent>
    <div class="static-content-item">
      <div>Zalo Authentication</div>
    </div>
    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <el-row>
      <el-col :span="12" class="grid-cell">
        <el-form-item label="App ID" label-width="120px" prop=" applicationID">
          <el-input v-model="zaloBuilderForm. applicationID"   type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12" class="grid-cell">
        <el-form-item label="Code Challenge" label-width="120px" prop="codeChallenge">
          <el-input type="password" v-model="zaloBuilderForm.codeChallenge"   show-password clearable>
          </el-input>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12" class="grid-cell">
        <el-form-item label="OA ID" label-width="120px" prop="oaId">
          <el-input v-model="zaloBuilderForm.oaId"   type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12" class="grid-cell">
        <el-form-item label="Code Verifier" label-width="120px" prop="codeVerifier">
          <el-input v-model="zaloBuilderForm.codeVerifier"   type="password" show-password clearable>
          </el-input>
        </el-form-item>
      </el-col>
    </el-row>
    <el-form-item label="Code" label-width="120px" prop="codeAuthorization">
      <el-input type="textarea" v-model="zaloBuilderForm.codeAuthorization"   rows="6">
      </el-input>
    </el-form-item>
    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <el-row>
      <el-col :span="12" class="grid-cell">
        <div class="static-content-item">
          <el-button type="primary" @click="zaloAuthClick">Chứng thực</el-button>
        </div>
      </el-col>
      <el-col :span="12" class="grid-cell">
        <div class="static-content-item">
          <el-button type="success" @click="submitForm">Save</el-button>
        </div>
      </el-col>
    </el-row>
  </el-form>

</template>

<script>
export default {
  name: 'zaloAuth',
}
</script>

<script setup>
import {
  findZaloApplication,
  updateZaloApplication
} from '@/api/zaloApplication'

import { useRoute } from 'vue-router'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const route = useRoute()

const searchInfo = ref({ zaloAuthID: Number(route.params.id) })

// const zaloApplication = ref({
//    applicationID: "", 
//   codeChallenge: "", 
//   oaId: "",
//   codeVerifier: "",
//   codeAuthorization: "",

// })

const getZaloApplicationID = async () => {
  const id = searchInfo.value.zaloAuthID;
  const response = await findZaloApplication({ ID: id });
  console.log(response);
  zaloBuilderForm.value = response.data.rezaloApplication;
}

getZaloApplicationID()


const zaloAuthClick = () => {
   window.location.href = zaloBuilderForm.value.authorizationURL
}

const zaloBuilderForm = ref({
   applicationID: "",
  codeChallenge: "",
  oaId: "",
  codeCerifier: "",
  codeAuthorization: "",
})

const submitForm = async()=> {
  var res = await updateZaloApplication(zaloBuilderForm.value)
  if(res.code === 0){
    ElMessage({
        type: 'success',
        message: 'Update success'
      })
  }else{

  }
}
const resetForm = () => {
  instance.ctx.$refs['vForm'].resetFields()
}

</script>

<style lang="scss">
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
</style>