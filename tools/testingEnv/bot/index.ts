import * as puppeteer from "puppeteer";
import * as path from "path";
import * as fs from "fs";
import { spawn } from "child_process";

const { PendingXHR } = require("pending-xhr-puppeteer");

interface OutputJSON {
    pageHasBeenLoaded: boolean | null;
}

(async () => {
    const objWithOutputJSON: OutputJSON = {
        pageHasBeenLoaded: null
    };
    const pathToServer: string = path.join(__dirname, "../server/index.js");
    const server = await spawn("node", [pathToServer], {
        detached: true
    });
    if(!fs.existsSync(path.join(__dirname, "../screens"))) {
        await fs.mkdirSync(path.join(__dirname, "../screens"));
    }
    if(!fs.existsSync(path.join(__dirname, "../temp"))) {
        await fs.mkdirSync(path.join(__dirname, "../temp"));
    }
    const browser = await puppeteer.launch();
    const page = await browser.newPage();
    const pendingXHR = new PendingXHR(page);
    try {
        await page.goto("http://localhost:4000");
        await pendingXHR.waitForAllXhrFinished();
        await page.screenshot({ path: path.join(__dirname, "../screens/init-load.png") });
        objWithOutputJSON.pageHasBeenLoaded = true;
    } catch {
        objWithOutputJSON.pageHasBeenLoaded = false;
    }
    try {
        await page.screenshot({ path: path.join(__dirname, "../screens/load-end.png") });
        await browser.close();
    } catch {
        console.error("Cannot kill browser");
    }
    await server.kill();
    const convertedJSON = await JSON.stringify(objWithOutputJSON);
    try {
        await fs.writeFile(path.join(__dirname, "../temp/outputTest.json"), convertedJSON, "utf8", () => {
            console.log("Sucessfully saved JSON to file")
        });
    } catch(err) {
        console.log(err);
    }
    let npm: string;
    if(process.platform === "win32") {
        npm = 'npm.cmd'
    } else {
        npm = 'npm'
    }
    const jest = await spawn(npm, ["run", "jest"]);
    await jest.stderr.pipe(process.stderr);
    await jest.stdout.pipe(process.stdout);
    await jest.stdout.setEncoding("utf8");
    await jest.stdout.on("data", async (data) => {
        await console.log(data.toString());
    });
})();