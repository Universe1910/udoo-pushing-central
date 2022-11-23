<template>
      <q-card flat bordered class="my-card bg-white q-pa-md">
        <div class="row items-center no-wrap">
          <div class="col">
            <div class="text-h6">{{ title }}</div>
          </div>

          <div class="col-auto">
            <flowy-drag-handle>
              <q-btn size="sm" color="grey-7" round flat icon="drag_handle" />
            </flowy-drag-handle>
          </div>
        </div>

        <div class="q-py-md" v-html="description"/>
        <q-btn color="primary" class="q-pa-none" no-caps @click="remove()">Remove</q-btn>
      </q-card>
</template>

<script>
export default {
    name: 'DemoNode'
}
</script>

<script setup>
const props = defineProps(['remove', 'node', 'title', 'description'])
</script>