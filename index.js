// @ts-check

const path = require("path");
const { name, version } = require("./package.json");

/** @type {import('caz').Template} */
module.exports = {
  name,
  version,
  prompts: [
    // TODO: custom template prompts
    {
      name: "name",
      type: "text",
      message: "Project name",
    },
    {
      name: "go_module_name",
      type: "text",
      message: "Go module name",
      initial: "app",
    },
  ],
  init: true,
  complete: async (ctx) => {
    // TODO: custom complete callback
    console.clear();
    console.log(
      `Created a new project in ${ctx.project} by the ${ctx.template} template.\n`,
    );
    console.log("Getting Started:");
    if (ctx.dest !== process.cwd()) {
      console.log(`  $ cd ${path.relative(process.cwd(), ctx.dest)}`);
    }
    console.log(`  $ make init`);
    console.log(`  $ make all`);
    console.log(`  $ make dev`);
    console.log("\nHappy hacking :)\n");
  },
};
