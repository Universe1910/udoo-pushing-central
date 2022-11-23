
export function makeTriggerActions(edges) {
  var actions = [];
  for (var index = 0; index < edges.length ; index++) {
    var edge = edges[index];
    if (!isNodeExistAction(edge.source, actions)) {
      var nextVals = composeEdges(edge.source, edges);
      actions.push({
        source: edge.source,
        node: edge.sourceNode,
        nexts: nextVals,
      });
    }
  }
  return actions;
}

function isNodeExistAction(sourceId, actions) {
  var vals = actions.filter((e) => {
    return e.source == sourceId;
  });
  return vals.length >= 1;
}

function composeEdges(sourceId, edges) {
  var nexts = edges.filter((e) => {
    return e.source == sourceId;
  });
  var vals = nexts.map((e) => {
    return {
      target: e.target,
      node: e.targetNode,
    };
  });
  return vals;
}
