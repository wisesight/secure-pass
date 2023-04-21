const { spawn } = require("child_process");

const binaryPath = "./secure-pass-0.0.1-darwin-amd64";
const args = ["-d", "BO8dDZrsOFE3rODo2NFtIg=="];

const child = spawn(binaryPath, args);

child.stdout.on("data", (data) => {
  console.log(`stdout: ${data}`);
});

child.stderr.on("data", (data) => {
  console.error(`stderr: ${data}`);
});

child.on("close", (code) => {
  console.log(`child process exited with code ${code}`);
});
