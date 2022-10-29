<template>
  <div id="emailBuilder">
    <div class="container">
      <!-- <div id="bar">
        <h1>Vue Email Editor</h1>

        <button v-on:click="saveDesign">Save Design</button>
        <button v-on:click="exportHtml">Export HTML</button>
      </div> -->
      <el-row class="mb-4">
        <el-button type="primary" @click="saveDesign">Save</el-button>
        <el-button type="success" @click="exportHtml">Export HTML</el-button>
      </el-row>

      <EmailEditor ref="emailEditor" v-on:load="editorLoaded" v-on:ready="editorReady" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'emailBuilder',
  components: {
    EmailEditor
  },
  data() {
    return {
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
    };
  },

  methods: {
    editorLoaded() {
      console.log('editorLoaded');
      // this.$refs.emailEditor.editor.loadDesign({});
      // this.emailPanel = this.$refs.emailEditor;
      let emailPanel = this.$refs.emailEditor;
    },
    // called when the editor has finished loading
    editorReady() {
      console.log('editorReady');
    },
    saveDesign() {
      this.$refs.emailEditor.editor.saveDesign((design) => {
        console.log('saveDesign', design);
      });
    },
    exportHtml() {
      this.$refs.emailEditor.editor.exportHtml((data) => {
        
        console.log('exportHtml', data);
        this.updateEmailContent(data);
      });
    },

  }
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


const route = useRoute()

const email = ref(null);

const searchInfo = ref({ emailTemplateId: Number(route.params.id) })

const getEmail = async () => {
  console.log('getEmail');
  const id = searchInfo.value.emailTemplateId;
  const response = await findEmailTemplate({ ID: id });
  email.value = response.data.reemail_template;
  const emailContent = JSON.parse(email.value.content);
  console.log(emailContent);
  // console.log(this.$refs)
}

getEmail()

onBeforeRouteUpdate((to, form) => {
console.log('onBeforeRouteUpdate');
})


// watch(() => {
  
//   console.log('watch');
//   // console.log(this.$refs);
// })


// watch(() => this.$refs.emailEditor, (editor) => {
//   // searchInfo.value.ID = Number(id)
//   // getDocument()
//   console.log('watch');
// })

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



</script>

<style>
iframe {
  height: 770px !important;
}
</style>

