<template>
  <div class="system">
    <el-form ref="form" :model="config" label-width="240px">
      <!--  System start  -->
      <el-collapse v-model="activeNames">
        <el-collapse-item title="System Configuration" name="1">
          <el-form-item label="Environmental value">
            <!-- <el-input v-model="config.system.env" />-->
            <el-select v-model="config.system.env" style="width:100%">
              <el-option value="public" />
              <el-option value="develop" />
            </el-select>
          </el-form-item>
          <el-form-item label="Port">
            <el-input v-model.number="config.system.addr" />
          </el-form-item>
          <el-form-item label="Database Type">
            <el-select v-model="config.system['db-type']" style="width:100%">
              <el-option value="mysql" />
              <el-option value="pgsql" />
            </el-select>
          </el-form-item>
          <el-form-item label="Oss Type">
            <el-select v-model="config.system['oss-type']" style="width:100%">
              <el-option value="local" />
              <el-option value="qiniu" />
              <el-option value="tencent-cos" />
              <el-option value="aliyun-oss" />
              <el-option value="huawei-obs" />
            </el-select>
          </el-form-item>
          <el-form-item label="Multi-point login interception">
            <el-checkbox v-model="config.system['use-multipoint']">Open</el-checkbox>
          </el-form-item>
          <el-form-item label="Open Redis">
            <el-checkbox v-model="config.system['use-redis']">Open</el-checkbox>
          </el-form-item>
          <el-form-item label="IP Limit count">
            <el-input-number v-model.number="config.system['iplimit-count']" />
          </el-form-item>
          <el-form-item label="IP Limit time">
            <el-input-number v-model.number="config.system['iplimit-time']" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="Jwt signature" name="2">
          <el-form-item label="Jwt signature">
            <el-input v-model="config.jwt['signing-key']" />
          </el-form-item>
          <el-form-item label="Expires Time">
            <el-input v-model="config.jwt['expires-time']" />
          </el-form-item>
          <el-form-item label="Buffer Time">
            <el-input v-model="config.jwt['buffer-time']" />
          </el-form-item>
          <el-form-item label="Issuer">
            <el-input v-model="config.jwt.issuer" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="ZAP log configuration" name="3">
          <el-form-item label="Level">
            <el-input v-model.number="config.zap.level" />
          </el-form-item>
          <el-form-item label="Format">
            <el-input v-model="config.zap.format" />
          </el-form-item>
          <el-form-item label="Prefix">
            <el-input v-model="config.zap.prefix" />
          </el-form-item>
          <el-form-item label="Director">
            <el-input v-model="config.zap.director" />
          </el-form-item>
          <el-form-item label="Encode Level">
            <el-input v-model="config.zap['encode-level']" />
          </el-form-item>
          <el-form-item label="Stracktrace Key">
            <el-input v-model="config.zap['stacktrace-key']" />
          </el-form-item>
          <el-form-item label="Log retention time (the default is based on heaven)">
            <el-input v-model.number="config.zap['max-age']" />
          </el-form-item>
          <el-form-item label="Show line">
            <el-checkbox v-model="config.zap['show-line']" />
          </el-form-item>
          <el-form-item label="Log in console">
            <el-checkbox v-model="config.zap['log-in-console']" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="Redis admin database configuration" name="4">
          <el-form-item label="Libary">
            <el-input v-model.number="config.redis.db" />
          </el-form-item>
          <el-form-item label="Address">
            <el-input v-model="config.redis.addr" />
          </el-form-item>
          <el-form-item label="Password">
            <el-input v-model="config.redis.password" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item title="Mailbox configuration" name="5">
          <el-form-item label="Receiver mailbox">
            <el-input v-model="config.email.to" placeholder="Can be separated by commas" />
          </el-form-item>
          <el-form-item label="Port">
            <el-input v-model.number="config.email.port" />
          </el-form-item>
          <el-form-item label="Sender mailbox">
            <el-input v-model="config.email.from" />
          </el-form-item>
          <el-form-item label="host">
            <el-input v-model="config.email.host" />
          </el-form-item>
          <el-form-item label="Is SSL">
            <el-checkbox v-model="config.email['is-ssl']" />
          </el-form-item>
          <el-form-item label="secret">
            <el-input v-model="config.email.secret" />
          </el-form-item>
          <el-form-item label="Test email">
            <el-button @click="email">Test email</el-button>
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="Verification Code configuration" name="7">
          <el-form-item label="Character length">
            <el-input v-model.number="config.captcha['key-long']" />
          </el-form-item>
          <el-form-item label="Image width">
            <el-input v-model.number="config.captcha['img-width']" />
          </el-form-item>
          <el-form-item label="Image Height">
            <el-input v-model.number="config.captcha['img-height']" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="Database configuration" name="9">
          <template v-if="config.system['db-type'] === 'mysql'">
            <el-form-item label="Username">
              <el-input v-model="config.mysql.username" />
            </el-form-item>
            <el-form-item label="Password">
              <el-input v-model="config.mysql.password" />
            </el-form-item>
            <el-form-item label="Path">
              <el-input v-model="config.mysql.path" />
            </el-form-item>
            <el-form-item label="Database name">
              <el-input v-model="config.mysql['db-name']" />
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input v-model.number="config.mysql['max-idle-conns']" />
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input v-model.number="config.mysql['max-open-conns']" />
            </el-form-item>
            <el-form-item label="Write Log">
              <el-checkbox v-model="config.mysql['log-zap']" />
            </el-form-item>
            <el-form-item label="Log Mode">
              <el-input v-model="config.mysql['log-mode']" />
            </el-form-item>
          </template>
          <template v-if="config.system.dbType === 'pgsql'">
            <el-form-item label="User name">
              <el-input v-model="config.pgsql.username" />
            </el-form-item>
            <el-form-item label="Password">
              <el-input v-model="config.pgsql.password" />
            </el-form-item>
            <el-form-item label="Path">
              <el-input v-model="config.pgsql.path" />
            </el-form-item>
            <el-form-item label="Database name">
              <el-input v-model="config.pgsql.dbname" />
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input v-model.number="config.pgsql['max-idle-conns']" />
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input v-model.number="config.pgsql['max-open-conns']" />
            </el-form-item>
            <el-form-item label="Write Log">
              <el-checkbox v-model="config.pgsql['log-zap']" />
            </el-form-item>
            <el-form-item label="Log mode">
              <el-input v-model="config.pgsql['log-mode']" />
            </el-form-item>
          </template>
        </el-collapse-item>

        <el-collapse-item title="OSS configuration" name="10">
          <template v-if="config.system['oss-type'] === 'local'">
            <h2>Local file configuration</h2>
            <el-form-item label="Local file access path">
              <el-input v-model="config.local.path" />
            </el-form-item>
            <el-form-item label="Local file strogare path">
              <el-input v-model="config.local['store-path']" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'qiniu'">
            <h2>qiniu upload configuration</h2>
            <el-form-item label="Storage area">
              <el-input v-model="config.qiniu.zone" />
            </el-form-item>
            <el-form-item label="Space name">
              <el-input v-model="config.qiniu.bucket" />
            </el-form-item>
            <el-form-item label="CDN acceleration domain name">
              <el-input v-model="config.qiniu['img-path']" />
            </el-form-item>
            <el-form-item label="Use HTTPS">
              <el-checkbox v-model="config.qiniu['use-https']">Open</el-checkbox>
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model="config.qiniu['access-key']" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config.qiniu['secret-key']" />
            </el-form-item>
            <el-form-item label="Use CDN domains">
              <el-checkbox v-model="config.qiniu['use-cdn-domains']">Open</el-checkbox>
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'tencent-cos'">
            <h2>Tencent Cloud COS upload configuration</h2>
            <el-form-item label="Reserve bucket name">
              <el-input v-model="config['tencent-cos']['bucket']" />
            </el-form-item>
            <el-form-item label="Affiliated area">
              <el-input v-model="config['tencent-cos'].region" />
            </el-form-item>
            <el-form-item label="secretID">
              <el-input v-model="config['tencent-cos']['secret-id']" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config['tencent-cos']['secret-key']" />
            </el-form-item>
            <el-form-item label="Affiliated area">
              <el-input v-model="config['tencent-cos']['path-prefix']" />
            </el-form-item>
            <el-form-item label="Access domain name">
              <el-input v-model="config['tencent-cos']['base-url']" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'aliyun-oss'">
            <h2>Alibaba Cloud OSS upload configuration</h2>
            <el-form-item label="Area">
              <el-input v-model="config['aliyun-oss'].endpoint" />
            </el-form-item>
            <el-form-item label="accessKeyId">
              <el-input v-model="config['aliyun-oss']['access-key-id']" />
            </el-form-item>
            <el-form-item label="accessKeySecret">
              <el-input v-model="config['aliyun-oss']['access-key-secret']" />
            </el-form-item>
            <el-form-item label="Reserve bucket name">
              <el-input v-model="config['aliyun-oss']['bucket-name']" />
            </el-form-item>
            <el-form-item label="Access domain name">
              <el-input v-model="config['aliyun-oss']['bucket-url']" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'huawei-obs'">
            <h2>Huawei Cloud OBS upload configuration</h2>
            <el-form-item label="Path">
              <el-input v-model="config['hua-wei-obs'].path" />
            </el-form-item>
            <el-form-item label="Reserve bucket name">
              <el-input v-model="config['hua-wei-obs'].bucket" />
            </el-form-item>
            <el-form-item label="Area">
              <el-input v-model="config['hua-wei-obs'].endpoint" />
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model="config['hua-wei-obs']['access-key']" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config['hua-wei-obs']['secret-key']" />
            </el-form-item>
          </template>

        </el-collapse-item>

        <el-collapse-item title="Excel upload configuration" name="11">
          <el-form-item label="Synthetic target address">
            <el-input v-model="config.excel.dir" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item title="Automated code configuration" name="12">
          <el-form-item label="Transfer restart(linux)">
            <el-checkbox v-model="config.autocode['transfer-restart']" />
          </el-form-item>
          <el-form-item label="root(Project root path)">
            <el-input v-model="config.autocode.root" disabled />
          </el-form-item>
          <el-form-item label="Server(Back -end code address)">
            <el-input v-model="config.autocode['transfer-restart']" />
          </el-form-item>
          <el-form-item label="SApi(Back -end API folder address)">
            <el-input v-model="config.autocode['server-api']" />
          </el-form-item>
          <el-form-item label="SInitialize(Back end Initialize folder)">
            <el-input v-model="config.autocode['server-initialize']" />
          </el-form-item>
          <el-form-item label="SModel(Back -end Model file address)">
            <el-input v-model="config.autocode['server-model']" />
          </el-form-item>
          <el-form-item label="SRequest(Request folder address)">
            <el-input v-model="config.autocode['server-request']" />
          </el-form-item>
          <el-form-item label="SRouter(Back end Router folder address)">
            <el-input v-model="config.autocode['server-router']" />
          </el-form-item>
          <el-form-item label="SService(The back -end service folder address)">
            <el-input v-model="config.autocode['server-service']" />
          </el-form-item>
          <el-form-item label="Web(Folder folder address)">
            <el-input v-model="config.autocode.web" />
          </el-form-item>
          <el-form-item label="WApi(Back -end wapi folder address)">
            <el-input v-model="config.autocode['web-api']" />
          </el-form-item>
          <el-form-item label="WForm(Back -end wform folder address)">
            <el-input v-model="config.autocode['web-form']" />
          </el-form-item>
          <el-form-item label="WTable(Back end wtable folder address)">
            <el-input v-model="config.autocode['web-table']" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item title="Timer (Timing Task)" name="13">
          <el-form-item label="Start（Is Enable）">
            <el-checkbox v-model="config.timer['start']" />
          </el-form-item>
          <el-form-item label="Spec(Cron expression)">
            <el-input v-model="config.timer.spec" />
          </el-form-item>
          <template v-for="(item,k) in config.timer.detail">
            <div v-for="(key,k2) in item" :key="k2">
              <el-form-item :key="k+k2" :label="k2">
                <el-input v-model="item[k2]" />
              </el-form-item>
            </div>
          </template>
        </el-collapse-item>
      </el-collapse>
    </el-form>
    <div class="gva-btn-list">
      <el-button type="primary" size="small" @click="update">Update immediately</el-button>
      <el-button type="primary" size="small" @click="reload">Restart service (under development)</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Config'
}
</script>
<script setup>
import { getSystemConfig, setSystemConfig } from '@/api/system'
import { emailTest } from '@/api/email'
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

