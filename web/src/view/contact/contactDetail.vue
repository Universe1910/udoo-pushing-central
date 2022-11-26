<template>
    <el-row :gutter="20">
        <el-col :span="16">
            <div class="grid-content ep-bg-purple">
                <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
                    <el-tab-pane label="Campaings" name="campaign">Campaings</el-tab-pane>
                    <el-tab-pane label="History" name="History">History</el-tab-pane>
                    <el-tab-pane label="Knowledge" name="Knowledge">Knowledge</el-tab-pane>

                </el-tabs>
            </div>
        </el-col>
        <el-col :span="8">
            <div class="grid-content ep-bg-purple">
                <div class="block-avatar">
                    <el-avatar :size="140" :src="circleUrl" />
                </div>
                <el-divider></el-divider>
                <el-descriptions title="Contact Infomation" direction="vertical" :column="1" :size="size" border>
                    <el-descriptions-item label="First name">{{contact.firstname}}</el-descriptions-item>
                    <el-descriptions-item label="Last name">{{contact.lastname}}</el-descriptions-item>
                    <el-descriptions-item label="Email" :span="2">{{contact.email}}</el-descriptions-item>
                    <el-descriptions-item label="Phone" :span="2">{{contact.phone}}</el-descriptions-item>
                    <el-descriptions-item label="Address" :span="2">{{contact.address}}</el-descriptions-item>
                    <el-descriptions-item label="State" :span="2">{{contact.state}}</el-descriptions-item>
                    <el-descriptions-item label="City" :span="2">{{contact.city}}</el-descriptions-item>
                    <el-descriptions-item label="Province" :span="2">{{contact.province}}</el-descriptions-item>
                    <el-descriptions-item label="Country" :span="2">{{contact.country}}</el-descriptions-item>
                    <el-descriptions-item label="Tags">
                        <el-tag v-for="tag in contact.tags" :key="tag.ID" size="small">{{tag.name}}</el-tag>
                    </el-descriptions-item>
                </el-descriptions>
            </div>
        </el-col>
    </el-row>
</template>

<script>
export default {
    name: "contactDetail",
};
</script>

<script setup>
import { ref, reactive } from "vue";
import { useRoute } from "vue-router";

import { findContact } from "@/api/contact";

const route = useRoute();
const searchInfo = ref({ contactID: route.params.id });
const contact = ref({
    CreatedAt: "2022-11-01T09:59:30.428774+07:00",
    ID: 2,
    UpdatedAt: "2022-11-01T10:02:17.675172+07:00",
    address: "29/23 Trần Khánh Dư, Phường 8",
    city: "Lâm Đồng",
    country: "Vietnam",
    email: "tinhpt.38@gmail.com",
    fbid: "1.tinhtd",
    firstname: "Phan",
    lastname: "Tính",
    phone: "+84865802659",
    province: "",
    state: "Lam Dong",
    tags: [],
    uuid: "00000000-0000-0000-0000-000000000000",
    zid: "0865802659",
    zipcode: "67000",
});

const circleUrl = reactive(
    "https://scontent.fdad3-4.fna.fbcdn.net/v/t39.30808-6/315211574_1810117899321918_6492145080887812975_n.jpg?_nc_cat=104&ccb=1-7&_nc_sid=09cbfe&_nc_ohc=Dm1t4lwIQbgAX-skWuS&tn=5ILtGfFjaJS3Eyyu&_nc_ht=scontent.fdad3-4.fna&oh=00_AfDKTGZ2Rhbceu-dJeB8xCykL-pWrfIU7llHzjQlYmApPA&oe=6385E6C5"
);

const getContact = async () => {
    var id = searchInfo.value.contactID;
    var res = await findContact({ ID: id });
    if (res.code === 0) {
        contact.value = res.data.recontact;
        console.log(contact.value);
    }
};

getContact();
</script>
<style scoped>
.el-row {
    margin-bottom: 20px;
}

.el-row:last-child {
    margin-bottom: 0;
}

.el-col {
    border-radius: 4px;
}

.grid-content {
    border-radius: 4px;
    min-height: 36px;
}

.demo-tabs>.el-tabs__content {
    padding: 32px;
    color: #6b778c;
    font-size: 32px;
    font-weight: 600;
}

.block-avatar {
    display: flex;
    flex-direction: row;
    justify-content: center;
}
</style>
