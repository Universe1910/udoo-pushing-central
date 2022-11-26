<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card gva-top-card">
        <div class="gva-top-card-left">
          <div class="gva-top-card-left-title">Good morning, Tinh TD, please start a day of work</div>
          <div class="gva-top-card-left-dot">{{ weatherInfo }}</div>
          <div class="gva-top-card-left-rows">
            <el-row>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dashboard-icon">
                    <sort />
                  </el-icon>
                  Today's traffic (1231231)
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dashboard-icon">
                    <avatar />
                  </el-icon>
                  Total number of users (24001)
                </div>
              </el-col>
              <el-col :span="8" :xs="24" :sm="8">
                <div class="flex-center">
                  <el-icon class="dashboard-icon">
                    <comment />
                  </el-icon>
                  Praise rate (99%)
                </div>
              </el-col>
            </el-row>
          </div>
          <div>
            <div class="gva-top-card-left-item">
              Udoo Platform:
              <a style="color:#409EFF" target="view_window" href="https://www.udoo.ooo">https://www.udoo.ooo</a>
            </div>
            <div class="gva-top-card-left-item">
              Udoo Work
              <a style="color:#409EFF" target="view_window" href="https://udoo.work">https://udoo.work</a>
            </div>
          </div>
        </div>
        <img src="@/assets/dashboard.png" class="gva-top-card-right" alt>
      </div>
    </div>
    <div class="gva-card-box">
      <el-card class="gva-card quick-entrance">
        <template #header>
          <div class="card-header">
            <span>Stats</span>
          </div>
        </template>
        <el-row :gutter="20" class="box-stats">
          <el-col :span="4" :xs="8">
            <div class="box-stats-container contact">
              <div class="box-stats-label">Contacts</div>
              <div class="box-stats-number">{{totalContact}}</div>
            </div>
          </el-col>
          <el-col :span="4" :xs="8" class="box-stats">
            <div class="box-stats-container campaign">
              <div class="box-stats-label">Campaigns</div>
              <div class="box-stats-number">{{totalCampaign}}</div>
            </div>
          </el-col>
          <el-col :span="4" :xs="8" class="box-stats">
            <div class="box-stats-container event">
              <div class="box-stats-label">Events</div>
              <div class="box-stats-number">{{totalEvents}}</div>
            </div>
          </el-col>
        </el-row>
      </el-card>
    </div>
    <div class="gva-card-box">
      <el-card class="gva-card quick-entrance">
        <template #header>
          <div class="card-header">
            <span>Navigation</span>
          </div>
        </template>
        <el-row :gutter="20">
          <el-col v-for="(card, key) in toolCards" :key="key" :span="4" :xs="8" class="quick-entrance-items"
            @click="toTarget(card.name)">
            <div class="quick-entrance-item">
              <div class="quick-entrance-item-icon" :style="{ backgroundColor: card.bg }">
                <el-icon>
                  <component :is="card.icon" :style="{ color: card.color }" />
                </el-icon>
              </div>
              <p>{{ card.label }}</p>
            </div>
          </el-col>
        </el-row>
      </el-card>
    </div>

    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>Statistics</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <el-col :xs="24" :sm="18">
              <echarts-line />
            </el-col>
            <el-col :xs="24" :sm="6">
              <dashboard-table />
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import EchartsLine from '@/view/dashboard/dashboardCharts/echartsLine.vue'
import DashboardTable from '@/view/dashboard/dashboardTable/dashboardTable.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useWeatherInfo } from '@/view/dashboard/weather.js'

import {
  getContactList
} from '@/api/contact'

import {
  getCampaignList
} from '@/api/campaign'


import {
  getCampaignLogList
} from '@/api/campaignLog'

const totalContact = ref(0)
const totalCampaign= ref(0)
const totalEvents = ref(0)

const getTotalContact = async () => {
  var res = await  getContactList()
  if(res.code == 0){
    totalContact.value = res.data.total
  }
}
getTotalContact()

const getTotalCampaign = async () => {
  var res = await  getCampaignList()
  if(res.code == 0){
    totalCampaign.value = res.data.total
  }
}

getTotalCampaign()

const getTotalEvents = async () => {
  var res = await  getCampaignLogList()
  if(res.code == 0){
    totalEvents.value = res.data.total
  }
}
getTotalEvents()

const weatherInfo = useWeatherInfo()

