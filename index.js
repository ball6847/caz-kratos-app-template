// @ts-check

import { relative } from "path";
import { name, version } from "./package.json";

/** @type {import('caz').Template} */
export default {
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
  complete: async (ctx) => {
    // TODO: custom complete callback
    console.clear();
    console.log(
      `Created a new project in ${ctx.project} by the ${ctx.template} template.\n`,
    );
    console.log("Getting Started:");
    if (ctx.dest !== process.cwd()) {
      console.log(`  $ cd ${relative(process.cwd(), ctx.dest)}`);
    }
    if (ctx.config.install === false) {
      console.log(`  $ npm install`);
    }
    console.log(`  $ ${ctx.config.install ? ctx.config.install : "npm"} test`);
    console.log("\nHappy hacking :)\n");
  },
};
