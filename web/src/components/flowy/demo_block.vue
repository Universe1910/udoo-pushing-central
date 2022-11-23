<template>
    <q-card flat bordered class="q-px-md q-py-sm row items-center justify-between">
  <div class="text-subtitle1">{{ title }}</div>
  <flowy-drag-handle>
    <q-btn color="grey-7" size="sm" class="q-ml-md" round flat icon="drag_handle" />
  </flowy-drag-handle>
</q-card>
</template>

<script>
export default {
    name: 'DemoBlock'
}
</script>

<script setup>
const props = defineProps(['remove', 'node', 'title', 'description'])
</script>