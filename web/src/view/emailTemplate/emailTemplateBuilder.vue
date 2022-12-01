<template>
  <div id="emailBuilder">
    <div class="container">
      <el-row class="mb-4">
        <el-button type="primary" @click="exportHtml">Apply Tempalte</el-button>
        <el-button type="success" @click="openVariableDialog">Variables</el-button>
        <el-button type="info" @click="openSendDialog">Send Test</el-button>
      </el-row>

      <EmailEditor ref="emailEditor" v-on:load="editorLoaded" v-on:ready="editorReady" />
    </div>
  </div>

  <el-dialog v-model="dialogSendTestVisible" :before-close="closeDialog" title="Pop-up">
    <el-form :model="testEmailData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
      <el-form-item label="Name:" prop="name">
        <el-input v-model="testEmailData.to" :clearable="true" placeholder="Please enter" />
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button size="small" @click="closeDialog">Cancel</el-button>
        <el-button size="small" type="primary" @click="sendEmailTest">Send</el-button>
      </div>
    </template>
  </el-dialog>

  <el-dialog v-model="dialogVariableVisible" :before-close="closeDialog" title="Pop-up">
    <el-form label-width="120px">
      <el-form-item label="Variables">
        <div class="flex justify-space-between mb-4 flex-wrap gap-4">
          <el-tag class="el-tag-margin-4" v-for="variable in variables" :key="variable" :type="variable" text bg>
            {{ variable }}</el-tag>
        </div>
        
      </el-form-item>
      <el-form-item>
        <div>Variable only use in string content</div>
      </el-form-item>

    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button size="small" @click="closeDialog">Cancel</el-button>
      </div>
    </template>
  </el-dialog>

</template>

<script>
export default {
  name: 'emailBuilder',
}
</script>


<script setup>
import {
  updateEmailTemplate,
  findEmailTemplate,
} from '@/api/emailTemplate'

import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRoute, onBeforeRouteUpdate } from 'vue-router'
import { EmailEditor } from 'vue-email-editor';
import { ref, watch } from 'vue'
import { sendEmail } from '@/api/emailTemplate.js'


const route = useRoute()


const emailEditor = ref({
  height: '800px',
  locale: 'en',
  projectId: 0, // replace with your project id
  tools: {
    // disable image tool
    image: {
      enabled: true,
    },
  },
  options: {},
  appearance: {
    theme: 'dark',
    panels: {
      tools: {
        dock: 'right',
      },
    },
  },
})

const email = ref(null);

const searchInfo = ref({ emailTemplateId: Number(route.params.id) })
const variables = ref(
  [
    "__CID__", "__LASTNAME__", "__FIRSTNAME__", "__PHONE__", "__EMAIL__",
    "__CAMPAIGN_NAME__", "__CAMPAIGN_START__", "__CAMPAIGN_END__",
    "__ZALO_PHONE__", "__FB_ID__", "__ADDRESS__", "__CITY__", "__STATE__", "__ZIPCODE__", "__COUNTRY__",
    "__DATE__", "__TIME__", "__CREATED_AT__"
  ]
)
const editorLoaded = async () => {
  console.log('editorLoaded');
  if (email.value.content) {
    const emailContent = JSON.parse(email.value.content);
    emailEditor.value.editor.loadDesign(emailContent.design);
  }

}
// called when the editor has finished loading
const editorReady = () => {
  console.log('editorReady');
}
// const saveDesign = () => {
//   emailEditor.value.editor.saveDesign((design) => {
//     console.log('saveDesign', design);
//   });
// }
const exportHtml = () => {
  emailEditor.value.editor.exportHtml((data) => {
    updateEmailContent(data);
  });
}

const getEmail = async () => {
  // console.log('getEmail');
  const id = searchInfo.value.emailTemplateId;
  const response = await findEmailTemplate({ ID: id });
  email.value = response.data.reemail_template;
  // console.log(email.value);
  // console.log(emailEditor.value);
  // emailEditor.value.projectId = email.value.ID 
  editorLoaded();
}

getEmail()



const updateEmailContent = async (data) => {
  const id = email.value.ID;
  debugger;
  email.value.content = JSON.stringify(data);
  const res = await updateEmailTemplate(email.value);
  console.log(res);
  debugger;
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


const testEmailData = ref({
  to: 'tinh.phan@udoo.ooo'
})
const dialogSendTestVisible = ref(false);
const dialogVariableVisible = ref(false);
const openSendDialog = () => {
  dialogSendTestVisible.value = true;
}

const openVariableDialog = () => {
  dialogVariableVisible.value = true;
}
const sendEmailTest = async () => {
  dialogSendTestVisible.value = false;
  //send email test

  var mailTo = testEmailData.value.to;
  var emailContent = JSON.parse(email.value.content);
  var contentHTML = emailContent.html;
  var emailData = {
    to: mailTo,
    subject: email.value.subject,
    body: contentHTML
  }
  var res = await sendEmail(emailData)
  // var res = await emailTest()
  if (res.code == 200) {
    ElMessage({
      type: 'success',
      message: 'Send Email successfully'
    })
  } else {
    ElMessage({
      type: 'warning',
      message: 'Send email failed'
    })
  }

}

const closeDialog = () => {
  dialogSendTestVisible.value = false
  dialogVariableVisible.value = false
}


</script>

<style>
iframe {
  height: 770px !important;
}
</style>

