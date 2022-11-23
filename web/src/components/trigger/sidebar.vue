<script setup>

const onDragStart = (event, nodeType, nodeLabel,actionName) => {
  if (event.dataTransfer) {
    event.dataTransfer.setData('application/vueflow', nodeType)
    event.dataTransfer.setData('nodeLabel', nodeLabel)
    event.dataTransfer.setData("actionName", actionName)
    event.dataTransfer.effectAllowed = 'move'
  }
}
</script>

<template>
  <aside>
    <div class="description">You can drag these nodes to the pane.</div>
    <div class="nodes">
      <div class="vue-flow__node-input input-node" id="input-add-contact" :draggable="true"
        @dragstart="onDragStart($event, 'input', 'Contact Added', 'event-add-contact')">
        <div class="input-label-group">Event</div>
        <div class="input-title">Contact Added</div>
        <div class="input-description">The trigger will starting when contact added</div>
      </div>
      <div class="vue-flow__node-input input-node" id="input-immediately" :draggable="true"
        @dragstart="onDragStart($event, 'input', 'Immediately', 'input-immediately')">
        <div class="input-label-group">Event</div>
        <div class="input-title">Immediately</div>
        <div class="input-description">The trigger will starting when called</div>
      </div>
      <div class="vue-flow__node-action action-node" id="action-send-zns" :draggable="true"
        @dragstart="onDragStart($event, 'action', 'Send ZNS', 'action-send-zns')">
        <div class="action-label-group">Action</div>
        <div class="action-title">Send ZNS</div>
        <div class="action-description">Send message with ZNS Template</div>
      </div>
      <div class="vue-flow__node-action action-node" id="action-send-email" :draggable="true"
        @dragstart="onDragStart($event, 'action', 'Send Email','action-send-email')">
        <div class="action-label-group">Action</div>
        <div class="action-title">Send Email</div>
        <div class="action-description">Send email to contact</div>
      </div>

      <div class="vue-flow__node-action action-node" id="action-send-messenger" :draggable="true"
        @dragstart="onDragStart($event, 'action', 'Send Messenger','action-send-messenger')">
        <div class="action-label-group">Action</div>
        <div class="action-title">Send Messenger</div>
        <div class="action-description">Send message by Facebook Messenger</div>
      </div>

      <div class="vue-flow__node-utils utils-node" id="utils-wait" :draggable="true"
        @dragstart="onDragStart($event, 'utils', 'Wait','utils-wait')">
        <div class="utils-label-group">Utils</div>
        <div class="utils-title">Wait</div>
        <div class="utils-description">Wait n second before action</div>
      </div>

    </div>
  </aside>
</template>


<style scoped>
.vue-flow__node-default,
.vue-flow__node-output{
  padding: 10px;
  border-radius: 3px;
  width: 150px;
  font-size: 12px;
  text-align: center;
  border-width: 1px;
  border-style: solid;
  color: var(--vf-node-text);
  background-color: var(--vf-node-bg);
  border-color: var(--vf-node-color);
}

.input-node, .action-node, .utils-node {
  background-color: #fff;
  color: black;
  border-radius: 4px;
  line-height: 1.5em;
}


.input-node .input-label-group {
  font-size: 16px;
  text-transform: uppercase;
  border-bottom: 1px dashed gray;
  background-color: #4e70ff;
  ;
  text-align: left;
  padding: 4px;
  color: #fff;
}

.input-node .input-title {
  font-size: 16px;
  text-transform: uppercase;
  border-bottom: 1px dashed gray;
  text-align: left;
  padding: 8px;
}

.input-node .input-description {
  font-size: 13px;
  font-style: italic;
  border-bottom: 1px dashed gray;
  text-align: left;
  padding: 8px;
}

.action-node .action-label-group {
  font-size: 16px;
  text-transform: uppercase;
  border-bottom: 1px dashed gray;
  background-color: #FF6464;
  ;
  text-align: left;
  padding: 4px;
  color: #fff;
}

.action-node .action-title {
  font-size: 16px;
  text-transform: uppercase;
  border-bottom: 1px dashed gray;
  text-align: left;
  padding: 8px;
}

.action-node .action-description {
  font-size: 13px;
  font-style: italic;
  border-bottom: 1px dashed gray;
  text-align: left;
  padding: 8px;
}

.utils-node .utils-label-group {
  font-size: 16px;
  text-transform: uppercase;
  border-bottom: 1px dashed gray;
  background-color: #F49D1A;
  ;
  text-align: left;
  padding: 4px;
  color: #fff;
}

.utils-node .utils-title {
  font-size: 16px;
  text-transform: uppercase;
  border-bottom: 1px dashed gray;
  text-align: left;
  padding: 8px;
}

.utils-node .utils-description {
  font-size: 13px;
  font-style: italic;
  border-bottom: 1px dashed gray;
  text-align: left;
  padding: 8px;
}
</style>
