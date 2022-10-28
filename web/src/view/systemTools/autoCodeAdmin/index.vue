<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="goAutoCode(null)">Add new</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="id" width="60" prop="ID" />
        <el-table-column align="left" label="Date" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Struct Name" min-width="150" prop="structName" />
        <el-table-column align="left" label="Describe" min-width="150" prop="structCNName" />
        <el-table-column align="left" label="Table Name" min-width="150" prop="tableName" />
        <el-table-column align="left" label="Rollback" min-width="150" prop="flag">
          <template #default="scope">
            <el-tag v-if="scope.row.flag" type="danger" size="small" effect="dark">
              Roll back
            </el-tag>
            <el-tag v-else size="small" type="success" effect="dark">
              Not rolling back
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Operate" min-width="240">
          <template #default="scope">
            <div>
              <el-button size="small" type="primary" link :disabled="scope.row.flag === 1"
                @click="rollbackFunc(scope.row, true)">Roll back (delete tables)</el-button>
              <el-button size="small" type="primary" link :disabled="scope.row.flag === 1"
                @click="rollbackFunc(scope.row, false)">Roll back (not deleted tables)</el-button>
              <el-button size="small" type="primary" link @click="goAutoCode(scope.row)">Reuse</el-button>
              <el-button size="small" type="primary" link @click="deleteRow(scope.row)">Delete</el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
          layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AutoCodeAdmin',
}
</script>

<script setup>
import { getSysHistory, rollback, delSysHistory } from '@/api/autoCode.js'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { formatDate } from '@/utils/format'
const router = useRouter()

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getSysHistory({
    page: page.value,
    pageSize: pageSize.value
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const deleteRow = async (row) => {
  ElMessageBox.confirm('Will this operation continue the history of this history?', 'Hint', {
    confirmButtonText: 'Sure',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(async () => {
    const res = await delSysHistory({ id: Number(row.ID) })
    if (res.code === 0) {
      ElMessage.success('Delete success')
      getTableData()
    }
  })
}
const rollbackFunc = async (row, flag) => {
  if (flag) {
    ElMessageBox.confirm(`This operation will automatically create files and API (delete table !!!), will it continue?,`, 'hint', {
      confirmButtonText: 'Sure',
      cancelButtonText: 'Cancel',
      type: 'warning'
    }).then(async () => {
      ElMessageBox.confirm(`This operation will automatically create files and APIs (delete table !!!), please continue to confirm! Intersection Intersection`, 'Will delete table', {
        confirmButtonText: 'Sure',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(async () => {
        ElMessageBox.confirm(`This operation will automatically create files and APIs (delete table !!!), please continue to confirm! Intersection Intersection`, 'Will delete table', {
          confirmButtonText: 'Sure',
          cancelButtonText: 'Cancel',
          type: 'warning'
        }).then(async () => {
          const res = await rollback({ id: Number(row.ID), deleteTable: flag })
          if (res.code === 0) {
            ElMessage.success('Rollback success')
            getTableData()
          }
        })
      })
    })
  } else {
    ElMessageBox.confirm(`Does this operation continue the files and API created by Delete?`, 'hint', {
      confirmButtonText: 'Sure',
      cancelButtonText: 'Cancel',
      type: 'warning'
    }).then(async () => {
      const res = await rollback({ id: Number(row.ID), deleteTable: flag })
      if (res.code === 0) {
        ElMessage.success('Rollback success')
        getTableData()
      }
    })
  }
}
const goAutoCode = (row) => {
  if (row) {
    router.push({
      name: 'autoCodeEdit', params: {
        id: row.ID
      }
    })
  } else {
    router.push({ name: 'autoCode' })
  }
}

</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;

  .el-button {
    float: right;
  }
}

.el-tag--mini {
  margin-left: 5px;
}

.warning {
  color: #dc143c;
}
</style>