const activeNames = reactive([])
const config = ref({
  system: {
    'iplimit-count': 0,
    'iplimit-time': 0
  },
  jwt: {},
  mysql: {},
  pgsql: {},
  excel: {},
  autocode: {},
  redis: {},
  qiniu: {},
  'tencent-cos': {},
  'aliyun-oss': {},
  'hua-wei-obs': {},
  captcha: {},
  zap: {},
  local: {},
  email: {},
  timer: {
    detail: {}
  }
})

const initForm = async() => {
  const res = await getSystemConfig()
  if (res.code === 0) {
    config.value = res.data.config
  }
}
initForm()
const reload = () => {}
const update = async() => {
  const res = await setSystemConfig({ config: config.value })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Update successfully!'
    })
    await initForm()
  }
}
const email = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Mail sent successfully'
    })
    await initForm()
  } else {
    ElMessage({
      type: 'error',
      message: 'Email sending failuer'
    })
  }
}

</script>

<style lang="scss">
.system {
  background: #fff;
  padding:36px;
  border-radius: 2px;
  h2 {
    padding: 10px;
    margin: 10px 0;
    font-size: 16px;
    box-shadow: -4px 0px 0px 0px #e7e8e8;
  }
  ::v-deep(.el-input-number__increase){
    top:5px !important;
  }
  .gva-btn-list{
    margin-top:16px;
  }
}
</style>