const toolCards = ref([
  {
    label: 'Users',
    icon: 'monitor',
    name: 'user',
    color: '#ff9c6e',
    bg: 'rgba(255, 156, 110,.3)'
  },
  {
    label: 'Role Management',
    icon: 'setting',
    name: 'authority',
    color: '#69c0ff',
    bg: 'rgba(105, 192, 255,.3)'
  },
  {
    label: 'Menu',
    icon: 'menu',
    name: 'menu',
    color: '#b37feb',
    bg: 'rgba(179, 127, 235,.3)'
  },
  {
    label: 'Code Generator',
    icon: 'cpu',
    name: 'autoCode',
    color: '#ffd666',
    bg: 'rgba(255, 214, 102,.3)'
  },
  {
    label: 'Form Generator',
    icon: 'document-checked',
    name: 'formCreate',
    color: '#ff85c0',
    bg: 'rgba(255, 133, 192,.3)'
  },
  {
    label: 'About Us',
    icon: 'user',
    name: 'about',
    color: '#5cdbd3',
    bg: 'rgba(92, 219, 211,.3)'
  }
])

const router = useRouter()

const toTarget = (name) => {
  router.push({ name })
}


</script>
<script>
export default {
  name: 'Dashboard'
}
</script>

<style lang="scss" scoped>
.box-stats-container {

  padding: 18px;
  border-radius: 4px;
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  align-content: center;
  align-items: center;
}

.contact {
  background-color: #FFE2D4;
}

.campaign {
  background-color: #D3ECFF;
}

.event {
  background-color: #FF4F3E;
}

.box-stats-number {
  color: #fff;
  font-size: 48px;
  font-weight: bold;
}

.box-stats-label {
  color: #191A23;
  font-size: 22px;
  font-weight: bold;
}

.box-stats {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

@mixin flex-center {
  display: flex;
  align-items: center;
}

.page {
  background: #f0f2f5;
  padding: 0;

  .gva-card-box {
    padding: 12px 16px;

    &+.gva-card-box {
      padding-top: 0px;
    }
  }

  .gva-card {
    box-sizing: border-box;
    background-color: #fff;
    border-radius: 2px;
    height: auto;
    padding: 26px 30px;
    overflow: hidden;
    box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
  }

  .gva-top-card {
    height: 260px;
    @include flex-center;
    justify-content: space-between;
    color: #777;

    &-left {
      height: 100%;
      display: flex;
      flex-direction: column;

      &-title {
        font-size: 22px;
        color: #343844;
      }

      &-dot {
        font-size: 16px;
        color: #6B7687;
        margin-top: 24px;
      }

      &-rows {
        // margin-top: 15px;
        margin-top: 18px;
        color: #6B7687;
        width: 600px;
        align-items: center;
      }

      &-item {
        +.gva-top-card-left-item {
          margin-top: 24px;
        }

        margin-top: 14px;
      }
    }

    &-right {
      height: 600px;
      width: 600px;
      margin-top: 28px;
    }
  }

  ::v-deep(.el-card__header) {
    padding: 0;
    border-bottom: none;
  }

  .card-header {
    padding-bottom: 20px;
    border-bottom: 1px solid #e8e8e8;
  }

  .quick-entrance-title {
    height: 30px;
    font-size: 22px;
    color: #333;
    width: 100%;
    border-bottom: 1px solid #eee;
  }

  .quick-entrance-items {
    @include flex-center;
    justify-content: center;
    text-align: center;
    color: #333;

    .quick-entrance-item {
      padding: 16px 28px;
      margin-top: -16px;
      margin-bottom: -16px;
      border-radius: 4px;
      transition: all 0.2s;

      &:hover {
        box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
      }

      cursor: pointer;
      height: auto;
      text-align: center;

      // align-items: center;
      &-icon {
        width: 50px;
        height: 50px !important;
        border-radius: 8px;
        @include flex-center;
        justify-content: center;
        margin: 0 auto;

        i {
          font-size: 24px;
        }
      }

      p {
        margin-top: 10px;
      }
    }
  }

  .echart-box {
    padding: 14px;
  }
}

.dashboard-icon {
  font-size: 20px;
  color: rgb(85, 160, 248);
  width: 30px;
  height: 30px;
  margin-right: 10px;
  @include flex-center;
}

.flex-center {
  @include flex-center;
}

//小屏幕不显示右侧，将登录框居中
@media (max-width: 750px) {
  .gva-card {
    padding: 20px 10px !important;

    .gva-top-card {
      height: auto;

      &-left {
        &-title {
          font-size: 20px !important;
        }

        &-rows {
          margin-top: 15px;
          align-items: center;
        }
      }

      &-right {
        display: none;
      }
    }

    .gva-middle-card {
      &-item {
        line-height: 20px;
      }
    }

    .dashboard-icon {
      font-size: 18px;
    }
  }
}
</style>
