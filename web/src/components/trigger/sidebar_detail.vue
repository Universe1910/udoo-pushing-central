<script setup>

const onNodeClick = (event, nodeType, nodeLabel,nodeId) => {
  if (event.dataTransfer) {
    event.dataTransfer.setData('application/vueflow', nodeType)
    event.dataTransfer.setData('nodeLabel', nodeLabel)
    event.dataTransfer.setData("nodeId", nodeId)
    // event.dataTransfer.effectAllowed = 'move'
  }
}
</script>

<template>
  <aside>
    <div class="description">Node Detail</div>
  </aside>
</template>


<style scoped>

</style>
