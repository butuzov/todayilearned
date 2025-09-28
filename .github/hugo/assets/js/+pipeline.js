function Pipeline(presetStages) {
  this.stages = presetStages || [];
}

Pipeline.prototype.pipe = function (stage) {
  this.stages.push(stage);

  return this;
};

Pipeline.prototype.process = function (args) {
  // Output is same as the passed args, if
  // there are no stages in the pipeline
  if (this.stages.length === 0) {
    return args;
  }

  // Set the stageOutput to be args
  // as there is no output to start with
  var stageOutput = args;

  this.stages.forEach(function (stage, counter) {
    // Output from the last stage was promise
    if (stageOutput && typeof stageOutput.then === "function") {
      // Call the next stage only when the promise is fulfilled
      stageOutput = stageOutput.then(stage);
    } else {
      // Otherwise, call the next stage with the last stage output
      if (typeof stage === "function") {
        stageOutput = stage(stageOutput);
      } else {
        stageOutput = stage;
      }
    }
  });

  return stageOutput;
};
